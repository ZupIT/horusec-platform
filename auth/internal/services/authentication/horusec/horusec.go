package horusec

import (
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	horusecAuthEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/horusec"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	authRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication"
	authUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

type Service struct {
	accountRepository accountRepository.IRepository
	authUseCases      authUseCases.IUseCases
	cache             *cache.Cache
	authRepository    authRepository.IRepository
	appConfig         app.IConfig
}

func NewHorusecAuthenticationService(repositoryAccount accountRepository.IRepository, appConfig app.IConfig,
	useCasesAuth authUseCases.IUseCases, repositoryAuth authRepository.IRepository) authentication.IService {
	return &Service{
		cache:             cache.New(authEnums.TokenDuration, authEnums.TokenCheckExpiredDuration),
		authUseCases:      useCasesAuth,
		accountRepository: repositoryAccount,
		authRepository:    repositoryAuth,
		appConfig:         appConfig,
	}
}

func (s *Service) Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error) {
	account, err := s.accountRepository.GetAccountByEmail(credentials.Username)
	if err != nil {
		return nil, horusecAuthEnums.ErrorWrongEmailOrPassword
	}

	if err := s.authUseCases.CheckLoginData(credentials, account); err != nil {
		return nil, err
	}

	return s.setTokensAndResponse(account)
}

func (s *Service) setTokensAndResponse(account *accountEntities.Account) (*authEntities.LoginResponse, error) {
	refreshToken := jwt.CreateRefreshToken()
	accessToken, expireAt, _ := jwt.CreateToken(account.ToTokenData(), nil)

	s.setRefreshTokenCache(account.AccountID.String(), refreshToken)

	return account.ToLoginResponse(accessToken, refreshToken, expireAt), nil
}

func (s *Service) setRefreshTokenCache(accountID, refreshToken string) {
	s.cache.Delete(refreshToken)
	_ = s.cache.Add(refreshToken, accountID, authEnums.TokenDuration)
}

func (s *Service) IsAuthorized(data *authEntities.AuthorizationData) (bool, error) {
	return s.authorizeByRole()[data.Type](data)
}

func (s *Service) authorizeByRole() map[auth.AuthorizationType]func(*authEntities.AuthorizationData) (bool, error) {
	return map[auth.AuthorizationType]func(*authEntities.AuthorizationData) (bool, error){
		auth.WorkspaceMember:      s.isWorkspaceMember,
		auth.WorkspaceAdmin:       s.isWorkspaceAdmin,
		auth.RepositoryMember:     s.isRepositoryMember,
		auth.RepositorySupervisor: s.isRepositorySupervisor,
		auth.RepositoryAdmin:      s.isRepositoryAdmin,
		auth.ApplicationAdmin:     s.isApplicationAdmin,
	}
}

func (s *Service) checkForMember(role accountEnums.Role, accountID uuid.UUID) bool {
	isApplicationAdmin, _ := s.checkForApplicationAdmin(accountID)

	return isApplicationAdmin || role == accountEnums.Admin ||
		role == accountEnums.Supervisor || role == accountEnums.Member
}

func (s *Service) checkForSupervisor(role accountEnums.Role, accountID uuid.UUID) bool {
	isApplicationAdmin, _ := s.checkForApplicationAdmin(accountID)

	return isApplicationAdmin || role == accountEnums.Admin ||
		role == accountEnums.Supervisor
}

func (s *Service) checkForAdmin(role accountEnums.Role, accountID uuid.UUID) bool {
	isApplicationAdmin, _ := s.checkForApplicationAdmin(accountID)

	return isApplicationAdmin || role == accountEnums.Admin
}

func (s *Service) checkForApplicationAdmin(accountID uuid.UUID) (bool, error) {
	if !s.appConfig.IsApplicationAdminEnabled() {
		return false, horusecAuthEnums.ErrorApplicationAdminNotEnabled
	}

	account, err := s.accountRepository.GetAccount(accountID)
	if err != nil {
		return false, errors.Wrap(err, horusecAuthEnums.ErrorFailedToGetAccountAppAdmin)
	}

	return account.IsApplicationAdmin, nil
}

func (s *Service) isWorkspaceMember(data *authEntities.AuthorizationData) (bool, error) {
	accountID, err := jwt.GetAccountIDByJWTToken(data.Token)
	if err != nil {
		return false, errors.Wrap(err, horusecAuthEnums.ErrorFailedToGetAccountIDFromToken)
	}

	role, err := s.authRepository.GetWorkspaceRole(accountID, data.WorkspaceID)
	if err != nil {
		return false, errors.Wrap(err, horusecAuthEnums.ErrorFailedToGetWorkspaceRole)
	}

	return s.checkForMember(role, accountID), nil
}

func (s *Service) isWorkspaceAdmin(data *authEntities.AuthorizationData) (bool, error) {
	accountID, err := jwt.GetAccountIDByJWTToken(data.Token)
	if err != nil {
		return false, errors.Wrap(err, horusecAuthEnums.ErrorFailedToGetAccountIDFromToken)
	}

	role, err := s.authRepository.GetWorkspaceRole(accountID, data.WorkspaceID)
	if err != nil {
		return false, errors.Wrap(err, horusecAuthEnums.ErrorFailedToGetWorkspaceRole)
	}

	return s.checkForAdmin(role, accountID), nil
}

func (s *Service) isRepositoryMember(data *authEntities.AuthorizationData) (bool, error) {
	accountID, err := jwt.GetAccountIDByJWTToken(data.Token)
	if err != nil {
		return false, errors.Wrap(err, horusecAuthEnums.ErrorFailedToGetAccountIDFromToken)
	}

	role, err := s.authRepository.GetRepositoryRole(accountID, data.WorkspaceID)
	if err != nil {
		return s.checkRepositoryRequestForWorkspaceAdmin(data, err)
	}

	return s.checkForMember(role, accountID), nil
}

func (s *Service) isRepositorySupervisor(data *authEntities.AuthorizationData) (bool, error) {
	accountID, err := jwt.GetAccountIDByJWTToken(data.Token)
	if err != nil {
		return false, errors.Wrap(err, horusecAuthEnums.ErrorFailedToGetAccountIDFromToken)
	}

	role, err := s.authRepository.GetRepositoryRole(accountID, data.WorkspaceID)
	if err != nil {
		return s.checkRepositoryRequestForWorkspaceAdmin(data, err)
	}

	return s.checkForSupervisor(role, accountID), nil
}

func (s *Service) isRepositoryAdmin(data *authEntities.AuthorizationData) (bool, error) {
	accountID, err := jwt.GetAccountIDByJWTToken(data.Token)
	if err != nil {
		return false, errors.Wrap(err, horusecAuthEnums.ErrorFailedToGetAccountIDFromToken)
	}

	role, err := s.authRepository.GetRepositoryRole(accountID, data.WorkspaceID)
	if err != nil {
		return s.checkRepositoryRequestForWorkspaceAdmin(data, err)
	}

	return s.checkForAdmin(role, accountID), nil
}

func (s *Service) checkRepositoryRequestForWorkspaceAdmin(data *authEntities.AuthorizationData,
	err error) (bool, error) {
	isWorkspaceAdmin, workspaceErr := s.isWorkspaceAdmin(data)
	if workspaceErr != nil {
		return isWorkspaceAdmin, errors.Wrap(workspaceErr, err.Error())
	}

	return isWorkspaceAdmin, err
}

func (s *Service) isApplicationAdmin(data *authEntities.AuthorizationData) (bool, error) {
	accountID, err := jwt.GetAccountIDByJWTToken(data.Token)
	if err != nil {
		return false, errors.Wrap(err, horusecAuthEnums.ErrorFailedToGetAccountIDFromToken)
	}

	return s.checkForApplicationAdmin(accountID)
}

func (s *Service) GetAccountFromToken(token string) (*proto.GetAccountDataResponse, error) {
	claims, err := jwt.DecodeToken(token)
	if err != nil {
		return nil, err
	}

	account, err := s.accountRepository.GetAccount(parser.ParseStringToUUID(claims.Subject))
	if err != nil {
		return nil, err
	}

	return account.ToGetAccountDataResponse(claims.Permissions), nil
}

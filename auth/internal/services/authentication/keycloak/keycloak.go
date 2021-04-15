package keycloak

import (
	"github.com/Nerzal/gocloak/v7"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	
	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	horusecAuthEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/horusec"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	authRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/authentication"
	keycloak "github.com/ZupIT/horusec-platform/auth/internal/services/authentication/keycloak/client"
	authUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

type IService interface {
	Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error)
}

type Service struct {
	accountRepository accountRepository.IRepository
	authUseCases      authUseCases.IUseCases
	cache             *cache.Cache
	authRepository    authRepository.IRepository
	appConfig         app.IConfig
	keycloak          keycloak.IClient
}

func NewKeycloakAuthenticationService(repositoryAccount accountRepository.IRepository, appConfig app.IConfig,
	useCasesAuth authUseCases.IUseCases, repositoryAuth authRepository.IRepository) IService {
	return &Service{
		keycloak:          keycloak.NewKeycloakClient(),
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

	token, err := s.keycloak.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		return nil, err
	}

	return s.setLoginResponse(account, token), nil
}

func (s *Service) setLoginResponse(account *accountEntities.Account, token *gocloak.JWT) *authEntities.LoginResponse {
	return &authEntities.LoginResponse{
		AccountID:          account.AccountID,
		AccessToken:        token.AccessToken,
		RefreshToken:       token.RefreshToken,
		Username:           account.Username,
		Email:              account.Email,
		ExpiresIn:          token.ExpiresIn,
		RefreshExpiresIn:   token.RefreshExpiresIn,
		IsApplicationAdmin: account.IsApplicationAdmin,
	}
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

	return isWorkspaceAdmin, nil
}

func (s *Service) isApplicationAdmin(data *authEntities.AuthorizationData) (bool, error) {
	accountID, err := jwt.GetAccountIDByJWTToken(data.Token)
	if err != nil {
		return false, errors.Wrap(err, horusecAuthEnums.ErrorFailedToGetAccountIDFromToken)
	}

	return s.checkForApplicationAdmin(accountID)
}

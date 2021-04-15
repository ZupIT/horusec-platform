package ldap

import (
	"strings"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"

	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/ldap/client"
	authUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/ldap"
	authRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/authentication"
	"github.com/ZupIT/horusec-platform/auth/config/app"
)

type IService interface {
	Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error)
}

type Service struct {
	ldap              client.ILdapClient
	accountRepository accountRepository.IRepository
	authRepository    authRepository.IRepository
	authUseCases      authUseCases.IUseCases
	appConfig         app.IConfig
}

func NewLDAPAuthenticationService(repositoryAccount accountRepository.IRepository, useCasesAuth authUseCases.IUseCases,
	appConfig app.IConfig, repositoryAuth authRepository.IRepository) IService {
	return &Service{
		ldap:              client.NewLdapClient(),
		accountRepository: repositoryAccount,
		authUseCases:      useCasesAuth,
		appConfig:         appConfig,
		authRepository:    repositoryAuth,
	}
}

func (s *Service) Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error) {
	isAuthenticated, userData, err := s.ldap.Authenticate(credentials.Username, credentials.Password)
	if err != nil || !isAuthenticated {
		return nil, s.verifyAuthenticateErrors(err)
	}

	account, err := s.getAccountOrCreateIfNotExist(userData)
	if err != nil {
		return nil, s.verifyAuthenticateErrors(err)
	}

	defer s.ldap.Close()
	return s.setTokenAndResponse(account, userData["dn"])
}

func (s *Service) verifyAuthenticateErrors(err error) error {
	if err != nil && err == ldap.ErrorUserDoesNotExist {
		return err
	}

	return ldap.ErrorLdapUnauthorized
}

func (s *Service) getAccountOrCreateIfNotExist(userData map[string]string) (*accountEntities.Account, error) {
	account, err := s.accountRepository.GetAccountByUsername(userData["sAMAccountName"])
	if account == nil || err != nil {
		return s.accountRepository.CreateAccount(s.authUseCases.SetLdapAccountData(userData))
	}

	return account, nil
}

func (s *Service) setTokenAndResponse(account *accountEntities.Account,
	userDN string) (*authEntities.LoginResponse, error) {
	userGroups, err := s.ldap.GetUserGroups(userDN)
	if err != nil {
		return nil, err
	}

	accessToken, expiresAt, _ := jwt.CreateToken(account.ToTokenData(), userGroups)
	return &authEntities.LoginResponse{
		AccessToken:        accessToken,
		ExpiresAt:          expiresAt,
		Username:           account.Username,
		Email:              account.Email,
		IsApplicationAdmin: s.isApplicationAdmin(userGroups),
	}, nil
}

func (s *Service) isApplicationAdmin(userGroups []string) bool {
	applicationAdminGroup, _ := s.getApplicationAdminAuthzGroupName()
	return s.checkIsAuthorized(applicationAdminGroup, userGroups)
}

func (s *Service) getApplicationAdminAuthzGroupName() ([]string, error) {
	applicationAdminGroup := env.GetEnvOrDefault(ldap.EnvLdapAdminGroup, "")

	if applicationAdminGroup == "" && s.appConfig.IsApplicationAdminEnabled() {
		return []string{}, ldap.ErrorLdapApplicationAdminGroupNotSet
	}

	return []string{applicationAdminGroup}, nil
}

func (s *Service) checkIsAuthorized(tokenGroups, horusecGroups []string) bool {
	for _, tokenGroup := range tokenGroups {
		if s.contains(horusecGroups, tokenGroup) {
			return true
		}
	}

	return false
}

func (s *Service) contains(horusecGroups []string, tokenGroup string) bool {
	for _, horusecGroup := range horusecGroups {
		if strings.TrimSpace(horusecGroup) == tokenGroup {
			return true
		}
	}

	return false
}

func (s *Service) IsAuthorized(data *authEntities.AuthorizationData) (bool, error) {
	tokenGroups, err := s.getUserGroupsFromJWT(data.Token)
	if err != nil {
		return false, err
	}

	horusecGroups, err := s.getHorusecAuthzGroups(data)
	if err != nil {
		return false, err
	}

	return s.checkIsAuthorized(tokenGroups, horusecGroups), nil
}

func (s *Service) getUserGroupsFromJWT(tokenStr string) ([]string, error) {
	token, err := jwt.DecodeToken(tokenStr)
	if err != nil {
		return nil, err
	}

	return token.Permissions, nil
}

func (s *Service) getHorusecAuthzGroups(data *authEntities.AuthorizationData) ([]string, error) {
	switch data.Type {
	case auth.WorkspaceAdmin, auth.WorkspaceMember:
		return s.getWorkspaceAuthzGroups(data)
	case auth.RepositoryAdmin, auth.RepositorySupervisor, auth.RepositoryMember:
		return s.getRepositoryAuthzGroups(data)
	}

	return s.getGroupsByAuthorizationType(data)
}

func (s *Service) getWorkspaceAuthzGroups(data *authEntities.AuthorizationData) ([]string, error) {
	workspaceGroups, err := s.authRepository.GetWorkspaceGroups(data.WorkspaceID)
	if err != nil {
		return nil, err
	}

	return s.getGroupsByAuthorizationType(data.SetGroups(workspaceGroups))
}

func (s *Service) getRepositoryAuthzGroups(data *authEntities.AuthorizationData) ([]string, error) {
	workspaceGroups, err := s.authRepository.GetWorkspaceGroups(data.WorkspaceID)
	if err != nil {
		return nil, err
	}

	repositoryGroups, err := s.authRepository.GetRepositoryGroups(data.RepositoryID)
	if err != nil {
		return nil, err
	}

	groups, err := s.getGroupsByAuthorizationType(data.SetGroups(repositoryGroups))
	return append(groups, workspaceGroups.AuthzAdmin...), err
}

func (s *Service) getGroupsByAuthorizationType(data *authEntities.AuthorizationData) (groups []string, err error) {
	appAdminAuthz, err := s.getApplicationAdminAuthzGroupName()
	if err != nil {
		return nil, err
	}

	switch data.Type {
	case auth.ApplicationAdmin:
		groups = appAdminAuthz
	case auth.RepositoryAdmin, auth.WorkspaceAdmin:
		groups = append(appAdminAuthz, data.AuthzAdmin...)
	case auth.RepositorySupervisor:
		groups = append(appAdminAuthz, append(data.AuthzAdmin, data.AuthzSupervisor...)...)
	case auth.RepositoryMember, auth.WorkspaceMember:
		groups = append(appAdminAuthz, append(data.AuthzAdmin, append(data.AuthzSupervisor, data.AuthzMember...)...)...)
	}

	return groups, err
}

// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ldap

import (
	"strings"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/cache"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	ldapEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/ldap"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	authRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/ldap/client"
	authUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

type Service struct {
	ldap              client.ILdapClient
	accountRepository accountRepository.IRepository
	authRepository    authRepository.IRepository
	authUseCases      authUseCases.IUseCases
	appConfig         app.IConfig
	cache             cache.ICache
}

type IService interface {
	Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error)
	IsAuthorized(data *authEntities.AuthorizationData) (bool, error)
	GetAccountDataFromToken(token string) (*proto.GetAccountDataResponse, error)
}

func NewLDAPAuthenticationService(repositoryAccount accountRepository.IRepository, useCasesAuth authUseCases.IUseCases,
	appConfig app.IConfig, repositoryAuth authRepository.IRepository, cacheLib cache.ICache) IService {
	return &Service{
		cache:             cacheLib,
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
		return nil, s.verifyAuthenticateErrors(err, isAuthenticated)
	}

	account, err := s.getAccountOrCreateIfNotExist(userData)
	if err != nil {
		return nil, err
	}

	defer s.ldap.Close()
	return s.setTokenAndResponse(account, userData["dn"])
}

func (s *Service) verifyAuthenticateErrors(err error, isAuthenticated bool) error {
	if err == nil && !isAuthenticated {
		return ldapEnums.ErrorLdapUnauthorized
	}

	return err
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

	return s.newLoginResponse(account, userGroups)
}

func (s *Service) newLoginResponse(account *accountEntities.Account,
	userGroups []string) (*authEntities.LoginResponse, error) {
	refreshToken := jwt.CreateRefreshToken()
	s.setRefreshTokenCache(account.AccountID.String(), refreshToken)

	accessToken, expiresAt, _ := jwt.CreateToken(account.ToTokenData(), userGroups)
	return &authEntities.LoginResponse{
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		ExpiresAt:          expiresAt,
		Username:           account.Username,
		Email:              account.Email,
		IsApplicationAdmin: s.isApplicationAdmin(userGroups),
	}, nil
}

func (s *Service) setRefreshTokenCache(accountID, refreshToken string) {
	s.cache.Delete(refreshToken)
	s.cache.Set(refreshToken, accountID, authEnums.TokenDuration)
}

func (s *Service) isApplicationAdmin(userGroups []string) bool {
	applicationAdminGroup, _ := s.getApplicationAdminAuthzGroupName()
	return s.checkIsAuthorized(applicationAdminGroup, userGroups)
}

func (s *Service) getApplicationAdminAuthzGroupName() ([]string, error) {
	applicationAdminGroup := env.GetEnvOrDefault(ldapEnums.EnvLdapAdminGroup, "")

	if applicationAdminGroup == "" && s.appConfig.IsApplicationAdmEnabled() {
		return []string{}, ldapEnums.ErrorLdapApplicationAdminGroupNotSet
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
	case auth.ApplicationAdmin:
		return s.getGroupsByAuthorizationType(data)
	case auth.WorkspaceAdmin, auth.WorkspaceMember:
		return s.getWorkspaceAuthzGroups(data)
	case auth.RepositoryAdmin, auth.RepositorySupervisor, auth.RepositoryMember:
		return s.getRepositoryAuthzGroups(data)
	}

	return nil, ldapEnums.ErrorInvalidAuthorizationType
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

	return s.getGroupsByType(appAdminAuthz, data), err
}

func (s *Service) getGroupsByType(appAdminAuthz []string, data *authEntities.AuthorizationData) (groups []string) {
	switch data.Type {
	case auth.ApplicationAdmin:
		groups = appAdminAuthz
	case auth.RepositoryAdmin, auth.WorkspaceAdmin:
		groups = s.appendAdmin(appAdminAuthz, data)
	case auth.RepositorySupervisor:
		groups = s.appendSupervisor(appAdminAuthz, data)
	case auth.RepositoryMember, auth.WorkspaceMember:
		groups = s.appendMember(appAdminAuthz, data)
	}

	return groups
}

func (s *Service) appendAdmin(appAdminAuthz []string, data *authEntities.AuthorizationData) []string {
	return append(appAdminAuthz, data.AuthzAdmin...)
}

func (s *Service) appendSupervisor(appAdminAuthz []string, data *authEntities.AuthorizationData) []string {
	return append(appAdminAuthz, append(data.AuthzAdmin, data.AuthzSupervisor...)...)
}

func (s *Service) appendMember(appAdminAuthz []string, data *authEntities.AuthorizationData) []string {
	return append(appAdminAuthz, append(data.AuthzAdmin, append(data.AuthzSupervisor, data.AuthzMember...)...)...)
}

func (s *Service) GetAccountDataFromToken(token string) (*proto.GetAccountDataResponse, error) {
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

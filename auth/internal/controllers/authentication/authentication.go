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

package authentication

import (
	authTypes "github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/horusec"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/keycloak"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/ldap"
)

type IController interface {
	Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error)
	IsAuthorized(data *authEntities.AuthorizationData) (bool, error)
	GetAccountInfo(token string) (*proto.GetAccountDataResponse, error)
	GetAccountInfoByEmail(email string) (*proto.GetAccountDataResponse, error)
}

type Controller struct {
	appConfig         app.IConfig
	horusecAuth       horusec.IService
	keycloakAuth      keycloak.IService
	ldapAuth          ldap.IService
	accountRepository accountRepository.IRepository
}

func NewAuthenticationController(appConfig app.IConfig, authHorusec horusec.IService, ldapAuth ldap.IService,
	keycloakAuth keycloak.IService, repositoryAccount accountRepository.IRepository) IController {
	return &Controller{
		appConfig:         appConfig,
		horusecAuth:       authHorusec,
		ldapAuth:          ldapAuth,
		keycloakAuth:      keycloakAuth,
		accountRepository: repositoryAccount,
	}
}

func (c *Controller) Login(credentials *authEntities.LoginCredentials) (*authEntities.LoginResponse, error) {
	switch c.appConfig.GetAuthenticationType() {
	case authTypes.Horusec:
		return c.horusecAuth.Login(credentials)
	case authTypes.Keycloak:
		return c.keycloakAuth.Login(credentials)
	case authTypes.Ldap:
		return c.ldapAuth.Login(credentials)
	}

	return nil, authEnums.ErrorAuthTypeInvalid
}

func (c *Controller) IsAuthorized(data *authEntities.AuthorizationData) (bool, error) {
	switch c.appConfig.GetAuthenticationType() {
	case authTypes.Horusec:
		return c.horusecAuth.IsAuthorized(data)
	case authTypes.Keycloak:
		return c.keycloakAuth.IsAuthorized(data)
	case authTypes.Ldap:
		return c.ldapAuth.IsAuthorized(data)
	}

	return false, authEnums.ErrorAuthTypeInvalid
}

func (c *Controller) GetAccountInfo(token string) (*proto.GetAccountDataResponse, error) {
	switch c.appConfig.GetAuthenticationType() {
	case authTypes.Horusec:
		return c.horusecAuth.GetAccountDataFromToken(token)
	case authTypes.Keycloak:
		return c.keycloakAuth.GetAccountDataFromToken(token)
	case authTypes.Ldap:
		return c.ldapAuth.GetAccountDataFromToken(token)
	}

	return nil, authEnums.ErrorAuthTypeInvalid
}

func (c *Controller) GetAccountInfoByEmail(email string) (*proto.GetAccountDataResponse, error) {
	if email == "" {
		return nil, authEnums.ErrorEmailEmpty
	}
	account, err := c.accountRepository.GetAccountByEmail(email)
	if err != nil {
		return nil, err
	}

	return account.ToGetAccountDataResponse(nil), nil
}

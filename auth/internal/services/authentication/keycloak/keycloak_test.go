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

package keycloak

import (
	"errors"
	"testing"

	"github.com/Nerzal/gocloak/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	keycloakEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/keycloak"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	authRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/authentication"
	keycloak "github.com/ZupIT/horusec-platform/auth/internal/services/authentication/keycloak/client"
	"github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

func TestNewKeycloakAuthenticationService(t *testing.T) {
	t.Run("should success create a new service", func(t *testing.T) {
		assert.NotNil(t, NewKeycloakAuthenticationService(
			nil, nil, nil, nil))
	})
}

func TestLogin(t *testing.T) {
	t.Run("should success create a new service", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		account := &accountEntities.Account{
			IsConfirmed:        true,
			Email:              "test@test.com",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(account, nil)

		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("Authenticate").Return(&gocloak.JWT{}, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		content, err := service.Login(&authEntities.LoginCredentials{})

		assert.NoError(t, err)
		assert.NotEmpty(t, content)
	})

	t.Run("should return error when failed to authenticate in keycloak", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		account := &accountEntities.Account{
			IsConfirmed:        true,
			Email:              "test@test.com",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(account, nil)

		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("Authenticate").Return(&gocloak.JWT{}, errors.New("test"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		content, err := service.Login(&authEntities.LoginCredentials{})

		assert.Error(t, err)
		assert.Nil(t, content)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}
		keycloakMock := &keycloak.Mock{}

		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        true,
			Email:              "test@test.com",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(account, errors.New("test"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		content, err := service.Login(&authEntities.LoginCredentials{})

		assert.Error(t, err)
		assert.Nil(t, content)
	})
}

func TestIsAuthorizedApplicationAdmin(t *testing.T) {
	t.Run("should return true and no error when application admin", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, nil)

		appConfig := &app.Config{EnableApplicationAdmin: true}

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, errors.New("test"))

		appConfig := &app.Config{EnableApplicationAdmin: true}

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when application admin not enabled", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, errors.New("test"))

		appConfig := &app.Config{EnableApplicationAdmin: false}

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.Nil, errors.New("error"))

		appConfig := &app.Config{EnableApplicationAdmin: true}

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedWorkspaceAdmin(t *testing.T) {
	t.Run("should return true and no error when workspace admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Admin, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.WorkspaceAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get workspace role", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Admin, errors.New("test"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.WorkspaceAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.Nil, errors.New("errornil"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.WorkspaceAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedWorkspaceMember(t *testing.T) {
	t.Run("should return true and no error when workspace member", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}

		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.WorkspaceMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get workspace role", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}

		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, errors.New("test"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.WorkspaceMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.Nil, errors.New("error"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.WorkspaceMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedRepositoryMember(t *testing.T) {
	t.Run("should return true and no error when repository member", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}

		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Member, nil)
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositoryMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get repository role", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}

		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Member, errors.New("test"))
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, errors.New("test"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositoryMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.Nil, errors.New("error"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.RepositoryMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedRepositorySupervisor(t *testing.T) {
	t.Run("should return true and no error when repository supervisor", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}

		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Supervisor, nil)
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Supervisor, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositorySupervisor,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get repository role", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}

		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Supervisor, errors.New("test"))
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Supervisor, errors.New("test"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositorySupervisor,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}
		keycloakMock := &keycloak.Mock{}

		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.Nil, errors.New("error"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.RepositorySupervisor,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedRepositoryAdmin(t *testing.T) {
	t.Run("should return true and no error when repository admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}

		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Admin, nil)
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Admin, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositoryAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return error when failed to get repository role", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}

		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			IsConfirmed:        false,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: false,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetRepositoryRole").Return(accountEnums.Admin, errors.New("test"))
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         auth.RepositoryAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return error when failed to get account id from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}
		keycloakMock := &keycloak.Mock{}

		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.Nil, errors.New("error"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		data := &authEntities.AuthorizationData{
			Token:        "test",
			Type:         auth.RepositoryAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestGetAccountDataFromToken(t *testing.T) {
	t.Run("should return account data without errors", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		sub := uuid.NewString()
		userInfo := &gocloak.UserInfo{Sub: &sub}

		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetUserInfo").Return(userInfo, nil)

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		result, err := service.GetAccountDataFromToken("test")
		assert.NoError(t, err)
		assert.Equal(t, account.AccountID.String(), result.AccountID)
		assert.Equal(t, account.IsApplicationAdmin, result.IsApplicationAdmin)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		sub := uuid.NewString()
		userInfo := &gocloak.UserInfo{Sub: &sub}

		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetUserInfo").Return(userInfo, nil)

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, errors.New("test"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		result, err := service.GetAccountDataFromToken("test")
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when failed to get user info", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}
		accountRepositoryMock := &accountRepository.Mock{}

		sub := uuid.NewString()
		userInfo := &gocloak.UserInfo{Sub: &sub}

		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetUserInfo").Return(userInfo, errors.New("test"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		result, err := service.GetAccountDataFromToken("test")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestGetUserInfo(t *testing.T) {
	t.Run("should success get user info", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}
		accountRepositoryMock := &accountRepository.Mock{}

		test := "test"
		userInfo := &gocloak.UserInfo{
			Sub:   &test,
			Email: &test,
		}

		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetUserInfo").Return(userInfo, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		result, err := service.GetUserInfo("test")
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when missing username or email", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}
		accountRepositoryMock := &accountRepository.Mock{}

		userInfo := &gocloak.UserInfo{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetUserInfo").Return(userInfo, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		result, err := service.GetUserInfo("test")
		assert.Error(t, err)
		assert.Equal(t, keycloakEnums.ErrorKeycloakMissingUsernameOrSub, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when missing username or email", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}
		accountRepositoryMock := &accountRepository.Mock{}

		userInfo := &gocloak.UserInfo{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetUserInfo").Return(userInfo, errors.New("test"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		result, err := service.GetUserInfo("test")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestCheckRepositoryRequestForWorkspaceAdmin(t *testing.T) {
	t.Run("should return true for workspace admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Admin, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})
		data := &authEntities.AuthorizationData{
			Token: token,
		}

		result, err := service.checkRepositoryRequestForWorkspaceAdmin(data, enums.ErrorNotFoundRecords)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should return false for workspace admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, nil)

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})
		data := &authEntities.AuthorizationData{
			Token: token,
		}

		result, err := service.checkRepositoryRequestForWorkspaceAdmin(data, enums.ErrorNotFoundRecords)
		assert.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("should return false and error when failed to check for workspace admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceRole").Return(accountEnums.Member, errors.New("test"))

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})
		data := &authEntities.AuthorizationData{
			Token: token,
		}

		result, err := service.checkRepositoryRequestForWorkspaceAdmin(data, enums.ErrorNotFoundRecords)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should return false and error when error different than not found records", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}
		keycloakMock := &keycloak.Mock{}
		keycloakMock.On("GetAccountIDByJWTToken").Return(uuid.New(), nil)

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		service := Service{
			accountRepository: accountRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			authRepository:    authRepositoryMock,
			appConfig:         appConfig,
			keycloak:          keycloakMock,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})
		data := &authEntities.AuthorizationData{
			Token: token,
		}

		result, err := service.checkRepositoryRequestForWorkspaceAdmin(data, errors.New("test"))
		assert.Error(t, err)
		assert.False(t, result)
	})
}

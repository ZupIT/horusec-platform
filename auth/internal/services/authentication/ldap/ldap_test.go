package ldap

import (
	"errors"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	authorization "github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/cache"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	ldapEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/ldap"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	authRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/ldap/client"
	"github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

func TestNewLDAPAuthenticationService(t *testing.T) {
	t.Run("should success create a new service", func(t *testing.T) {
		assert.NotNil(t, NewLDAPAuthenticationService(nil, nil, nil,
			nil, nil))
	})
}

func TestLogin(t *testing.T) {
	t.Run("should login creating account without errors", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		account := &accountEntities.Account{
			Username: "test",
			Email:    "test@test.com",
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByUsername").Return(account, errors.New("test"))
		accountRepositoryMock.On("CreateAccount").Return(account, nil)

		ldapMock := &client.Mock{}
		ldapMock.On("Authenticate").Return(true, map[string]string{}, nil)
		ldapMock.On("GetUserGroups").Return([]string{}, nil)
		ldapMock.On("Close")

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		result, err := service.Login(&authEntities.LoginCredentials{})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.AccessToken)
		assert.NotEmpty(t, result.RefreshToken)
		assert.False(t, result.IsApplicationAdmin)
		assert.Equal(t, "test", result.Username)
		assert.Equal(t, "test@test.com", result.Email)
	})

	t.Run("should login with existing account without errors", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		account := &accountEntities.Account{
			Username: "test",
			Email:    "test@test.com",
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByUsername").Return(account, nil)

		ldapMock := &client.Mock{}
		ldapMock.On("Authenticate").Return(true, map[string]string{}, nil)
		ldapMock.On("GetUserGroups").Return([]string{}, nil)
		ldapMock.On("Close")

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		result, err := service.Login(&authEntities.LoginCredentials{})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.AccessToken)
		assert.NotEmpty(t, result.RefreshToken)
		assert.False(t, result.IsApplicationAdmin)
		assert.Equal(t, "test", result.Username)
		assert.Equal(t, "test@test.com", result.Email)
	})

	t.Run("should return error when failed to get user groups", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		account := &accountEntities.Account{
			Username: "test",
			Email:    "test@test.com",
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByUsername").Return(account, nil)

		ldapMock := &client.Mock{}
		ldapMock.On("Authenticate").Return(true, map[string]string{}, nil)
		ldapMock.On("GetUserGroups").Return([]string{}, errors.New("test"))
		ldapMock.On("Close")

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		result, err := service.Login(&authEntities.LoginCredentials{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when failed to create account", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		account := &accountEntities.Account{
			Username: "test",
			Email:    "test@test.com",
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByUsername").Return(account, errors.New("test"))
		accountRepositoryMock.On("CreateAccount").Return(account, errors.New("test"))

		ldapMock := &client.Mock{}
		ldapMock.On("Authenticate").Return(true, map[string]string{}, nil)
		ldapMock.On("GetUserGroups").Return([]string{}, errors.New("test"))
		ldapMock.On("Close")

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		result, err := service.Login(&authEntities.LoginCredentials{})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error user does no exist when failed to authenticate", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		ldapMock := &client.Mock{}
		ldapMock.On("Authenticate").Return(
			true, map[string]string{}, ldapEnums.ErrorUserDoesNotExist)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		result, err := service.Login(&authEntities.LoginCredentials{})
		assert.Error(t, err)
		assert.Equal(t, ldapEnums.ErrorUserDoesNotExist, err)
		assert.Nil(t, result)
	})

	t.Run("should return error failed to authenticate", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		ldapMock := &client.Mock{}
		ldapMock.On("Authenticate").Return(
			true, map[string]string{}, ldapEnums.ErrorLdapUnauthorized)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		result, err := service.Login(&authEntities.LoginCredentials{})
		assert.Error(t, err)
		assert.Equal(t, ldapEnums.ErrorLdapUnauthorized, err)
		assert.Nil(t, result)
	})

	t.Run("should return error ldap unauthorized", func(t *testing.T) {
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		account := &accountEntities.Account{
			Username: "test",
			Email:    "test@test.com",
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByUsername").Return(account, errors.New("test"))
		accountRepositoryMock.On("CreateAccount").Return(account, nil)

		ldapMock := &client.Mock{}
		ldapMock.On("Authenticate").Return(false, map[string]string{}, nil)
		ldapMock.On("GetUserGroups").Return([]string{}, nil)
		ldapMock.On("Close")

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		result, err := service.Login(&authEntities.LoginCredentials{})
		assert.Error(t, err)
		assert.Equal(t, ldapEnums.ErrorLdapUnauthorized, err)
		assert.Nil(t, result)
	})
}

func TestIsAuthorizedApplicationAdmin(t *testing.T) {
	t.Run("should should return true and no error when application admin", func(t *testing.T) {
		_ = os.Setenv(ldapEnums.EnvLdapAdminGroup, "test")

		accountRepositoryMock := &accountRepository.Mock{}
		authRepositoryMock := &authRepository.Mock{}
		ldapMock := &client.Mock{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		appConfig := &app.Config{
			EnableApplicationAdmin: true,
		}

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should should return false and no error when application admin group not check", func(t *testing.T) {
		_ = os.Setenv(ldapEnums.EnvLdapAdminGroup, "test2")

		authRepositoryMock := &authRepository.Mock{}
		ldapMock := &client.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{
			EnableApplicationAdmin: true,
		}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("should should return false and error when application admin group not set", func(t *testing.T) {
		_ = os.Setenv(ldapEnums.EnvLdapAdminGroup, "")

		accountRepositoryMock := &accountRepository.Mock{}
		authRepositoryMock := &authRepository.Mock{}
		ldapMock := &client.Mock{}

		appConfig := &app.Config{
			EnableApplicationAdmin: true,
		}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.False(t, result)
		assert.Error(t, err)
		assert.Equal(t, err, ldapEnums.ErrorLdapApplicationAdminGroupNotSet)
	})

	t.Run("should should return error when failed to get account id from token", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		ldapMock := &client.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{
			EnableApplicationAdmin: false,
		}

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		data := &authEntities.AuthorizationData{
			Token:        "",
			Type:         authorization.ApplicationAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.False(t, result)
		assert.Error(t, err)
	})
}

func TestIsAuthorizedWorkspaceMember(t *testing.T) {
	t.Run("should should true and no error for workspace member", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzMember: []string{"test"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.WorkspaceMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should should false and no error when groups not check", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzMember: []string{"test2"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.WorkspaceMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("should should false and error when failed to get workspace groups", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, errors.New("test"))

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.WorkspaceMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should should error when failed to get groups from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		data := &authEntities.AuthorizationData{
			Token:        "",
			Type:         authorization.WorkspaceMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedWorkspaceAdmin(t *testing.T) {
	t.Run("should should true and no error for workspace admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzAdmin: []string{"test"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.WorkspaceAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should should false and no error when groups not check", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzAdmin: []string{"test2"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.WorkspaceAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("should should false and error when failed to get workspace groups", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, errors.New("test"))

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.WorkspaceAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should should error when failed to get groups from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		data := &authEntities.AuthorizationData{
			Token:        "",
			Type:         authorization.WorkspaceAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedRepositoryAdmin(t *testing.T) {
	t.Run("should should true and no error for repository admin", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzAdmin: []string{"test"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)
		authRepositoryMock.On("GetRepositoryGroups").Return(groups, nil)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositoryAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should should false and no error when groups not check", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzAdmin: []string{"test2"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)
		authRepositoryMock.On("GetRepositoryGroups").Return(groups, nil)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositoryAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("should should false and error when failed to get repository groups", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzAdmin: []string{"test"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)
		authRepositoryMock.On("GetRepositoryGroups").Return(groups, errors.New("test"))

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositoryAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should should false and error when failed to get workspace groups", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, errors.New("test"))

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositoryAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should should error when failed to get groups from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		data := &authEntities.AuthorizationData{
			Token:        "",
			Type:         authorization.RepositoryAdmin,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedRepositorySupervisor(t *testing.T) {
	t.Run("should should true and no error for repository supervisor", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzAdmin: []string{"test"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)
		authRepositoryMock.On("GetRepositoryGroups").Return(groups, nil)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositorySupervisor,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should should false and no error when groups not check", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzAdmin: []string{"test2"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)
		authRepositoryMock.On("GetRepositoryGroups").Return(groups, nil)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositorySupervisor,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("should should false and error when failed to get repository groups", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzAdmin: []string{"test"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)
		authRepositoryMock.On("GetRepositoryGroups").Return(groups, errors.New("test"))

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositorySupervisor,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should should false and error when failed to get workspace groups", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, errors.New("test"))

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositorySupervisor,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should should error when failed to get groups from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		data := &authEntities.AuthorizationData{
			Token:        "",
			Type:         authorization.RepositorySupervisor,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestIsAuthorizedRepositoryMember(t *testing.T) {
	t.Run("should should true and no error for repository member", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzAdmin: []string{"test"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)
		authRepositoryMock.On("GetRepositoryGroups").Return(groups, nil)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositoryMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("should should false and no error when groups not check", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzAdmin: []string{"test2"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)
		authRepositoryMock.On("GetRepositoryGroups").Return(groups, nil)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositoryMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("should should false and error when failed to get repository groups", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{AuthzAdmin: []string{"test"}}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, nil)
		authRepositoryMock.On("GetRepositoryGroups").Return(groups, errors.New("test"))

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositoryMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should should false and error when failed to get workspace groups", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}

		account := &accountEntities.Account{
			Username:           "test",
			Email:              "test@test.com",
			IsApplicationAdmin: true,
		}

		groups := &authEntities.AuthzGroups{}
		authRepositoryMock := &authRepository.Mock{}
		authRepositoryMock.On("GetWorkspaceGroups").Return(groups, errors.New("test"))

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		data := &authEntities.AuthorizationData{
			Token:        token,
			Type:         authorization.RepositoryMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("should should error when failed to get groups from token", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}
		appConfig := &app.Config{}
		authRepositoryMock := &authRepository.Mock{}

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		data := &authEntities.AuthorizationData{
			Token:        "",
			Type:         authorization.RepositoryMember,
			WorkspaceID:  uuid.New(),
			RepositoryID: uuid.New(),
		}

		result, err := service.IsAuthorized(data)
		assert.Error(t, err)
		assert.False(t, result)
	})
}

func TestGetHorusecAuthzGroups(t *testing.T) {
	t.Run("should error when invalid authorization type", func(t *testing.T) {
		service := &Service{}

		data := &authEntities.AuthorizationData{
			Type: "test",
		}

		groups, err := service.getHorusecAuthzGroups(data)

		assert.Nil(t, groups)
		assert.Error(t, err)
	})
}

func TestGetAccountDataFromToken(t *testing.T) {
	t.Run("should return account data without errors", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}
		ldapMock := &client.Mock{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, nil)

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		result, err := service.GetAccountDataFromToken(token)
		assert.NoError(t, err)
		assert.Equal(t, account.AccountID.String(), result.AccountID)
		assert.Equal(t, account.IsApplicationAdmin, result.IsApplicationAdmin)
		assert.Equal(t, []string{"test"}, result.Permissions)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}
		ldapMock := &client.Mock{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, errors.New("test"))

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		result, err := service.GetAccountDataFromToken(token)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}
		ldapMock := &client.Mock{}

		account := &accountEntities.Account{
			AccountID:          uuid.New(),
			IsConfirmed:        true,
			Email:              "test",
			Username:           "test",
			IsApplicationAdmin: true,
		}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(account, errors.New("test"))

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		token, _, _ := jwt.CreateToken(account.ToTokenData(), []string{"test"})

		result, err := service.GetAccountDataFromToken(token)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when failed to decode token", func(t *testing.T) {
		authRepositoryMock := &authRepository.Mock{}
		appConfig := &app.Config{}
		accountRepositoryMock := &accountRepository.Mock{}
		ldapMock := &client.Mock{}

		service := Service{
			cache:             cache.NewCache(),
			ldap:              ldapMock,
			accountRepository: accountRepositoryMock,
			authRepository:    authRepositoryMock,
			authUseCases:      authentication.NewAuthenticationUseCases(),
			appConfig:         appConfig,
		}

		result, err := service.GetAccountDataFromToken("")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

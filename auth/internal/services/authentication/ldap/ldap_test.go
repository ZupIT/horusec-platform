package ldap

import (
	"errors"
	"testing"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	authRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/authentication"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication/ldap/client"
	"github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
	ldapEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/ldap"
)

func TestNewLDAPAuthenticationService(t *testing.T) {
	t.Run("should success create a new service", func(t *testing.T) {
		assert.NotNil(t, NewLDAPAuthenticationService(
			nil, nil, nil, nil))
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
			cache:             cache.New(authEnums.TokenDuration, authEnums.TokenCheckExpiredDuration),
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
			cache:             cache.New(authEnums.TokenDuration, authEnums.TokenCheckExpiredDuration),
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
			cache:             cache.New(authEnums.TokenDuration, authEnums.TokenCheckExpiredDuration),
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
			cache:             cache.New(authEnums.TokenDuration, authEnums.TokenCheckExpiredDuration),
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
			cache:             cache.New(authEnums.TokenDuration, authEnums.TokenCheckExpiredDuration),
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
			cache:             cache.New(authEnums.TokenDuration, authEnums.TokenCheckExpiredDuration),
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

package authentication

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	"github.com/ZupIT/horusec-platform/auth/internal/services/authentication"
)

func TestNewAuthenticationController(t *testing.T) {
	t.Run("should success create a new controller", func(t *testing.T) {
		assert.NotNil(t, NewAuthenticationController(nil, nil,
			nil, nil, nil))
	})
}

func TestLogin(t *testing.T) {
	t.Run("should success login with horusec auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: auth.Horusec}

		authenticationMock := &authentication.Mock{}
		authenticationMock.On("Login").Return(&authEntities.LoginResponse{}, nil)

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.Login(&authEntities.LoginCredentials{})
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("should success login with keycloak auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: auth.Keycloak}

		authenticationMock := &authentication.Mock{}
		authenticationMock.On("Login").Return(&authEntities.LoginResponse{}, nil)

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.Login(&authEntities.LoginCredentials{})
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("should success login with ldap auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: auth.Ldap}

		authenticationMock := &authentication.Mock{}
		authenticationMock.On("Login").Return(&authEntities.LoginResponse{}, nil)

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.Login(&authEntities.LoginCredentials{})
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("should return error when invalid auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: "test"}
		authenticationMock := &authentication.Mock{}

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.Login(&authEntities.LoginCredentials{})
		assert.Error(t, err)
		assert.Equal(t, authEnums.ErrorAuthTypeInvalid, err)
		assert.Nil(t, response)
	})
}

func TestIsAuthorized(t *testing.T) {
	t.Run("should success return authorized with horusec auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: auth.Horusec}

		authenticationMock := &authentication.Mock{}
		authenticationMock.On("IsAuthorized").Return(true, nil)

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.IsAuthorized(&authEntities.AuthorizationData{})
		assert.NoError(t, err)
		assert.True(t, response)
	})

	t.Run("should success return authorized with keycloak auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: auth.Keycloak}

		authenticationMock := &authentication.Mock{}
		authenticationMock.On("IsAuthorized").Return(true, nil)

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.IsAuthorized(&authEntities.AuthorizationData{})
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("should success return authorized with ldap auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: auth.Ldap}

		authenticationMock := &authentication.Mock{}
		authenticationMock.On("IsAuthorized").Return(true, nil)

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.IsAuthorized(&authEntities.AuthorizationData{})
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("should return error when invalid auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: "test"}
		authenticationMock := &authentication.Mock{}

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.IsAuthorized(&authEntities.AuthorizationData{})
		assert.Error(t, err)
		assert.Equal(t, authEnums.ErrorAuthTypeInvalid, err)
		assert.False(t, response)
	})
}

func TestGetAccountInfo(t *testing.T) {
	t.Run("should success get account info with horusec auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: auth.Horusec}

		authenticationMock := &authentication.Mock{}
		authenticationMock.On("GetAccountDataFromToken").Return(&proto.GetAccountDataResponse{}, nil)

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.GetAccountInfo("")
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("should success get account info with keycloak auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: auth.Keycloak}

		authenticationMock := &authentication.Mock{}
		authenticationMock.On("GetAccountDataFromToken").Return(&proto.GetAccountDataResponse{}, nil)

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.GetAccountInfo("")
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("should success get account info with ldap auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: auth.Ldap}

		authenticationMock := &authentication.Mock{}
		authenticationMock.On("GetAccountDataFromToken").Return(&proto.GetAccountDataResponse{}, nil)

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.GetAccountInfo("")
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("should return error when invalid auth type", func(t *testing.T) {
		accountRepositoryMock := &accountRepository.Mock{}

		appConfig := &app.Config{AuthType: "test"}
		authenticationMock := &authentication.Mock{}

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.GetAccountInfo("")
		assert.Error(t, err)
		assert.Equal(t, authEnums.ErrorAuthTypeInvalid, err)
		assert.Nil(t, response)
	})
}

func TestGetAccountInfoByEmail(t *testing.T) {
	t.Run("should success get account info by email", func(t *testing.T) {
		authenticationMock := &authentication.Mock{}

		appConfig := &app.Config{AuthType: auth.Horusec}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(&accountEntities.Account{}, nil)

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		response, err := controller.GetAccountInfoByEmail("test@test.com")
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		authenticationMock := &authentication.Mock{}

		appConfig := &app.Config{AuthType: auth.Horusec}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(
			&accountEntities.Account{}, errors.New("test"))

		controller := NewAuthenticationController(appConfig, authenticationMock, authenticationMock,
			authenticationMock, accountRepositoryMock)

		_, err := controller.GetAccountInfoByEmail("test@test.com")
		assert.Error(t, err)
	})
}

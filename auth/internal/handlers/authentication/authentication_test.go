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
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	authController "github.com/ZupIT/horusec-platform/auth/internal/controllers/authentication"
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	horusecAuthEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/horusec"
	ldapEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication/ldap"
	authUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

func TestNewAuthenticationHandler(t *testing.T) {
	t.Run("should success create a new handler", func(t *testing.T) {
		assert.NotNil(t, NewAuthenticationHandler(nil, nil, nil))
	})
}

func TestGetConfig(t *testing.T) {
	t.Run("should return 200 hen success get config", func(t *testing.T) {
		appConfig := &app.Config{}
		controllerMock := &authController.Mock{}

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodGet, "test", nil)
		w := httptest.NewRecorder()

		handler.GetConfig(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestLogin(t *testing.T) {
	credentials := &authEntities.LoginCredentials{
		Username: "test",
		Password: "test",
	}

	t.Run("should return 200 when login with success", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Horusec}

		controllerMock := &authController.Mock{}
		controllerMock.On("Login").Return(&authEntities.LoginResponse{}, nil)

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(credentials.ToBytes()))
		w := httptest.NewRecorder()

		handler.Login(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when something went wrong auth type horusec", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Horusec}

		controllerMock := &authController.Mock{}
		controllerMock.On("Login").Return(&authEntities.LoginResponse{}, errors.New("test"))

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(credentials.ToBytes()))
		w := httptest.NewRecorder()

		handler.Login(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 403 when email not confirmed auth type horusec", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Horusec}

		controllerMock := &authController.Mock{}
		controllerMock.On("Login").Return(
			&authEntities.LoginResponse{}, horusecAuthEnums.ErrorAccountEmailNotConfirmed)

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(credentials.ToBytes()))
		w := httptest.NewRecorder()

		handler.Login(w, r)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("should return 403 when wrong email or password auth type horusec", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Horusec}

		controllerMock := &authController.Mock{}
		controllerMock.On("Login").Return(
			&authEntities.LoginResponse{}, horusecAuthEnums.ErrorWrongEmailOrPassword)

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(credentials.ToBytes()))
		w := httptest.NewRecorder()

		handler.Login(w, r)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("should return 403 when not found and auth type horusec", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Horusec}

		controllerMock := &authController.Mock{}
		controllerMock.On("Login").Return(
			&authEntities.LoginResponse{}, databaseEnums.ErrorNotFoundRecords)

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(credentials.ToBytes()))
		w := httptest.NewRecorder()

		handler.Login(w, r)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("should return 403 user does not exist auth type ldap", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Ldap}

		controllerMock := &authController.Mock{}
		controllerMock.On("Login").Return(
			&authEntities.LoginResponse{}, ldapEnums.ErrorUserDoesNotExist)

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(credentials.ToBytes()))
		w := httptest.NewRecorder()

		handler.Login(w, r)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("should return 403 when unauthorized error auth type ldap", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Ldap}

		controllerMock := &authController.Mock{}
		controllerMock.On("Login").Return(
			&authEntities.LoginResponse{}, ldapEnums.ErrorLdapUnauthorized)

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(credentials.ToBytes()))
		w := httptest.NewRecorder()

		handler.Login(w, r)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("should return 500 when something went wrong auth type ldap", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Ldap}

		controllerMock := &authController.Mock{}
		controllerMock.On("Login").Return(&authEntities.LoginResponse{}, errors.New("test"))

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(credentials.ToBytes()))
		w := httptest.NewRecorder()

		handler.Login(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 500 when something went wrong auth type keycloak", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Keycloak}

		controllerMock := &authController.Mock{}
		controllerMock.On("Login").Return(&authEntities.LoginResponse{}, errors.New("test"))

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(credentials.ToBytes()))
		w := httptest.NewRecorder()

		handler.Login(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 500 when something went wrong auth type unknown", func(t *testing.T) {
		appConfig := &app.Config{}

		controllerMock := &authController.Mock{}
		controllerMock.On("Login").Return(&authEntities.LoginResponse{}, errors.New("test"))

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader(credentials.ToBytes()))
		w := httptest.NewRecorder()

		handler.Login(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid request body", func(t *testing.T) {
		appConfig := &app.Config{}
		controllerMock := &authController.Mock{}

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodPost, "test", bytes.NewReader([]byte("")))
		w := httptest.NewRecorder()

		handler.Login(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestIsAuthorized(t *testing.T) {
	t.Run("should success verify authorization and no error", func(t *testing.T) {
		appConfig := &app.Config{}

		controllerMock := &authController.Mock{}
		controllerMock.On("IsAuthorized").Return(true, nil)

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		result, err := handler.IsAuthorized(context.Background(), &proto.IsAuthorizedData{})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetAccountInfo(t *testing.T) {
	t.Run("should success get account info with token", func(t *testing.T) {
		appConfig := &app.Config{}

		controllerMock := &authController.Mock{}
		controllerMock.On("GetAccountInfo").Return(&proto.GetAccountDataResponse{}, nil)

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		result, err := handler.GetAccountInfo(context.Background(), &proto.GetAccountData{Token: "test"})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success get account info with email", func(t *testing.T) {
		appConfig := &app.Config{}

		controllerMock := &authController.Mock{}
		controllerMock.On("GetAccountInfoByEmail").Return(&proto.GetAccountDataResponse{}, nil)

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		result, err := handler.GetAccountInfo(context.Background(), &proto.GetAccountData{})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetAuthConfig(t *testing.T) {
	t.Run("should success return auth config", func(t *testing.T) {
		appConfig := &app.Config{}
		controllerMock := &authController.Mock{}

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		result, err := handler.GetAuthConfig(context.Background(), &proto.GetAuthConfigData{})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestOptions(t *testing.T) {
	t.Run("should return 204 when options", func(t *testing.T) {
		appConfig := &app.Config{}
		controllerMock := &authController.Mock{}

		handler := NewAuthenticationHandler(appConfig, authUseCases.NewAuthenticationUseCases(), controllerMock)

		r, _ := http.NewRequest(http.MethodOptions, "test", nil)
		w := httptest.NewRecorder()

		handler.Options(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

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

package app

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	"github.com/ZupIT/horusec-platform/auth/config/app/enums"
)

func getMockedConnection() *database.Connection {
	databaseMock := &database.Mock{}
	databaseMock.On("Create").Return(&response.Response{})

	return &database.Connection{Read: databaseMock, Write: databaseMock}
}

func setDefaultEnvs() {
	_ = os.Setenv(enums.EnvAuthURL, "http://localhost:8006")
	_ = os.Setenv(enums.EnvAuthType, auth.Horusec.ToString())
	_ = os.Setenv(enums.EnvDisableEmails, "false")
	_ = os.Setenv(enums.EnvEnableApplicationAdmin, "false")
	_ = os.Setenv(enums.EnvApplicationAdminData, enums.ApplicationAdminDefaultData)
	_ = os.Setenv(enums.EnvDefaultUserData, enums.DefaultUserData)
	_ = os.Setenv(enums.EnvHorusecManager, "http://localhost:8043")
}

func TestNewAuthAppConfig(t *testing.T) {
	t.Run("should success create a new config without default users", func(t *testing.T) {
		_ = os.Setenv(enums.EnvEnableDefaultUser, "false")
		_ = os.Setenv(enums.EnvEnableApplicationAdmin, "false")

		assert.NotNil(t, NewAuthAppConfig(&database.Connection{Read: &database.Mock{}, Write: &database.Mock{}}))
	})

	t.Run("should success create a new config with default users", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("Find").Return(
			response.NewResponse(0, databaseEnums.ErrorNotFoundRecords, nil))
		databaseMock.On("Delete").Return(&response.Response{})
		databaseMock.On("CommitTransaction").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)

		_ = os.Setenv(enums.EnvEnableDefaultUser, "true")
		_ = os.Setenv(enums.EnvEnableApplicationAdmin, "true")

		assert.NotNil(t, NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock}))
	})

	t.Run("should panic when failed to create app admin", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(
			response.NewResponse(0, errors.New("test"), nil))
		databaseMock.On("Find").Return(
			response.NewResponse(0, databaseEnums.ErrorNotFoundRecords, nil))
		databaseMock.On("Delete").Return(&response.Response{})
		databaseMock.On("CommitTransaction").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)

		_ = os.Setenv(enums.EnvEnableDefaultUser, "false")
		_ = os.Setenv(enums.EnvEnableApplicationAdmin, "true")

		assert.Panics(t, func() {
			NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock})
		})
	})

	t.Run("should panic when failed to delete app admin", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Return(
			response.NewResponse(0, errors.New("test"), nil))
		databaseMock.On("Find").Return(
			response.NewResponse(0, databaseEnums.ErrorNotFoundRecords, nil))
		databaseMock.On("CommitTransaction").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)

		_ = os.Setenv(enums.EnvEnableDefaultUser, "false")
		_ = os.Setenv(enums.EnvEnableApplicationAdmin, "true")

		assert.Panics(t, func() {
			NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock})
		})
	})

	t.Run("should return errors when failed to update app admin", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(
			response.NewResponse(0, errors.New("test"), nil))
		databaseMock.On("Find").Return(&response.Response{})
		databaseMock.On("Delete").Return(&response.Response{})
		databaseMock.On("RollbackTransaction").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)

		_ = os.Setenv(enums.EnvEnableDefaultUser, "false")
		_ = os.Setenv(enums.EnvEnableApplicationAdmin, "true")

		assert.Panics(t, func() {
			NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock})
		})
	})

	t.Run("should return errors when failed to delete app admin", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("Find").Return(&response.Response{})
		databaseMock.On("Delete").Return(
			response.NewResponse(0, errors.New("test"), nil))
		databaseMock.On("RollbackTransaction").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)

		_ = os.Setenv(enums.EnvEnableDefaultUser, "false")
		_ = os.Setenv(enums.EnvEnableApplicationAdmin, "true")

		assert.Panics(t, func() {
			NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock})
		})
	})

	t.Run("should success create a new config with existing users", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Create").Once().Return(
			response.NewResponse(0, errors.New(enums.DuplicatedAccount), nil))
		databaseMock.On("Find").Return(&response.Response{})
		databaseMock.On("Delete").Return(&response.Response{})
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("CommitTransaction").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)

		_ = os.Setenv(enums.EnvEnableApplicationAdmin, "true")
		_ = os.Setenv(enums.EnvEnableDefaultUser, "true")

		assert.NotNil(t, NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock}))
	})

	t.Run("should panic when failed to create account", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(
			response.NewResponse(0, errors.New("test"), nil))

		_ = os.Setenv(enums.EnvEnableApplicationAdmin, "true")
		_ = os.Setenv(enums.EnvEnableDefaultUser, "true")

		assert.Panics(t, func() {
			NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock})
		})
	})

	t.Run("should panic when failed to get default user data", func(t *testing.T) {
		databaseMock := &database.Mock{}

		_ = os.Setenv(enums.EnvDefaultUserData, "test")
		_ = os.Setenv(enums.EnvEnableDefaultUser, "true")

		assert.Panics(t, func() {
			NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock})
		})
	})

	t.Run("should panic when failed to get app admin user data", func(t *testing.T) {
		databaseMock := &database.Mock{}

		_ = os.Setenv(enums.EnvApplicationAdminData, "test")
		_ = os.Setenv(enums.EnvEnableApplicationAdmin, "true")
		_ = os.Setenv(enums.EnvEnableDefaultUser, "false")

		assert.Panics(t, func() {
			NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock})
		})
	})

	t.Run("should log warning message when another auth type", func(t *testing.T) {
		databaseMock := &database.Mock{}

		_ = os.Setenv(enums.EnvEnableApplicationAdmin, "false")
		_ = os.Setenv(enums.EnvEnableDefaultUser, "true")
		_ = os.Setenv(enums.EnvAuthType, "ldap")

		assert.NotNil(t, NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock}))
	})

	t.Run("should panic when failed to get admin account by email", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").Return(
			response.NewResponse(0, errors.New("test"), nil))

		_ = os.Setenv(enums.EnvApplicationAdminData, enums.ApplicationAdminDefaultData)
		_ = os.Setenv(enums.EnvEnableApplicationAdmin, "true")
		_ = os.Setenv(enums.EnvEnableDefaultUser, "false")

		assert.Panics(t, func() {
			NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock})
		})
	})

	setDefaultEnvs()
}

func TestGetAuthType(t *testing.T) {
	t.Run("should success get auth type", func(t *testing.T) {
		appConfig := NewAuthAppConfig(getMockedConnection())

		assert.Equal(t, auth.Horusec, appConfig.GetAuthenticationType())
	})
}

func TestToConfigResponse(t *testing.T) {
	t.Run("should success parse config to response", func(t *testing.T) {
		appConfig := NewAuthAppConfig(getMockedConnection())

		result := appConfig.ToConfigResponse()
		assert.NotPanics(t, func() {
			assert.Equal(t, false, result["enableApplicationAdmin"])
			assert.Equal(t, auth.Horusec, result["authType"])
			assert.Equal(t, false, result["disableEmails"])
		})
	})
}

func TestIsApplicationAdminEnabled(t *testing.T) {
	t.Run("should return false when not active", func(t *testing.T) {
		appConfig := NewAuthAppConfig(getMockedConnection())

		assert.False(t, appConfig.IsApplicationAdmEnabled())
	})
}

func TestIsDisableEmails(t *testing.T) {
	t.Run("should return false when not active", func(t *testing.T) {
		appConfig := NewAuthAppConfig(getMockedConnection())

		assert.False(t, appConfig.IsEmailsDisabled())
	})
}

func TestToGetAuthConfigResponse(t *testing.T) {
	t.Run("should return false when not active", func(t *testing.T) {
		appConfig := NewAuthAppConfig(getMockedConnection())

		result := appConfig.ToGetAuthConfigResponse()
		assert.Equal(t, false, result.EnableApplicationAdmin)
		assert.Equal(t, false, result.DisableEmails)
		assert.Equal(t, auth.Horusec.ToString(), result.AuthType)
	})
}

func TestGetHorusecAuthURL(t *testing.T) {
	t.Run("should success get auth url", func(t *testing.T) {
		appConfig := NewAuthAppConfig(getMockedConnection())

		assert.Equal(t, "http://localhost:8006", appConfig.GetHorusecAuthURL())
	})
}

func TestGetHorusecManagerURL(t *testing.T) {
	t.Run("should success get manager url", func(t *testing.T) {
		appConfig := NewAuthAppConfig(getMockedConnection())

		assert.Equal(t, "http://localhost:8043", appConfig.GetHorusecManagerURL())
	})
}

func TestGetEnableApplicationAdmin(t *testing.T) {
	t.Run("should success get if app admin is enabled", func(t *testing.T) {
		appConfig := NewAuthAppConfig(getMockedConnection())

		assert.False(t, appConfig.GetEnableApplicationAdmin())
	})
}

func TestGetEnableDefaultUser(t *testing.T) {
	t.Run("should success get if default user is enabled", func(t *testing.T) {
		_ = os.Setenv(enums.EnvEnableDefaultUser, "true")

		appConfig := NewAuthAppConfig(getMockedConnection())

		assert.True(t, appConfig.GetEnableDefaultUser())
	})
}

func TestGetDefaultUserData(t *testing.T) {
	t.Run("should success get if default user is enabled", func(t *testing.T) {
		appConfig := NewAuthAppConfig(getMockedConnection())

		account, err := appConfig.GetDefaultUserData()
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestGetApplicationAdminData(t *testing.T) {
	t.Run("should success get application admin data", func(t *testing.T) {
		appConfig := NewAuthAppConfig(getMockedConnection())

		account, err := appConfig.GetApplicationAdminData()
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestCheckForNilError(t *testing.T) {
	t.Run("should return error message when not empty ", func(t *testing.T) {
		appConfig := &Config{}

		assert.Equal(t, "test", appConfig.checkForNilError(errors.New("test")))
	})

	t.Run("should return empty message when nil error ", func(t *testing.T) {
		appConfig := &Config{}

		assert.Empty(t, appConfig.checkForNilError(nil))
	})
}

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

package account

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/Nerzal/gocloak/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	emailEnums "github.com/ZupIT/horusec-devkit/pkg/enums/email"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	accountEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/account"
)

func getAppConfig() app.IConfig {
	databaseMock := &database.Mock{}
	databaseMock.On("Create").Return(&response.Response{})

	return app.NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock})
}

func TestNewAccountUseCases(t *testing.T) {
	t.Run("should success create a new use cases", func(t *testing.T) {
		assert.NotNil(t, NewAccountUseCases(getAppConfig()))
	})
}

func TestFilterAccountByID(t *testing.T) {
	t.Run("should success create a filter by account id", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		id := uuid.New()

		filter := useCases.FilterAccountByID(id)
		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["account_id"])
		})
	})
}

func TestFilterAccountByEmail(t *testing.T) {
	t.Run("should success create a filter by account id", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		filter := useCases.FilterAccountByEmail("test@test.com")
		assert.NotPanics(t, func() {
			assert.Equal(t, "test@test.com", filter["email"])
		})
	})
}

func TestFilterAccountByUsername(t *testing.T) {
	t.Run("should success create a filter by account id", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		filter := useCases.FilterAccountByUsername("test")
		assert.NotPanics(t, func() {
			assert.Equal(t, "test", filter["username"])
		})
	})
}

func TestAccessTokenFromIOReadCloser(t *testing.T) {
	t.Run("should success get data from request body", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		data := map[string]string{"accessToken": "test"}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		result, err := useCases.AccessTokenFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "test", result.AccessToken)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		result, err := useCases.AccessTokenFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestNewAccountFromKeycloakUserInfo(t *testing.T) {
	t.Run("should success create account from keycloak info with preferred username", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		test := "test"
		id := uuid.NewString()
		userInfo := &gocloak.UserInfo{
			PreferredUsername: &test,
			Email:             &test,
			Sub:               &id,
		}

		account := useCases.NewAccountFromKeycloakUserInfo(userInfo)
		assert.Equal(t, parser.ParseStringToUUID(id), account.AccountID)
		assert.Equal(t, "test", account.Email)
		assert.Equal(t, "test", account.Username)
		assert.Equal(t, true, account.IsConfirmed)
		assert.NotEqual(t, time.Time{}, account.CreatedAt)
		assert.NotEqual(t, time.Time{}, account.UpdatedAt)
	})

	t.Run("should success create account from keycloak info with name", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		test := "test"
		testEmpty := ""
		id := uuid.NewString()
		userInfo := &gocloak.UserInfo{
			PreferredUsername: &testEmpty,
			Name:              &test,
			Email:             &test,
			Sub:               &id,
		}

		account := useCases.NewAccountFromKeycloakUserInfo(userInfo)
		assert.Equal(t, parser.ParseStringToUUID(id), account.AccountID)
		assert.Equal(t, "test", account.Email)
		assert.Equal(t, "test", account.Username)
		assert.Equal(t, true, account.IsConfirmed)
		assert.NotEqual(t, time.Time{}, account.CreatedAt)
		assert.NotEqual(t, time.Time{}, account.UpdatedAt)
	})
}

func TestCheckCreateAccountErrors(t *testing.T) {
	t.Run("should return same error when not specified", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		err := useCases.CheckCreateAccountErrors(errors.New("test"))
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
	})

	t.Run("should return same error username already in use", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		err := useCases.CheckCreateAccountErrors(errors.New(accountEnums.DuplicatedConstraintPrimaryKey))
		assert.Error(t, err)
		assert.Equal(t, accountEnums.ErrorUsernameAlreadyInUse, err)
	})

	t.Run("should return same error username already in use", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		err := useCases.CheckCreateAccountErrors(errors.New(accountEnums.DuplicatedConstraintUsername))
		assert.Error(t, err)
		assert.Equal(t, accountEnums.ErrorUsernameAlreadyInUse, err)
	})

	t.Run("should return same error email already in use", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		err := useCases.CheckCreateAccountErrors(errors.New(accountEnums.DuplicatedConstraintEmail))
		assert.Error(t, err)
		assert.Equal(t, accountEnums.ErrorEmailAlreadyInUse, err)
	})
}

func TestAccountDataFromIOReadCloser(t *testing.T) {
	t.Run("should success get data from request body", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		data := &accountEntities.Data{
			Email:    "test@test.com",
			Password: "test@123S",
			Username: "test",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		result, err := useCases.AccountDataFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, data, result)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		result, err := useCases.AccountDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestNewAccountValidationEmail(t *testing.T) {
	t.Run("should success create a new validation email", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		account := &accountEntities.Account{
			Email:    "test@test.com",
			Username: "test",
		}

		emailBytes := useCases.NewAccountValidationEmail(account)
		assert.NotNil(t, emailBytes)
		assert.NotEmpty(t, emailBytes)

		email := &emailEntities.Message{}
		assert.NoError(t, json.Unmarshal(emailBytes, email))

		assert.Equal(t, "test@test.com", email.To)
		assert.Equal(t, emailEnums.AccountConfirmation, email.TemplateName)
		assert.Equal(t, "[Horusec] Account Confirmation Email", email.Subject)

		assert.NotPanics(t, func() {
			data := email.Data.(map[string]interface{})

			assert.Equal(t, "http://localhost:8006/auth/account/validate/"+
				"00000000-0000-0000-0000-000000000000", data["URL"])
			assert.Equal(t, "test", data["Username"])

		})
	})
}

func TestEmailFromIOReadCloser(t *testing.T) {
	t.Run("should success get data from request body", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		data := &accountEntities.Email{
			Email: "test@test.com",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		result, err := useCases.EmailFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, data, result)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		result, err := useCases.EmailFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestGenerateResetPasswordCode(t *testing.T) {
	t.Run("should generate a reset password code", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		result := useCases.GenerateResetPasswordCode()
		assert.NotEmpty(t, result)
		assert.Len(t, result, 6)
	})
}

func TestNewResetPasswordCodeEmail(t *testing.T) {
	t.Run("should success create a new password code email", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		account := &accountEntities.Account{
			Email:    "test@test.com",
			Username: "test",
		}

		emailBytes := useCases.NewResetPasswordCodeEmail(account, "123456")
		assert.NotNil(t, emailBytes)
		assert.NotEmpty(t, emailBytes)

		email := &emailEntities.Message{}
		assert.NoError(t, json.Unmarshal(emailBytes, email))

		assert.Equal(t, "test@test.com", email.To)
		assert.Equal(t, emailEnums.ResetPassword, email.TemplateName)
		assert.Equal(t, "[Horusec] Reset Password", email.Subject)

		assert.NotPanics(t, func() {
			data := email.Data.(map[string]interface{})

			assert.Equal(t, "http://localhost:8043/auth/recovery-password/"+
				"check-code?email=test@test.com&code=123456", data["URL"])
			assert.Equal(t, "test", data["Username"])
			assert.Equal(t, "123456", data["Code"])

		})
	})
}

func TestResetCodeDataFromIOReadCloser(t *testing.T) {
	t.Run("should success get data from request body", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		data := &accountEntities.ResetCodeData{
			Email: "test@test.com",
			Code:  "123456",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		result, err := useCases.ResetCodeDataFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, data, result)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		result, err := useCases.ResetCodeDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestChangePasswordDataFromIOReadCloser(t *testing.T) {
	t.Run("should success get data from request body", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		data := &accountEntities.ChangePasswordData{
			Password: "Test@123",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		result, err := useCases.ChangePasswordDataFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, data, result)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		result, err := useCases.ChangePasswordDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestRefreshTokenFromIOReadCloser(t *testing.T) {
	t.Run("should success get data from request body", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		data := &accountEntities.RefreshToken{
			RefreshToken: "test",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		result, err := useCases.RefreshTokenFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, data, result)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		result, err := useCases.RefreshTokenFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestCheckEmailAndUsernameFromIOReadCloser(t *testing.T) {
	t.Run("should success get data from request body", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		data := &accountEntities.CheckEmailAndUsername{
			Email:    "test@test.com",
			Username: "test",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		result, err := useCases.CheckEmailAndUsernameFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, data, result)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		result, err := useCases.CheckEmailAndUsernameFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestUpdateAccountFromIOReadCloser(t *testing.T) {
	t.Run("should success get data from request body", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		data := &accountEntities.UpdateAccount{
			Email:    "test@test.com",
			Username: "test",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		result, err := useCases.UpdateAccountFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, data, result)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewAccountUseCases(getAppConfig())

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		result, err := useCases.UpdateAccountFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

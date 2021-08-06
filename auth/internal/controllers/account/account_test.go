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
	"errors"
	"testing"

	"github.com/Nerzal/gocloak/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/cache"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/crypto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	accountEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/account"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	accountRepository "github.com/ZupIT/horusec-platform/auth/internal/repositories/account"
	authServices "github.com/ZupIT/horusec-platform/auth/internal/services/authentication"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
)

func getAppConfig() app.IConfig {
	databaseMock := &database.Mock{}
	databaseMock.On("Create").Return(&response.Response{})

	return app.NewAuthAppConfig(&database.Connection{Read: databaseMock, Write: databaseMock}, nil)
}

func TestNewAccountController(t *testing.T) {
	t.Run("should success create a new controller", func(t *testing.T) {
		assert.NotNil(t, NewAccountController(nil, nil, nil,
			nil, nil, nil))
	})
}

func TestCreateAccountKeycloak(t *testing.T) {
	t.Run("should success create a new account", func(t *testing.T) {
		appConfig := getAppConfig()
		brokerMock := &broker.Mock{}

		test := ""
		userInfo := &gocloak.UserInfo{
			Sub:               &test,
			Name:              &test,
			PreferredUsername: &test,
			Email:             &test,
		}

		serviceMock := &authServices.Mock{}
		serviceMock.On("GetUserInfo").Return(userInfo, nil)

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("CreateAccount").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cache.NewCache())

		result, err := controller.CreateAccountKeycloak("test")
		assert.NotNil(t, result)
		assert.NoError(t, err)
	})

	t.Run("should return already existing account", func(t *testing.T) {
		appConfig := getAppConfig()
		brokerMock := &broker.Mock{}

		test := ""
		userInfo := &gocloak.UserInfo{
			Sub:               &test,
			Name:              &test,
			PreferredUsername: &test,
			Email:             &test,
		}

		serviceMock := &authServices.Mock{}
		serviceMock.On("GetUserInfo").Return(userInfo, nil)

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("CreateAccount").Return(&accountEntities.Account{},
			errors.New(accountEnums.DuplicatedConstraintPrimaryKey))

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cache.NewCache())

		_, err := controller.CreateAccountKeycloak("test")
		assert.Error(t, err)
	})

	t.Run("should return error when failed to get user info ", func(t *testing.T) {
		appConfig := getAppConfig()
		brokerMock := &broker.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		userInfo := &gocloak.UserInfo{}

		serviceMock := &authServices.Mock{}
		serviceMock.On("GetUserInfo").Return(userInfo, errors.New("test"))

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cache.NewCache())

		_, err := controller.CreateAccountKeycloak("test")
		assert.Error(t, err)
	})
}

func TestCreateAccountHorusec(t *testing.T) {
	t.Run("should success create a new account with email enabled", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("CreateAccount").Return(&accountEntities.Account{}, nil)

		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cache.NewCache())

		data := &accountEntities.Data{}

		result, err := controller.CreateAccountHorusec(data)
		assert.NotNil(t, result)
		assert.NoError(t, err)
	})

	t.Run("should success create a new account with email disabled", func(t *testing.T) {
		appConfig := &app.Config{DisableEmails: true}
		serviceMock := &authServices.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("CreateAccount").Return(&accountEntities.Account{}, nil)

		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cache.NewCache())

		data := &accountEntities.Data{}

		result, err := controller.CreateAccountHorusec(data)
		assert.NotNil(t, result)
		assert.NoError(t, err)
	})

	t.Run("should return error when failed to create account", func(t *testing.T) {
		appConfig := &app.Config{DisableEmails: true}
		serviceMock := &authServices.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("CreateAccount").Return(&accountEntities.Account{}, errors.New("test"))

		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cache.NewCache())

		data := &accountEntities.Data{}

		result, err := controller.CreateAccountHorusec(data)
		assert.Nil(t, result)
		assert.Error(t, err)
	})
}

func TestValidateAccountEmail(t *testing.T) {
	t.Run("should success validate account", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{}, nil)
		accountRepositoryMock.On("Update").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cache.NewCache())

		assert.NoError(t, controller.ValidateAccountEmail(uuid.New()))
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{}, errors.New("test"))

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cache.NewCache())

		assert.Error(t, controller.ValidateAccountEmail(uuid.New()))
	})
}

func TestSendResetPasswordCode(t *testing.T) {
	t.Run("should success send reset password code", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}

		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(nil)

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cache.NewCache())

		assert.NoError(t, controller.SendResetPasswordCode("test"))
	})

	t.Run("should not send email when it is disabled", func(t *testing.T) {
		appConfig := &app.Config{DisableEmails: true}
		serviceMock := &authServices.Mock{}

		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(nil)

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cache.NewCache())

		assert.NoError(t, controller.SendResetPasswordCode("test"))
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(
			&accountEntities.Account{}, errors.New("test"))

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cache.NewCache())

		assert.Error(t, controller.SendResetPasswordCode("test"))
	})
}

func TestCheckResetPasswordCode(t *testing.T) {
	t.Run("should success verify code without errors", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}

		cacheMock := &cache.Mock{}
		cacheMock.On("GetString").Return("123456", nil)
		cacheMock.On("Delete")

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(&accountEntities.Account{}, nil)
		accountRepositoryMock.On("Update").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.ResetCodeData{
			Email: "test@test.com",
			Code:  "123456",
		}

		result, err := controller.CheckResetPasswordCode(data)
		assert.NotEmpty(t, result)
		assert.NoError(t, err)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}

		cacheMock := &cache.Mock{}
		cacheMock.On("GetString").Return("123456", nil)

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(
			&accountEntities.Account{}, errors.New("test"))
		accountRepositoryMock.On("Update").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.ResetCodeData{
			Email: "test@test.com",
			Code:  "123456",
		}

		result, err := controller.CheckResetPasswordCode(data)
		assert.Empty(t, result)
		assert.Error(t, err)
	})

	t.Run("should return error when wrong code", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}

		cacheMock := &cache.Mock{}
		cacheMock.On("GetString").Return("123456", nil)

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(
			&accountEntities.Account{}, errors.New("test"))
		accountRepositoryMock.On("Update").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.ResetCodeData{
			Email: "test@test.com",
			Code:  "654321",
		}

		result, err := controller.CheckResetPasswordCode(data)
		assert.Empty(t, result)
		assert.Error(t, err)
		assert.Equal(t, accountEnums.ErrorIncorrectRetrievePasswordCode, err)
	})

	t.Run("should return error when failed to get stored code", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		cacheMock := &cache.Mock{}
		cacheMock.On("GetString").Return("123456", errors.New("test"))

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.ResetCodeData{
			Email: "test@test.com",
			Code:  "654321",
		}

		result, err := controller.CheckResetPasswordCode(data)
		assert.Empty(t, result)
		assert.Error(t, err)
	})
}

func TestChangePassword(t *testing.T) {
	t.Run("should success update account password", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{}, nil)
		accountRepositoryMock.On("Update").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.ChangePasswordData{
			Password: "test",
		}

		assert.NoError(t, controller.ChangePassword(data))
	})

	t.Run("should return error when same password", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}

		password, _ := crypto.HashPasswordBcrypt("test")

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{Password: password}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.ChangePasswordData{
			Password: "test",
		}

		err := controller.ChangePassword(data)
		assert.Error(t, err)
		assert.Equal(t, accountEnums.ErrorPasswordEqualPrevious, err)
	})

	t.Run("should return error failed to get account", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{}, errors.New("test"))

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.ChangePasswordData{
			Password: "test",
		}

		assert.Error(t, controller.ChangePassword(data))
	})
}

func TestRefreshToken(t *testing.T) {
	t.Run("should success refresh token", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}

		cacheMock := &cache.Mock{}
		cacheMock.On("GetString").Return("test", nil)
		cacheMock.On("Delete")
		cacheMock.On("Set")

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		result, err := controller.RefreshToken("test")
		assert.NotNil(t, result)
		assert.NoError(t, err)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}

		cacheMock := &cache.Mock{}
		cacheMock.On("GetString").Return("test", nil)

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{}, errors.New("test"))

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		result, err := controller.RefreshToken("test")
		assert.Nil(t, result)
		assert.Error(t, err)
	})

	t.Run("should return error when failed to get refresh token", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		cacheMock := &cache.Mock{}
		cacheMock.On("GetString").Return("", errors.New("test"))

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		result, err := controller.RefreshToken("test")
		assert.Nil(t, result)
		assert.Error(t, err)
	})
}

func TestLogout(t *testing.T) {
	t.Run("should success logout user", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		cacheMock := &cache.Mock{}
		cacheMock.On("Delete")

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		assert.NotPanics(t, func() {
			controller.Logout("")
		})
	})
}

func TestCheckExistingEmailOrUsername(t *testing.T) {
	t.Run("should return no error when not in use", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(&accountEntities.Account{}, nil)
		accountRepositoryMock.On("GetAccountByUsername").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.CheckEmailAndUsername{}

		assert.NoError(t, controller.CheckExistingEmailOrUsername(data))
	})

	t.Run("should return error username already in use", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(&accountEntities.Account{}, nil)
		accountRepositoryMock.On("GetAccountByUsername").Return(
			&accountEntities.Account{Username: "test"}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.CheckEmailAndUsername{}

		err := controller.CheckExistingEmailOrUsername(data)
		assert.Error(t, err)
		assert.Equal(t, accountEnums.ErrorUsernameAlreadyInUse, err)
	})

	t.Run("should return error email already in use", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccountByEmail").Return(
			&accountEntities.Account{Email: "test"}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.CheckEmailAndUsername{}

		err := controller.CheckExistingEmailOrUsername(data)
		assert.Error(t, err)
		assert.Equal(t, accountEnums.ErrorEmailAlreadyInUse, err)
	})
}

func TestDeleteAccount(t *testing.T) {
	t.Run("should success delete account", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("Delete").Return(nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		assert.NoError(t, controller.DeleteAccount(uuid.New()))
	})
}

func TestGetAccountID(t *testing.T) {
	t.Run("should success get account id horusec", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Horusec}
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		account := &accountEntities.Account{AccountID: uuid.New()}
		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		result, err := controller.GetAccountID(token)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success get account id ldap", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Ldap}
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		account := &accountEntities.Account{AccountID: uuid.New()}
		token, _, _ := jwt.CreateToken(account.ToTokenData(), nil)

		result, err := controller.GetAccountID(token)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success get account id keycloak", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Keycloak}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		serviceMock := &authServices.Mock{}
		serviceMock.On("GetAccountDataFromToken").Return(&proto.GetAccountDataResponse{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		result, err := controller.GetAccountID("")
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get account data keycloak", func(t *testing.T) {
		appConfig := &app.Config{AuthType: auth.Keycloak}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}

		serviceMock := &authServices.Mock{}
		serviceMock.On("GetAccountDataFromToken").Return(
			&proto.GetAccountDataResponse{}, errors.New("test"))

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		result, err := controller.GetAccountID("")
		assert.Error(t, err)
		assert.Equal(t, uuid.Nil, result)
	})

	t.Run("should return error when invalid auth type", func(t *testing.T) {
		appConfig := &app.Config{AuthType: "test"}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}
		accountRepositoryMock := &accountRepository.Mock{}
		serviceMock := &authServices.Mock{}

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		result, err := controller.GetAccountID("")
		assert.Error(t, err)
		assert.Equal(t, uuid.Nil, result)
		assert.Equal(t, authEnums.ErrorAuthTypeInvalid, err)
	})
}

func TestUpdateAccount(t *testing.T) {
	t.Run("should success update account without email change", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		brokerMock := &broker.Mock{}
		cacheMock := &cache.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{}, nil)
		accountRepositoryMock.On("Update").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.UpdateAccount{}

		result, err := controller.UpdateAccount(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success update account with email change", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		cacheMock := &cache.Mock{}

		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(nil)

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{}, nil)
		accountRepositoryMock.On("Update").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.UpdateAccount{Email: "test"}

		result, err := controller.UpdateAccount(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to send email", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		cacheMock := &cache.Mock{}

		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(errors.New("test"))

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{}, nil)
		accountRepositoryMock.On("Update").Return(&accountEntities.Account{}, nil)

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.UpdateAccount{Email: "test"}

		result, err := controller.UpdateAccount(data)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when failed to update", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		cacheMock := &cache.Mock{}

		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(nil)

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{}, nil)
		accountRepositoryMock.On("Update").Return(&accountEntities.Account{}, errors.New("test"))

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.UpdateAccount{Email: "test"}

		result, err := controller.UpdateAccount(data)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when failed to get account", func(t *testing.T) {
		appConfig := getAppConfig()
		serviceMock := &authServices.Mock{}
		cacheMock := &cache.Mock{}
		brokerMock := &broker.Mock{}

		accountRepositoryMock := &accountRepository.Mock{}
		accountRepositoryMock.On("GetAccount").Return(&accountEntities.Account{}, errors.New("test"))

		controller := NewAccountController(accountRepositoryMock, serviceMock,
			accountUseCases.NewAccountUseCases(appConfig), appConfig, brokerMock, cacheMock)

		data := &accountEntities.UpdateAccount{Email: "test"}

		result, err := controller.UpdateAccount(data)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

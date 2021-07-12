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
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	"github.com/ZupIT/horusec-platform/auth/config/app"
	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
)

func TestNewAccountRepository(t *testing.T) {
	t.Run("should create account repository", func(t *testing.T) {
		assert.NotNil(t, NewAccountRepository(&database.Connection{}, nil))
	})
}

func TestGetAccount(t *testing.T) {
	t.Run("should success get a account by id", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").Return(&response.Response{})

		repository := NewAccountRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			accountUseCases.NewAccountUseCases(&app.Config{}))

		account, err := repository.GetAccount(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestGetAccountByEmail(t *testing.T) {
	t.Run("should success get a account by email", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").Return(&response.Response{})

		repository := NewAccountRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			accountUseCases.NewAccountUseCases(&app.Config{}))

		account, err := repository.GetAccountByEmail("test@test.com")
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestGetAccountByUsername(t *testing.T) {
	t.Run("should success get a account by username", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").Return(&response.Response{})

		repository := NewAccountRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			accountUseCases.NewAccountUseCases(&app.Config{}))

		account, err := repository.GetAccountByUsername("test")
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestCreateAccount(t *testing.T) {
	t.Run("should success create account without errors", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})

		repository := NewAccountRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			accountUseCases.NewAccountUseCases(&app.Config{}))

		account, err := repository.CreateAccount(&accountEntities.Account{})
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should success update account", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		repository := NewAccountRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			accountUseCases.NewAccountUseCases(&app.Config{}))

		account, err := repository.Update(&accountEntities.Account{})
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should success delete account", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Return(&response.Response{})

		repository := NewAccountRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			accountUseCases.NewAccountUseCases(&app.Config{}))

		assert.NoError(t, repository.Delete(uuid.New()))
	})
}

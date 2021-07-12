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
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	authUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

func TestNewAuthenticationRepository(t *testing.T) {
	t.Run("should success create new repository", func(t *testing.T) {
		assert.NotNil(t, NewAuthenticationRepository(&database.Connection{}, authUseCases.NewAuthenticationUseCases()))
	})
}

func TestGetWorkspaceGroups(t *testing.T) {
	t.Run("should success get workspace groups", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").Return(&response.Response{})

		repository := NewAuthenticationRepository(&database.Connection{
			Read: databaseMock, Write: databaseMock}, authUseCases.NewAuthenticationUseCases())

		account, err := repository.GetWorkspaceGroups(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestGetRepositoryGroups(t *testing.T) {
	t.Run("should success get repository groups", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").Return(&response.Response{})

		repository := NewAuthenticationRepository(&database.Connection{
			Read: databaseMock, Write: databaseMock}, authUseCases.NewAuthenticationUseCases())

		account, err := repository.GetRepositoryGroups(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestGetWorkspaceRole(t *testing.T) {
	t.Run("should success get workspace role", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").Return(&response.Response{})

		repository := NewAuthenticationRepository(&database.Connection{
			Read: databaseMock, Write: databaseMock}, authUseCases.NewAuthenticationUseCases())

		account, err := repository.GetWorkspaceRole(uuid.New(), uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestGetRepositoryRole(t *testing.T) {
	t.Run("should success get repository role", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").Return(&response.Response{})

		repository := NewAuthenticationRepository(&database.Connection{
			Read: databaseMock, Write: databaseMock}, authUseCases.NewAuthenticationUseCases())

		account, err := repository.GetRepositoryRole(uuid.New(), uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

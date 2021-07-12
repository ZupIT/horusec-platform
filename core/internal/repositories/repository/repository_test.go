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

package repository

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	workspaceRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/workspace"
	repositoryUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
)

func TestNewRepositoryRepository(t *testing.T) {
	t.Run("should success create a repository repository", func(t *testing.T) {
		assert.NotNil(t, NewRepositoryRepository(&database.Connection{}, repositoryUseCases.NewRepositoryUseCases(),
			&workspaceRepository.Mock{}))
	})
}

func TestGetRepositoryByName(t *testing.T) {
	t.Run("should success get repository by name", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").
			Return(response.NewResponse(1, nil, &repositoryEntities.Repository{}))

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), &workspaceRepository.Mock{})

		result, err := repository.GetRepositoryByName(uuid.New(), "test")
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetRepository(t *testing.T) {
	t.Run("should success get a repository", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").
			Return(response.NewResponse(1, nil, &repositoryEntities.Repository{}))

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), &workspaceRepository.Mock{})

		result, err := repository.GetRepository(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetAccountRepository(t *testing.T) {
	t.Run("should success get a account repository", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").
			Return(response.NewResponse(1, nil, &repositoryEntities.Repository{}))

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), &workspaceRepository.Mock{})

		result, err := repository.GetAccountRepository(uuid.New(), uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestListRepositoriesAuthTypeHorusec(t *testing.T) {
	t.Run("should success list repositories when admin", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Raw").Return(&response.Response{})

		workspaceRepositoryMock := &workspaceRepository.Mock{}
		workspaceRepositoryMock.On("GetAccountWorkspace").Return(
			&workspaceEntities.AccountWorkspace{Role: account.Admin}, nil)

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), workspaceRepositoryMock)

		result, err := repository.ListRepositoriesAuthTypeHorusec(uuid.New(), uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success list repositories by role", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Raw").Return(&response.Response{})

		workspaceRepositoryMock := &workspaceRepository.Mock{}
		workspaceRepositoryMock.On("GetAccountWorkspace").Return(
			&workspaceEntities.AccountWorkspace{Role: account.Member}, nil)

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), workspaceRepositoryMock)

		result, err := repository.ListRepositoriesAuthTypeHorusec(uuid.New(), uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get account workspace", func(t *testing.T) {
		databaseMock := &database.Mock{}

		workspaceRepositoryMock := &workspaceRepository.Mock{}
		workspaceRepositoryMock.On("GetAccountWorkspace").Return(
			&workspaceEntities.AccountWorkspace{}, errors.New("test"))

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), workspaceRepositoryMock)

		result, err := repository.ListRepositoriesAuthTypeHorusec(uuid.New(), uuid.New())
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestListRepositoriesAuthTypeLdap(t *testing.T) {
	t.Run("should success list repositories", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Raw").Return(&response.Response{})

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), &workspaceRepository.Mock{})

		result, err := repository.ListRepositoriesAuthTypeLdap(uuid.New(), []string{"test"})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestIsNotMemberOfWorkspace(t *testing.T) {
	t.Run("should return true when user does not belong to workspace", func(t *testing.T) {
		databaseMock := &database.Mock{}

		workspaceRepositoryMock := &workspaceRepository.Mock{}
		workspaceRepositoryMock.On("GetAccountWorkspace").Return(
			&workspaceEntities.AccountWorkspace{}, errors.New("test"))

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), workspaceRepositoryMock)

		assert.True(t, repository.IsNotMemberOfWorkspace(uuid.New(), uuid.New()))
	})

	t.Run("should return false when user belong to workspace", func(t *testing.T) {
		databaseMock := &database.Mock{}

		workspaceRepositoryMock := &workspaceRepository.Mock{}
		workspaceRepositoryMock.On("GetAccountWorkspace").Return(&workspaceEntities.AccountWorkspace{}, nil)

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), workspaceRepositoryMock)

		assert.False(t, repository.IsNotMemberOfWorkspace(uuid.New(), uuid.New()))
	})
}

func TestListAllRepositoryUsers(t *testing.T) {
	t.Run("should success get all repositories users", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Raw").Return(&response.Response{})

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), workspaceRepositoryMock)

		result, err := repository.ListAllRepositoryUsers(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetWorkspace(t *testing.T) {
	t.Run("should success get workspace", func(t *testing.T) {
		databaseMock := &database.Mock{}

		workspaceRepositoryMock := &workspaceRepository.Mock{}
		workspaceRepositoryMock.On("GetWorkspace").Return(&workspaceEntities.Workspace{}, nil)

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), workspaceRepositoryMock)

		result, err := repository.GetWorkspace(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestListRepositoriesWhenApplicationAdmin(t *testing.T) {
	t.Run("should success list repositories", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Raw").Return(&response.Response{})

		repository := NewRepositoryRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			repositoryUseCases.NewRepositoryUseCases(), workspaceRepositoryMock)

		result, err := repository.ListRepositoriesWhenApplicationAdmin()
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

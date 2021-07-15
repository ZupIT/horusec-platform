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
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	tokenEntities "github.com/ZupIT/horusec-platform/core/internal/entities/token"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	repositoryRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/repository"
	workspaceRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/workspace"
	repositoryUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
	tokenUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/token"
)

func TestCreate(t *testing.T) {
	data := &repositoryEntities.Data{
		AccountID:   uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		Permissions: []string{"test"},
	}

	workspace := &workspaceEntities.Workspace{
		WorkspaceID: uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test2"},
		AuthzAdmin:  []string{"test2"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("should success create a new repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, databaseEnums.ErrorNotFoundRecords)
		repositoryMock.On("GetWorkspace").Return(&workspaceEntities.Workspace{}, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.Create(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success create a new repository with the workspace groups", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		data := &repositoryEntities.Data{
			AccountID:   uuid.New(),
			Name:        "test",
			Description: "test",
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
			Permissions: []string{"test"},
		}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, databaseEnums.ErrorNotFoundRecords)
		repositoryMock.On("GetWorkspace").Return(workspace, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		data.AuthzAdmin = []string{}
		data.AuthzMember = []string{}
		data.AuthzSupervisor = []string{}

		result, err := controller.Create(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, pq.StringArray([]string{"test2"}), result.AuthzAdmin)
		assert.Equal(t, pq.StringArray([]string{"test2"}), result.AuthzMember)
		assert.Equal(t, pq.StringArray([]string{"test2"}), result.AuthzSupervisor)
	})

	t.Run("should return error when creating account repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, databaseEnums.ErrorNotFoundRecords)
		repositoryMock.On("GetWorkspace").Return(&workspaceEntities.Workspace{}, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Once().Return(&response.Response{})
		databaseMock.On("Create").Return(
			response.NewResponse(0, errors.New("test"), nil))
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("RollbackTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		_, err := controller.Create(data)
		assert.Error(t, err)
	})

	t.Run("should return error when creating repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, databaseEnums.ErrorNotFoundRecords)
		repositoryMock.On("GetWorkspace").Return(&workspaceEntities.Workspace{}, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(
			response.NewResponse(0, errors.New("test"), nil))
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("RollbackTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		_, err := controller.Create(data)
		assert.Error(t, err)
	})

	t.Run("should return error when failed to get workspace", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, databaseEnums.ErrorNotFoundRecords)
		repositoryMock.On("GetWorkspace").Return(&workspaceEntities.Workspace{}, errors.New("test"))

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		_, err := controller.Create(data)
		assert.Error(t, err)
	})

	t.Run("should return error name already in use", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, nil)

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		_, err := controller.Create(data)
		assert.Error(t, err)
	})
}

func TestGet(t *testing.T) {
	data := &repositoryEntities.Data{
		AccountID:    uuid.New(),
		RepositoryID: uuid.New(),
	}

	t.Run("should success get a repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		workspaceRepositoryMock.On("IsWorkspaceAdmin").Return(false)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, nil)
		repositoryMock.On("GetAccountRepository").Return(&repositoryEntities.AccountRepository{}, nil)

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.Get(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		workspaceRepositoryMock.On("IsWorkspaceAdmin").Return(false)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, errors.New("test"))
		repositoryMock.On("GetAccountRepository").Return(&repositoryEntities.AccountRepository{}, nil)

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		_, err := controller.Get(data)
		assert.Error(t, err)
	})

	t.Run("should return error when failed to get account repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		workspaceRepositoryMock.On("IsWorkspaceAdmin").Return(false)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetAccountRepository").Return(
			&repositoryEntities.AccountRepository{}, errors.New("test"))

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		_, err := controller.Get(data)
		assert.Error(t, err)
	})

	t.Run("should success get a repository when application admin", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		data.IsApplicationAdmin = true
		result, err := controller.Get(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get repository and user is application admin", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, errors.New("test"))

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		data.IsApplicationAdmin = true
		_, err := controller.Get(data)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
	})

	t.Run("should success get a repository when workspace admin", func(t *testing.T) {
		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		workspaceRepositoryMock := &workspaceRepository.Mock{}
		workspaceRepositoryMock.On("IsWorkspaceAdmin").Return(true)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		data.IsApplicationAdmin = false
		result, err := controller.Get(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get repository and user is workspace admin", func(t *testing.T) {
		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		workspaceRepositoryMock := &workspaceRepository.Mock{}
		workspaceRepositoryMock.On("IsWorkspaceAdmin").Return(true)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, errors.New("test"))

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		data.IsApplicationAdmin = false
		_, err := controller.Get(data)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
	})
}

func TestUpdate(t *testing.T) {
	data := &repositoryEntities.Data{
		AccountID:       uuid.New(),
		Name:            "test",
		Description:     "test",
		AuthzMember:     []string{"test"},
		AuthzAdmin:      []string{"test"},
		AuthzSupervisor: []string{"test"},
		Permissions:     []string{"test"},
	}

	t.Run("should success update repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{Name: "test2"}, nil)
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, databaseEnums.ErrorNotFoundRecords)

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.Update(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error name already in use", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, nil)
		repositoryMock.On("GetRepositoryByName").Return(
			&repositoryEntities.Repository{}, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		_, err := controller.Update(data)
		assert.Error(t, err)
	})

	t.Run("should return error while getting repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, errors.New("test"))

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		_, err := controller.Update(data)
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should success delete repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		repositoryMock := &repositoryRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		assert.NoError(t, controller.Delete(uuid.New()))
	})
}

func TestList(t *testing.T) {
	data := &repositoryEntities.Data{
		AccountID:   uuid.New(),
		WorkspaceID: uuid.New(),
		Permissions: []string{"test"},
	}

	t.Run("should success list repositories when horusec auth type", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		databaseMock := &database.Mock{}

		appConfig := &app.Mock{}
		appConfig.On("GetAuthenticationType").Return(auth.Horusec)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("ListRepositoriesAuthTypeHorusec").Return(&[]repositoryEntities.Response{}, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.List(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success list repositories when ldap auth type", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		databaseMock := &database.Mock{}

		appConfig := &app.Mock{}
		appConfig.On("GetAuthenticationType").Return(auth.Ldap)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("ListRepositoriesAuthTypeLdap").Return(&[]repositoryEntities.Response{}, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.List(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success list repositories when application admin", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		databaseMock := &database.Mock{}

		appConfig := &app.Mock{}
		appConfig.On("GetAuthenticationType").Return(auth.Horusec)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("ListRepositoriesWhenApplicationAdmin").Return(
			&[]repositoryEntities.Response{}, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		data.IsApplicationAdmin = true
		result, err := controller.List(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestUpdateRole(t *testing.T) {
	data := &roleEntities.Data{
		Role:         account.Member,
		AccountID:    uuid.New(),
		WorkspaceID:  uuid.New(),
		RepositoryID: uuid.New(),
	}

	t.Run("should success update user role", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("IsNotMemberOfWorkspace").Return(false)
		repositoryMock.On("GetAccountRepository").Return(&repositoryEntities.AccountRepository{}, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.UpdateRole(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to update", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("IsNotMemberOfWorkspace").Return(false)
		repositoryMock.On("GetAccountRepository").Return(&repositoryEntities.AccountRepository{}, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(
			response.NewResponse(0, errors.New("test"), nil))

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		_, err := controller.UpdateRole(data)
		assert.Error(t, err)
	})

	t.Run("should return error when failed to get account repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("IsNotMemberOfWorkspace").Return(false)
		repositoryMock.On("GetAccountRepository").Return(
			&repositoryEntities.AccountRepository{}, errors.New("test"))

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.UpdateRole(data)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when user does not belong to workspace", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("IsNotMemberOfWorkspace").Return(true)

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.UpdateRole(data)
		assert.Error(t, err)
		assert.Equal(t, repositoryEnums.ErrorUserDoesNotBelongToWorkspace, err)
		assert.Nil(t, result)
	})
}

func TestInviteUser(t *testing.T) {
	data := &roleEntities.UserData{
		Role:         account.Member,
		AccountID:    uuid.New(),
		WorkspaceID:  uuid.New(),
		RepositoryID: uuid.New(),
	}

	t.Run("should success invite user with broker enabled", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		appConfig := &app.Mock{}
		appConfig.On("IsEmailsDisabled").Return(false)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("IsNotMemberOfWorkspace").Return(false)
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})

		brokerMock := &broker.Mock{}
		brokerMock.On("Publish").Return(nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(brokerMock, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.InviteUser(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success invite user with broker disabled", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}

		appConfig := &app.Mock{}
		appConfig.On("IsEmailsDisabled").Return(true)

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("IsNotMemberOfWorkspace").Return(false)
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.InviteUser(data)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to create", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("IsNotMemberOfWorkspace").Return(false)
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(
			response.NewResponse(0, errors.New("test"), nil))

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.InviteUser(data)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error when failed to get repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}
		databaseMock := &database.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("IsNotMemberOfWorkspace").Return(false)
		repositoryMock.On("GetRepository").Return(&repositoryEntities.Repository{}, errors.New("test"))

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.InviteUser(data)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("should return error member not member of repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}
		databaseMock := &database.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("IsNotMemberOfWorkspace").Return(true)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.InviteUser(data)
		assert.Error(t, err)
		assert.Equal(t, repositoryEnums.ErrorUserDoesNotBelongToWorkspace, err)
		assert.Nil(t, result)
	})
}

func TestGetUsers(t *testing.T) {
	t.Run("should success get repository users", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}
		databaseMock := &database.Mock{}

		repositoryMock := &repositoryRepository.Mock{}
		repositoryMock.On("ListAllRepositoryUsers").Return(&[]roleEntities.Response{}, nil)

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.GetUsers(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestRemoveUser(t *testing.T) {
	data := &roleEntities.Data{
		Role:         account.Member,
		AccountID:    uuid.New(),
		WorkspaceID:  uuid.New(),
		RepositoryID: uuid.New(),
	}

	t.Run("should success remove user from repository", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		appConfig := &app.Mock{}
		repositoryMock := &repositoryRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		assert.NoError(t, controller.RemoveUser(data))
	})
}

func TestCreateToken(t *testing.T) {
	data := &tokenEntities.Data{}

	t.Run("should success create a new repository token ", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		repositoryMock := &repositoryRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.CreateToken(data)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})
}

func TestDeleteToken(t *testing.T) {
	t.Run("should success delete a repository token ", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		repositoryMock := &repositoryRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		assert.NoError(t, controller.DeleteToken(&tokenEntities.Data{}))
	})
}

func TestListTokens(t *testing.T) {
	t.Run("should success list repository tokens", func(t *testing.T) {
		workspaceRepositoryMock := &workspaceRepository.Mock{}
		repositoryMock := &repositoryRepository.Mock{}
		appConfig := &app.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Find").Return(&response.Response{})

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewRepositoryController(&broker.Mock{}, databaseConnection, appConfig,
			repositoryUseCases.NewRepositoryUseCases(), repositoryMock, &tokenUseCases.UseCases{},
			workspaceRepositoryMock)

		result, err := controller.ListTokens(&tokenEntities.Data{})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

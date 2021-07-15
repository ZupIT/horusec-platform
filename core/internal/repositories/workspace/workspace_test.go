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

package workspace

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

func TestNewWorkspaceRepository(t *testing.T) {
	t.Run("should success create a workspace repository", func(t *testing.T) {
		assert.NotNil(t, NewWorkspaceRepository(&database.Connection{}, workspaceUseCases.NewWorkspaceUseCases()))
	})
}

func TestGetWorkspace(t *testing.T) {
	t.Run("should success get a workspace", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").
			Return(response.NewResponse(1, nil, &workspaceEntities.Workspace{}))

		repository := NewWorkspaceRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			workspaceUseCases.NewWorkspaceUseCases())

		result, err := repository.GetWorkspace(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetAccountWorkspace(t *testing.T) {
	t.Run("should success get a account workspace", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Find").
			Return(response.NewResponse(1, nil, &workspaceEntities.AccountWorkspace{}))

		repository := NewWorkspaceRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			workspaceUseCases.NewWorkspaceUseCases())

		result, err := repository.GetAccountWorkspace(uuid.New(), uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestListWorkspacesAuthTypeHorusec(t *testing.T) {
	t.Run("should success get all user workspaces when auth type horusec", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Raw").
			Return(response.NewResponse(1, nil, &[]workspaceEntities.Workspace{}))

		repository := NewWorkspaceRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			workspaceUseCases.NewWorkspaceUseCases())

		result, err := repository.ListWorkspacesAuthTypeHorusec(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestListWorkspacesAuthTypeLdap(t *testing.T) {
	t.Run("should success get all user workspaces when auth type ldap", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Raw").
			Return(response.NewResponse(1, nil, &[]workspaceEntities.Workspace{}))

		repository := NewWorkspaceRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			workspaceUseCases.NewWorkspaceUseCases())

		result, err := repository.ListWorkspacesAuthTypeLdap([]string{"test"})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestListAllWorkspaceUsers(t *testing.T) {
	t.Run("should success get all workspace users", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Raw").
			Return(response.NewResponse(1, nil, &[]roleEntities.Response{}))

		repository := NewWorkspaceRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			workspaceUseCases.NewWorkspaceUseCases())

		result, err := repository.ListAllWorkspaceUsers(uuid.New())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestListWorkspacesApplicationAdmin(t *testing.T) {
	t.Run("should success list workspaces", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Raw").
			Return(response.NewResponse(1, nil, &[]roleEntities.Response{}))

		repository := NewWorkspaceRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			workspaceUseCases.NewWorkspaceUseCases())

		result, err := repository.ListWorkspacesApplicationAdmin()
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestIsWorkspaceAdmin(t *testing.T) {
	t.Run("should return false for admin", func(t *testing.T) {
		accountWorkspace := &workspaceEntities.AccountWorkspace{Role: accountEnums.Admin}

		databaseMock := &database.Mock{}
		databaseMock.On("Find").Return(response.NewResponse(1, nil, accountWorkspace))

		repository := NewWorkspaceRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			workspaceUseCases.NewWorkspaceUseCases())

		assert.False(t, repository.IsWorkspaceAdmin(uuid.New(), uuid.New()))
	})
}

func TestGetWorkspaceLdap(t *testing.T) {
	t.Run("should success get workspace", func(t *testing.T) {
		databaseMock := &database.Mock{}
		databaseMock.On("Raw").Return(&response.Response{})

		repository := NewWorkspaceRepository(&database.Connection{Read: databaseMock, Write: databaseMock},
			workspaceUseCases.NewWorkspaceUseCases())

		result, err := repository.GetWorkspaceLdap(uuid.New(), []string{})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

package workspace

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

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

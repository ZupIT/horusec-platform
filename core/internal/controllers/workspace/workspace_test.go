package workspace

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	workspaceRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/workspace"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

func TestNewWorkspaceController(t *testing.T) {
	t.Run("should success create a new workspace controller", func(t *testing.T) {
		assert.NotNil(t, NewWorkspaceController(&broker.Broker{}, &database.Connection{},
			&app.Config{}, workspaceUseCases.NewWorkspaceUseCases(), &workspaceRepository.Repository{}))
	})
}

func TestCreate(t *testing.T) {
	workspaceData := &workspaceEntities.Data{
		AccountID:   uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		Permissions: []string{"test"},
	}

	t.Run("should success create a new workspace with horusec auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Create(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success create a new workspace with ldap auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Create(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success create a new workspace with keycloak auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Create(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should success create a new workspace with keycloak auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(&response.Response{})
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("CommitTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Create(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when creating account workspace relation", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Once().Return(&response.Response{})
		databaseMock.On("Create").Return(response.NewResponse(1,
			errors.New("test"), nil))
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("RollbackTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Create(workspaceData)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
		assert.Nil(t, result)
	})

	t.Run("should return error when creating workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Create").Return(response.NewResponse(1,
			errors.New("test"), nil))
		databaseMock.On("StartTransaction").Return(databaseMock)
		databaseMock.On("RollbackTransaction").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Create(workspaceData)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
		assert.Nil(t, result)
	})
}

func TestGet(t *testing.T) {
	workspaceData := &workspaceEntities.Data{
		AccountID:   uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		Permissions: []string{"test"},
	}

	accountWorkspace := &workspaceEntities.AccountWorkspace{
		WorkspaceID: uuid.New(),
		AccountID:   uuid.New(),
		Role:        account.Admin,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	workspace := &workspaceEntities.Workspace{
		WorkspaceID: uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("should success get workspace with role", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		repositoryMock.On("GetAccountWorkspace").Return(accountWorkspace, nil)
		repositoryMock.On("GetWorkspace").Return(workspace, nil)

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Get(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		repositoryMock.On("GetAccountWorkspace").Return(accountWorkspace, nil)
		repositoryMock.On("GetWorkspace").Return(workspace, errors.New("test"))

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		_, err := controller.Get(workspaceData)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
	})

	t.Run("should return error when failed to get workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		repositoryMock.On("GetAccountWorkspace").Return(accountWorkspace, errors.New("test"))

		databaseMock := &database.Mock{}
		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		_, err := controller.Get(workspaceData)
		assert.Error(t, err)
		assert.Equal(t, errors.New("test"), err)
	})
}

func TestUpdate(t *testing.T) {
	workspaceData := &workspaceEntities.Data{
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
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("should success get workspace with role", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetWorkspace").Return(workspace, nil)

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.Update(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to get workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("GetWorkspace").Return(workspace, errors.New("test"))

		databaseMock := &database.Mock{}
		databaseMock.On("Update").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		_, err := controller.Update(workspaceData)
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should success delete workspace", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}

		databaseMock := &database.Mock{}
		databaseMock.On("Delete").Return(&response.Response{})

		appConfig := &app.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		err := controller.Delete(uuid.New())
		assert.NoError(t, err)
	})
}

func TestList(t *testing.T) {
	workspaceData := &workspaceEntities.Data{
		AccountID:   uuid.New(),
		Name:        "test",
		Description: "test",
		AuthzMember: []string{"test"},
		AuthzAdmin:  []string{"test"},
		Permissions: []string{"test"},
	}

	workspaceResponse := &[]workspaceEntities.Response{
		{
			WorkspaceID: uuid.New(),
			Name:        "test",
			Role:        account.Admin,
			Description: "test",
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	t.Run("should list workspaces when horusec auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("ListWorkspacesAuthTypeHorusec").Return(workspaceResponse, nil)

		appConfig := &app.Mock{}
		appConfig.On("GetAuthorizationType").Return(auth.Horusec)

		databaseMock := &database.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.List(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should list workspaces when ldap auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("ListWorkspacesAuthTypeLdap").Return(workspaceResponse, nil)

		appConfig := &app.Mock{}
		appConfig.On("GetAuthorizationType").Return(auth.Ldap)

		databaseMock := &database.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		result, err := controller.List(workspaceData)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when failed to list with horusec auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("ListWorkspacesAuthTypeHorusec").Return(workspaceResponse, errors.New("test"))

		appConfig := &app.Mock{}
		appConfig.On("GetAuthorizationType").Return(auth.Horusec)

		databaseMock := &database.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		_, err := controller.List(workspaceData)
		assert.Error(t, err)
	})

	t.Run("should return error when failed to list with ldap auth type", func(t *testing.T) {
		repositoryMock := &workspaceRepository.Mock{}
		repositoryMock.On("ListWorkspacesAuthTypeLdap").Return(workspaceResponse, errors.New("test"))

		appConfig := &app.Mock{}
		appConfig.On("GetAuthorizationType").Return(auth.Ldap)

		databaseMock := &database.Mock{}

		databaseConnection := &database.Connection{Read: databaseMock, Write: databaseMock}
		controller := NewWorkspaceController(&broker.Broker{}, databaseConnection, appConfig,
			workspaceUseCases.NewWorkspaceUseCases(), repositoryMock)

		_, err := controller.List(workspaceData)
		assert.Error(t, err)
	})
}

package workspace

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	brokerService "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
	"github.com/google/uuid"

	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	workspaceEnums "github.com/ZupIT/horusec-platform/core/internal/enums/workspace"
	workspaceRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/workspace"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

type IController interface {
	Create(data *workspaceEntities.Data) (*workspaceEntities.Workspace, error)
	Get(data *workspaceEntities.Data) (*workspaceEntities.Response, error)
	Update(data *workspaceEntities.Data) (*workspaceEntities.Workspace, error)
	Delete(workspaceID uuid.UUID) error
	List(data *workspaceEntities.Data) (*[]workspaceEntities.Response, error)
}

type Controller struct {
	broker        brokerService.IBroker
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	appConfig     app.IConfig
	useCases      workspaceUseCases.IUseCases
	repository    workspaceRepository.IRepository
}

func NewWorkspaceController(broker brokerService.IBroker, databaseConnection *database.Connection, appConfig app.IConfig,
	useCases workspaceUseCases.IUseCases, repository workspaceRepository.IRepository) IController {
	return &Controller{
		broker:        broker,
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
		appConfig:     appConfig,
		useCases:      useCases,
		repository:    repository,
	}
}

func (c *Controller) Create(data *workspaceEntities.Data) (*workspaceEntities.Workspace, error) {
	transaction := c.databaseWrite.StartTransaction()
	workspace := data.ToWorkspace()

	if err := transaction.Create(workspace, workspaceEnums.DatabaseWorkspaceTable).GetError(); err != nil {
		logger.LogError(workspaceEnums.ErrorRollbackCreate, transaction.RollbackTransaction().GetError())
		return nil, err
	}

	if err := transaction.Create(workspace.ToAccountWorkspace(data.AccountID, account.Admin),
		workspaceEnums.DatabaseAccountWorkspaceTable).GetError(); err != nil {
		logger.LogError(workspaceEnums.ErrorRollbackCreate, transaction.RollbackTransaction().GetError())
		return nil, err
	}

	return workspace, transaction.CommitTransaction().GetError()
}

func (c *Controller) Get(data *workspaceEntities.Data) (*workspaceEntities.Response, error) {
	accountWorkspace := &workspaceEntities.AccountWorkspace{}

	if err := c.databaseRead.Find(accountWorkspace,
		c.useCases.FilterAccountWorkspaceByID(data.AccountID, data.WorkspaceID),
		workspaceEnums.DatabaseAccountWorkspaceTable).GetError(); err != nil {
		return nil, err
	}

	workspace, err := c.getWorkspace(data)
	if err != nil {
		return nil, err
	}

	return workspace.ToWorkspaceResponse(accountWorkspace.Role), nil
}

func (c *Controller) getWorkspace(data *workspaceEntities.Data) (*workspaceEntities.Workspace, error) {
	workspace := &workspaceEntities.Workspace{}

	return workspace, c.databaseRead.Find(workspace, c.useCases.FilterWorkspaceByID(data.WorkspaceID),
		workspaceEnums.DatabaseWorkspaceTable).GetError()
}

func (c *Controller) Update(data *workspaceEntities.Data) (*workspaceEntities.Workspace, error) {
	workspace, err := c.getWorkspace(data)
	if err != nil {
		return nil, err
	}

	return workspace, c.databaseWrite.Update(workspace.Update(data), c.useCases.FilterWorkspaceByID(data.WorkspaceID),
		workspaceEnums.DatabaseWorkspaceTable).GetError()
}

func (c *Controller) Delete(workspaceID uuid.UUID) error {
	return c.databaseWrite.Delete(c.useCases.FilterWorkspaceByID(workspaceID),
		workspaceEnums.DatabaseWorkspaceTable).GetError()
}

func (c *Controller) List(data *workspaceEntities.Data) (*[]workspaceEntities.Response, error) {
	if c.appConfig.GetAuthorizationType() == auth.Ldap {
		return c.repository.ListWorkspacesAuthTypeLdap(data.Permissions)
	}

	return c.repository.ListWorkspacesAuthTypeHorusec(data.AccountID)
}

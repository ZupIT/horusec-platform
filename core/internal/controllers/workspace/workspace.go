package workspace

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	brokerService "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	workspaceEnums "github.com/ZupIT/horusec-platform/core/internal/enums/workspace"
)

type IController interface {
	Create(data *workspaceEntities.CreateWorkspaceData) (*workspaceEntities.Workspace, error)
}

type Controller struct {
	broker        brokerService.IBroker
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	appConfig     app.IConfig
}

func NewWorkspaceController(broker brokerService.IBroker, databaseConnection *database.Connection,
	appConfig app.IConfig) IController {
	return &Controller{
		broker:        broker,
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
		appConfig:     appConfig,
	}
}

func (c *Controller) Create(data *workspaceEntities.CreateWorkspaceData) (*workspaceEntities.Workspace, error) {
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

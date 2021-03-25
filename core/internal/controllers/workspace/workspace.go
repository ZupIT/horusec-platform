package workspace

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	workspaceEnums "github.com/ZupIT/horusec-platform/core/internal/enums/workspace"
	"github.com/ZupIT/horusec-platform/core/internal/utils"
)

type IController interface {
	Create(data *workspaceEntities.CreateWorkspaceData) (*workspaceEntities.Workspace, error)
}

type Controller struct {
	broker        broker.IBroker
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	appConfig     app.IConfig
}

func NewWorkspaceController(broker broker.IBroker, databaseConnection *database.Connection,
	appConfig app.IConfig) IController {
	return &Controller{
		broker:        broker,
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
		appConfig:     appConfig,
	}
}

func (c *Controller) Create(data *workspaceEntities.CreateWorkspaceData) (*workspaceEntities.Workspace, error) {
	if utils.IsInvalidLdapGroups(c.appConfig.GetAuthorizationType(), data.AuthzAdmin, data.Permissions) {
		return nil, workspaceEnums.ErrorInvalidLdapGroup
	}

	return c.createWorkspaceTransaction(data)
}

func (c *Controller) createWorkspaceTransaction(
	data *workspaceEntities.CreateWorkspaceData) (*workspaceEntities.Workspace, error) {
	transaction := c.databaseWrite.StartTransaction()
	workspace := data.ToWorkspace()

	if err := transaction.Create(workspace, workspaceEnums.DatabaseWorkspaceTable).GetError(); err != nil {
		return nil, transaction.RollbackTransaction().GetError()
	}

	if err := transaction.Create(workspace.ToAccountWorkspace(data.AccountID, account.Admin),
		workspaceEnums.DatabaseAccountWorkspaceTable).GetError(); err != nil {
		return nil, transaction.RollbackTransaction().GetError()
	}

	return workspace, transaction.CommitTransaction().GetError()
}

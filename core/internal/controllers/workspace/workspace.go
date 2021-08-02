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
	"github.com/google/uuid"

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	brokerService "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	tokenEntities "github.com/ZupIT/horusec-platform/core/internal/entities/token"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	tokenEnums "github.com/ZupIT/horusec-platform/core/internal/enums/token"
	workspaceEnums "github.com/ZupIT/horusec-platform/core/internal/enums/workspace"
	workspaceRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/workspace"
	tokenUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/token"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

type IController interface {
	Create(data *workspaceEntities.Data) (*workspaceEntities.Response, error)
	Get(data *workspaceEntities.Data) (*workspaceEntities.Response, error)
	Update(data *workspaceEntities.Data) (*workspaceEntities.Response, error)
	Delete(workspaceID uuid.UUID) error
	List(data *workspaceEntities.Data) (*[]workspaceEntities.Response, error)
	UpdateRole(data *roleEntities.Data) (*roleEntities.Response, error)
	InviteUser(data *roleEntities.UserData) (*roleEntities.Response, error)
	GetUsers(workspaceID uuid.UUID, noBelongRepositoryID uuid.UUID) (*[]roleEntities.Response, error)
	RemoveUser(data *roleEntities.Data) error
	CreateToken(data *tokenEntities.Data) (string, error)
	DeleteToken(data *tokenEntities.Data) error
	ListTokens(workspaceID uuid.UUID) (*[]tokenEntities.Response, error)
}

type Controller struct {
	broker        brokerService.IBroker
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	appConfig     app.IConfig
	useCases      workspaceUseCases.IUseCases
	repository    workspaceRepository.IRepository
	tokenUseCases tokenUseCases.IUseCases
}

func NewWorkspaceController(broker brokerService.IBroker, databaseConnection *database.Connection,
	appConfig app.IConfig, useCases workspaceUseCases.IUseCases, repository workspaceRepository.IRepository,
	useCasesToken tokenUseCases.IUseCases) IController {
	return &Controller{
		broker:        broker,
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
		appConfig:     appConfig,
		useCases:      useCases,
		repository:    repository,
		tokenUseCases: useCasesToken,
	}
}

func (c *Controller) Create(data *workspaceEntities.Data) (*workspaceEntities.Response, error) {
	transaction := c.databaseWrite.StartTransaction()
	workspace := data.ToWorkspace()

	if err := transaction.Create(workspace, workspaceEnums.DatabaseWorkspaceTable).GetError(); err != nil {
		logger.LogError(workspaceEnums.ErrorRollbackCreate, transaction.RollbackTransaction().GetError())
		return nil, err
	}

	if err := transaction.Create(workspace.ToAccountWorkspace(data.AccountID, accountEnums.Admin),
		workspaceEnums.DatabaseAccountWorkspaceTable).GetError(); err != nil {
		logger.LogError(workspaceEnums.ErrorRollbackCreate, transaction.RollbackTransaction().GetError())
		return nil, err
	}

	return workspace.ToWorkspaceResponse(accountEnums.Admin), transaction.CommitTransaction().GetError()
}

func (c *Controller) Get(data *workspaceEntities.Data) (*workspaceEntities.Response, error) {
	if data.IsApplicationAdmin {
		return c.getWorkspaceWhenAppAdmin(data)
	}

	if c.appConfig.GetAuthenticationType() == auth.Ldap {
		return c.repository.GetWorkspaceLdap(data.WorkspaceID, data.Permissions)
	}

	return c.getWorkspace(data)
}

func (c *Controller) getWorkspaceWhenAppAdmin(data *workspaceEntities.Data) (*workspaceEntities.Response, error) {
	workspace, err := c.repository.GetWorkspace(data.WorkspaceID)
	if err != nil {
		return nil, err
	}

	return workspace.ToWorkspaceResponse(accountEnums.ApplicationAdmin), nil
}

func (c *Controller) getWorkspace(data *workspaceEntities.Data) (*workspaceEntities.Response, error) {
	accountWorkspace, err := c.repository.GetAccountWorkspace(data.AccountID, data.WorkspaceID)
	if err != nil {
		return nil, err
	}

	workspace, err := c.repository.GetWorkspace(data.WorkspaceID)
	if err != nil {
		return nil, err
	}

	return workspace.ToWorkspaceResponse(accountWorkspace.Role), nil
}

func (c *Controller) Update(data *workspaceEntities.Data) (*workspaceEntities.Response, error) {
	workspace, err := c.repository.GetWorkspace(data.WorkspaceID)
	if err != nil {
		return nil, err
	}

	workspace.Update(data)
	return workspace.ToWorkspaceResponse(accountEnums.Admin), c.databaseWrite.Update(
		workspace, c.useCases.FilterWorkspaceByID(data.WorkspaceID), workspaceEnums.DatabaseWorkspaceTable).GetError()
}

func (c *Controller) Delete(workspaceID uuid.UUID) error {
	return c.databaseWrite.Delete(c.useCases.FilterWorkspaceByID(workspaceID),
		workspaceEnums.DatabaseWorkspaceTable).GetError()
}

func (c *Controller) List(data *workspaceEntities.Data) (*[]workspaceEntities.Response, error) {
	if data.IsApplicationAdmin {
		return c.repository.ListWorkspacesApplicationAdmin()
	}

	if c.appConfig.GetAuthenticationType() == auth.Ldap {
		return c.repository.ListWorkspacesAuthTypeLdap(data.Permissions)
	}

	return c.repository.ListWorkspacesAuthTypeHorusec(data.AccountID)
}

func (c *Controller) UpdateRole(data *roleEntities.Data) (*roleEntities.Response, error) {
	accountWorkspace, err := c.repository.GetAccountWorkspace(data.AccountID, data.WorkspaceID)
	if err != nil {
		return nil, err
	}

	accountWorkspace.Update(data)
	return accountWorkspace.ToResponse(), c.databaseWrite.Update(accountWorkspace,
		c.useCases.FilterAccountWorkspaceByID(data.AccountID, data.WorkspaceID),
		workspaceEnums.DatabaseAccountWorkspaceTable).GetError()
}

func (c *Controller) InviteUser(data *roleEntities.UserData) (*roleEntities.Response, error) {
	workspace, err := c.repository.GetWorkspace(data.WorkspaceID)
	if err != nil {
		return nil, err
	}

	accountWorkspace := workspace.ToAccountWorkspace(data.AccountID, data.Role)
	if err := c.databaseWrite.Create(accountWorkspace,
		workspaceEnums.DatabaseAccountWorkspaceTable).GetError(); err != nil {
		return nil, err
	}

	return accountWorkspace.ToResponseWithEmailAndUsername(data.Email, data.Username),
		c.sendInviteUserEmail(data.Email, data.Username, workspace.Name)
}

func (c *Controller) sendInviteUserEmail(email, username, workspaceName string) error {
	if c.appConfig.IsEmailsDisabled() {
		return nil
	}

	return c.broker.Publish(queues.HorusecEmail.ToString(), "", "",
		c.useCases.NewOrganizationInviteEmail(email, username, workspaceName))
}

func (c *Controller) GetUsers(workspaceID, noBelongRepositoryID uuid.UUID) (*[]roleEntities.Response, error) {
	if noBelongRepositoryID != uuid.Nil {
		return c.repository.ListWorkspaceUsersNoBelong(workspaceID, noBelongRepositoryID)
	}

	return c.repository.ListAllWorkspaceUsers(workspaceID)
}

func (c *Controller) RemoveUser(data *roleEntities.Data) error {
	if err := c.databaseWrite.Delete(c.useCases.FilterAccountWorkspaceByID(data.AccountID, data.WorkspaceID),
		repositoryEnums.DatabaseAccountRepositoryTable).GetError(); err != nil {
		return err
	}

	return c.databaseWrite.Delete(c.useCases.FilterAccountWorkspaceByID(data.AccountID, data.WorkspaceID),
		workspaceEnums.DatabaseAccountWorkspaceTable).GetError()
}

func (c *Controller) CreateToken(data *tokenEntities.Data) (string, error) {
	token, tokenString := data.ToToken()

	return tokenString, c.databaseWrite.Create(token, tokenEnums.DatabaseTokens).GetError()
}

func (c *Controller) DeleteToken(data *tokenEntities.Data) error {
	return c.databaseWrite.Delete(c.tokenUseCases.FilterWorkspaceTokenByID(data.TokenID, data.WorkspaceID),
		tokenEnums.DatabaseTokens).GetError()
}

func (c *Controller) ListTokens(workspaceID uuid.UUID) (*[]tokenEntities.Response, error) {
	tokens := &[]tokenEntities.Response{}

	return tokens, c.databaseRead.Find(tokens, c.tokenUseCases.FilterListWorkspaceTokens(workspaceID),
		tokenEnums.DatabaseTokens).GetErrorExceptNotFound()
}

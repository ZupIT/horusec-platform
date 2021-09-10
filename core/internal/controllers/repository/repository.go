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
	"github.com/google/uuid"

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	brokerService "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	tokenEntities "github.com/ZupIT/horusec-platform/core/internal/entities/token"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	tokenEnums "github.com/ZupIT/horusec-platform/core/internal/enums/token"
	repositoryRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/repository"
	workspaceRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/workspace"
	repositoriesUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
	tokenUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/token"
)

type IController interface {
	Create(data *repositoryEntities.Data) (*repositoryEntities.Response, error)
	Get(data *repositoryEntities.Data) (*repositoryEntities.Response, error)
	Update(data *repositoryEntities.Data) (*repositoryEntities.Response, error)
	Delete(repositoryID uuid.UUID) error
	List(repositoryData *repositoryEntities.Data,
		paginatedData *repositoryEntities.PaginatedContent) (*[]repositoryEntities.Response, error)
	UpdateRole(data *roleEntities.Data) (*roleEntities.Response, error)
	InviteUser(data *roleEntities.UserData) (*roleEntities.Response, error)
	GetUsers(repositoryID uuid.UUID) (*[]roleEntities.Response, error)
	RemoveUser(data *roleEntities.Data) error
	CreateToken(data *tokenEntities.Data) (string, error)
	DeleteToken(data *tokenEntities.Data) error
	ListTokens(data *tokenEntities.Data) (*[]tokenEntities.Response, error)
}

type Controller struct {
	broker              brokerService.IBroker
	databaseRead        database.IDatabaseRead
	databaseWrite       database.IDatabaseWrite
	appConfig           app.IConfig
	useCases            repositoriesUseCases.IUseCases
	repository          repositoryRepository.IRepository
	workspaceRepository workspaceRepository.IRepository
	tokenUseCases       tokenUseCases.IUseCases
}

func NewRepositoryController(broker brokerService.IBroker, databaseConnection *database.Connection,
	appConfig app.IConfig, useCases repositoriesUseCases.IUseCases, repository repositoryRepository.IRepository,
	useCasesToken tokenUseCases.IUseCases, repositoryWorkspace workspaceRepository.IRepository) IController {
	return &Controller{
		databaseRead:        databaseConnection.Read,
		databaseWrite:       databaseConnection.Write,
		appConfig:           appConfig,
		useCases:            useCases,
		repository:          repository,
		broker:              broker,
		tokenUseCases:       useCasesToken,
		workspaceRepository: repositoryWorkspace,
	}
}

func (c *Controller) Create(data *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	_, err := c.repository.GetRepositoryByName(data.WorkspaceID, data.Name)
	if !c.useCases.IsNotFoundError(err) {
		return nil, repositoryEnums.ErrorRepositoryNameAlreadyInUse
	}

	workspace, err := c.repository.GetWorkspace(data.WorkspaceID)
	if err != nil {
		return nil, err
	}

	return c.createTransaction(data.AccountID, c.useCases.InheritWorkspaceGroups(data.ToRepository(), workspace))
}

func (c *Controller) createTransaction(accountID uuid.UUID,
	repository *repositoryEntities.Repository) (*repositoryEntities.Response, error) {
	transaction := c.databaseWrite.StartTransaction()

	if err := transaction.Create(repository, repositoryEnums.DatabaseRepositoryTable).GetError(); err != nil {
		logger.LogError(repositoryEnums.ErrorRollbackCreate, transaction.RollbackTransaction().GetError())
		return nil, err
	}

	if err := transaction.Create(repository.ToAccountRepository(accountID, accountEnums.Admin),
		repositoryEnums.DatabaseAccountRepositoryTable).GetError(); err != nil {
		logger.LogError(repositoryEnums.ErrorRollbackCreate, transaction.RollbackTransaction().GetError())
		return nil, err
	}

	return repository.ToRepositoryResponse(accountEnums.Admin), transaction.CommitTransaction().GetError()
}

func (c *Controller) Get(data *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	if data.IsApplicationAdmin {
		return c.getRepositoryWhenAdmin(data, accountEnums.ApplicationAdmin)
	}

	if c.appConfig.GetAuthenticationType() == auth.Ldap {
		return c.repository.GetRepositoryLdap(data.RepositoryID, data.Permissions)
	}

	if c.workspaceRepository.IsWorkspaceAdmin(data.AccountID, data.WorkspaceID) {
		return c.getRepositoryWhenAdmin(data, accountEnums.Admin)
	}

	return c.getRepository(data)
}

func (c *Controller) getRepositoryWhenAdmin(
	data *repositoryEntities.Data, role accountEnums.Role) (*repositoryEntities.Response, error) {
	repository, err := c.repository.GetRepository(data.RepositoryID)
	if err != nil {
		return nil, err
	}

	return repository.ToRepositoryResponse(role), nil
}

func (c *Controller) getRepository(data *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	accountRepository, err := c.repository.GetAccountRepository(data.AccountID, data.RepositoryID)
	if err != nil {
		return nil, err
	}

	repository, err := c.repository.GetRepository(data.RepositoryID)
	if err != nil {
		return nil, err
	}

	return repository.ToRepositoryResponse(accountRepository.Role), nil
}

func (c *Controller) Update(data *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	repository, err := c.repository.GetRepository(data.RepositoryID)
	if err != nil {
		return nil, err
	}

	_, err = c.repository.GetRepositoryByName(data.WorkspaceID, data.Name)
	if repository.Name != data.Name && !c.useCases.IsNotFoundError(err) {
		return nil, repositoryEnums.ErrorRepositoryNameAlreadyInUse
	}

	return repository.Update(data).ToRepositoryResponse(accountEnums.Admin),
		c.databaseWrite.Update(repository.ToUpdateMap(data), c.useCases.FilterRepositoryByID(data.RepositoryID),
			repositoryEnums.DatabaseRepositoryTable).GetError()
}

func (c *Controller) Delete(repositoryID uuid.UUID) error {
	return c.databaseWrite.Delete(c.useCases.FilterRepositoryByID(repositoryID),
		repositoryEnums.DatabaseRepositoryTable).GetError()
}

func (c *Controller) List(repositoryData *repositoryEntities.Data,
	paginatedData *repositoryEntities.PaginatedContent) (*[]repositoryEntities.Response, error) {
	if repositoryData.IsApplicationAdmin {
		return c.repository.ListRepositoriesWhenApplicationAdmin(paginatedData, repositoryData.WorkspaceID)
	}

	if c.appConfig.GetAuthenticationType() == auth.Ldap {
		return c.repository.ListRepositoriesAuthTypeLdap(
			repositoryData.WorkspaceID, repositoryData.Permissions, paginatedData)
	}

	return c.repository.ListRepositoriesAuthTypeHorusec(
		repositoryData.AccountID, repositoryData.WorkspaceID, paginatedData)
}

func (c *Controller) UpdateRole(data *roleEntities.Data) (*roleEntities.Response, error) {
	if c.repository.IsNotMemberOfWorkspace(data.AccountID, data.WorkspaceID) {
		return nil, repositoryEnums.ErrorUserDoesNotBelongToWorkspace
	}

	accountRepository, err := c.repository.GetAccountRepository(data.AccountID, data.RepositoryID)
	if err != nil {
		return nil, err
	}

	accountRepository.Update(data.Role)
	return accountRepository.ToResponse(), c.databaseWrite.Update(accountRepository,
		c.useCases.FilterAccountRepositoryByID(data.AccountID, data.RepositoryID),
		repositoryEnums.DatabaseAccountRepositoryTable).GetError()
}

func (c *Controller) InviteUser(data *roleEntities.UserData) (*roleEntities.Response, error) {
	if c.repository.IsNotMemberOfWorkspace(data.AccountID, data.WorkspaceID) {
		return nil, repositoryEnums.ErrorUserDoesNotBelongToWorkspace
	}

	repository, err := c.repository.GetRepository(data.RepositoryID)
	if err != nil {
		return nil, err
	}

	return c.createRepositoryRelationAndSendEmail(data, repository)
}

func (c *Controller) createRepositoryRelationAndSendEmail(data *roleEntities.UserData,
	repository *repositoryEntities.Repository) (*roleEntities.Response, error) {
	accountRepository := repository.ToAccountRepository(data.AccountID, data.Role)
	if err := c.databaseWrite.Create(accountRepository,
		repositoryEnums.DatabaseAccountRepositoryTable).GetError(); err != nil {
		return nil, err
	}

	return accountRepository.ToResponseWithEmailAndUsername(data.Email, data.Username),
		c.sendInviteUserEmail(data.Email, data.Username, repository.Name)
}

func (c *Controller) sendInviteUserEmail(email, username, repositoryName string) error {
	if c.appConfig.IsEmailsDisabled() {
		return nil
	}

	return c.broker.Publish(queues.HorusecEmail.ToString(), "", "",
		c.useCases.NewRepositoryInviteEmail(email, username, repositoryName))
}

func (c *Controller) GetUsers(repositoryID uuid.UUID) (*[]roleEntities.Response, error) {
	return c.repository.ListAllRepositoryUsers(repositoryID)
}

func (c *Controller) RemoveUser(data *roleEntities.Data) error {
	return c.databaseWrite.Delete(c.useCases.FilterAccountRepositoryByID(data.AccountID, data.RepositoryID),
		repositoryEnums.DatabaseAccountRepositoryTable).GetError()
}

func (c *Controller) CreateToken(data *tokenEntities.Data) (string, error) {
	token, tokenString := data.ToToken()

	return tokenString, c.databaseWrite.Create(token, tokenEnums.DatabaseTokens).GetError()
}

func (c *Controller) DeleteToken(data *tokenEntities.Data) error {
	return c.databaseWrite.Delete(c.tokenUseCases.FilterRepositoryTokenByID(
		data.TokenID, data.WorkspaceID, data.RepositoryID), tokenEnums.DatabaseTokens).GetError()
}

func (c *Controller) ListTokens(data *tokenEntities.Data) (*[]tokenEntities.Response, error) {
	tokens := &[]tokenEntities.Response{}

	return tokens, c.databaseRead.Find(tokens, c.tokenUseCases.FilterListRepositoryTokens(
		data.WorkspaceID, data.RepositoryID), tokenEnums.DatabaseTokens).GetErrorExceptNotFound()
}

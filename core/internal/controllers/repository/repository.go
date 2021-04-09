package repository

import (
	"github.com/google/uuid"

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	repositoryRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/repository"
	repositoriesUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
)

type IController interface {
	Create(data *repositoryEntities.Data) (*repositoryEntities.Response, error)
	Get(data *repositoryEntities.Data) (*repositoryEntities.Response, error)
	Update(data *repositoryEntities.Data) (*repositoryEntities.Response, error)
	Delete(repositoryID uuid.UUID) error
	List(data *repositoryEntities.Data) (*[]repositoryEntities.Response, error)
	UpdateRole(data *roleEntities.Data) (*roleEntities.Response, error)
}

type Controller struct {
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	appConfig     app.IConfig
	useCases      repositoriesUseCases.IUseCases
	repository    repositoryRepository.IRepository
}

func NewRepositoryController(databaseConnection *database.Connection, appConfig app.IConfig,
	useCases repositoriesUseCases.IUseCases, repository repositoryRepository.IRepository) IController {
	return &Controller{
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
		appConfig:     appConfig,
		useCases:      useCases,
		repository:    repository,
	}
}

func (c *Controller) Create(data *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	_, err := c.repository.GetRepositoryByName(data.WorkspaceID, data.Name)
	if c.useCases.IsNotFoundError(err) {
		return c.createTransaction(data)
	}

	return nil, repositoryEnums.ErrorRepositoryNameAlreadyInUse
}

func (c *Controller) createTransaction(data *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	transaction := c.databaseWrite.StartTransaction()
	repository := data.ToRepository()

	if err := transaction.Create(repository, repositoryEnums.DatabaseRepositoryTable).GetError(); err != nil {
		logger.LogError(repositoryEnums.ErrorRollbackCreate, transaction.RollbackTransaction().GetError())
		return nil, err
	}

	if err := transaction.Create(repository.ToAccountRepository(data.AccountID, accountEnums.Admin),
		repositoryEnums.DatabaseAccountRepositoryTable).GetError(); err != nil {
		logger.LogError(repositoryEnums.ErrorRollbackCreate, transaction.RollbackTransaction().GetError())
		return nil, err
	}

	return repository.ToRepositoryResponse(accountEnums.Admin), transaction.CommitTransaction().GetError()
}

func (c *Controller) Get(data *repositoryEntities.Data) (*repositoryEntities.Response, error) {
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

	repository.Update(data)
	return repository.ToRepositoryResponse(accountEnums.Admin), c.databaseWrite.Update(repository,
		c.useCases.FilterRepositoryByID(data.RepositoryID), repositoryEnums.DatabaseRepositoryTable).GetError()
}

func (c *Controller) Delete(repositoryID uuid.UUID) error {
	return c.databaseWrite.Delete(c.useCases.FilterRepositoryByID(repositoryID),
		repositoryEnums.DatabaseRepositoryTable).GetError()
}

func (c *Controller) List(data *repositoryEntities.Data) (*[]repositoryEntities.Response, error) {
	if c.appConfig.GetAuthorizationType() == auth.Ldap {
		return c.repository.ListRepositoriesAuthTypeLdap(data.WorkspaceID, data.Permissions)
	}

	return c.repository.ListRepositoriesAuthTypeHorusec(data.AccountID, data.WorkspaceID)
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

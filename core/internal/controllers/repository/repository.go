package repository

import (
	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	repositoryRepository "github.com/ZupIT/horusec-platform/core/internal/repositories/repository"
	repositoriesUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
)

type IController interface {
	Create(data *repositoryEntities.Data) (*repositoryEntities.Response, error)
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

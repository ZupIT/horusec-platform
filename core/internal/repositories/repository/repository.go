package repository

import (
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	repositoriesUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
)

type IRepository interface {
	GetRepositoryByName(workspaceID uuid.UUID, name string) (*repositoryEntities.Repository, error)
}

type Repository struct {
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	useCases      repositoriesUseCases.IUseCases
}

func NewRepositoryRepository(databaseConnection *database.Connection,
	useCases repositoriesUseCases.IUseCases) IRepository {
	return &Repository{
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
		useCases:      useCases,
	}
}

func (r *Repository) GetRepositoryByName(workspaceID uuid.UUID, name string) (*repositoryEntities.Repository, error) {
	repository := &repositoryEntities.Repository{}

	return repository, r.databaseRead.Find(repository, r.useCases.FilterRepositoryByName(workspaceID, name),
		repositoryEnums.DatabaseRepositoryTable).GetError()
}

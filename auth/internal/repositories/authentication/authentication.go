package authentication

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/google/uuid"

	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
	authEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/authentication"
	authUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/authentication"
)

type IRepository interface {
	GetWorkspaceGroups(workspaceID uuid.UUID) (*authEntities.AuthzGroups, error)
	GetRepositoryGroups(repositoryID uuid.UUID) (*authEntities.AuthzGroups, error)
}

type Repository struct {
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	useCases      authUseCases.IUseCases
}

func NewAuthenticationRepository(connection *database.Connection, useCases authUseCases.IUseCases) IRepository {
	return &Repository{
		databaseRead:  connection.Read,
		databaseWrite: connection.Write,
		useCases:      useCases,
	}
}

func (r *Repository) GetWorkspaceGroups(workspaceID uuid.UUID) (*authEntities.AuthzGroups, error) {
	authzGroups := &authEntities.AuthzGroups{}

	return authzGroups, r.databaseRead.Find(authzGroups, r.useCases.FilterWorkspaceByID(workspaceID),
		authEnums.TableWorkspaces).GetError()
}

func (r *Repository) GetRepositoryGroups(repositoryID uuid.UUID) (*authEntities.AuthzGroups, error) {
	authzGroups := &authEntities.AuthzGroups{}

	return authzGroups, r.databaseRead.Find(authzGroups, r.useCases.FilterRepositoryByID(repositoryID),
		authEnums.TableRepositories).GetError()
}

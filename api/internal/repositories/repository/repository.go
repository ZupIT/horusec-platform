package repository

import (
	"time"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
)

type IRepository interface {
	CreateRepository(ID, workspaceID uuid.UUID, name string) error
}

type Repository struct {
	databaseWrite       database.IDatabaseWrite
	databaseRead        database.IDatabaseRead
	repositoryTableName string
}

func NewRepositoriesRepository(connection *database.Connection) IRepository {
	return &Repository{
		databaseWrite:       connection.Write,
		databaseRead:        connection.Read,
		repositoryTableName: "repositories",
	}
}

func (r *Repository) CreateRepository(repositoryID, workspaceID uuid.UUID, name string) error {
	entity := map[string]interface{}{
		"repository_id":    repositoryID,
		"workspace_id":     workspaceID,
		"name":             name,
		"created_at":       time.Now(),
		"updated_at":       time.Now(),
		"authz_member":     "{}",
		"authz_admin":      "{}",
		"authz_supervisor": "{}",
	}
	return r.databaseWrite.Create(entity, r.repositoryTableName).GetError()
}

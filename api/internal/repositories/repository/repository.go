package repository

import (
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
)

type IRepository interface {
	CreateRepository(ID, workspaceID uuid.UUID, name string) error
	FindRepository(workspaceID uuid.UUID, name string) (uuid.UUID, error)
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

func (r *Repository) FindRepository(workspaceID uuid.UUID, name string) (uuid.UUID, error) {
	entity := map[string]interface{}{}
	condition := map[string]interface{}{
		"workspace_id": workspaceID,
		"name":         name,
	}
	res := r.databaseRead.Find(&entity, condition, r.repositoryTableName)
	if res.GetError() != nil {
		return uuid.Nil, res.GetError()
	}
	if res.GetData() == nil {
		return uuid.Nil, enums.ErrorNotFoundRecords
	}
	entity = *res.GetData().(*map[string]interface{})
	repositoryID := entity["repository_id"].(string)
	return uuid.Parse(repositoryID)
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

package webhook

import "github.com/google/uuid"

type Repository struct {
	RepositoryID uuid.UUID `json:"repositoryID" gorm:"Column:repository_id"`
	Name         string    `json:"name"`
}

package webhook

import "github.com/google/uuid"

type Repository struct {
	RepositoryID uuid.UUID `json:"repositoryID"`
	Name string `json:"name"`
}
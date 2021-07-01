package core

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Repository struct {
	RepositoryID    uuid.UUID      `json:"repositoryID" gorm:"primary_key"`
	WorkspaceID     uuid.UUID      `json:"workspaceID"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	AuthzMember     pq.StringArray `json:"authzMember" gorm:"type:text[]"`
	AuthzAdmin      pq.StringArray `json:"authzAdmin" gorm:"type:text[]"`
	AuthzSupervisor pq.StringArray `json:"authzSupervisor" gorm:"type:text[]"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
}

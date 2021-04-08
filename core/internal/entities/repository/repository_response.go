package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

type Response struct {
	WorkspaceID     uuid.UUID      `json:"workspaceID"`
	RepositoryID    uuid.UUID      `json:"repositoryID"`
	Name            string         `json:"name"`
	Role            account.Role   `json:"role"`
	Description     string         `json:"description"`
	AuthzMember     pq.StringArray `json:"authzMember" gorm:"type:text[]"`
	AuthzAdmin      pq.StringArray `json:"authzAdmin" gorm:"type:text[]"`
	AuthzSupervisor pq.StringArray `json:"authzSupervisor" gorm:"type:text[]"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
}

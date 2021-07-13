package core

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Workspace struct {
	WorkspaceID uuid.UUID      `json:"workspaceID" gorm:"primary_key"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	AuthzMember pq.StringArray `json:"authzMember" gorm:"type:text[]"`
	AuthzAdmin  pq.StringArray `json:"authzAdmin" gorm:"type:text[]"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

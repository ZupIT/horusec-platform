package token

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	TokenID        uuid.UUID `gorm:"Column:token_id"`
	RepositoryID   uuid.UUID `gorm:"Column:repository_id"`
	RepositoryName string    `gorm:"Column:repository_name"`
	WorkspaceID    uuid.UUID `gorm:"Column:workspace_id"`
	WorkspaceName  string    `gorm:"Column:workspace_name"`
	ExpiresAt      time.Time `gorm:"Column:expires_at"`
	IsExpirable    bool      `gorm:"Column:is_expirable"`
}

package token

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	TokenID      uuid.UUID `json:"tokenID"`
	WorkspaceID  uuid.UUID `json:"workspaceID"`
	RepositoryID string `json:"repositoryID"`
	Description  string    `json:"description"`
	SuffixValue  string    `json:"suffixValue"`
	IsExpirable  bool      `json:"isExpirable"`
	CreatedAt    time.Time `json:"createdAt"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

package token

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	TokenID      uuid.UUID  `json:"tokenID"`
	WorkspaceID  uuid.UUID  `json:"workspaceID"`
	RepositoryID *uuid.UUID `json:"repositoryID"`
	Description  string     `json:"description"`
	SuffixValue  string     `json:"suffixValue"`
	Value        string     `json:"value"`
	IsExpirable  bool       `json:"isExpirable"`
	CreatedAt    time.Time  `json:"createdAt"`
	ExpiresAt    time.Time  `json:"expiresAt"`
}

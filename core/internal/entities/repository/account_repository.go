package repository

import (
	"time"

	"github.com/google/uuid"

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

type AccountRepository struct {
	RepositoryID uuid.UUID         `json:"repositoryID"`
	AccountID    uuid.UUID         `json:"accountID"`
	WorkspaceID  uuid.UUID         `json:"workspaceID"`
	Role         accountEnums.Role `json:"role"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
}

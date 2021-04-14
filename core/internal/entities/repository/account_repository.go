package repository

import (
	"time"

	"github.com/google/uuid"

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"

	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
)

type AccountRepository struct {
	RepositoryID uuid.UUID         `json:"repositoryID"`
	AccountID    uuid.UUID         `json:"accountID"`
	WorkspaceID  uuid.UUID         `json:"workspaceID"`
	Role         accountEnums.Role `json:"role"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
}

func (a *AccountRepository) Update(role accountEnums.Role) {
	a.Role = role
	a.UpdatedAt = time.Now()
}

func (a *AccountRepository) ToResponse() *roleEntities.Response {
	return &roleEntities.Response{
		AccountID: a.AccountID,
		Role:      a.Role,
	}
}

func (a *AccountRepository) ToResponseWithEmailAndUsername(email, username string) *roleEntities.Response {
	return &roleEntities.Response{
		AccountID: a.AccountID,
		Email:     email,
		Username:  username,
		Role:      a.Role,
	}
}

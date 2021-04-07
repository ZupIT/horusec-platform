package workspace

import (
	"time"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"

	"github.com/ZupIT/horusec-platform/core/internal/entities/role"
)

type AccountWorkspace struct {
	WorkspaceID uuid.UUID    `json:"workspaceID"`
	AccountID   uuid.UUID    `json:"accountID"`
	Role        account.Role `json:"role"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

func (a *AccountWorkspace) Update(data *role.Data) {
	a.Role = data.Role
	a.UpdatedAt = time.Now()
}

func (a *AccountWorkspace) ToResponse() *role.Response {
	return &role.Response{
		AccountID: a.AccountID,
		Role:      a.Role,
	}
}

func (a *AccountWorkspace) ToResponseWithEmailAndUsername(email, username string) *role.Response {
	return &role.Response{
		AccountID: a.AccountID,
		Email:     email,
		Username:  username,
		Role:      a.Role,
	}
}

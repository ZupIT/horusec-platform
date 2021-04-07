package workspace

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
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

func (w *Workspace) ToAccountWorkspace(accountID uuid.UUID, role account.Role) *AccountWorkspace {
	return &AccountWorkspace{
		WorkspaceID: w.WorkspaceID,
		AccountID:   accountID,
		Role:        role,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (w *Workspace) ToWorkspaceResponse(role account.Role) *Response {
	return &Response{
		WorkspaceID: w.WorkspaceID,
		Name:        w.Name,
		Role:        role,
		Description: w.Description,
		AuthzAdmin:  w.AuthzAdmin,
		AuthzMember: w.AuthzMember,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
	}
}

func (w *Workspace) Update(data *Data) {
	w.Name = data.Name
	w.Description = data.Description
	w.AuthzMember = data.AuthzMember
	w.AuthzAdmin = data.AuthzAdmin
	w.UpdatedAt = time.Now()
}

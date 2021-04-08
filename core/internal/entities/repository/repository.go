package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
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

func (r *Repository) ToAccountRepository(accountID uuid.UUID, role account.Role) *AccountRepository {
	return &AccountRepository{
		RepositoryID: r.RepositoryID,
		AccountID:    accountID,
		WorkspaceID:  r.WorkspaceID,
		Role:         role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (r *Repository) ToRepositoryResponse(role account.Role) *Response {
	return &Response{
		WorkspaceID:     r.WorkspaceID,
		RepositoryID:    r.RepositoryID,
		Name:            r.Name,
		Role:            role,
		Description:     r.Description,
		AuthzMember:     r.AuthzMember,
		AuthzAdmin:      r.AuthzAdmin,
		AuthzSupervisor: r.AuthzSupervisor,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
	}
}

func (r *Repository) Update(data *Data) {
	r.Name = data.Name
	r.Description = data.Description
	r.AuthzMember = data.AuthzMember
	r.AuthzSupervisor = data.AuthzSupervisor
	r.AuthzAdmin = data.AuthzAdmin
	r.UpdatedAt = time.Now()
}

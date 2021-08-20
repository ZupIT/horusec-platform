// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func (r *Repository) Update(data *Data) *Repository {
	r.Name = data.Name
	r.Description = data.Description
	r.AuthzMember = data.AuthzMember
	r.AuthzSupervisor = data.AuthzSupervisor
	r.AuthzAdmin = data.AuthzAdmin
	r.UpdatedAt = time.Now()
	return r
}

func (r *Repository) ContainsAllAuthzGroups() bool {
	if len(r.AuthzAdmin) == 0 || len(r.AuthzMember) == 0 || len(r.AuthzSupervisor) == 0 {
		return false
	}

	return true
}

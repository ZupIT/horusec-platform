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

func (w *Workspace) Update(data *Data) *Workspace {
	w.Name = data.Name
	w.Description = data.Description
	w.AuthzMember = data.AuthzMember
	w.AuthzAdmin = data.AuthzAdmin
	w.UpdatedAt = time.Now()
	return w
}

func (w *Workspace) ToUpdateMap(data *Data) map[string]interface{} {
	return map[string]interface{}{
		"name":         data.Name,
		"description":  data.Description,
		"authz_member": pq.Array(data.AuthzMember),
		"authz_admin":  pq.Array(data.AuthzAdmin),
		"updated_at":   time.Now(),
	}
}

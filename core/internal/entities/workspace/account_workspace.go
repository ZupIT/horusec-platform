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

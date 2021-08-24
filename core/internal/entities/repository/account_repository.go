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

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"

	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
)

//TODO add unity tests
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

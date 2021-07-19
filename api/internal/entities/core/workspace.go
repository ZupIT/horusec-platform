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

package core

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
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

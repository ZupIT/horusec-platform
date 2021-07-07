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

package token

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	TokenID        uuid.UUID `gorm:"Column:token_id"`
	RepositoryID   uuid.UUID `gorm:"Column:repository_id"`
	RepositoryName string    `gorm:"Column:repository_name"`
	WorkspaceID    uuid.UUID `gorm:"Column:workspace_id"`
	WorkspaceName  string    `gorm:"Column:workspace_name"`
	ExpiresAt      time.Time `gorm:"Column:expires_at"`
	IsExpirable    bool      `gorm:"Column:is_expirable"`
}

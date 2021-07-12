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
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"

	"github.com/ZupIT/horusec-platform/core/internal/entities/role"
)

func TestUpdateAccountWorkspace(t *testing.T) {
	t.Run("should success update account workspace data", func(t *testing.T) {
		expectedTime := time.Now()

		accountWorkspace := &AccountWorkspace{
			WorkspaceID: uuid.New(),
			AccountID:   uuid.New(),
			Role:        account.Admin,
			CreatedAt:   expectedTime,
			UpdatedAt:   expectedTime,
		}

		accountWorkspace.Update(&role.Data{Role: account.Member})
		assert.Equal(t, account.Member, accountWorkspace.Role)
		assert.NotEqual(t, expectedTime, accountWorkspace.UpdatedAt)
	})
}

func TestToResponse(t *testing.T) {
	t.Run("should parse to response", func(t *testing.T) {
		expectedTime := time.Now()

		accountWorkspace := &AccountWorkspace{
			WorkspaceID: uuid.New(),
			AccountID:   uuid.New(),
			Role:        account.Admin,
			CreatedAt:   expectedTime,
			UpdatedAt:   expectedTime,
		}

		response := accountWorkspace.ToResponse()
		assert.Equal(t, accountWorkspace.Role, response.Role)
		assert.Equal(t, accountWorkspace.AccountID, response.AccountID)
	})
}

func TestToResponseWithEmailAndUsername(t *testing.T) {
	t.Run("should parse to response with email and username", func(t *testing.T) {
		expectedTime := time.Now()

		accountWorkspace := &AccountWorkspace{
			WorkspaceID: uuid.New(),
			AccountID:   uuid.New(),
			Role:        account.Admin,
			CreatedAt:   expectedTime,
			UpdatedAt:   expectedTime,
		}

		response := accountWorkspace.ToResponseWithEmailAndUsername("test@test.com", "test")
		assert.Equal(t, accountWorkspace.Role, response.Role)
		assert.Equal(t, accountWorkspace.AccountID, response.AccountID)
		assert.Equal(t, "test@test.com", response.Email)
		assert.Equal(t, "test", response.Username)
	})
}

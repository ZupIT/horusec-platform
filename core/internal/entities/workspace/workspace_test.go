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
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

func TestToAccountWorkspace(t *testing.T) {
	t.Run("should success parse to account workspace", func(t *testing.T) {
		accountID := uuid.New()

		workspace := &Workspace{
			WorkspaceID: uuid.New(),
			Name:        "test",
			Description: "test",
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		accountWorkspace := workspace.ToAccountWorkspace(accountID, account.Admin)
		assert.Equal(t, workspace.WorkspaceID, accountWorkspace.WorkspaceID)
		assert.Equal(t, accountID, accountWorkspace.AccountID)
		assert.Equal(t, account.Admin, accountWorkspace.Role)
		assert.NotEqual(t, time.Time{}, accountWorkspace.CreatedAt)
		assert.NotEqual(t, time.Time{}, accountWorkspace.UpdatedAt)
	})
}

func TestToWorkspaceResponse(t *testing.T) {
	t.Run("should success parse to workspace response", func(t *testing.T) {
		workspace := &Workspace{
			WorkspaceID: uuid.New(),
			Name:        "test",
			Description: "test",
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		response := workspace.ToWorkspaceResponse(account.Admin)
		assert.Equal(t, workspace.WorkspaceID, response.WorkspaceID)
		assert.Equal(t, workspace.Name, response.Name)
		assert.Equal(t, workspace.Description, response.Description)
		assert.Equal(t, workspace.AuthzMember, response.AuthzMember)
		assert.Equal(t, workspace.AuthzAdmin, response.AuthzAdmin)
		assert.Equal(t, workspace.CreatedAt, response.CreatedAt)
		assert.Equal(t, workspace.UpdatedAt, response.UpdatedAt)
		assert.Equal(t, account.Admin, response.Role)
	})
}

func TestUpdateWorkspace(t *testing.T) {
	t.Run("should success update workspace data", func(t *testing.T) {
		expectedTime := time.Now()

		workspace := &Workspace{
			WorkspaceID: uuid.New(),
			Name:        "test",
			Description: "test",
			AuthzMember: []string{""},
			AuthzAdmin:  []string{""},
			CreatedAt:   expectedTime,
			UpdatedAt:   expectedTime,
		}

		data := &Data{
			Name:        "test2",
			Description: "test2",
			AuthzMember: []string{"test2"},
			AuthzAdmin:  []string{"test2"},
		}

		workspace.Update(data)
		assert.Equal(t, data.Name, workspace.Name)
		assert.Equal(t, data.Description, workspace.Description)
		assert.Equal(t, pq.StringArray(data.AuthzAdmin), workspace.AuthzAdmin)
		assert.Equal(t, pq.StringArray(data.AuthzMember), workspace.AuthzMember)
		assert.NotEqual(t, expectedTime, workspace.UpdatedAt)
	})
}

func TestToUpdateMap(t *testing.T) {
	t.Run("should success parse to update map", func(t *testing.T) {
		workspace := &Workspace{
			WorkspaceID: uuid.New(),
			Name:        "test1",
			Description: "test1",
			AuthzMember: []string{"test1"},
			AuthzAdmin:  []string{"test1"},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		data := &Data{
			Name:        "test2",
			Description: "test2",
			AuthzMember: []string{"test2"},
			AuthzAdmin:  []string{"test2"},
		}

		assert.NotPanics(t, func() {
			result := *workspace.ToUpdateMap(data)
			assert.Equal(t, "test2", result["name"])
			assert.Equal(t, "test2", result["description"])
			assert.Equal(t, pq.Array([]string{"test2"}), result["authz_member"])
			assert.Equal(t, pq.Array([]string{"test2"}), result["authz_admin"])
			assert.NotEqual(t, workspace.UpdatedAt, result["updated_at"])
		})
	})
}

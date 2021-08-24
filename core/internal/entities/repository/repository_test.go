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
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

func TestToAccountRepository(t *testing.T) {
	t.Run("should success parse to account repository", func(t *testing.T) {
		repository := &Repository{
			RepositoryID:    uuid.New(),
			WorkspaceID:     uuid.New(),
			Name:            "test",
			Description:     "test",
			AuthzMember:     []string{"test"},
			AuthzAdmin:      []string{"test"},
			AuthzSupervisor: []string{"test"},
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		id := uuid.New()
		accountRepository := repository.ToAccountRepository(id, account.Member)
		assert.Equal(t, repository.RepositoryID, accountRepository.RepositoryID)
		assert.Equal(t, repository.WorkspaceID, accountRepository.WorkspaceID)
		assert.Equal(t, id, accountRepository.AccountID)
		assert.NotEqual(t, time.Time{}, accountRepository.CreatedAt)
		assert.NotEqual(t, time.Time{}, accountRepository.UpdatedAt)
	})
}

func TestToRepositoryResponse(t *testing.T) {
	t.Run("should success parse to account repository", func(t *testing.T) {
		repository := &Repository{
			RepositoryID:    uuid.New(),
			WorkspaceID:     uuid.New(),
			Name:            "test",
			Description:     "test",
			AuthzMember:     []string{"test"},
			AuthzAdmin:      []string{"test"},
			AuthzSupervisor: []string{"test"},
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		response := repository.ToRepositoryResponse(account.Member)
		assert.Equal(t, repository.CreatedAt, response.CreatedAt)
		assert.Equal(t, repository.UpdatedAt, response.UpdatedAt)
		assert.Equal(t, repository.WorkspaceID, response.WorkspaceID)
		assert.Equal(t, repository.RepositoryID, response.RepositoryID)
		assert.Equal(t, repository.AuthzSupervisor, response.AuthzSupervisor)
		assert.Equal(t, repository.AuthzAdmin, response.AuthzAdmin)
		assert.Equal(t, repository.AuthzMember, response.AuthzMember)
		assert.Equal(t, repository.Name, response.Name)
		assert.Equal(t, account.Member, response.Role)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("should success update repository data", func(t *testing.T) {
		expectedTime := time.Now()

		repository := &Repository{
			UpdatedAt: expectedTime,
		}

		data := &Data{
			Name:            "test",
			Description:     "test",
			AuthzMember:     []string{"test"},
			AuthzSupervisor: []string{"test"},
			AuthzAdmin:      []string{"test"},
		}

		repository.Update(data)
		assert.Equal(t, data.Name, repository.Name)
		assert.Equal(t, data.Description, repository.Description)
		assert.Equal(t, pq.StringArray(data.AuthzMember), repository.AuthzMember)
		assert.Equal(t, pq.StringArray(data.AuthzSupervisor), repository.AuthzSupervisor)
		assert.Equal(t, pq.StringArray(data.AuthzAdmin), repository.AuthzAdmin)
		assert.NotEqual(t, expectedTime, repository.UpdatedAt)
	})
}

func TestContainsAllAuthzGroups(t *testing.T) {
	t.Run("should return true when repository contains all groups", func(t *testing.T) {
		repository := &Repository{
			AuthzMember:     []string{"test"},
			AuthzAdmin:      []string{"test"},
			AuthzSupervisor: []string{"test"},
		}

		assert.True(t, repository.ContainsAllAuthzGroups())
	})

	t.Run("should return false when repository do not contains authz member", func(t *testing.T) {
		repository := &Repository{
			AuthzAdmin:      []string{"test"},
			AuthzSupervisor: []string{"test"},
		}

		assert.False(t, repository.ContainsAllAuthzGroups())
	})

	t.Run("should return false when repository do not contains authz admin", func(t *testing.T) {
		repository := &Repository{
			AuthzMember:     []string{"test"},
			AuthzSupervisor: []string{"test"},
		}

		assert.False(t, repository.ContainsAllAuthzGroups())
	})

	t.Run("should return false when repository do not contains authz supervisor", func(t *testing.T) {
		repository := &Repository{
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
		}

		assert.False(t, repository.ContainsAllAuthzGroups())
	})
}

func TestToUpdateMap(t *testing.T) {
	t.Run("should success parse to update map", func(t *testing.T) {
		repository := &Repository{
			RepositoryID:    uuid.New(),
			WorkspaceID:     uuid.New(),
			Name:            "test1",
			Description:     "test1",
			AuthzMember:     []string{"test1"},
			AuthzAdmin:      []string{"test1"},
			AuthzSupervisor: []string{"test1"},
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		data := &Data{
			AccountID:       uuid.Nil,
			WorkspaceID:     uuid.Nil,
			Name:            "test2",
			Description:     "test2",
			AuthzMember:     []string{"test2"},
			AuthzAdmin:      []string{"test2"},
			AuthzSupervisor: []string{"test2"},
			Permissions:     nil,
		}

		assert.NotPanics(t, func() {
			result := *repository.ToUpdateMap(data)
			assert.Equal(t, "test2", result["name"])
			assert.Equal(t, "test2", result["description"])
			assert.Equal(t, pq.Array([]string{"test2"}), result["authz_member"])
			assert.Equal(t, pq.Array([]string{"test2"}), result["authz_supervisor"])
			assert.Equal(t, pq.Array([]string{"test2"}), result["authz_admin"])
			assert.NotEqual(t, repository.UpdatedAt, result["updated_at"])
		})
	})
}

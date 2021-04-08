package repository

import (
	"testing"
	"time"

	"github.com/google/uuid"
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

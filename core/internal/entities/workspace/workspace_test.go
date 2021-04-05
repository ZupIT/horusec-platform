package workspace

import (
	"testing"
	"time"

	"github.com/google/uuid"
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
		assert.NotEqual(t, &time.Time{}, accountWorkspace.CreatedAt)
		assert.Equal(t, time.Time{}, accountWorkspace.UpdatedAt)
	})
}

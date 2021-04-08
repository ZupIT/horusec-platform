package role

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

func TestValidateRoleData(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &Data{}

		data.Role = account.Admin
		assert.NoError(t, data.Validate())

		data.Role = account.Supervisor
		assert.NoError(t, data.Validate())

		data.Role = account.Member
		assert.NoError(t, data.Validate())
	})

	t.Run("should return when invalid role value", func(t *testing.T) {
		data := &Data{
			Role: "test",
		}

		assert.Error(t, data.Validate())
	})
}

func TestSetAccountAndWorkspaceID(t *testing.T) {
	t.Run("should success set account and workspace id", func(t *testing.T) {
		data := &Data{}
		id := uuid.New()

		_ = data.SetAccountAndWorkspaceID(id, id)
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, id, data.AccountID)
	})
}

func TestToBytesRoleData(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		data := Data{}

		assert.NotEmpty(t, data.ToBytes())
	})
}

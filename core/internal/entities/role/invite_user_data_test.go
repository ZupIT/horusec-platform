package role

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
)

func TestValidateInviteUserData(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := InviteUserData{
			Role:      account.Admin,
			Email:     "test@test.com",
			AccountID: uuid.New(),
			Username:  "test",
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when invalid role", func(t *testing.T) {
		data := InviteUserData{
			Role:      "test",
			Email:     "test@test.com",
			AccountID: uuid.New(),
			Username:  "test",
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return no error when invalid email", func(t *testing.T) {
		data := InviteUserData{
			Role:      account.Admin,
			Email:     "test",
			AccountID: uuid.New(),
			Username:  "test",
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return no error when missing username", func(t *testing.T) {
		data := InviteUserData{
			Role:      account.Admin,
			Email:     "test@test.com",
			AccountID: uuid.New(),
			Username:  "",
		}

		assert.Error(t, data.Validate())
	})
}

func TestSetWorkspaceID(t *testing.T) {
	t.Run("should success set workspace id", func(t *testing.T) {
		data := InviteUserData{}
		id := uuid.New()

		data.SetWorkspaceID(id)
		assert.Equal(t, id, data.WorkspaceID)
	})
}

func TestToBytesInviteUserData(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		data := InviteUserData{}

		assert.NotEmpty(t, data.ToBytes())
	})
}

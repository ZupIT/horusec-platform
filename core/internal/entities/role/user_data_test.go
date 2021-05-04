package role

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
)

func TestValidateInviteUserData(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := UserData{
			Role:      account.Admin,
			Email:     "test@test.com",
			AccountID: uuid.New(),
			Username:  "test",
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when invalid role", func(t *testing.T) {
		data := UserData{
			Role:      "test",
			Email:     "test@test.com",
			AccountID: uuid.New(),
			Username:  "test",
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return no error when invalid email", func(t *testing.T) {
		data := UserData{
			Role:      account.Admin,
			Email:     "test",
			AccountID: uuid.New(),
			Username:  "test",
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return no error when missing username", func(t *testing.T) {
		data := UserData{
			Role:      account.Admin,
			Email:     "test@test.com",
			AccountID: uuid.New(),
			Username:  "",
		}

		assert.NoError(t, data.Validate())
	})
}

func TestSetWorkspaceID(t *testing.T) {
	t.Run("should success set workspace id", func(t *testing.T) {
		data := UserData{}
		id := uuid.New()

		data.SetIDs(id.String(), id.String())
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, id, data.RepositoryID)
	})
}

func TestToBytesInviteUserData(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		data := UserData{}

		assert.NotEmpty(t, data.ToBytes())
	})
}

func TestSetWorkspaceIDAndAccountData(t *testing.T) {
	t.Run("should success set workspace id and account data", func(t *testing.T) {
		data := UserData{}

		id := uuid.New()
		accountData := &proto.GetAccountDataResponse{
			AccountID:          id.String(),
			IsApplicationAdmin: false,
			Permissions:        nil,
			Email:              "test@test.com",
			Username:           "test",
		}

		_ = data.SetWorkspaceIDAndAccountData(id.String(), accountData)
		assert.Equal(t, id, data.AccountID)
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, "test@test.com", data.Email)
		assert.Equal(t, "test", data.Username)
	})
}

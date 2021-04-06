package workspace

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
)

const (
	MaxCharacters255 = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
)

func TestValidate(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		workspaceData := &Data{
			AccountID:   uuid.Nil,
			Name:        "test",
			Description: "test",
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
			Permissions: nil,
		}

		assert.NoError(t, workspaceData.Validate())
	})

	t.Run("should return error when name is bigger than 255", func(t *testing.T) {
		workspaceData := &Data{
			AccountID:   uuid.Nil,
			Name:        MaxCharacters255,
			Description: "test",
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
			Permissions: nil,
		}

		assert.Error(t, workspaceData.Validate())
	})

	t.Run("should return error when description is bigger than 255", func(t *testing.T) {
		workspaceData := &Data{
			AccountID:   uuid.Nil,
			Name:        "test",
			Description: MaxCharacters255,
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
			Permissions: nil,
		}

		assert.Error(t, workspaceData.Validate())
	})

	t.Run("should return error when more than 5 authz member permissions", func(t *testing.T) {
		workspaceData := &Data{
			AccountID:   uuid.Nil,
			Name:        "test",
			Description: "test",
			AuthzMember: []string{"test", "test", "test", "test", "test", "test"},
			AuthzAdmin:  []string{"test"},
			Permissions: nil,
		}

		assert.Error(t, workspaceData.Validate())
	})

	t.Run("should return error when more than 5 authz admin permissions", func(t *testing.T) {
		workspaceData := &Data{
			AccountID:   uuid.Nil,
			Name:        "test",
			Description: "test",
			AuthzAdmin:  []string{"test", "test", "test", "test", "test", "test"},
			AuthzMember: []string{"test"},
			Permissions: nil,
		}

		assert.Error(t, workspaceData.Validate())
	})

	t.Run("should return error when permissions are different than empty", func(t *testing.T) {
		workspaceData := &Data{
			AccountID:   uuid.Nil,
			Name:        "test",
			Description: "test",
			AuthzAdmin:  []string{"test"},
			AuthzMember: []string{"test"},
			Permissions: []string{"test"},
		}

		assert.Error(t, workspaceData.Validate())
	})
}

func TestToWorkspace(t *testing.T) {
	t.Run("should success parse to workspace", func(t *testing.T) {
		workspaceData := &Data{
			AccountID:   uuid.New(),
			Name:        "test",
			Description: "test",
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
			Permissions: []string{"test"},
		}

		workspace := workspaceData.ToWorkspace()
		assert.NotNil(t, workspace)
		assert.NotNil(t, workspace.WorkspaceID)
		assert.Equal(t, "test", workspace.Name)
		assert.Equal(t, "test", workspace.Description)
		assert.Len(t, workspace.AuthzAdmin, 1)
		assert.Len(t, workspace.AuthzMember, 1)
		assert.NotEqual(t, &time.Time{}, &workspace.CreatedAt)
		assert.Equal(t, &time.Time{}, &workspace.UpdatedAt)
	})
}

func TestSetAccountData(t *testing.T) {
	t.Run("should success set account data", func(t *testing.T) {
		workspaceData := &Data{}
		accountData := &proto.GetAccountDataResponse{
			AccountID:   uuid.New().String(),
			Permissions: []string{"test"},
		}

		workspaceData.SetAccountData(accountData)
		assert.Equal(t, uuid.MustParse(accountData.AccountID), workspaceData.AccountID)
		assert.Equal(t, accountData.Permissions, workspaceData.Permissions)
	})
}

func TestToBytes(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		workspaceData := &Data{AccountID: uuid.New()}

		assert.NotEmpty(t, workspaceData.ToBytes())
	})
}

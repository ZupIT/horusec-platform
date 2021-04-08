package repository

import (
	"testing"

	"github.com/lib/pq"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
)

const (
	MaxCharacters255 = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
)

func TestValidate(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &Data{
			AccountID:       uuid.Nil,
			WorkspaceID:     uuid.Nil,
			Name:            "test",
			Description:     "test",
			AuthzMember:     nil,
			AuthzAdmin:      nil,
			AuthzSupervisor: nil,
			Permissions:     nil,
		}

		assert.NoError(t, data.Validate())
	})

	t.Run("should return error when name is bigger than 255", func(t *testing.T) {
		data := &Data{
			AccountID:       uuid.Nil,
			WorkspaceID:     uuid.Nil,
			Name:            MaxCharacters255,
			Description:     "test",
			AuthzMember:     []string{"test"},
			AuthzAdmin:      []string{"test"},
			AuthzSupervisor: []string{"test"},
			Permissions:     nil,
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return error when description is bigger than 255", func(t *testing.T) {
		data := &Data{
			AccountID:       uuid.Nil,
			WorkspaceID:     uuid.Nil,
			Name:            "test",
			Description:     MaxCharacters255,
			AuthzMember:     []string{"test"},
			AuthzAdmin:      []string{"test"},
			AuthzSupervisor: []string{"test"},
			Permissions:     nil,
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return error when more than 5 authz member permissions", func(t *testing.T) {
		data := &Data{
			AccountID:       uuid.Nil,
			Name:            "test",
			Description:     "test",
			AuthzMember:     []string{"test", "test", "test", "test", "test", "test"},
			AuthzAdmin:      []string{"test"},
			AuthzSupervisor: []string{"test"},
			Permissions:     nil,
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return error when more than 5 authz admin permissions", func(t *testing.T) {
		data := &Data{
			AccountID:       uuid.Nil,
			WorkspaceID:     uuid.Nil,
			Name:            "test",
			Description:     "test",
			AuthzAdmin:      []string{"test", "test", "test", "test", "test", "test"},
			AuthzMember:     []string{"test"},
			AuthzSupervisor: []string{"test"},
			Permissions:     nil,
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return error when more than 5 authz supervisor permissions", func(t *testing.T) {
		data := &Data{
			AccountID:       uuid.Nil,
			WorkspaceID:     uuid.Nil,
			Name:            "test",
			Description:     "test",
			AuthzSupervisor: []string{"test", "test", "test", "test", "test", "test"},
			AuthzAdmin:      []string{"test"},
			AuthzMember:     []string{"test"},
			Permissions:     nil,
		}

		assert.Error(t, data.Validate())
	})

	t.Run("should return error when permissions are different than empty", func(t *testing.T) {
		data := &Data{
			AccountID:       uuid.Nil,
			WorkspaceID:     uuid.Nil,
			Name:            "test",
			Description:     "test",
			AuthzAdmin:      []string{"test"},
			AuthzMember:     []string{"test"},
			AuthzSupervisor: []string{"test"},
			Permissions:     []string{"test"},
		}

		assert.Error(t, data.Validate())
	})
}

func TestCheckLdapGroups(t *testing.T) {
	t.Run("should should return no error when valid groups", func(t *testing.T) {
		data := &Data{AuthzAdmin: []string{"test"}, Permissions: []string{"test"}}

		assert.NoError(t, data.CheckLdapGroups(auth.Ldap))
	})

	t.Run("should should return error when invalid groups", func(t *testing.T) {
		data := &Data{AuthzAdmin: []string{""}, Permissions: []string{""}}

		assert.Error(t, data.CheckLdapGroups(auth.Ldap))
	})
}

func TestToBytes(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		data := &Data{AccountID: uuid.New()}

		assert.NotEmpty(t, data.ToBytes())
	})
}

func TestSetWorkspaceIDAndAccountData(t *testing.T) {
	t.Run("should success set account data and workspace id", func(t *testing.T) {
		data := &Data{}
		id := uuid.New()
		accountData := &proto.GetAccountDataResponse{Permissions: []string{"test"}, AccountID: id.String()}

		_ = data.SetWorkspaceIDAndAccountData(id, accountData)
		assert.Equal(t, id, data.AccountID)
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, []string{"test"}, data.Permissions)
	})
}

func TestToRepository(t *testing.T) {
	t.Run("should success parse to repository", func(t *testing.T) {
		data := &Data{
			WorkspaceID:     uuid.New(),
			AccountID:       uuid.New(),
			Name:            "test",
			Description:     "test",
			AuthzMember:     []string{"test"},
			AuthzAdmin:      []string{"test"},
			AuthzSupervisor: []string{"test"},
		}

		repository := data.ToRepository()
		assert.Equal(t, data.WorkspaceID, repository.WorkspaceID)
		assert.Equal(t, data.Name, repository.Name)
		assert.Equal(t, data.Description, repository.Description)
		assert.Equal(t, pq.StringArray(data.AuthzMember), repository.AuthzMember)
		assert.Equal(t, pq.StringArray(data.AuthzAdmin), repository.AuthzAdmin)
		assert.Equal(t, pq.StringArray(data.AuthzSupervisor), repository.AuthzSupervisor)
	})
}

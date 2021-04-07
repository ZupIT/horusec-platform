package workspace

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	"github.com/ZupIT/horusec-devkit/pkg/enums/account"
	emailEnums "github.com/ZupIT/horusec-devkit/pkg/enums/email"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/role"
	"github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

func TestNewWorkspaceUseCases(t *testing.T) {
	t.Run("should success create a new use cases", func(t *testing.T) {
		assert.NotNil(t, NewWorkspaceUseCases())
	})
}

func TestWorkspaceDataFromIOReadCloser(t *testing.T) {
	t.Run("should success get workspace data from request body", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()

		data := &workspace.Data{
			AccountID: uuid.New(),
			Name:      "test",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		response, err := useCases.WorkspaceDataFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, data.AccountID, response.AccountID)
		assert.Equal(t, data.Name, response.Name)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		response, err := useCases.WorkspaceDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestFilterAccountWorkspaceByID(t *testing.T) {
	t.Run("should success create a account workspace filter by id", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()
		id := uuid.New()

		filter := useCases.FilterAccountWorkspaceByID(id, id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
			assert.Equal(t, id, filter["account_id"])
		})
	})
}

func TestFilterWorkspaceByID(t *testing.T) {
	t.Run("should success create a account workspace filter by id", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()
		id := uuid.New()

		filter := useCases.FilterWorkspaceByID(id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
		})
	})
}

func TestNewWorkspaceData(t *testing.T) {
	t.Run("should success create a new workspace data", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()
		id := uuid.New()

		accountData := &proto.GetAccountDataResponse{
			AccountID:   id.String(),
			Permissions: []string{"test"},
		}

		data := useCases.NewWorkspaceData(id, accountData)

		assert.Equal(t, id, data.AccountID)
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, []string{"test"}, data.Permissions)
	})
}

func TestRoleDataFromIOReadCloser(t *testing.T) {
	t.Run("should success get role data from request body", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()
		id := uuid.New()

		data := &role.Data{
			AccountID:    id,
			WorkspaceID:  id,
			RepositoryID: id,
			Role:         "admin",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		response, err := useCases.RoleDataFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, data.AccountID, response.AccountID)
		assert.Equal(t, data.WorkspaceID, response.WorkspaceID)
		assert.Equal(t, data.RepositoryID, response.RepositoryID)
		assert.Equal(t, data.Role, response.Role)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		response, err := useCases.RoleDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestNewOrganizationInviteEmail(t *testing.T) {
	t.Run("should success create a new organization invite email", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()

		emailBytes := useCases.NewOrganizationInviteEmail("test@test.com", "test", "test")
		assert.NotNil(t, emailBytes)
		assert.NotEmpty(t, emailBytes)

		email := &emailEntities.Message{}
		assert.NoError(t, json.Unmarshal(emailBytes, email))

		assert.Equal(t, "test@test.com", email.To)
		assert.Equal(t, emailEnums.OrganizationInvite, email.TemplateName)
		assert.Equal(t, "[Horusec] Organization invite", email.Subject)
		assert.Equal(t, "test@test.com", email.To)

		assert.NotPanics(t, func() {
			data := email.Data.(map[string]interface{})

			assert.Equal(t, "test", data["WorkspaceName"])
			assert.Equal(t, "test", data["Username"])
			assert.Equal(t, "http://localhost:8043", data["URL"])
		})
	})
}

func TestNewRoleData(t *testing.T) {
	t.Run("should success create a new workspace data", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()
		id := uuid.New()

		data := useCases.NewRoleData(id, id)

		assert.Equal(t, id, data.AccountID)
		assert.Equal(t, id, data.WorkspaceID)
	})
}

func TestInviteUserDataFromIOReadCloser(t *testing.T) {
	t.Run("should success invite user data from request body", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()
		id := uuid.New()

		data := &role.InviteUserData{
			Role:         "admin",
			Email:        "test@test.com",
			AccountID:    id,
			WorkspaceID:  id,
			RepositoryID: id,
			Username:     "test",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		data, err = useCases.InviteUserDataFromIOReadCloser(readCloser)

		assert.NoError(t, err)
		assert.Equal(t, account.Admin, data.Role)
		assert.Equal(t, "test@test.com", data.Email)
		assert.Equal(t, "test", data.Username)
		assert.Equal(t, id, data.AccountID)
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, id, data.RepositoryID)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		response, err := useCases.InviteUserDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

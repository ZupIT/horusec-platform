package repository

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	emailEnums "github.com/ZupIT/horusec-devkit/pkg/enums/email"
	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

func TestNewRepositoryUseCases(t *testing.T) {
	t.Run("should success create a new use cases", func(t *testing.T) {
		assert.NotNil(t, NewRepositoryUseCases())
	})
}

func TestRepositoryDataFromIOReadCloser(t *testing.T) {
	t.Run("should success get repository data from request body", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		data := &repositoryEntities.Data{
			AccountID: uuid.New(),
			Name:      "test",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		response, err := useCases.RepositoryDataFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, data.AccountID, response.AccountID)
		assert.Equal(t, data.Name, response.Name)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		response, err := useCases.RepositoryDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestFilterRepositoryByName(t *testing.T) {
	t.Run("should success create a repository filter by name", func(t *testing.T) {
		useCases := NewRepositoryUseCases()
		id := uuid.New()

		filter := useCases.FilterRepositoryByNameAndWorkspace(id, "test")

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
			assert.Equal(t, "test", filter["name"])
		})
	})
}

func TestIsNotFoundError(t *testing.T) {
	t.Run("should return false when it is an error different than not found", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		assert.False(t, useCases.IsNotFoundError(errors.New("test")))
	})

	t.Run("should return true when it is not found error", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		assert.True(t, useCases.IsNotFoundError(databaseEnums.ErrorNotFoundRecords))
	})
}

func TestNewRepositoryData(t *testing.T) {
	t.Run("should success create a new repository data with account and repository id", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		id := uuid.New()
		accountData := &proto.GetAccountDataResponse{
			AccountID:          id.String(),
			Permissions:        []string{"test"},
			IsApplicationAdmin: true,
		}

		data := useCases.NewRepositoryData(id, id, accountData)
		assert.Equal(t, id, data.RepositoryID)
		assert.Equal(t, id, data.WorkspaceID)
		assert.Equal(t, id, data.AccountID)
		assert.Equal(t, []string{"test"}, data.Permissions)
		assert.Equal(t, true, data.IsApplicationAdmin)
	})
}

func TestFilterRepositoryByID(t *testing.T) {
	t.Run("should success create a repository filter by id", func(t *testing.T) {
		useCases := NewRepositoryUseCases()
		id := uuid.New()

		filter := useCases.FilterRepositoryByID(id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["repository_id"])
		})
	})
}

func TestFilterAccountRepositoryByID(t *testing.T) {
	t.Run("should success create a account repository filter by id", func(t *testing.T) {
		useCases := NewRepositoryUseCases()
		id := uuid.New()

		filter := useCases.FilterAccountRepositoryByID(id, id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["repository_id"])
			assert.Equal(t, id, filter["account_id"])
		})
	})
}

func TestNewOrganizationInviteEmail(t *testing.T) {
	t.Run("should success create a new repository invite email", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		emailBytes := useCases.NewRepositoryInviteEmail("test@test.com", "test", "test")
		assert.NotNil(t, emailBytes)
		assert.NotEmpty(t, emailBytes)

		email := &emailEntities.Message{}
		assert.NoError(t, json.Unmarshal(emailBytes, email))

		assert.Equal(t, "test@test.com", email.To)
		assert.Equal(t, emailEnums.RepositoryInvite, email.TemplateName)
		assert.Equal(t, "[Horusec] Repository invite", email.Subject)

		assert.NotPanics(t, func() {
			data := email.Data.(map[string]interface{})

			assert.Equal(t, "test", data["repositoryName"])
			assert.Equal(t, "test", data["username"])

		})
	})
}

func TestInheritWorkspaceGroups(t *testing.T) {
	t.Run("should success inherit workspace groups", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		workspace := &workspaceEntities.Workspace{
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
		}

		repository := &repositoryEntities.Repository{}
		_ = useCases.InheritWorkspaceGroups(repository, workspace)

		assert.Equal(t, workspace.AuthzAdmin, repository.AuthzAdmin)
		assert.Equal(t, workspace.AuthzAdmin, repository.AuthzSupervisor)
		assert.Equal(t, workspace.AuthzMember, repository.AuthzMember)
	})

	t.Run("should not inherit workspace groups", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		workspace := &workspaceEntities.Workspace{
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
		}

		repository := &repositoryEntities.Repository{
			AuthzMember:     []string{"test2"},
			AuthzAdmin:      []string{"test2"},
			AuthzSupervisor: []string{"test2"},
		}
		_ = useCases.InheritWorkspaceGroups(repository, workspace)

		assert.NotEqual(t, workspace.AuthzAdmin, repository.AuthzAdmin)
		assert.NotEqual(t, workspace.AuthzAdmin, repository.AuthzSupervisor)
		assert.NotEqual(t, workspace.AuthzMember, repository.AuthzMember)
	})

	t.Run("should inherit workspace groups for authz member", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		workspace := &workspaceEntities.Workspace{
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
		}

		repository := &repositoryEntities.Repository{
			AuthzAdmin:      []string{"test2"},
			AuthzSupervisor: []string{"test2"},
		}
		_ = useCases.InheritWorkspaceGroups(repository, workspace)

		assert.NotEqual(t, workspace.AuthzAdmin, repository.AuthzAdmin)
		assert.NotEqual(t, workspace.AuthzAdmin, repository.AuthzSupervisor)
		assert.Equal(t, workspace.AuthzMember, repository.AuthzMember)
	})

	t.Run("should inherit workspace groups for authz supervisor", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		workspace := &workspaceEntities.Workspace{
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
		}

		repository := &repositoryEntities.Repository{
			AuthzMember: []string{"test2"},
			AuthzAdmin:  []string{"test2"},
		}
		_ = useCases.InheritWorkspaceGroups(repository, workspace)

		assert.NotEqual(t, workspace.AuthzAdmin, repository.AuthzAdmin)
		assert.Equal(t, workspace.AuthzAdmin, repository.AuthzSupervisor)
		assert.NotEqual(t, workspace.AuthzMember, repository.AuthzMember)
	})

	t.Run("should inherit workspace groups for authz admin", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		workspace := &workspaceEntities.Workspace{
			AuthzMember: []string{"test"},
			AuthzAdmin:  []string{"test"},
		}

		repository := &repositoryEntities.Repository{
			AuthzMember:     []string{"test2"},
			AuthzSupervisor: []string{"test2"},
		}
		_ = useCases.InheritWorkspaceGroups(repository, workspace)

		assert.Equal(t, workspace.AuthzAdmin, repository.AuthzAdmin)
		assert.NotEqual(t, workspace.AuthzAdmin, repository.AuthzSupervisor)
		assert.NotEqual(t, workspace.AuthzMember, repository.AuthzMember)
	})
}

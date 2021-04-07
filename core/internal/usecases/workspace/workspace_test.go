package workspace

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

func TestNewWorkspaceUseCases(t *testing.T) {
	t.Run("should success create a new use cases", func(t *testing.T) {
		assert.NotNil(t, NewWorkspaceUseCases())
	})
}

func TestGetCreateWorkspaceData(t *testing.T) {
	t.Run("should success get workspace data from request body", func(t *testing.T) {
		useCases := NewWorkspaceUseCases()

		createWorkspaceData := &workspace.Data{
			AccountID: uuid.New(),
			Name:      "test",
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(createWorkspaceData)
		assert.NoError(t, err)

		response, err := useCases.WorkspaceDataFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, createWorkspaceData.AccountID, response.AccountID)
		assert.Equal(t, createWorkspaceData.Name, response.Name)
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

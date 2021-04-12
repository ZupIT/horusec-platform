package token

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/token"
)

func TestNewTokenUseCases(t *testing.T) {
	t.Run("should success create a new use cases", func(t *testing.T) {
		assert.NotNil(t, NewTokenUseCases())
	})
}

func TestWorkspaceDataFromIOReadCloser(t *testing.T) {
	t.Run("should success get workspace data from request body", func(t *testing.T) {
		useCases := NewTokenUseCases()

		data := &token.Data{
			Description: "test",
			IsExpirable: true,
			ExpiresAt:   time.Date(9999, 1, 1, 1, 1, 1, 1, time.UTC),
		}

		readCloser, err := parser.ParseEntityToIOReadCloser(data)
		assert.NoError(t, err)

		response, err := useCases.TokenDataFromIOReadCloser(readCloser)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, data.Description, response.Description)
		assert.Equal(t, data.IsExpirable, response.IsExpirable)
		assert.Equal(t, data.ExpiresAt, response.ExpiresAt)
	})

	t.Run("should return error when failed to parse body to entity", func(t *testing.T) {
		useCases := NewTokenUseCases()

		readCloser, err := parser.ParseEntityToIOReadCloser("")
		assert.NoError(t, err)

		response, err := useCases.TokenDataFromIOReadCloser(readCloser)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestFilterWorkspaceTokenByID(t *testing.T) {
	t.Run("should success create a token workspace filter by id", func(t *testing.T) {
		useCases := NewTokenUseCases()
		id := uuid.New()

		filter := useCases.FilterWorkspaceTokenByID(id, id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
			assert.Equal(t, id, filter["token_id"])
			assert.Equal(t, nil, filter["repository_id"])
		})
	})
}

func TestFilterRepositoryTokenByID(t *testing.T) {
	t.Run("should success create a repository token filter by id", func(t *testing.T) {
		useCases := NewTokenUseCases()
		id := uuid.New()

		filter := useCases.FilterRepositoryTokenByID(id, id, id)

		assert.NotPanics(t, func() {
			assert.Equal(t, id, filter["workspace_id"])
			assert.Equal(t, id, filter["repository_id"])
			assert.Equal(t, id, filter["token_id"])
		})
	})
}

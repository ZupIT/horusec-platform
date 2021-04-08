package repository

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/core/internal/entities/repository"
)

func TestNewRepositoryUseCases(t *testing.T) {
	t.Run("should success create a new use cases", func(t *testing.T) {
		assert.NotNil(t, NewRepositoryUseCases())
	})
}

func TestRepositoryDataFromIOReadCloser(t *testing.T) {
	t.Run("should success get repository data from request body", func(t *testing.T) {
		useCases := NewRepositoryUseCases()

		data := &repository.Data{
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

		filter := useCases.FilterRepositoryByName(id, "test")

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

		data := useCases.NewRepositoryData(id, id)
		assert.Equal(t, id, data.RepositoryID)
		assert.Equal(t, id, data.AccountID)
	})
}

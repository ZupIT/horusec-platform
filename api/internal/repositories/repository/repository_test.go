package repository

import (
	"errors"
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

func TestRepository_CreateRepository(t *testing.T) {
	t.Run("Should create repository with success", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockWrite.On("Create").Return(response.NewResponse(1, nil, nil))
		connectionMock := &database.Connection{
			Write: mockWrite,
		}
		err := NewRepositoriesRepository(connectionMock).CreateRepository(uuid.New(), uuid.New(), uuid.New().String())
		assert.NoError(t, err)
	})
	t.Run("Should create repository with error", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockWrite.On("Create").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		connectionMock := &database.Connection{
			Write: mockWrite,
		}
		err := NewRepositoriesRepository(connectionMock).CreateRepository(uuid.New(), uuid.New(), uuid.New().String())
		assert.Error(t, err)
	})
}

func TestRepository_FindRepository(t *testing.T) {
	t.Run("Should find repository existing and return RepositoryID", func(t *testing.T) {
		mockRead := &database.Mock{}
		mockRead.On("Find").Return(response.NewResponse(0, nil, &map[string]interface{}{
			"repository_id": uuid.NewString(),
		}))
		connectionMock := &database.Connection{
			Read: mockRead,
		}
		repositoryID, err := NewRepositoriesRepository(connectionMock).FindRepository(uuid.New(), uuid.New().String())
		assert.NotEqual(t, repositoryID, uuid.Nil)
		assert.NoError(t, err)
	})
	t.Run("Should find repository existing and return records not found because not exists data", func(t *testing.T) {
		mockRead := &database.Mock{}
		mockRead.On("Find").Return(response.NewResponse(0, nil, nil))
		connectionMock := &database.Connection{
			Read: mockRead,
		}
		repositoryID, err := NewRepositoriesRepository(connectionMock).FindRepository(uuid.New(), uuid.New().String())
		assert.Equal(t, repositoryID, uuid.Nil)
		assert.Equal(t, enums.ErrorNotFoundRecords, err)
	})
	t.Run("Should find repository existing and return records not found", func(t *testing.T) {
		mockRead := &database.Mock{}
		mockRead.On("Find").Return(response.NewResponse(0, enums.ErrorNotFoundRecords, nil))
		connectionMock := &database.Connection{
			Read: mockRead,
		}
		repositoryID, err := NewRepositoriesRepository(connectionMock).FindRepository(uuid.New(), uuid.New().String())
		assert.Equal(t, repositoryID, uuid.Nil)
		assert.Equal(t, enums.ErrorNotFoundRecords, err)
	})
}

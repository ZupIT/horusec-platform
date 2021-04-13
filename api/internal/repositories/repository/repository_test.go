package repository

import (
	"errors"
	"testing"

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

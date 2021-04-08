package analysis

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	analysisEnums "github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

func TestAnalysis_FindAnalysisByID(t *testing.T) {
	t.Run("Should find analysis by id with success", func(t *testing.T) {
		data := &analysis.Analysis{
			ID:             uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.New().String(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.New().String(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		}
		mockRead := &database.Mock{}
		mockRead.On("Find").Return(response.NewResponse(0, nil, data))
		connectionMock := &database.Connection{
			Read: mockRead,
		}
		res := NewRepositoriesAnalysis(connectionMock).FindAnalysisByID(uuid.New())
		assert.NoError(t, res.GetError())
		assert.NotEmpty(t, res.GetData())
		assert.NotEqual(t, res.GetData().(*analysis.Analysis).ID, uuid.Nil)
	})
	t.Run("Should find analysis by id with success", func(t *testing.T) {
		data := &analysis.Analysis{
			ID:             uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.New().String(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.New().String(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		}
		mockRead := &database.Mock{}
		mockRead.On("Find").Return(response.NewResponse(0, nil, data))
		connectionMock := &database.Connection{
			Read: mockRead,
		}
		res := NewRepositoriesAnalysis(connectionMock).FindAnalysisByID(uuid.New())
		assert.NoError(t, res.GetError())
		assert.NotEmpty(t, res.GetData())
		assert.NotEqual(t, res.GetData().(*analysis.Analysis).ID, uuid.Nil)
	})
}

func TestAnalysis_CreateAnalysis(t *testing.T) {
	t.Run("Should create analysis with success", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockWrite.On("Create").Return(response.NewResponse(0, nil, nil))
		connectionMock := &database.Connection{
			Write: mockWrite,
		}
		data := &analysis.Analysis{
			ID:             uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.New().String(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.New().String(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		}
		err := NewRepositoriesAnalysis(connectionMock).CreateAnalysis(data)
		assert.NoError(t, err)
	})
	t.Run("Should create analysis with error", func(t *testing.T) {
		mockWrite := &database.Mock{}
		mockWrite.On("Create").Return(response.NewResponse(0, errors.New("unexpected error"), nil))
		connectionMock := &database.Connection{
			Write: mockWrite,
		}
		data := &analysis.Analysis{
			ID:             uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: uuid.New().String(),
			WorkspaceID:    uuid.New(),
			WorkspaceName:  uuid.New().String(),
			Status:         analysisEnums.Success,
			Errors:         "",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
		}
		err := NewRepositoriesAnalysis(connectionMock).CreateAnalysis(data)
		assert.Error(t, err)
	})
}

package analysis

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	analysisEnum "github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
	appConfiguration "github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

func TestController_GetAnalysis(t *testing.T) {
	t.Run("Should return analysis existing from database", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		dbMockRead := &database.Mock{}
		dbMockRead.On("Find").Return(response.NewResponse(0, nil, &analysis.Analysis{
			ID:         uuid.New(),
			Status:     analysisEnum.Success,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
		}))
		dbMockWrite := &database.Mock{}
		mockAppConfig := &appConfiguration.Mock{}
		controller := NewAnalysisController(
			brokerMock,
			&database.Connection{Read: dbMockRead, Write: dbMockWrite},
			mockAppConfig)
		res, err := controller.GetAnalysis(uuid.New())
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})
	t.Run("Should return error when get analysis from database", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		dbMockRead := &database.Mock{}
		dbMockRead.On("Find").Return(response.NewResponse(0, errors.New("unknown error"), nil))
		dbMockWrite := &database.Mock{}
		mockAppConfig := &appConfiguration.Mock{}
		controller := &Controller{
			broker:        brokerMock,
			databaseRead:  dbMockRead,
			databaseWrite: dbMockWrite,
			appConfig:     mockAppConfig,
		}
		res, err := controller.GetAnalysis(uuid.New())
		assert.Error(t, err)
		assert.Nil(t, res)
	})
	t.Run("Should return error not found records when get analysis from database", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		dbMockRead := &database.Mock{}
		dbMockRead.On("Find").Return(response.NewResponse(0, enums.ErrorNotFoundRecords, nil))
		dbMockWrite := &database.Mock{}
		mockAppConfig := &appConfiguration.Mock{}
		controller := NewAnalysisController(
			brokerMock,
			&database.Connection{Read: dbMockRead, Write: dbMockWrite},
			mockAppConfig)
		res, err := controller.GetAnalysis(uuid.New())
		assert.Equal(t, err, enums.ErrorNotFoundRecords)
		assert.Nil(t, res)
	})
	t.Run("Should not found data return not found records error", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		dbMockRead := &database.Mock{}
		dbMockRead.On("Find").Return(response.NewResponse(0, nil, nil))
		dbMockWrite := &database.Mock{}
		mockAppConfig := &appConfiguration.Mock{}
		controller := NewAnalysisController(
			brokerMock,
			&database.Connection{Read: dbMockRead, Write: dbMockWrite},
			mockAppConfig)
		res, err := controller.GetAnalysis(uuid.New())
		assert.Equal(t, err, enums.ErrorNotFoundRecords)
		assert.Nil(t, res)
	})
}

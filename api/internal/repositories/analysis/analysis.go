package analysis

import (
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

type IAnalysis interface {
	FindAnalysisByID(analysisID uuid.UUID) response.IResponse
	CreateAnalysis(newAnalysis *analysis.Analysis) error
}

type Analysis struct {
	databaseWrite database.IDatabaseWrite
	databaseRead  database.IDatabaseRead
}

func NewRepositoriesAnalysis(connection *database.Connection) IAnalysis {
	return &Analysis{
		databaseWrite: connection.Write,
		databaseRead:  connection.Read,
	}
}

func (a *Analysis) FindAnalysisByID(analysisID uuid.UUID) response.IResponse {
	return a.databaseRead.Find(
		&analysis.Analysis{},
		map[string]interface{}{"analysis_id": analysisID},
		(&analysis.Analysis{}).GetTable())
}

func (a *Analysis) CreateAnalysis(newAnalysis *analysis.Analysis) error {
	return a.databaseWrite.Create(newAnalysis, newAnalysis.GetTable()).GetError()
}

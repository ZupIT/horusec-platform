package analysis

import (
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/api/internal/repositories/analysis/enums"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseEnums "github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type IAnalysis interface {
	FindAnalysisByID(analysisID uuid.UUID) response.IResponse
	CreateFullAnalysis(newAnalysis *analysis.Analysis) error
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

func (a *Analysis) CreateFullAnalysis(newAnalysis *analysis.Analysis) error {
	a.databaseWrite = a.databaseWrite.StartTransaction()
	if err := a.createAnalysis(newAnalysis); err != nil {
		logger.LogError(enums.ErrorRollbackCreate, a.databaseWrite.RollbackTransaction().GetError())
		return err
	}
	if err := a.createManyToManyAnalysisAndVulnerabilities(newAnalysis); err != nil {
		logger.LogError(enums.ErrorRollbackCreate, a.databaseWrite.RollbackTransaction().GetError())
		return err
	}
	err := a.databaseWrite.CommitTransaction().GetError()
	logger.LogError(enums.ErrorCommitCreate, err)
	return err
}

func (a *Analysis) createAnalysis(newAnalysis *analysis.Analysis) error {
	analysisToCreate := &analysis.Analysis{
		ID:             newAnalysis.ID,
		RepositoryID:   newAnalysis.RepositoryID,
		RepositoryName: newAnalysis.RepositoryName,
		WorkspaceID:    newAnalysis.WorkspaceID,
		WorkspaceName:  newAnalysis.WorkspaceName,
		Status:         newAnalysis.Status,
		Errors:         newAnalysis.Errors,
		CreatedAt:      newAnalysis.CreatedAt,
		FinishedAt:     newAnalysis.FinishedAt,
	}
	return a.databaseWrite.Create(analysisToCreate, analysisToCreate.GetTable()).GetError()
}

func (a *Analysis) createManyToManyAnalysisAndVulnerabilities(newAnalysis *analysis.Analysis) error {
	for index := range newAnalysis.AnalysisVulnerabilities {
		manyToMany := newAnalysis.AnalysisVulnerabilities[index]
		vulnerabilityID, err := a.createVulnerabilityIfNotExists(&manyToMany.Vulnerability)
		if err != nil {
			return err
		}
		manyToMany.VulnerabilityID = vulnerabilityID
		if err := a.createManyToMany(&manyToMany); err != nil {
			return err
		}
	}
	return nil
}

func (a *Analysis) createVulnerabilityIfNotExists(vuln *vulnerability.Vulnerability) (uuid.UUID, error) {
	res := a.findVulnerabilityByHash(vuln.VulnHash)
	exists, err := a.checkIfAlreadyExistsVulnerability(res)
	if err == nil {
		if !exists {
			return vuln.VulnerabilityID, a.databaseWrite.Create(vuln, vuln.GetTable()).GetError()
		}
		return uuid.Parse(res.GetData().(map[string]interface{})["vulnerability_id"].(string))
	}
	return uuid.Nil, err
}

func (a *Analysis) checkIfAlreadyExistsVulnerability(res response.IResponse) (bool, error) {
	if res.GetError() != nil {
		if res.GetError() == databaseEnums.ErrorNotFoundRecords {
			return false, nil
		}
		return true, res.GetError()
	}
	return res.GetData() != nil, nil
}

func (a *Analysis) createManyToMany(manyToMany *analysis.AnalysisVulnerabilities) error {
	manyToManyForCreate := &analysis.AnalysisVulnerabilities{
		VulnerabilityID: manyToMany.VulnerabilityID,
		AnalysisID:      manyToMany.AnalysisID,
		CreatedAt:       manyToMany.CreatedAt,
	}
	return a.databaseWrite.Create(manyToManyForCreate, manyToManyForCreate.GetTable()).GetError()
}

func (a *Analysis) findVulnerabilityByHash(vulnHash string) response.IResponse {
	query := `
		SELECT vulnerabilities.vulnerability_id as vulnerability_id
		FROM vulnerabilities
		WHERE vulnerabilities.vuln_hash = ?
	`
	return a.databaseRead.Raw(query, map[string]interface{}{}, vulnHash)
}

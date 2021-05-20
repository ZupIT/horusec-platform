package analysis

import (
	"time"

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
	entity := &analysis.Analysis{}
	condition := map[string]interface{}{"analysis_id": analysisID}
	preloads := map[string][]interface{}{
		"AnalysisVulnerabilities":               {},
		"AnalysisVulnerabilities.Vulnerability": {},
	}
	return a.databaseRead.FindPreload(entity, condition, preloads, entity.GetTable())
}

func (a *Analysis) CreateFullAnalysis(newAnalysis *analysis.Analysis) error {
	tsx := a.databaseWrite.StartTransaction()
	if err := a.createAnalysis(newAnalysis, tsx); err != nil {
		logger.LogError(enums.ErrorRollbackCreate, tsx.RollbackTransaction().GetError())
		return err
	}
	if err := a.createManyToManyAnalysisAndVulnerabilities(newAnalysis, tsx); err != nil {
		logger.LogError(enums.ErrorRollbackCreate, tsx.RollbackTransaction().GetError())
		return err
	}
	err := tsx.CommitTransaction().GetError()
	logger.LogError(enums.ErrorCommitCreate, err)
	return err
}

func (a *Analysis) createAnalysis(newAnalysis *analysis.Analysis, tsx database.IDatabaseWrite) error {
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
	return tsx.Create(analysisToCreate, analysisToCreate.GetTable()).GetError()
}

func (a *Analysis) createManyToManyAnalysisAndVulnerabilities(newAnalysis *analysis.Analysis,
	tsx database.IDatabaseWrite) error {
	for index := range newAnalysis.AnalysisVulnerabilities {
		manyToMany := newAnalysis.AnalysisVulnerabilities[index]
		vulnerabilityID, err := a.createVulnerabilityIfNotExists(&manyToMany.Vulnerability, newAnalysis.WorkspaceID, tsx)
		if err != nil {
			return err
		}
		manyToMany.VulnerabilityID = vulnerabilityID
		if err := a.createManyToMany(&manyToMany, tsx); err != nil {
			return err
		}
	}
	return nil
}

func (a *Analysis) createVulnerabilityIfNotExists(vuln *vulnerability.Vulnerability, workspaceID uuid.UUID,
	tsx database.IDatabaseWrite) (uuid.UUID, error) {
	res := a.findVulnerabilityByHashInWorkspace(vuln.VulnHash, workspaceID)
	exists, err := a.checkIfAlreadyExistsVulnerability(res)
	if err == nil {
		if !exists {
			return vuln.VulnerabilityID, tsx.Create(vuln, vuln.GetTable()).GetError()
		}
		return a.updateCommitAuthors(vuln, res.GetData(), tsx)
	}
	return uuid.Nil, err
}

func (a *Analysis) updateCommitAuthors(vuln *vulnerability.Vulnerability, resFindVuln interface{},
	tsx database.IDatabaseWrite) (uuid.UUID, error) {
	vulnID, err := uuid.Parse(resFindVuln.(map[string]interface{})["vulnerability_id"].(string))
	if err != nil {
		return uuid.Nil, err
	}
	tableName := (&vulnerability.Vulnerability{}).GetTable()
	condition := map[string]interface{}{"vulnerability_id": vulnID}
	entity := map[string]interface{}{
		"commit_author":  vuln.CommitAuthor,
		"commit_email":   vuln.CommitEmail,
		"commit_hash":    vuln.CommitHash,
		"commit_message": vuln.CommitMessage,
		"commit_date":    vuln.CommitDate,
	}
	return vulnID, tsx.Update(entity, condition, tableName).GetErrorExceptNotFound()
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

func (a *Analysis) createManyToMany(manyToMany *analysis.AnalysisVulnerabilities, tsx database.IDatabaseWrite) error {
	manyToManyForCreate := &analysis.AnalysisVulnerabilities{
		VulnerabilityID: manyToMany.VulnerabilityID,
		AnalysisID:      manyToMany.AnalysisID,
		CreatedAt:       time.Now(),
	}
	return tsx.Create(manyToManyForCreate, manyToManyForCreate.GetTable()).GetError()
}

func (a *Analysis) findVulnerabilityByHashInWorkspace(vulnHash string, workspaceID uuid.UUID) response.IResponse {
	query := `
		SELECT vulnerabilities.vulnerability_id as vulnerability_id
		FROM vulnerabilities
		INNER JOIN analysis_vulnerabilities ON vulnerabilities.vulnerability_id = analysis_vulnerabilities.vulnerability_id 
		INNER JOIN analysis ON analysis_vulnerabilities.analysis_id = analysis.analysis_id 
		WHERE vulnerabilities.vuln_hash = ?
		AND analysis.workspace_id = ?
	`
	return a.databaseRead.Raw(query, map[string]interface{}{}, vulnHash, workspaceID)
}

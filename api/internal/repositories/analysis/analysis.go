// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package analysis

import (
	"context"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/services/tracer"

	"github.com/opentracing/opentracing-go"

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
	FindAnalysisByID(ctx context.Context, analysisID uuid.UUID) response.IResponse
	CreateFullAnalysis(ctx context.Context, newAnalysis *analysis.Analysis) error
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

func (a *Analysis) FindAnalysisByID(ctx context.Context, analysisID uuid.UUID) response.IResponse {
	span, _ := opentracing.StartSpanFromContext(ctx, "FindAnalysisByID")
	defer span.Finish()
	entity := &analysis.Analysis{}
	condition := map[string]interface{}{"analysis_id": analysisID}
	preloads := map[string][]interface{}{
		"AnalysisVulnerabilities":               {},
		"AnalysisVulnerabilities.Vulnerability": {},
	}
	return a.databaseRead.FindPreload(entity, condition, preloads, entity.GetTable())
}

func (a *Analysis) CreateFullAnalysis(ctx context.Context, newAnalysis *analysis.Analysis) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateFullAnalysis")
	defer span.Finish()
	tsx := a.databaseWrite.StartTransaction()
	if err := a.createAnalysis(ctx, newAnalysis, tsx); err != nil {
		return setAnalysisRollbackCreateError(span, err, tsx)
	}
	if err := a.createManyToManyAnalysisAndVulnerabilities(ctx, newAnalysis, tsx); err != nil {
		return setAnalysisRollbackCreateError(span, err, tsx)
	}
	err := tsx.CommitTransaction().GetError()
	logger.LogError(enums.ErrorCommitCreate, err)
	return err
}

func setAnalysisRollbackCreateError(span opentracing.Span, err error, tsx database.IDatabaseWrite) error {
	tracer.SetSpanError(span, err)
	logger.LogError(enums.ErrorRollbackCreate, tsx.RollbackTransaction().GetError())
	return err
}

func (a *Analysis) createAnalysis(ctx context.Context, newAnalysis *analysis.Analysis,
	tsx database.IDatabaseWrite) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "createAnalysis")
	defer span.Finish()
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

func (a *Analysis) createManyToManyAnalysisAndVulnerabilities(ctx context.Context,
	newAnalysis *analysis.Analysis, tsx database.IDatabaseWrite) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createManyToManyAnalysisAndVulnerabilities")
	defer span.Finish()
	for index := range newAnalysis.AnalysisVulnerabilities {
		manyToMany := newAnalysis.AnalysisVulnerabilities[index]
		err := a.createManyToManyAndVulnerabilities(ctx, newAnalysis, tsx, &manyToMany, span)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Analysis) createManyToManyAndVulnerabilities(ctx context.Context, newAnalysis *analysis.Analysis,
	tsx database.IDatabaseWrite, manyToMany *analysis.AnalysisVulnerabilities, span opentracing.Span) error {
	vulnerabilityID, err := a.createVulnerabilityIfNotExists(ctx, &manyToMany.Vulnerability,
		newAnalysis.RepositoryID, tsx)
	if err != nil {
		tracer.SetSpanError(span, err)
		return err
	}
	manyToMany.VulnerabilityID = vulnerabilityID
	if err = a.createManyToMany(ctx, manyToMany, tsx); err != nil {
		tracer.SetSpanError(span, err)
		return err
	}
	return nil
}

func (a *Analysis) createVulnerabilityIfNotExists(ctx context.Context, vuln *vulnerability.Vulnerability,
	repositoryID uuid.UUID, tsx database.IDatabaseWrite) (uuid.UUID, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createVulnerabilityIfNotExists")
	defer span.Finish()
	res := a.findVulnerabilityByHashInRepository(ctx, vuln.VulnHash, repositoryID)
	exists, err := a.checkIfAlreadyExistsVulnerability(res)
	if err == nil {
		if !exists {
			return vuln.VulnerabilityID, tsx.Create(vuln, vuln.GetTable()).GetError()
		}
		return a.updateCommitAuthors(ctx, vuln, res.GetData(), tsx)
	}
	return uuid.Nil, err
}

// nolint
func (a *Analysis) updateCommitAuthors(ctx context.Context, vuln *vulnerability.Vulnerability, resFindVuln interface{},
	tsx database.IDatabaseWrite) (uuid.UUID, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "updateCommitAuthors")
	defer span.Finish()
	vulnID, err := uuid.Parse(resFindVuln.(map[string]interface{})["vulnerability_id"].(string))
	if err != nil {
		tracer.SetSpanError(span, err)
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

func (a *Analysis) createManyToMany(ctx context.Context, manyToMany *analysis.AnalysisVulnerabilities,
	tsx database.IDatabaseWrite) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "createManyToMany")
	defer span.Finish()
	manyToManyForCreate := &analysis.AnalysisVulnerabilities{
		VulnerabilityID: manyToMany.VulnerabilityID,
		AnalysisID:      manyToMany.AnalysisID,
		CreatedAt:       time.Now(),
	}
	return tsx.Create(manyToManyForCreate, manyToManyForCreate.GetTable()).GetError()
}

func (a *Analysis) findVulnerabilityByHashInRepository(ctx context.Context, vulnHash string,
	repositoryID uuid.UUID) response.IResponse {
	span, _ := opentracing.StartSpanFromContext(ctx, "createManyToMany")
	defer span.Finish()
	query := `
		SELECT vulnerabilities.vulnerability_id as vulnerability_id
		FROM vulnerabilities
		INNER JOIN analysis_vulnerabilities ON vulnerabilities.vulnerability_id = analysis_vulnerabilities.vulnerability_id 
		INNER JOIN analysis ON analysis_vulnerabilities.analysis_id = analysis.analysis_id 
		WHERE vulnerabilities.vuln_hash = ?
		AND analysis.repository_id = ?
	`
	return a.databaseRead.Raw(query, map[string]interface{}{}, vulnHash, repositoryID)
}

package analysisv1

import (
	"strings"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/cli"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/enums/confidence"
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"github.com/ZupIT/horusec-devkit/pkg/enums/tools"
	"github.com/google/uuid"
)

type AnalysisV1 struct {
	analysis.Analysis
	CompanyID   uuid.UUID `json:"companyID"`
	CompanyName string    `json:"companyName"`
}

type AnalysisCLIDataV1 struct {
	Analysis       *AnalysisV1 `json:"analysis"`
	RepositoryName string      `json:"repositoryName"`
}

func (a *AnalysisCLIDataV1) ParseDataV1ToV2() (analysisData *cli.AnalysisData) {
	return &cli.AnalysisData{
		Analysis:       a.getAnalysisToV2(),
		RepositoryName: a.RepositoryName,
	}
}

func (a *AnalysisCLIDataV1) getAnalysisToV2() *analysis.Analysis {
	return &analysis.Analysis{
		ID:                      a.Analysis.ID,
		RepositoryID:            a.Analysis.RepositoryID,
		RepositoryName:          a.Analysis.RepositoryName,
		WorkspaceID:             a.Analysis.CompanyID,
		WorkspaceName:           a.Analysis.CompanyName,
		Status:                  a.Analysis.Status,
		Errors:                  a.Analysis.Errors,
		CreatedAt:               a.Analysis.CreatedAt,
		FinishedAt:              a.Analysis.FinishedAt,
		AnalysisVulnerabilities: a.getAnalysisVulnerabilitiesToV2(),
	}
}

func (a *AnalysisCLIDataV1) getAnalysisVulnerabilitiesToV2() (manyToMany []analysis.AnalysisVulnerabilities) {
	for key := range a.Analysis.AnalysisVulnerabilities {
		item := a.Analysis.AnalysisVulnerabilities[key]
		manyToMany = append(manyToMany, analysis.AnalysisVulnerabilities{
			VulnerabilityID: item.VulnerabilityID,
			AnalysisID:      item.AnalysisID,
			CreatedAt:       item.CreatedAt,
			Vulnerability:   a.getVulnerabilityToV2(&item),
		})
	}
	return manyToMany
}

func (a *AnalysisCLIDataV1) getVulnerabilityToV2(
	vuln *analysis.AnalysisVulnerabilities) (newVuln vulnerability.Vulnerability) {
	newVuln = vuln.Vulnerability
	newVuln.Confidence = a.setValidConfidenceToV2(newVuln.Confidence)
	newVuln.Severity = a.setValidSeverityToV2(newVuln.Severity)
	newVuln.SecurityTool = a.setValidTool(newVuln.SecurityTool)
	return newVuln
}

func (a *AnalysisCLIDataV1) setValidConfidenceToV2(currentConfidence confidence.Confidence) confidence.Confidence {
	for _, conf := range confidence.Values() {
		if strings.EqualFold(string(currentConfidence), string(conf)) {
			return conf
		}
	}
	return confidence.Low
}

func (a *AnalysisCLIDataV1) setValidSeverityToV2(currentSeverity severities.Severity) severities.Severity {
	for _, sev := range severities.Values() {
		if strings.EqualFold(string(currentSeverity), string(sev)) {
			return sev
		}
	}
	return severities.Unknown
}

func (a *AnalysisCLIDataV1) setValidTool(currentTool tools.Tool) tools.Tool {
	for _, tool := range tools.Values() {
		if strings.EqualFold(string(currentTool), string(tool)) {
			return tool
		}
	}
	return tools.HorusecEngine
}

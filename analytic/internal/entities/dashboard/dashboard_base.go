package dashboard

import (
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"

	"github.com/google/uuid"
)

type Vulnerability struct {
	VulnerabilityID       uuid.UUID `json:"vulnerabilityID" gorm:"Column:vulnerability_id"`
	CreatedAt             time.Time `json:"createdAt" gorm:"Column:created_at"`
	Active                bool      `json:"active" gorm:"Column:active"`
	WorkspaceID           uuid.UUID `json:"workspaceID" gorm:"Column:workspace_id"`
	RepositoryID          uuid.UUID `json:"repositoryID" gorm:"Column:repository_id"`
	CriticalVulnerability int       `json:"criticalVulnerability" gorm:"Column:critical_vulnerability"`
	CriticalFalsePositive int       `json:"criticalFalsePositive" gorm:"Column:critical_false_positive"`
	CriticalRiskAccepted  int       `json:"criticalRiskAccepted" gorm:"Column:critical_risk_accepted"`
	CriticalCorrected     int       `json:"criticalCorrected" gorm:"Column:critical_corrected"`
	HighVulnerability     int       `json:"highVulnerability" gorm:"Column:high_vulnerability"`
	HighFalsePositive     int       `json:"highFalsePositive" gorm:"Column:high_false_positive"`
	HighRiskAccepted      int       `json:"highRiskAccepted" gorm:"Column:high_risk_accepted"`
	HighCorrected         int       `json:"highCorrected" gorm:"Column:high_corrected"`
	MediumVulnerability   int       `json:"mediumVulnerability" gorm:"Column:medium_vulnerability"`
	MediumFalsePositive   int       `json:"mediumFalsePositive" gorm:"Column:medium_false_positive"`
	MediumRiskAccepted    int       `json:"mediumRiskAccepted" gorm:"Column:medium_risk_accepted"`
	MediumCorrected       int       `json:"mediumCorrected" gorm:"Column:medium_corrected"`
	LowVulnerability      int       `json:"lowVulnerability" gorm:"Column:low_vulnerability"`
	LowFalsePositive      int       `json:"lowFalsePositive" gorm:"Column:low_false_positive"`
	LowRiskAccepted       int       `json:"lowRiskAccepted" gorm:"Column:low_risk_accepted"`
	LowCorrected          int       `json:"lowCorrected" gorm:"Column:low_corrected"`
	InfoVulnerability     int       `json:"infoVulnerability" gorm:"Column:info_vulnerability"`
	InfoFalsePositive     int       `json:"infoFalsePositive" gorm:"Column:info_false_positive"`
	InfoRiskAccepted      int       `json:"infoRiskAccepted" gorm:"Column:info_risk_accepted"`
	InfoCorrected         int       `json:"infoCorrected" gorm:"Column:info_corrected"`
	UnknownVulnerability  int       `json:"unknownVulnerability" gorm:"Column:unknown_vulnerability"`
	UnknownFalsePositive  int       `json:"unknownFalsePositive" gorm:"Column:unknown_false_positive"`
	UnknownRiskAccepted   int       `json:"unknownRiskAccepted" gorm:"Column:unknown_risk_accepted"`
	UnknownCorrected      int       `json:"unknownCorrected" gorm:"Column:unknown_corrected"`
}

func (v *Vulnerability) ToResponseSeverity() ResponseSeverity {
	responseSeverity := &ResponseSeverity{
		Critical: &ResponseSeverityContAndTypes{Count: 0, Types: &ResponseVulnTypes{}},
		High:     &ResponseSeverityContAndTypes{Count: 0, Types: &ResponseVulnTypes{}},
		Medium:   &ResponseSeverityContAndTypes{Count: 0, Types: &ResponseVulnTypes{}},
		Low:      &ResponseSeverityContAndTypes{Count: 0, Types: &ResponseVulnTypes{}},
		Info:     &ResponseSeverityContAndTypes{Count: 0, Types: &ResponseVulnTypes{}},
		Unknown:  &ResponseSeverityContAndTypes{Count: 0, Types: &ResponseVulnTypes{}},
	}
	responseSeverity = responseSeverity.SumVulnerabilityCritical(v)
	responseSeverity = responseSeverity.SumVulnerabilityHigh(v)
	responseSeverity = responseSeverity.SumVulnerabilityMedium(v)
	responseSeverity = responseSeverity.SumVulnerabilityLow(v)
	responseSeverity = responseSeverity.SumVulnerabilityInfo(v)
	responseSeverity = responseSeverity.SumVulnerabilityUnknown(v)
	return *responseSeverity
}

// nolint:funlen,gocyclo // is not necessary unknown type and factory of severity
func (v *Vulnerability) AddCountVulnerabilityBySeverity(
	count int, severity severities.Severity, vulnType vulnerability.Type) {
	switch severity {
	case severities.Critical:
		v.AddCountVulnerabilityCritical(count, vulnType)
	case severities.High:
		v.AddCountVulnerabilityHigh(count, vulnType)
	case severities.Medium:
		v.AddCountVulnerabilityMedium(count, vulnType)
	case severities.Low:
		v.AddCountVulnerabilityLow(count, vulnType)
	case severities.Info:
		v.AddCountVulnerabilityInfo(count, vulnType)
	case severities.Unknown:
		v.AddCountVulnerabilityUnknown(count, vulnType)
	}
}

// nolint:exhaustive // is not necessary unknown type
func (v *Vulnerability) AddCountVulnerabilityCritical(count int, vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.CriticalVulnerability += count
	case vulnerability.RiskAccepted:
		v.CriticalRiskAccepted += count
	case vulnerability.FalsePositive:
		v.CriticalFalsePositive += count
	case vulnerability.Corrected:
		v.CriticalCorrected += count
	}
}

// nolint:exhaustive // is not duplicate method
func (v *Vulnerability) AddCountVulnerabilityHigh(count int, vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.HighVulnerability += count
	case vulnerability.RiskAccepted:
		v.HighRiskAccepted += count
	case vulnerability.FalsePositive:
		v.HighFalsePositive += count
	case vulnerability.Corrected:
		v.HighCorrected += count
	}
}

// nolint:exhaustive // is not duplicate method
func (v *Vulnerability) AddCountVulnerabilityMedium(count int, vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.MediumVulnerability += count
	case vulnerability.RiskAccepted:
		v.MediumRiskAccepted += count
	case vulnerability.FalsePositive:
		v.MediumFalsePositive += count
	case vulnerability.Corrected:
		v.MediumCorrected += count
	}
}

// nolint:exhaustive // is not duplicate method
func (v *Vulnerability) AddCountVulnerabilityLow(count int, vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.LowVulnerability += count
	case vulnerability.RiskAccepted:
		v.LowRiskAccepted += count
	case vulnerability.FalsePositive:
		v.LowFalsePositive += count
	case vulnerability.Corrected:
		v.LowCorrected += count
	}
}

// nolint:exhaustive // is not duplicate method
func (v *Vulnerability) AddCountVulnerabilityInfo(count int, vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.InfoVulnerability += count
	case vulnerability.RiskAccepted:
		v.InfoRiskAccepted += count
	case vulnerability.FalsePositive:
		v.InfoFalsePositive += count
	case vulnerability.Corrected:
		v.InfoCorrected += count
	}
}

// nolint:exhaustive // is not duplicate method
func (v *Vulnerability) AddCountVulnerabilityUnknown(count int, vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.UnknownVulnerability += count
	case vulnerability.RiskAccepted:
		v.UnknownRiskAccepted += count
	case vulnerability.FalsePositive:
		v.UnknownFalsePositive += count
	case vulnerability.Corrected:
		v.UnknownCorrected += count
	}
}

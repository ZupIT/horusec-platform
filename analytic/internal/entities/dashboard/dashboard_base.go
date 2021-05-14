package dashboard

import (
	"time"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"
)

type Vulnerability struct {
	VulnerabilityID       uuid.UUID `json:"vulnerabilityID" gorm:"Column:vulnerability_id"`
	CreatedAt             time.Time `json:"createdAt" gorm:"Column:created_at"`
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

func (v *Vulnerability) ToResponseSeverity() response.BySeverities {
	responseSeverity := &response.BySeverities{
		Critical: &response.BySeverity{Count: 0, Types: &response.ByVulnerabilityTypes{}},
		High:     &response.BySeverity{Count: 0, Types: &response.ByVulnerabilityTypes{}},
		Medium:   &response.BySeverity{Count: 0, Types: &response.ByVulnerabilityTypes{}},
		Low:      &response.BySeverity{Count: 0, Types: &response.ByVulnerabilityTypes{}},
		Info:     &response.BySeverity{Count: 0, Types: &response.ByVulnerabilityTypes{}},
		Unknown:  &response.BySeverity{Count: 0, Types: &response.ByVulnerabilityTypes{}},
	}
	responseSeverity = responseSeverity.SumVulnerabilityCritical(v)
	responseSeverity = responseSeverity.SumVulnerabilityHigh(v)
	responseSeverity = responseSeverity.SumVulnerabilityMedium(v)
	responseSeverity = responseSeverity.SumVulnerabilityLow(v)
	responseSeverity = responseSeverity.SumVulnerabilityInfo(v)
	responseSeverity = responseSeverity.SumVulnerabilityUnknown(v)
	return *responseSeverity
}

func (v *Vulnerability) AddCountVulnerabilityBySeverity(severity severities.Severity, vulnType vulnerability.Type) {
	switch severity {
	case severities.Critical:
		v.AddCountVulnerabilityCritical(vulnType)
	case severities.High:
		v.AddCountVulnerabilityHigh(vulnType)
	case severities.Medium:
		v.AddCountVulnerabilityMedium(vulnType)
	case severities.Low:
		v.AddCountVulnerabilityLow(vulnType)
	case severities.Info:
		v.AddCountVulnerabilityInfo(vulnType)
	case severities.Unknown:
		v.AddCountVulnerabilityUnknown(vulnType)
	}
}

func (v *Vulnerability) AddCountVulnerabilityCritical(vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.CriticalVulnerability++
	case vulnerability.RiskAccepted:
		v.CriticalRiskAccepted++
	case vulnerability.FalsePositive:
		v.CriticalFalsePositive++
	case vulnerability.Corrected:
		v.CriticalCorrected++
	}
}

func (v *Vulnerability) AddCountVulnerabilityHigh(vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.HighVulnerability++
	case vulnerability.RiskAccepted:
		v.HighRiskAccepted++
	case vulnerability.FalsePositive:
		v.HighFalsePositive++
	case vulnerability.Corrected:
		v.HighCorrected++
	}
}

func (v *Vulnerability) AddCountVulnerabilityMedium(vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.MediumVulnerability++
	case vulnerability.RiskAccepted:
		v.MediumRiskAccepted++
	case vulnerability.FalsePositive:
		v.MediumFalsePositive++
	case vulnerability.Corrected:
		v.MediumCorrected++
	}
}

func (v *Vulnerability) AddCountVulnerabilityLow(vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.LowVulnerability++
	case vulnerability.RiskAccepted:
		v.LowRiskAccepted++
	case vulnerability.FalsePositive:
		v.LowFalsePositive++
	case vulnerability.Corrected:
		v.LowCorrected++
	}
}

func (v *Vulnerability) AddCountVulnerabilityInfo(vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.InfoVulnerability++
	case vulnerability.RiskAccepted:
		v.InfoRiskAccepted++
	case vulnerability.FalsePositive:
		v.InfoFalsePositive++
	case vulnerability.Corrected:
		v.InfoCorrected++
	}
}

func (v *Vulnerability) AddCountVulnerabilityUnknown(vulnType vulnerability.Type) {
	switch vulnType {
	case vulnerability.Vulnerability:
		v.UnknownVulnerability++
	case vulnerability.RiskAccepted:
		v.UnknownRiskAccepted++
	case vulnerability.FalsePositive:
		v.UnknownFalsePositive++
	case vulnerability.Corrected:
		v.UnknownCorrected++
	}
}

package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"

	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	enumsVulnerability "github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"
)

type VulnerabilityTypes struct {
	Vulnerability int `json:"vulnerability" gorm:"Column:vulnerability"`
	RiskAccepted  int `json:"riskAccept" gorm:"Column:risk_accept"`
	FalsePositive int `json:"falsePositive" gorm:"Column:false_positive"`
	Corrected     int `json:"corrected" gorm:"Column:corrected"`
	Unknown       int `json:"unknown" gorm:"Column:unknown"`
}

func (v *VulnerabilityTypes) setCountByType(count int, vulnType enumsVulnerability.Type) *VulnerabilityTypes {
	switch vulnType {
	case enumsVulnerability.Vulnerability:
		v.Vulnerability = count
	case enumsVulnerability.RiskAccepted:
		v.RiskAccepted = count
	case enumsVulnerability.FalsePositive:
		v.FalsePositive = count
	case enumsVulnerability.Corrected:
		v.Corrected = count
	case enumsVulnerability.Unknown:
		v.Unknown = count
	}
	return v
}

type VulnerabilityCount struct {
	Count int                `json:"count"`
	Types *VulnerabilityTypes `json:"types"`
}

type VulnerabilityBySeverity struct {
	Critical *VulnerabilityCount `json:"critical"`
	High     *VulnerabilityCount `json:"high"`
	Medium   *VulnerabilityCount `json:"medium"`
	Low      *VulnerabilityCount `json:"low"`
	Info     *VulnerabilityCount `json:"info"`
	Unknown  *VulnerabilityCount `json:"unknown"`
}

func NewVulnerabilityBySeverity() *VulnerabilityBySeverity {
	return &VulnerabilityBySeverity{
		Critical: &VulnerabilityCount{Types: &VulnerabilityTypes{}},
		High:     &VulnerabilityCount{Types: &VulnerabilityTypes{}},
		Medium:   &VulnerabilityCount{Types: &VulnerabilityTypes{}},
		Low:      &VulnerabilityCount{Types: &VulnerabilityTypes{}},
		Info:     &VulnerabilityCount{Types: &VulnerabilityTypes{}},
		Unknown:  &VulnerabilityCount{Types: &VulnerabilityTypes{}},
	}
}

func (v *VulnerabilityBySeverity) SetCountBySeverityAndCountType(countSeverity, countType int, severity severities.Severity, vulnType enumsVulnerability.Type) *VulnerabilityBySeverity {
	switch severity {
	case severities.Critical:
		v.Critical.Count = countSeverity
		v.Critical.Types = v.Critical.Types.setCountByType(countType, vulnType)
	case severities.High:
		v.High.Count = countSeverity
		v.High.Types = v.High.Types.setCountByType(countType, vulnType)
	case severities.Medium:
		v.Medium.Count = countSeverity
		v.Medium.Types = v.Medium.Types.setCountByType(countType, vulnType)
	case severities.Low:
		v.Low.Count = countSeverity
		v.Low.Types = v.Low.Types.setCountByType(countType, vulnType)
	case severities.Info:
		v.Info.Count = countSeverity
		v.Info.Types = v.Info.Types.setCountByType(countType, vulnType)
	case severities.Unknown:
		v.Unknown.Count = countSeverity
		v.Unknown.Types = v.Unknown.Types.setCountByType(countType, vulnType)
	}
	return v
}

type VulnerabilityByDeveloper struct {
	DeveloperEmail string `json:"developerEmail" gorm:"Column:commit_email"`
	VulnerabilityBySeverity `gorm:"-"`
}

type VulnerabilityByRepository struct {
	RepositoryID   string `json:"repository_id" gorm:"Column:repository_id"`
	RepositoryName string `json:"repository_name" gorm:"Column:repository_name"`
	VulnerabilityBySeverity `gorm:"-"`
}

// nolint:lll // annotations is necessary in one line
type VulnerabilityByLanguage struct {
	Language languages.Language `json:"language" example:"Leaks" enums:"Go,C#,Dart,Ruby,Python,Java,Kotlin,Javascript,Typescript,Leaks,HCL,C,PHP,HTML,Generic,YAML,Elixir,Shell"`
	VulnerabilityBySeverity `gorm:"-"`
}

type VulnerabilityByTime struct {
	Time time.Time `json:"time" gorm:"Column:created_at"`
	VulnerabilityBySeverity `gorm:"-"`
}

type VulnerabilityDetails struct {
	TotalItems      int                          `json:"totalItems"`
	Vulnerabilities []vulnerability.Vulnerability `json:"vulnerabilities"`
}

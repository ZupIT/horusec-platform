package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"
)

type VulnerabilitiesByLanguage struct {
	Language languages.Language `json:"language" gorm:"Column:language"`
	Vulnerability
}

func ParseListVulnByLanguageToListResponse(vulns []*VulnerabilitiesByLanguage) (result []response.ResponseByLanguage) {
	for index := range vulns {
		result = append(result, vulns[index].ToResponseByLanguage())
	}
	return result
}

func (v *VulnerabilitiesByLanguage) GetTable() string {
	return "vulnerabilities_by_language"
}

func (v *VulnerabilitiesByLanguage) ToResponseByLanguage() response.ResponseByLanguage {
	return response.ResponseByLanguage{
		Language:              v.Language,
		response.BySeverities: v.ToResponseSeverity(),
	}
}

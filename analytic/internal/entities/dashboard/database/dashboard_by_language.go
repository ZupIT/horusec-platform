package database

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"
)

type VulnerabilitiesByLanguage struct {
	Language languages.Language `json:"language" gorm:"Column:language"`
	response.Vulnerability
}

func ParseListVulnByLanguageToListResponse(vulns []*VulnerabilitiesByLanguage) (result []response.ByLanguage) {
	for index := range vulns {
		result = append(result, vulns[index].ToResponseByLanguage())
	}
	return result
}

func (v *VulnerabilitiesByLanguage) ToResponseByLanguage() response.ByLanguage {
	return response.ByLanguage{
		Language:     v.Language,
		BySeverities: v.ToResponseBySeverities(),
	}
}

package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
)

type VulnerabilitiesByLanguage struct {
	Language languages.Language `json:"language" gorm:"Column:language"`
	Vulnerability
}

func ParseListVulnByLanguageToListResponse(vulns []*VulnerabilitiesByLanguage) (result []ResponseByLanguage) {
	for index := range vulns {
		result = append(result, vulns[index].ToResponseByLanguage())
	}
	return result
}

func (v *VulnerabilitiesByLanguage) GetTable() string {
	return "vulnerabilities_by_language"
}

func (v *VulnerabilitiesByLanguage) ToResponseByLanguage() ResponseByLanguage {
	return ResponseByLanguage{
		Language:         v.Language,
		ResponseSeverity: v.ToResponseSeverity(),
	}
}

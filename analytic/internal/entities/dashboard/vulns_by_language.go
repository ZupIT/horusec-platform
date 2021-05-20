package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
)

type VulnerabilitiesByLanguage struct {
	Language languages.Language `json:"language" gorm:"Column:language"`
	Vulnerability
}

func (v *VulnerabilitiesByLanguage) ToResponseByLanguage() ByLanguage {
	return ByLanguage{
		Language:     v.Language,
		BySeverities: v.ToResponseBySeverities(),
	}
}

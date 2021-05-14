package repositories

import (
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"
)

type VulnerabilitiesByAuthor struct {
	Author string `json:"author" gorm:"Column:author"`
	response.Vulnerability
}

func ParseListVulnByAuthorToListResponse(vulns []*VulnerabilitiesByAuthor) (result []response.ByAuthor) {
	for index := range vulns {
		result = append(result, vulns[index].ToResponseByAuthor())
	}
	return result
}

func (v *VulnerabilitiesByAuthor) ToResponseByAuthor() response.ByAuthor {
	return response.ByAuthor{
		Author:       v.Author,
		BySeverities: v.ToResponseBySeverities(),
	}
}

package repositories

import (
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"
)

type VulnerabilitiesByRepository struct {
	RepositoryName string `json:"repository_name" gorm:"Column:repository_name"`
	response.Vulnerability
}

func ParseListVulnByRepositoryToListResponse(vulns []*VulnerabilitiesByRepository) (result []response.ByRepository) {
	for index := range vulns {
		result = append(result, vulns[index].ToResponseByRepository())
	}
	return result
}

func (v *VulnerabilitiesByRepository) ToResponseByRepository() response.ByRepository {
	return response.ByRepository{
		RepositoryName: v.RepositoryName,
		BySeverities:   v.ToResponseBySeverities(),
	}
}

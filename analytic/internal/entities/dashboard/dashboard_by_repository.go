package dashboard

import "github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"

type VulnerabilitiesByRepository struct {
	RepositoryName string `json:"repository_name" gorm:"Column:repository_name"`
	Vulnerability
}

func ParseListVulnByRepositoryToListResponse(vulns []*VulnerabilitiesByRepository) (result []response.ResponseByRepository) {
	for index := range vulns {
		result = append(result, vulns[index].ToResponseByRepository())
	}
	return result
}

func (v *VulnerabilitiesByRepository) GetTable() string {
	return "vulnerabilities_by_repository"
}

func (v *VulnerabilitiesByRepository) ToResponseByRepository() response.ResponseByRepository {
	return response.ResponseByRepository{
		RepositoryName:        v.RepositoryName,
		response.BySeverities: v.ToResponseSeverity(),
	}
}

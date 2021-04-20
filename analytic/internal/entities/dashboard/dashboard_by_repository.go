package dashboard

type VulnerabilitiesByRepository struct {
	RepositoryName string `json:"repository_name" gorm:"Column:repository_name"`
	IsLast         bool   `json:"isLast" gorm:"Column:is_last"`
	Vulnerability
}

func ParseListVulnByRepositoryToListResponse(vulns []VulnerabilitiesByRepository) (result []ResponseByRepository) {
	for _, vuln := range vulns {
		result = append(result, vuln.ToResponseByRepository())
	}
	return result
}

func (v *VulnerabilitiesByRepository) GetTable() string {
	return "vulnerabilities_by_repository"
}

func (v *VulnerabilitiesByRepository) ToResponseByRepository() ResponseByRepository {
	return ResponseByRepository{
		RepositoryName:   v.RepositoryName,
		ResponseSeverity: v.ToResponseSeverity(),
	}
}

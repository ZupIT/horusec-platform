package dashboard

type VulnerabilitiesByAuthor struct {
	Author string `json:"author" gorm:"Column:author"`
	Vulnerability
}

func ParseListVulnByAuthorToListResponse(vulns []VulnerabilitiesByAuthor) (result []ResponseByAuthor) {
	for _, vuln := range vulns {
		result = append(result, vuln.ToResponseByAuthor())
	}
	return result
}

func (v *VulnerabilitiesByAuthor) GetTable() string {
	return "vulnerabilities_by_author"
}

func (v *VulnerabilitiesByAuthor) ToResponseByAuthor() ResponseByAuthor {
	return ResponseByAuthor{
		Author:           v.Author,
		ResponseSeverity: v.ToResponseSeverity(),
	}
}

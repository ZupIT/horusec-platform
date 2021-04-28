package dashboard

type VulnerabilitiesByAuthor struct {
	Author string `json:"author" gorm:"Column:author"`
	Vulnerability
}

func ParseListVulnByAuthorToListResponse(vulns []*VulnerabilitiesByAuthor) (result []ResponseByAuthor) {
	for index := range vulns {
		result = append(result, vulns[index].ToResponseByAuthor())
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

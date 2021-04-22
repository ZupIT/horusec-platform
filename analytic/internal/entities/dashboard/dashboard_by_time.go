package dashboard

type VulnerabilitiesByTime struct {
	Vulnerability
}

func ParseListVulnByTimeToListResponse(vulns []VulnerabilitiesByTime) (result []ResponseByTime) {
	for index := range vulns {
		result = append(result, vulns[index].ToResponseByTime())
	}
	return result
}

func (v *VulnerabilitiesByTime) GetTable() string {
	return "vulnerabilities_by_time"
}

func (v *VulnerabilitiesByTime) ToResponseByTime() ResponseByTime {
	return ResponseByTime{
		Time:             v.CreatedAt,
		ResponseSeverity: v.ToResponseSeverity(),
	}
}

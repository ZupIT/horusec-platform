package dashboard

type VulnerabilitiesByTime struct {
	Vulnerability
}

func ParseListVulnByTimeToListResponse(vulns []VulnerabilitiesByTime) (result []ResponseByTime) {
	for _, vuln := range vulns {
		result = append(result, vuln.ToResponseByTime())
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

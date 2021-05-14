package dashboard

import "github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"

type VulnerabilitiesByTime struct {
	Vulnerability
}

func ParseListVulnByTimeToListResponse(vulns []*VulnerabilitiesByTime) (result []response.ResponseByTime) {
	for index := range vulns {
		result = append(result, vulns[index].ToResponseByTime())
	}
	return result
}

func (v *VulnerabilitiesByTime) GetTable() string {
	return "vulnerabilities_by_time"
}

func (v *VulnerabilitiesByTime) ToResponseByTime() response.ResponseByTime {
	return response.ResponseByTime{
		Time:                  v.CreatedAt,
		response.BySeverities: v.ToResponseSeverity(),
	}
}

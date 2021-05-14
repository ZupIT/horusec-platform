package database

import (
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"
)

type VulnerabilitiesByTime struct {
	response.Vulnerability
}

func ParseListVulnByTimeToListResponse(vulns []*VulnerabilitiesByTime) (result []response.ByTime) {
	for index := range vulns {
		result = append(result, vulns[index].ToResponseByTime())
	}
	return result
}

func (v *VulnerabilitiesByTime) ToResponseByTime() response.ByTime {
	return response.ByTime{
		Time:         v.CreatedAt,
		BySeverities: v.ToResponseBySeverities(),
	}
}

package database

import (
	"testing"

	response2 "github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"

	"github.com/stretchr/testify/assert"
)

func TestParseListVulnByRepositoryToListResponse(t *testing.T) {
	t.Run("Parse to response correctly", func(t *testing.T) {
		vulns := []*VulnerabilitiesByRepository{
			&VulnerabilitiesByRepository{
				RepositoryName: "my-repository",
				Vulnerability: response2.Vulnerability{
					CriticalVulnerability: 1,
				},
			},
		}
		response := ParseListVulnByRepositoryToListResponse(vulns)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, 1, response[0].Critical.Count)
	})
}

func TestVulnerabilitiesByRepository_GetTable(t *testing.T) {
	t.Run("Should get table correctly", func(t *testing.T) {
		assert.Equal(t, "vulnerabilities_by_repository", (&VulnerabilitiesByRepository{}).GetTable())
	})
}

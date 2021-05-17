package dashboard

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseListVulnByTimeToListResponse(t *testing.T) {
	t.Run("Parse to response correctly", func(t *testing.T) {
		vulns := []*VulnerabilitiesByTime{
			&VulnerabilitiesByTime{
				Vulnerability: Vulnerability{
					CreatedAt:             time.Now(),
					CriticalVulnerability: 1,
				},
			},
		}
		response := ParseListVulnByTimeToListResponse(vulns)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, 1, response[0].Critical.Count)
	})
}

func TestVulnerabilitiesByTime_GetTable(t *testing.T) {
	t.Run("Should get table correctly", func(t *testing.T) {
		assert.Equal(t, "vulnerabilities_by_time", (&VulnerabilitiesByTime{}).GetTable())
	})
}

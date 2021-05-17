package dashboard

import (
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/stretchr/testify/assert"
)

func TestParseListVulnByLanguageToListResponse(t *testing.T) {
	t.Run("Parse to response correctly", func(t *testing.T) {
		vulns := []*VulnerabilitiesByLanguage{
			&VulnerabilitiesByLanguage{
				Language: languages.Leaks,
				Vulnerability: Vulnerability{
					CriticalVulnerability: 1,
				},
			},
		}
		response := ParseListVulnByLanguageToListResponse(vulns)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, 1, response[0].Critical.Count)
	})
}

func TestVulnerabilitiesByLanguage_GetTable(t *testing.T) {
	t.Run("Should get table correctly", func(t *testing.T) {
		assert.Equal(t, "vulnerabilities_by_language", (&VulnerabilitiesByLanguage{}).GetTable())
	})
}

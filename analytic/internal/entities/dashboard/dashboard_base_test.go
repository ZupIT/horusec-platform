package dashboard

import (
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"
	"github.com/stretchr/testify/assert"
)

func TestVulnerability_ToResponseSeverity(t *testing.T) {
	t.Run("Should parse to response correctly", func(t *testing.T) {
		v := &Vulnerability{
			CriticalVulnerability: 2,
			CriticalRiskAccepted:  1,
			CriticalFalsePositive: 1,
			CriticalCorrected:     2,
			HighVulnerability:     2,
			HighFalsePositive:     1,
			HighRiskAccepted:      1,
			HighCorrected:         1,
			MediumVulnerability:   2,
			MediumFalsePositive:   1,
			MediumRiskAccepted:    1,
			MediumCorrected:       0,
			LowVulnerability:      2,
			LowFalsePositive:      1,
			LowRiskAccepted:       0,
			LowCorrected:          0,
			InfoVulnerability:     1,
			InfoFalsePositive:     0,
			InfoRiskAccepted:      1,
			InfoCorrected:         0,
			UnknownVulnerability:  0,
			UnknownFalsePositive:  1,
			UnknownRiskAccepted:   0,
			UnknownCorrected:      0,
		}
		response := v.ToResponseSeverity()
		assert.Equal(t, 6, response.Critical.Count)
		assert.Equal(t, 5, response.High.Count)
		assert.Equal(t, 4, response.Medium.Count)
		assert.Equal(t, 3, response.Low.Count)
		assert.Equal(t, 2, response.Info.Count)
		assert.Equal(t, 1, response.Unknown.Count)
	})
}

func TestVulnerability_AddCountVulnerabilityBySeverity(t *testing.T) {
	t.Run("Should add without error by analysis", func(t *testing.T) {
		v := &Vulnerability{}
		v.AddCountVulnerabilityBySeverity(6, severities.Critical, vulnerability.Vulnerability)
		v.AddCountVulnerabilityBySeverity(6, severities.Critical, vulnerability.FalsePositive)
		v.AddCountVulnerabilityBySeverity(6, severities.Critical, vulnerability.RiskAccepted)
		v.AddCountVulnerabilityBySeverity(6, severities.Critical, vulnerability.Corrected)
		v.AddCountVulnerabilityBySeverity(5, severities.High, vulnerability.Vulnerability)
		v.AddCountVulnerabilityBySeverity(5, severities.High, vulnerability.FalsePositive)
		v.AddCountVulnerabilityBySeverity(5, severities.High, vulnerability.RiskAccepted)
		v.AddCountVulnerabilityBySeverity(5, severities.High, vulnerability.Corrected)
		v.AddCountVulnerabilityBySeverity(4, severities.Medium, vulnerability.Vulnerability)
		v.AddCountVulnerabilityBySeverity(4, severities.Medium, vulnerability.FalsePositive)
		v.AddCountVulnerabilityBySeverity(4, severities.Medium, vulnerability.RiskAccepted)
		v.AddCountVulnerabilityBySeverity(4, severities.Medium, vulnerability.Corrected)
		v.AddCountVulnerabilityBySeverity(3, severities.Low, vulnerability.Vulnerability)
		v.AddCountVulnerabilityBySeverity(3, severities.Low, vulnerability.FalsePositive)
		v.AddCountVulnerabilityBySeverity(3, severities.Low, vulnerability.RiskAccepted)
		v.AddCountVulnerabilityBySeverity(3, severities.Low, vulnerability.Corrected)
		v.AddCountVulnerabilityBySeverity(2, severities.Info, vulnerability.Vulnerability)
		v.AddCountVulnerabilityBySeverity(2, severities.Info, vulnerability.FalsePositive)
		v.AddCountVulnerabilityBySeverity(2, severities.Info, vulnerability.RiskAccepted)
		v.AddCountVulnerabilityBySeverity(2, severities.Info, vulnerability.Corrected)
		v.AddCountVulnerabilityBySeverity(1, severities.Unknown, vulnerability.Vulnerability)
		v.AddCountVulnerabilityBySeverity(1, severities.Unknown, vulnerability.FalsePositive)
		v.AddCountVulnerabilityBySeverity(1, severities.Unknown, vulnerability.RiskAccepted)
		v.AddCountVulnerabilityBySeverity(1, severities.Unknown, vulnerability.Corrected)
		response := v.ToResponseSeverity()
		assert.Equal(t, 24, response.Critical.Count)
		assert.Equal(t, 20, response.High.Count)
		assert.Equal(t, 16, response.Medium.Count)
		assert.Equal(t, 12, response.Low.Count)
		assert.Equal(t, 8, response.Info.Count)
		assert.Equal(t, 4, response.Unknown.Count)
	})
}

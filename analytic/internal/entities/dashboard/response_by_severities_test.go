package dashboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumVulnerabilityCritical(t *testing.T) {
	t.Run("should success sum count and parse by severity", func(t *testing.T) {
		bySeverities := &BySeverities{
			Critical: &BySeverity{
				Types: &ByVulnerabilityTypes{},
			},
		}

		vuln := &Vulnerability{
			CriticalVulnerability: 5,
			CriticalRiskAccepted:  6,
			CriticalFalsePositive: 7,
			CriticalCorrected:     8,
		}

		bySeverities = bySeverities.SumVulnerabilityCritical(vuln)
		assert.Equal(t, 26, bySeverities.Critical.Count)
		assert.Equal(t, 5, bySeverities.Critical.Types.Vulnerability)
		assert.Equal(t, 6, bySeverities.Critical.Types.RiskAccepted)
		assert.Equal(t, 7, bySeverities.Critical.Types.FalsePositive)
		assert.Equal(t, 8, bySeverities.Critical.Types.Corrected)
	})
}

func TestSumVulnerabilityHigh(t *testing.T) {
	t.Run("should success sum count and parse by severity", func(t *testing.T) {
		bySeverities := &BySeverities{
			High: &BySeverity{
				Types: &ByVulnerabilityTypes{},
			},
		}

		vuln := &Vulnerability{
			HighVulnerability: 5,
			HighRiskAccepted:  6,
			HighFalsePositive: 7,
			HighCorrected:     8,
		}

		bySeverities = bySeverities.SumVulnerabilityHigh(vuln)
		assert.Equal(t, 26, bySeverities.High.Count)
		assert.Equal(t, 5, bySeverities.High.Types.Vulnerability)
		assert.Equal(t, 6, bySeverities.High.Types.RiskAccepted)
		assert.Equal(t, 7, bySeverities.High.Types.FalsePositive)
		assert.Equal(t, 8, bySeverities.High.Types.Corrected)
	})
}

func TestSumVulnerabilityMedium(t *testing.T) {
	t.Run("should success sum count and parse by severity", func(t *testing.T) {
		bySeverities := &BySeverities{
			Medium: &BySeverity{
				Types: &ByVulnerabilityTypes{},
			},
		}

		vuln := &Vulnerability{
			MediumVulnerability: 5,
			MediumRiskAccepted:  6,
			MediumFalsePositive: 7,
			MediumCorrected:     8,
		}

		bySeverities = bySeverities.SumVulnerabilityMedium(vuln)
		assert.Equal(t, 26, bySeverities.Medium.Count)
		assert.Equal(t, 5, bySeverities.Medium.Types.Vulnerability)
		assert.Equal(t, 6, bySeverities.Medium.Types.RiskAccepted)
		assert.Equal(t, 7, bySeverities.Medium.Types.FalsePositive)
		assert.Equal(t, 8, bySeverities.Medium.Types.Corrected)
	})
}

func TestSumVulnerabilityLow(t *testing.T) {
	t.Run("should success sum count and parse by severity", func(t *testing.T) {
		bySeverities := &BySeverities{
			Low: &BySeverity{
				Types: &ByVulnerabilityTypes{},
			},
		}

		vuln := &Vulnerability{
			LowVulnerability: 5,
			LowRiskAccepted:  6,
			LowFalsePositive: 7,
			LowCorrected:     8,
		}

		bySeverities = bySeverities.SumVulnerabilityLow(vuln)
		assert.Equal(t, 26, bySeverities.Low.Count)
		assert.Equal(t, 5, bySeverities.Low.Types.Vulnerability)
		assert.Equal(t, 6, bySeverities.Low.Types.RiskAccepted)
		assert.Equal(t, 7, bySeverities.Low.Types.FalsePositive)
		assert.Equal(t, 8, bySeverities.Low.Types.Corrected)
	})
}

func TestSumVulnerabilityUnknown(t *testing.T) {
	t.Run("should success sum count and parse by severity", func(t *testing.T) {
		bySeverities := &BySeverities{
			Unknown: &BySeverity{
				Types: &ByVulnerabilityTypes{},
			},
		}

		vuln := &Vulnerability{
			UnknownVulnerability: 5,
			UnknownRiskAccepted:  6,
			UnknownFalsePositive: 7,
			UnknownCorrected:     8,
		}

		bySeverities = bySeverities.SumVulnerabilityUnknown(vuln)
		assert.Equal(t, 26, bySeverities.Unknown.Count)
		assert.Equal(t, 5, bySeverities.Unknown.Types.Vulnerability)
		assert.Equal(t, 6, bySeverities.Unknown.Types.RiskAccepted)
		assert.Equal(t, 7, bySeverities.Unknown.Types.FalsePositive)
		assert.Equal(t, 8, bySeverities.Unknown.Types.Corrected)
	})
}

func TestSumVulnerabilityInfo(t *testing.T) {
	t.Run("should success sum count and parse by severity", func(t *testing.T) {
		bySeverities := &BySeverities{
			Info: &BySeverity{
				Types: &ByVulnerabilityTypes{},
			},
		}

		vuln := &Vulnerability{
			InfoVulnerability: 5,
			InfoRiskAccepted:  6,
			InfoFalsePositive: 7,
			InfoCorrected:     8,
		}

		bySeverities = bySeverities.SumVulnerabilityInfo(vuln)
		assert.Equal(t, 26, bySeverities.Info.Count)
		assert.Equal(t, 5, bySeverities.Info.Types.Vulnerability)
		assert.Equal(t, 6, bySeverities.Info.Types.RiskAccepted)
		assert.Equal(t, 7, bySeverities.Info.Types.FalsePositive)
		assert.Equal(t, 8, bySeverities.Info.Types.Corrected)
	})
}

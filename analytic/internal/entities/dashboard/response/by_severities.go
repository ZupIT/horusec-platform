package response

type BySeverities struct {
	Critical *BySeverity `json:"critical"`
	High     *BySeverity `json:"high"`
	Medium   *BySeverity `json:"medium"`
	Low      *BySeverity `json:"low"`
	Info     *BySeverity `json:"info"`
	Unknown  *BySeverity `json:"unknown"`
}

func (s *BySeverities) SumVulnerabilityCritical(vuln *Vulnerability) *BySeverities {
	s.Critical.Count = vuln.CriticalVulnerability + vuln.CriticalRiskAccepted +
		vuln.CriticalFalsePositive + vuln.CriticalCorrected

	s.Critical.Types.Vulnerability = vuln.CriticalVulnerability
	s.Critical.Types.RiskAccepted = vuln.CriticalRiskAccepted
	s.Critical.Types.FalsePositive = vuln.CriticalFalsePositive
	s.Critical.Types.Corrected = vuln.CriticalCorrected

	return s
}

func (s *BySeverities) SumVulnerabilityHigh(vuln *Vulnerability) *BySeverities {
	s.High.Count = vuln.HighVulnerability + vuln.HighRiskAccepted + vuln.HighFalsePositive + vuln.HighCorrected

	s.High.Types.Vulnerability = vuln.HighVulnerability
	s.High.Types.RiskAccepted = vuln.HighRiskAccepted
	s.High.Types.FalsePositive = vuln.HighFalsePositive
	s.High.Types.Corrected = vuln.HighCorrected

	return s
}

func (s *BySeverities) SumVulnerabilityMedium(vuln *Vulnerability) *BySeverities {
	s.Medium.Count = vuln.MediumVulnerability + vuln.MediumRiskAccepted +
		vuln.MediumFalsePositive + vuln.MediumCorrected

	s.Medium.Types.Vulnerability = vuln.MediumVulnerability
	s.Medium.Types.RiskAccepted = vuln.MediumRiskAccepted
	s.Medium.Types.FalsePositive = vuln.MediumFalsePositive
	s.Medium.Types.Corrected = vuln.MediumCorrected

	return s
}

func (s *BySeverities) SumVulnerabilityLow(vuln *Vulnerability) *BySeverities {
	s.Low.Count = vuln.LowVulnerability + vuln.LowRiskAccepted + vuln.LowFalsePositive + vuln.LowCorrected

	s.Low.Types.Vulnerability = vuln.LowVulnerability
	s.Low.Types.RiskAccepted = vuln.LowRiskAccepted
	s.Low.Types.FalsePositive = vuln.LowFalsePositive
	s.Low.Types.Corrected = vuln.LowCorrected

	return s
}

func (s *BySeverities) SumVulnerabilityInfo(vuln *Vulnerability) *BySeverities {
	s.Info.Count = vuln.InfoVulnerability + vuln.InfoRiskAccepted + vuln.InfoFalsePositive + vuln.InfoCorrected

	s.Info.Types.Vulnerability = vuln.InfoVulnerability
	s.Info.Types.RiskAccepted = vuln.InfoRiskAccepted
	s.Info.Types.FalsePositive = vuln.InfoFalsePositive
	s.Info.Types.Corrected = vuln.InfoCorrected

	return s
}

func (s *BySeverities) SumVulnerabilityUnknown(vuln *Vulnerability) *BySeverities {
	s.Unknown.Count = vuln.UnknownVulnerability + vuln.UnknownRiskAccepted +
		vuln.UnknownFalsePositive + vuln.UnknownCorrected

	s.Unknown.Types.Vulnerability = vuln.UnknownVulnerability
	s.Unknown.Types.RiskAccepted = vuln.UnknownRiskAccepted
	s.Unknown.Types.FalsePositive = vuln.UnknownFalsePositive
	s.Unknown.Types.Corrected = vuln.UnknownCorrected

	return s
}

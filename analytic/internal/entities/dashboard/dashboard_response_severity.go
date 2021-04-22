package dashboard

type ResponseVulnTypes struct {
	Vulnerability int `json:"vulnerability"`
	RiskAccepted  int `json:"riskAccepted"`
	FalsePositive int `json:"falsePositive"`
	Corrected     int `json:"corrected"`
}

type ResponseSeverityContAndTypes struct {
	Count int                `json:"count"`
	Types *ResponseVulnTypes `json:"types"`
}

type ResponseSeverity struct {
	Critical *ResponseSeverityContAndTypes `json:"critical"`
	High     *ResponseSeverityContAndTypes `json:"high"`
	Medium   *ResponseSeverityContAndTypes `json:"medium"`
	Low      *ResponseSeverityContAndTypes `json:"low"`
	Info     *ResponseSeverityContAndTypes `json:"info"`
}

func (r *ResponseSeverity) SumVulnerabilityCritical(vuln *Vulnerability) *ResponseSeverity {
	r.Critical.Count = vuln.CriticalVulnerability + vuln.CriticalRiskAccepted +
		vuln.CriticalFalsePositive + vuln.CriticalCorrected
	r.Critical.Types.Vulnerability = vuln.CriticalVulnerability
	r.Critical.Types.RiskAccepted = vuln.CriticalRiskAccepted
	r.Critical.Types.FalsePositive = vuln.CriticalFalsePositive
	r.Critical.Types.Corrected = vuln.CriticalCorrected
	return r
}

func (r *ResponseSeverity) SumVulnerabilityHigh(vuln *Vulnerability) *ResponseSeverity {
	r.High.Count = vuln.HighVulnerability + vuln.HighRiskAccepted +
		vuln.HighFalsePositive + vuln.HighCorrected
	r.High.Types.Vulnerability = vuln.HighVulnerability
	r.High.Types.RiskAccepted = vuln.HighRiskAccepted
	r.High.Types.FalsePositive = vuln.HighFalsePositive
	r.High.Types.Corrected = vuln.HighCorrected
	return r
}

func (r *ResponseSeverity) SumVulnerabilityMedium(vuln *Vulnerability) *ResponseSeverity {
	r.Medium.Count = vuln.MediumVulnerability + vuln.MediumRiskAccepted +
		vuln.MediumFalsePositive + vuln.MediumCorrected
	r.Medium.Types.Vulnerability = vuln.MediumVulnerability
	r.Medium.Types.RiskAccepted = vuln.MediumRiskAccepted
	r.Medium.Types.FalsePositive = vuln.MediumFalsePositive
	r.Medium.Types.Corrected = vuln.MediumCorrected
	return r
}

func (r *ResponseSeverity) SumVulnerabilityLow(vuln *Vulnerability) *ResponseSeverity {
	r.Low.Count = vuln.LowVulnerability + vuln.LowRiskAccepted +
		vuln.LowFalsePositive + vuln.LowCorrected
	r.Low.Types.Vulnerability = vuln.LowVulnerability
	r.Low.Types.RiskAccepted = vuln.LowRiskAccepted
	r.Low.Types.FalsePositive = vuln.LowFalsePositive
	r.Low.Types.Corrected = vuln.LowCorrected
	return r
}

func (r *ResponseSeverity) SumVulnerabilityInfo(vuln *Vulnerability) *ResponseSeverity {
	r.Info.Count = vuln.InfoVulnerability + vuln.InfoRiskAccepted +
		vuln.InfoFalsePositive + vuln.InfoCorrected
	r.Info.Types.Vulnerability = vuln.InfoVulnerability
	r.Info.Types.RiskAccepted = vuln.InfoRiskAccepted
	r.Info.Types.FalsePositive = vuln.InfoFalsePositive
	r.Info.Types.Corrected = vuln.InfoCorrected
	return r
}

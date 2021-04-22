package dashboard

type ResponseVulnTypes struct {
	Vulnerability int `json:"vulnerability"`
	RiskAccepted  int `json:"riskAccepted"`
	FalsePositive int `json:"falsePositive"`
	Corrected     int `json:"corrected"`
	Unknown       int `json:"unknown"`
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
	Unknown  *ResponseSeverityContAndTypes `json:"unknown"`
}

func (r *ResponseSeverity) SumVulnerabilityCritical(v *Vulnerability) *ResponseSeverity {
	r.Critical.Count = v.CriticalVulnerability + v.CriticalRiskAccepted + v.CriticalFalsePositive + v.CriticalCorrected + v.CriticalUnknown
	r.Critical.Types.Vulnerability = v.CriticalVulnerability
	r.Critical.Types.RiskAccepted = v.CriticalRiskAccepted
	r.Critical.Types.FalsePositive = v.CriticalFalsePositive
	r.Critical.Types.Corrected = v.CriticalCorrected
	r.Critical.Types.Unknown = v.CriticalUnknown
	return r
}

func (r *ResponseSeverity) SumVulnerabilityHigh(vuln *Vulnerability) *ResponseSeverity {
	r.High.Count = vuln.HighVulnerability + vuln.HighRiskAccepted + vuln.HighFalsePositive + vuln.HighCorrected + vuln.HighUnknown
	r.High.Types.Vulnerability = vuln.HighVulnerability
	r.High.Types.RiskAccepted = vuln.HighRiskAccepted
	r.High.Types.FalsePositive = vuln.HighFalsePositive
	r.High.Types.Corrected = vuln.HighCorrected
	r.High.Types.Unknown = vuln.HighUnknown
	return r
}

func (r *ResponseSeverity) SumVulnerabilityMedium(vuln *Vulnerability) *ResponseSeverity {
	r.Medium.Count = vuln.MediumVulnerability + vuln.MediumRiskAccepted + vuln.MediumFalsePositive + vuln.MediumCorrected + vuln.MediumUnknown
	r.Medium.Types.Vulnerability = vuln.MediumVulnerability
	r.Medium.Types.RiskAccepted = vuln.MediumRiskAccepted
	r.Medium.Types.FalsePositive = vuln.MediumFalsePositive
	r.Medium.Types.Corrected = vuln.MediumCorrected
	r.Medium.Types.Unknown = vuln.MediumUnknown
	return r
}

func (r *ResponseSeverity) SumVulnerabilityLow(vuln *Vulnerability) *ResponseSeverity {
	r.Low.Count = vuln.LowVulnerability + vuln.LowRiskAccepted + vuln.LowFalsePositive + vuln.LowCorrected + vuln.LowUnknown
	r.Low.Types.Vulnerability = vuln.LowVulnerability
	r.Low.Types.RiskAccepted = vuln.LowRiskAccepted
	r.Low.Types.FalsePositive = vuln.LowFalsePositive
	r.Low.Types.Corrected = vuln.LowCorrected
	r.Low.Types.Unknown = vuln.LowUnknown
	return r
}

func (r *ResponseSeverity) SumVulnerabilityInfo(vuln *Vulnerability) *ResponseSeverity {
	r.Info.Count = vuln.InfoVulnerability + vuln.InfoRiskAccepted + vuln.InfoFalsePositive + vuln.InfoCorrected + vuln.InfoUnknown
	r.Info.Types.Vulnerability = vuln.InfoVulnerability
	r.Info.Types.RiskAccepted = vuln.InfoRiskAccepted
	r.Info.Types.FalsePositive = vuln.InfoFalsePositive
	r.Info.Types.Corrected = vuln.InfoCorrected
	r.Info.Types.Unknown = vuln.InfoUnknown
	return r
}

func (r *ResponseSeverity) SumVulnerabilityUnknown(vuln *Vulnerability) *ResponseSeverity {
	r.Unknown.Count = vuln.UnknownVulnerability + vuln.UnknownRiskAccepted + vuln.UnknownFalsePositive + vuln.UnknownCorrected + vuln.UnknownUnknown
	r.Unknown.Types.Vulnerability = vuln.UnknownVulnerability
	r.Unknown.Types.RiskAccepted = vuln.UnknownRiskAccepted
	r.Unknown.Types.FalsePositive = vuln.UnknownFalsePositive
	r.Unknown.Types.Corrected = vuln.UnknownCorrected
	r.Unknown.Types.Unknown = vuln.UnknownUnknown
	return r
}

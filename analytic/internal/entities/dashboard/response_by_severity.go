package dashboard

type BySeverity struct {
	Count int                   `json:"count"`
	Types *ByVulnerabilityTypes `json:"types"`
}

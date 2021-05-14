package response

type Response struct {
	TotalAuthors                int            `json:"totalAuthors"`
	TotalRepositories           int            `json:"totalRepositories"`
	VulnerabilityBySeverity     *BySeverities  `json:"vulnerabilityBySeverity"`
	VulnerabilitiesByAuthor     []ByAuthor     `json:"vulnerabilitiesByAuthor"`
	VulnerabilitiesByRepository []ByRepository `json:"vulnerabilitiesByRepository"`
	VulnerabilitiesByLanguage   []ByLanguage   `json:"vulnerabilitiesByLanguage"`
	VulnerabilitiesByTime       []ByTime       `json:"vulnerabilityByTime"`
}

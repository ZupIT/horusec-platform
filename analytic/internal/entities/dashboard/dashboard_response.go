package dashboard

import (
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
)

type ResponseByAuthor struct {
	Author string `json:"author"`
	ResponseSeverity
}

type ResponseByRepository struct {
	RepositoryName string `json:"repositoryName"`
	ResponseSeverity
}

type ResponseByLanguage struct {
	Language languages.Language `json:"language"`
	ResponseSeverity
}

type ResponseByTime struct {
	Time time.Time `json:"time"`
	ResponseSeverity
}

type Response struct {
	TotalAuthors                int                    `json:"totalAuthors"`
	TotalRepositories           int                    `json:"totalRepositories"`
	VulnerabilityBySeverity     ResponseSeverity       `json:"vulnerabilityBySeverity"`
	VulnerabilitiesByAuthor     []ResponseByAuthor     `json:"vulnerabilitiesByAuthor"`
	VulnerabilitiesByRepository []ResponseByRepository `json:"vulnerabilitiesByRepository"`
	VulnerabilitiesByLanguage   []ResponseByLanguage   `json:"vulnerabilitiesByLanguage"`
	VulnerabilitiesByTime       []ResponseByTime       `json:"vulnerabilityByTime"`
}

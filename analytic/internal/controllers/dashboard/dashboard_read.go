package dashboard

import (
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
)

type IReadController interface {
	GetAllCharts(filter *dashboard.FilterDashboard) (*dashboard.Response, error)
}

type ControllerRead struct {
	repoDashboard repoDashboard.IRepoDashboard
}

func NewControllerDashboardRead(repositoryDashboard repoDashboard.IRepoDashboard) IReadController {
	return &ControllerRead{
		repoDashboard: repositoryDashboard,
	}
}

func (c *ControllerRead) GetAllCharts(filter *dashboard.FilterDashboard) (*dashboard.Response, error) {
	dashResponse := &dashboard.Response{}
	return c.getAllChartsFirst(filter, dashResponse)
}

func (c *ControllerRead) getAllChartsFirst(filter *dashboard.FilterDashboard,
	dashResponse *dashboard.Response) (*dashboard.Response, error) {
	totalAuthors, err := c.repoDashboard.GetDashboardTotalDevelopers(filter)
	if err != nil {
		return nil, err
	}
	dashResponse.TotalAuthors = totalAuthors

	totalRepositories, err := c.repoDashboard.GetDashboardTotalRepositories(filter)
	if err != nil {
		return nil, err
	}
	dashResponse.TotalRepositories = totalRepositories
	return c.getAllChartsSecond(filter, dashResponse)
}

func (c *ControllerRead) getAllChartsSecond(filter *dashboard.FilterDashboard,
	dashResponse *dashboard.Response) (*dashboard.Response, error) {
	vulnerabilityBySeverity, err := c.repoDashboard.GetDashboardVulnBySeverity(filter)
	if err != nil {
		return nil, err
	}
	dashResponse.VulnerabilityBySeverity = vulnerabilityBySeverity.ToResponseSeverity()

	vulnerabilitiesByAuthor, err := c.repoDashboard.GetDashboardVulnByAuthor(filter)
	if err != nil {
		return nil, err
	}
	dashResponse.VulnerabilitiesByAuthor = dashboard.ParseListVulnByAuthorToListResponse(vulnerabilitiesByAuthor)
	return c.getAllChartsThird(filter, dashResponse)
}

func (c *ControllerRead) getAllChartsThird(filter *dashboard.FilterDashboard,
	dashResponse *dashboard.Response) (*dashboard.Response, error) {
	vulnsByRepository, err := c.repoDashboard.GetDashboardVulnByRepository(filter)
	if err != nil {
		return nil, err
	}
	dashResponse.VulnerabilitiesByRepository = dashboard.ParseListVulnByRepositoryToListResponse(vulnsByRepository)
	vulnerabilitiesByLanguage, err := c.repoDashboard.GetDashboardVulnByLanguage(filter)
	if err != nil {
		return nil, err
	}
	dashResponse.VulnerabilitiesByLanguage = dashboard.ParseListVulnByLanguageToListResponse(vulnerabilitiesByLanguage)
	return c.getAllChartsFinal(filter, dashResponse)
}

func (c *ControllerRead) getAllChartsFinal(filter *dashboard.FilterDashboard,
	dashResponse *dashboard.Response) (*dashboard.Response, error) {
	vulnerabilitiesByTime, err := c.repoDashboard.GetDashboardVulnByTime(filter)
	if err != nil {
		return nil, err
	}
	dashResponse.VulnerabilitiesByTime = dashboard.ParseListVulnByTimeToListResponse(vulnerabilitiesByTime)
	return dashResponse, nil
}

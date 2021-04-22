package dashboard

import (
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
)

const (
	KeyTotalAuthors                = "totalAuthors"
	KeyTotalRepositories           = "totalRepositories"
	KeyVulnerabilityBySeverity     = "vulnerabilityBySeverity"
	KeyVulnerabilitiesByAuthor     = "vulnerabilitiesByAuthor"
	KeyVulnerabilitiesByRepository = "vulnerabilitiesByRepository"
	KeyVulnerabilitiesByLanguage   = "vulnerabilitiesByLanguage"
	KeyVulnerabilitiesByTime       = "vulnerabilitiesByTime"
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

// nolint:dupl // methods is not duplicate
func (c *ControllerRead) getAllChartsFirst(filter *dashboard.FilterDashboard,
	dashResponse *dashboard.Response) (*dashboard.Response, error) {
	res := c.repoDashboard.GetDashboardTotalDevelopers(filter)
	if res.GetErrorExceptNotFound() != nil {
		return nil, res.GetErrorExceptNotFound()
	}
	dashResponse.TotalAuthors = c.dataDashToTotalRepositoriesAndAuthors(res.GetData())

	res = c.repoDashboard.GetDashboardTotalRepositories(filter)
	if res.GetErrorExceptNotFound() != nil {
		return nil, res.GetErrorExceptNotFound()
	}
	dashResponse.TotalRepositories = c.dataDashToTotalRepositoriesAndAuthors(res.GetData())
	return c.getAllChartsSecond(filter, dashResponse)
}

// nolint:dupl // methods is not duplicate
func (c *ControllerRead) getAllChartsSecond(filter *dashboard.FilterDashboard,
	dashResponse *dashboard.Response) (*dashboard.Response, error) {
	res := c.repoDashboard.GetDashboardVulnBySeverity(filter)
	if res.GetErrorExceptNotFound() != nil {
		return nil, res.GetErrorExceptNotFound()
	}
	dashResponse.VulnerabilityBySeverity = c.dataDashToVulnBySeverity(res.GetData())
	res = c.repoDashboard.GetDashboardVulnByAuthor(filter)
	if res.GetErrorExceptNotFound() != nil {
		return nil, res.GetErrorExceptNotFound()
	}
	dashResponse.VulnerabilitiesByAuthor = c.dataDashToVulnByAuthor(res.GetData())
	return c.getAllChartsThird(filter, dashResponse)
}

// nolint:dupl // methods is not duplicate
func (c *ControllerRead) getAllChartsThird(filter *dashboard.FilterDashboard,
	dashResponse *dashboard.Response) (*dashboard.Response, error) {
	res := c.repoDashboard.GetDashboardVulnByRepository(filter)
	if res.GetErrorExceptNotFound() != nil {
		return nil, res.GetErrorExceptNotFound()
	}
	dashResponse.VulnerabilitiesByRepository = c.dataDashToVulnByRepository(res.GetData())
	res = c.repoDashboard.GetDashboardVulnByLanguage(filter)
	if res.GetErrorExceptNotFound() != nil {
		return nil, res.GetErrorExceptNotFound()
	}
	dashResponse.VulnerabilitiesByLanguage = c.dataDashToVulnByLanguage(res.GetData())
	return c.getAllChartsFinal(filter, dashResponse)
}

func (c *ControllerRead) getAllChartsFinal(filter *dashboard.FilterDashboard,
	dashResponse *dashboard.Response) (*dashboard.Response, error) {
	res := c.repoDashboard.GetDashboardVulnByTime(filter)
	if res.GetErrorExceptNotFound() != nil {
		return nil, res.GetErrorExceptNotFound()
	}
	dashResponse.VulnerabilitiesByTime = c.dataDashToVulnByTime(res.GetData())
	return dashResponse, nil
}

func (c *ControllerRead) dataDashToTotalRepositoriesAndAuthors(data interface{}) int {
	if data == nil {
		return 0
	}
	return data.(int)
}

func (c *ControllerRead) dataDashToVulnBySeverity(data interface{}) dashboard.ResponseSeverity {
	if data == nil {
		return dashboard.ResponseSeverity{}
	}
	return data.(*dashboard.VulnerabilitiesByTime).ToResponseSeverity()
}

func (c *ControllerRead) dataDashToVulnByAuthor(data interface{}) []dashboard.ResponseByAuthor {
	if data == nil {
		return []dashboard.ResponseByAuthor{}
	}
	return dashboard.ParseListVulnByAuthorToListResponse(data.([]*dashboard.VulnerabilitiesByAuthor))
}

func (c *ControllerRead) dataDashToVulnByRepository(data interface{}) []dashboard.ResponseByRepository {
	if data == nil {
		return []dashboard.ResponseByRepository{}
	}
	return dashboard.ParseListVulnByRepositoryToListResponse(data.([]*dashboard.VulnerabilitiesByRepository))
}

func (c *ControllerRead) dataDashToVulnByLanguage(data interface{}) []dashboard.ResponseByLanguage {
	if data == nil {
		return []dashboard.ResponseByLanguage{}
	}
	return dashboard.ParseListVulnByLanguageToListResponse(data.([]*dashboard.VulnerabilitiesByLanguage))
}

func (c *ControllerRead) dataDashToVulnByTime(data interface{}) []dashboard.ResponseByTime {
	if data == nil {
		return []dashboard.ResponseByTime{}
	}
	return dashboard.ParseListVulnByTimeToListResponse(data.([]*dashboard.VulnerabilitiesByTime))
}

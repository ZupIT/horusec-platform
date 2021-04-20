package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

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
	for key, dashFunc := range c.mapRepoDashboardFuncToDashboard() {
		res := dashFunc(filter)
		if res.GetErrorExceptNotFound() != nil {
			return nil, res.GetErrorExceptNotFound()
		}
		dashResponse = c.setDashDataByKey(key, dashResponse, res.GetData())
	}
	return dashResponse, nil
}

func (c *ControllerRead) mapRepoDashboardFuncToDashboard() map[string]func(*dashboard.FilterDashboard) response.IResponse {
	return map[string]func(*dashboard.FilterDashboard) response.IResponse{
		KeyTotalAuthors:                c.repoDashboard.GetDashboardTotalDevelopers,
		KeyTotalRepositories:           c.repoDashboard.GetDashboardTotalRepositories,
		KeyVulnerabilityBySeverity:     c.repoDashboard.GetDashboardVulnBySeverity,
		KeyVulnerabilitiesByAuthor:     c.repoDashboard.GetDashboardVulnByAuthor,
		KeyVulnerabilitiesByRepository: c.repoDashboard.GetDashboardVulnByRepository,
		KeyVulnerabilitiesByLanguage:   c.repoDashboard.GetDashboardVulnByLanguage,
		KeyVulnerabilitiesByTime:       c.repoDashboard.GetVulnByTime,
	}
}

func (c *ControllerRead) setDashDataByKey(key string, dash *dashboard.Response, data interface{}) *dashboard.Response {
	switch key {
	case KeyTotalAuthors:
		dash.TotalAuthors = data.(int)
	case KeyTotalRepositories:
		dash.TotalRepositories = data.(int)
	case KeyVulnerabilityBySeverity:
		dash.VulnerabilityBySeverity = data.(*dashboard.VulnerabilitiesByTime).ToResponseSeverity()
	case KeyVulnerabilitiesByAuthor:
		dash.VulnerabilitiesByAuthor = dashboard.ParseListVulnByAuthorToListResponse(*data.(*[]dashboard.VulnerabilitiesByAuthor))
	case KeyVulnerabilitiesByRepository:
		dash.VulnerabilitiesByRepository = dashboard.ParseListVulnByRepositoryToListResponse(*data.(*[]dashboard.VulnerabilitiesByRepository))
	case KeyVulnerabilitiesByLanguage:
		dash.VulnerabilitiesByLanguage = dashboard.ParseListVulnByLanguageToListResponse(*data.(*[]dashboard.VulnerabilitiesByLanguage))
	case KeyVulnerabilitiesByTime:
		dash.VulnerabilitiesByTime = dashboard.ParseListVulnByTimeToListResponse(*data.(*[]dashboard.VulnerabilitiesByTime))
	}
	return dash
}

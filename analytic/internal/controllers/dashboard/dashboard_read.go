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
	for key, dashFunc := range c.mapRepoDashFuncToDashboard() {
		res := dashFunc(filter)
		if res.GetErrorExceptNotFound() != nil {
			return nil, res.GetErrorExceptNotFound()
		}
		if data := res.GetData(); data != nil {
			dashResponse = c.setDashDataByKey(key, dashResponse, data)
		}
	}
	return dashResponse, nil
}

func (c *ControllerRead) mapRepoDashFuncToDashboard() map[string]func(*dashboard.FilterDashboard) response.IResponse {
	return map[string]func(*dashboard.FilterDashboard) response.IResponse{
		KeyTotalAuthors:                c.repoDashboard.GetDashboardTotalDevelopers,
		KeyTotalRepositories:           c.repoDashboard.GetDashboardTotalRepositories,
		KeyVulnerabilityBySeverity:     c.repoDashboard.GetDashboardVulnBySeverity,
		KeyVulnerabilitiesByAuthor:     c.repoDashboard.GetDashboardVulnByAuthor,
		KeyVulnerabilitiesByRepository: c.repoDashboard.GetDashboardVulnByRepository,
		KeyVulnerabilitiesByLanguage:   c.repoDashboard.GetDashboardVulnByLanguage,
		KeyVulnerabilitiesByTime:       c.repoDashboard.GetDashboardVulnByTime,
	}
}

// nolint:funlen,gocyclo // factory is necessary all switch case lines
func (c *ControllerRead) setDashDataByKey(key string, dash *dashboard.Response, data interface{}) *dashboard.Response {
	switch key {
	case KeyTotalAuthors:
		dash.TotalAuthors = data.(int)
	case KeyTotalRepositories:
		dash.TotalRepositories = data.(int)
	case KeyVulnerabilityBySeverity:
		dash.VulnerabilityBySeverity = c.dataDashToVulnBySeverity(data)
	case KeyVulnerabilitiesByAuthor:
		dash.VulnerabilitiesByAuthor = c.dataDashToVulnByAuthor(data)
	case KeyVulnerabilitiesByRepository:
		dash.VulnerabilitiesByRepository = c.dataDashToVulnByRepository(data)
	case KeyVulnerabilitiesByLanguage:
		dash.VulnerabilitiesByLanguage = c.dataDashToVulnByLanguage(data)
	case KeyVulnerabilitiesByTime:
		dash.VulnerabilitiesByTime = c.dataDashToVulnByTime(data)
	}
	return dash
}

func (c *ControllerRead) dataDashToVulnBySeverity(data interface{}) dashboard.ResponseSeverity {
	return data.(*dashboard.VulnerabilitiesByTime).ToResponseSeverity()
}

func (c *ControllerRead) dataDashToVulnByAuthor(data interface{}) []dashboard.ResponseByAuthor {
	return dashboard.ParseListVulnByAuthorToListResponse(*data.(*[]dashboard.VulnerabilitiesByAuthor))
}

func (c *ControllerRead) dataDashToVulnByRepository(data interface{}) []dashboard.ResponseByRepository {
	return dashboard.ParseListVulnByRepositoryToListResponse(*data.(*[]dashboard.VulnerabilitiesByRepository))
}

func (c *ControllerRead) dataDashToVulnByLanguage(data interface{}) []dashboard.ResponseByLanguage {
	return dashboard.ParseListVulnByLanguageToListResponse(*data.(*[]dashboard.VulnerabilitiesByLanguage))
}

func (c *ControllerRead) dataDashToVulnByTime(data interface{}) []dashboard.ResponseByTime {
	return dashboard.ParseListVulnByTimeToListResponse(*data.(*[]dashboard.VulnerabilitiesByTime))
}

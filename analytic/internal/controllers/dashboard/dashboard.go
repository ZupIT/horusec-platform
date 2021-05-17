package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"

	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	dashboardUseCases "github.com/ZupIT/horusec-platform/analytic/internal/usecases/dashboard"
)

type IController interface {
	GetAllDashboardCharts(filter *dashboard.Filter) (*dashboard.Response, error)
	AddVulnerabilitiesByAuthor(entity *analysis.Analysis) error
	AddVulnerabilitiesByRepository(entity *analysis.Analysis) error
	AddVulnerabilitiesByLanguage(entity *analysis.Analysis) error
	AddVulnerabilitiesByTime(entity *analysis.Analysis) error
}

type Controller struct {
	repository    repoDashboard.IRepoDashboard
	useCases      dashboardUseCases.IUseCases
	databaseWrite database.IDatabaseWrite
}

func NewControllerDashboardRead(repository repoDashboard.IRepoDashboard,
	connection *database.Connection, useCases dashboardUseCases.IUseCases) IController {
	return &Controller{
		repository:    repository,
		databaseWrite: connection.Write,
		useCases:      useCases,
	}
}

func (c *Controller) AddVulnerabilitiesByAuthor(entity *analysis.Analysis) error {
	vulnsByAuthor := c.useCases.ParseAnalysisToVulnerabilitiesByAuthor(entity)

	return c.databaseWrite.Create(vulnsByAuthor, dashboardEnums.TableVulnerabilitiesByAuthor).GetError()
}

func (c *Controller) AddVulnerabilitiesByRepository(entity *analysis.Analysis) error {
	vulnsByRepository := c.useCases.ParseAnalysisToVulnerabilitiesByRepository(entity)

	return c.databaseWrite.Create(vulnsByRepository, dashboardEnums.TableVulnerabilitiesByRepository).GetError()
}

func (c *Controller) AddVulnerabilitiesByLanguage(entity *analysis.Analysis) error {
	vulnsByLanguage := c.useCases.ParseAnalysisToVulnerabilitiesByLanguage(entity)

	return c.databaseWrite.Create(vulnsByLanguage, dashboardEnums.TableVulnerabilitiesByLanguage).GetError()
}

func (c *Controller) AddVulnerabilitiesByTime(entity *analysis.Analysis) error {
	vulnsByTime := c.useCases.ParseAnalysisToVulnerabilitiesByTime(entity)

	return c.databaseWrite.Create(vulnsByTime, dashboardEnums.TableVulnerabilitiesByTime).GetError()
}

func (c *Controller) GetAllDashboardCharts(filter *dashboard.Filter) (*dashboard.Response, error) {
	response := &dashboard.Response{}

	if err := response.SetTotalAuthors(c.repository.GetDashboardTotalDevelopers(filter)); err != nil {
		return nil, err
	}

	if err := response.SetTotalRepositories(c.repository.GetDashboardTotalRepositories(filter)); err != nil {
		return nil, err
	}

	return c.getChartsBySeverityAndAuthor(filter, response)
}

func (c *Controller) getChartsBySeverityAndAuthor(filter *dashboard.Filter,
	response *dashboard.Response) (*dashboard.Response, error) {
	if err := response.SetChartBySeverity(c.repository.GetDashboardVulnBySeverity(filter)); err != nil {
		return nil, err
	}

	if err := response.SetChartByAuthor(c.repository.GetDashboardVulnByAuthor(filter)); err != nil {
		return nil, err
	}

	return c.getChartsByRepositoryAndLanguage(filter, response)
}

func (c *Controller) getChartsByRepositoryAndLanguage(filter *dashboard.Filter,
	response *dashboard.Response) (*dashboard.Response, error) {
	if err := response.SetChartByRepository(c.repository.GetDashboardVulnByRepository(filter)); err != nil {
		return nil, err
	}

	if err := response.SetChartByLanguage(c.repository.GetDashboardVulnByLanguage(filter)); err != nil {
		return nil, err
	}

	return c.getChartByTime(filter, response)
}

func (c *Controller) getChartByTime(filter *dashboard.Filter,
	response *dashboard.Response) (*dashboard.Response, error) {
	if err := response.SetChartByTime(c.repository.GetDashboardVulnByTime(filter)); err != nil {
		return nil, err
	}

	return response, nil
}

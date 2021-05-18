package dashboard

import (
	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	dashboardUseCases "github.com/ZupIT/horusec-platform/analytic/internal/usecases/dashboard"
)

type IController interface {
	GetAllDashboardCharts(filter *dashboard.Filter) (*dashboard.Response, error)
	AddVulnerabilitiesByAuthor(entity *analysisEntities.Analysis) error
	AddVulnerabilitiesByRepository(entity *analysisEntities.Analysis) error
	AddVulnerabilitiesByLanguage(entity *analysisEntities.Analysis) error
	AddVulnerabilitiesByTime(entity *analysisEntities.Analysis) error
}

type Controller struct {
	repository    repoDashboard.IRepoDashboard
	useCases      dashboardUseCases.IUseCases
	databaseWrite database.IDatabaseWrite
}

func NewDashboardController(repository repoDashboard.IRepoDashboard,
	connection *database.Connection, useCases dashboardUseCases.IUseCases) IController {
	return &Controller{
		repository:    repository,
		databaseWrite: connection.Write,
		useCases:      useCases,
	}
}

func (c *Controller) AddVulnerabilitiesByAuthor(analysis *analysisEntities.Analysis) error {
	return c.databaseWrite.Create(c.useCases.ParseAnalysisToVulnerabilitiesByAuthor(analysis),
		dashboardEnums.TableVulnerabilitiesByAuthor).GetError()
}

func (c *Controller) AddVulnerabilitiesByRepository(analysis *analysisEntities.Analysis) error {
	return c.databaseWrite.Create(c.useCases.ParseAnalysisToVulnerabilitiesByRepository(analysis),
		dashboardEnums.TableVulnerabilitiesByRepository).GetError()
}

func (c *Controller) AddVulnerabilitiesByLanguage(analysis *analysisEntities.Analysis) error {
	return c.databaseWrite.Create(c.useCases.ParseAnalysisToVulnerabilitiesByLanguage(analysis),
		dashboardEnums.TableVulnerabilitiesByLanguage).GetError()
}

func (c *Controller) AddVulnerabilitiesByTime(analysis *analysisEntities.Analysis) error {
	return c.databaseWrite.Create(c.useCases.ParseAnalysisToVulnerabilitiesByTime(analysis),
		dashboardEnums.TableVulnerabilitiesByTime).GetError()
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

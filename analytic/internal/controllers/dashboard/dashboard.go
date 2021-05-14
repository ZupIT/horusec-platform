package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	repositoryEntities "github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/repository"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"
	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	dashboardUseCases "github.com/ZupIT/horusec-platform/analytic/internal/usecases/dashboard"
)

type IController interface {
	GetAllDashboardCharts(filter *repositoryEntities.Filter) (*response.Response, error)
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

func (c *Controller) GetAllDashboardCharts(filters *repositoryEntities.Filter) (*response.Response, error) {
	dashboardResponse := &response.Response{}

	return c.getTotalDevelopersAndRepositories(filters, dashboardResponse)
}

func (c *Controller) getTotalDevelopersAndRepositories(filter *repositoryEntities.Filter,
	dashboardResponse *response.Response) (*response.Response, error) {
	totalAuthors, err := c.repository.GetDashboardTotalDevelopers(filter)
	if err != nil {
		return nil, err
	}

	totalRepositories, err := c.repository.GetDashboardTotalRepositories(filter)
	if err != nil {
		return nil, err
	}

	dashboardResponse.TotalAuthors = totalAuthors
	dashboardResponse.TotalRepositories = totalRepositories
	return c.getAllChartsSecond(filter, dashboardResponse)
}

func (c *Controller) getAllChartsSecond(filter *repositoryEntities.Filter,
	dashboardResponse *response.Response) (*response.Response, error) {
	vulnerabilityBySeverity, err := c.repository.GetDashboardVulnBySeverity(filter)
	if err != nil {
		return nil, err
	}

	vulnerabilitiesByAuthor, err := c.repository.GetDashboardVulnByAuthor(filter)
	if err != nil {
		return nil, err
	}

	dashboardResponse.VulnerabilityBySeverity = vulnerabilityBySeverity.ToResponseBySeverities()
	dashboardResponse.VulnerabilitiesByAuthor = repositoryEntities.ParseListVulnByAuthorToListResponse(vulnerabilitiesByAuthor)
	return c.getAllChartsThird(filter, dashboardResponse)
}

func (c *Controller) getAllChartsThird(filter *repositoryEntities.Filter,
	dashboardResponse *response.Response) (*response.Response, error) {
	vulnsByRepository, err := c.repository.GetDashboardVulnByRepository(filter)
	if err != nil {
		return nil, err
	}

	vulnerabilitiesByLanguage, err := c.repository.GetDashboardVulnByLanguage(filter)
	if err != nil {
		return nil, err
	}

	dashboardResponse.VulnerabilitiesByRepository = repositoryEntities.ParseListVulnByRepositoryToListResponse(vulnsByRepository)
	dashboardResponse.VulnerabilitiesByLanguage = repositoryEntities.ParseListVulnByLanguageToListResponse(vulnerabilitiesByLanguage)
	return c.getAllChartsFinal(filter, dashboardResponse)
}

func (c *Controller) getAllChartsFinal(filter *repositoryEntities.Filter,
	dashboardResponse *response.Response) (*response.Response, error) {
	vulnerabilitiesByTime, err := c.repository.GetDashboardVulnByTime(filter)
	if err != nil {
		return nil, err
	}

	dashboardResponse.VulnerabilitiesByTime = repositoryEntities.ParseListVulnByTimeToListResponse(vulnerabilitiesByTime)
	return dashboardResponse, nil
}

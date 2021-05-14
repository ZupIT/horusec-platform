package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	database2 "github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/repositories"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"
	dashboardfilter "github.com/ZupIT/horusec-platform/analytic/internal/usecase/dashboard"

	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
)

type IController interface {
	GetAllDashboardCharts(filter *database2.Filter) (*response.Response, error)
	AddVulnerabilitiesByAuthor(entity *analysis.Analysis) error
	AddVulnerabilitiesByRepository(entity *analysis.Analysis) error
	AddVulnerabilitiesByLanguage(entity *analysis.Analysis) error
	AddVulnerabilitiesByTime(entity *analysis.Analysis) error
}

type Controller struct {
	repository    repoDashboard.IRepoDashboard
	useCases      dashboardfilter.IUseCases
	databaseWrite database.IDatabaseWrite
}

func NewControllerDashboardRead(repository repoDashboard.IRepoDashboard,
	connection *database.Connection, useCases dashboardfilter.IUseCases) IController {
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

func (c *Controller) GetAllDashboardCharts(filters *database2.Filter) (*response.Response, error) {
	dashResponse := &response.Response{}

	return c.getTotalDevelopersAndRepositories(filters, dashResponse)
}

func (c *Controller) getTotalDevelopersAndRepositories(filter *database2.Filter,
	dashResponse *response.Response) (*response.Response, error) {
	totalAuthors, err := c.repository.GetDashboardTotalDevelopers(filter)
	if err != nil {
		return nil, err
	}

	totalRepositories, err := c.repository.GetDashboardTotalRepositories(filter)
	if err != nil {
		return nil, err
	}

	dashResponse.TotalAuthors = totalAuthors
	dashResponse.TotalRepositories = totalRepositories
	return c.getAllChartsSecond(filter, dashResponse)
}

func (c *Controller) getAllChartsSecond(filter *database2.Filter,
	dashResponse *response.Response) (*response.Response, error) {
	vulnerabilityBySeverity, err := c.repository.GetDashboardVulnBySeverity(filter)
	if err != nil {
		return nil, err
	}

	vulnerabilitiesByAuthor, err := c.repository.GetDashboardVulnByAuthor(filter)
	if err != nil {
		return nil, err
	}

	dashResponse.VulnerabilityBySeverity = vulnerabilityBySeverity.ToResponseBySeverities()
	dashResponse.VulnerabilitiesByAuthor = database2.ParseListVulnByAuthorToListResponse(vulnerabilitiesByAuthor)
	return c.getAllChartsThird(filter, dashResponse)
}

func (c *Controller) getAllChartsThird(filter *database2.Filter,
	dashResponse *response.Response) (*response.Response, error) {
	vulnsByRepository, err := c.repository.GetDashboardVulnByRepository(filter)
	if err != nil {
		return nil, err
	}

	vulnerabilitiesByLanguage, err := c.repository.GetDashboardVulnByLanguage(filter)
	if err != nil {
		return nil, err
	}

	dashResponse.VulnerabilitiesByRepository = database2.ParseListVulnByRepositoryToListResponse(vulnsByRepository)
	dashResponse.VulnerabilitiesByLanguage = database2.ParseListVulnByLanguageToListResponse(vulnerabilitiesByLanguage)
	return c.getAllChartsFinal(filter, dashResponse)
}

func (c *Controller) getAllChartsFinal(filter *database2.Filter,
	dashResponse *response.Response) (*response.Response, error) {
	vulnerabilitiesByTime, err := c.repository.GetDashboardVulnByTime(filter)
	if err != nil {
		return nil, err
	}

	dashResponse.VulnerabilitiesByTime = database2.ParseListVulnByTimeToListResponse(vulnerabilitiesByTime)
	return dashResponse, nil
}

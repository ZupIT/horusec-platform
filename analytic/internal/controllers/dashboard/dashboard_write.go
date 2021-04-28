package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"

	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	useCaseAnalysis "github.com/ZupIT/horusec-platform/analytic/internal/usecase/analysis"
)

type IWriteController interface {
	AddVulnerabilitiesByAuthor(entity *analysis.Analysis) error
	AddVulnerabilitiesByRepository(entity *analysis.Analysis) error
	AddVulnerabilitiesByLanguage(entity *analysis.Analysis) error
	AddVulnerabilitiesByTime(entity *analysis.Analysis) error
}

type ControllerWrite struct {
	repoDashboard repoDashboard.IRepoDashboard
	useCase       useCaseAnalysis.IUseCase
}

func NewControllerDashboardWrite(repositoryDashboard repoDashboard.IRepoDashboard) IWriteController {
	return &ControllerWrite{
		repoDashboard: repositoryDashboard,
		useCase:       useCaseAnalysis.NewUseCaseAnalysis(),
	}
}

func (c *ControllerWrite) AddVulnerabilitiesByAuthor(entity *analysis.Analysis) error {
	tableName := (&dashboard.VulnerabilitiesByAuthor{}).GetTable()
	if err := c.inactiveVulnerabilities(entity.RepositoryID, tableName); err != nil {
		return err
	}
	vulnsByAuthor := c.useCase.ParseAnalysisToVulnerabilitiesByAuthor(entity)
	for index := range vulnsByAuthor {
		vuln := &vulnsByAuthor[index]
		if err := c.repoDashboard.Save(vuln, tableName); err != nil {
			return err
		}
	}
	return nil
}
func (c *ControllerWrite) AddVulnerabilitiesByRepository(entity *analysis.Analysis) error {
	tableName := (&dashboard.VulnerabilitiesByRepository{}).GetTable()
	if err := c.inactiveVulnerabilities(entity.RepositoryID, tableName); err != nil {
		return err
	}
	vulnsByRepository := c.useCase.ParseAnalysisToVulnerabilitiesByRepository(entity)
	for index := range vulnsByRepository {
		vuln := &vulnsByRepository[index]
		if err := c.repoDashboard.Save(vuln, tableName); err != nil {
			return err
		}
	}
	return nil
}
func (c *ControllerWrite) AddVulnerabilitiesByLanguage(entity *analysis.Analysis) error {
	tableName := (&dashboard.VulnerabilitiesByLanguage{}).GetTable()
	if err := c.inactiveVulnerabilities(entity.RepositoryID, tableName); err != nil {
		return err
	}
	vulnsByLanguage := c.useCase.ParseAnalysisToVulnerabilitiesByLanguage(entity)
	for index := range vulnsByLanguage {
		vuln := &vulnsByLanguage[index]
		if err := c.repoDashboard.Save(vuln, tableName); err != nil {
			return err
		}
	}
	return nil
}
func (c *ControllerWrite) AddVulnerabilitiesByTime(entity *analysis.Analysis) error {
	tableName := (&dashboard.VulnerabilitiesByTime{}).GetTable()
	vulnsByTime := c.useCase.ParseAnalysisToVulnerabilitiesByTime(entity)
	for index := range vulnsByTime {
		vuln := vulnsByTime[index]
		if err := c.repoDashboard.Save(&vuln, tableName); err != nil {
			return err
		}
	}
	return nil
}

func (c *ControllerWrite) inactiveVulnerabilities(repositoryID uuid.UUID, tableName string) error {
	conditionToUpdate := map[string]interface{}{
		"active":        true,
		"repository_id": repositoryID,
	}
	return c.repoDashboard.Inactive(conditionToUpdate, tableName)
}

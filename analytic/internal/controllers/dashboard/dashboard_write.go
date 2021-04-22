package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"

	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
	useCaseAnalysis "github.com/ZupIT/horusec-platform/analytic/internal/usecase/analysis"
)

type IWriteController interface {
	AddNewAnalysis(analysis *analysis.Analysis) error
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

func (c *ControllerWrite) AddNewAnalysis(entity *analysis.Analysis) error {
	if err := c.addVulnerabilitiesByAuthor(entity); err != nil {
		return err
	}
	if err := c.addVulnerabilitiesByRepository(entity); err != nil {
		return err
	}
	if err := c.addVulnerabilitiesByLanguage(entity); err != nil {
		return err
	}
	if err := c.addVulnerabilitiesByTime(entity); err != nil {
		return err
	}
	return nil
}
func (c *ControllerWrite) addVulnerabilitiesByAuthor(entity *analysis.Analysis) error {
	vulnsByAuthor := c.useCase.ParseAnalysisToVulnerabilitiesByAuthor(entity)
	for _, vuln := range vulnsByAuthor {
		return c.repoDashboard.SaveNewVulnByEntity(vuln, (&vuln).GetTable())
	}
	return nil
}
func (c *ControllerWrite) addVulnerabilitiesByRepository(entity *analysis.Analysis) error {
	vulnsByRepository := c.useCase.ParseAnalysisToVulnerabilitiesByRepository(entity)
	for _, vuln := range vulnsByRepository {
		return c.repoDashboard.SaveNewVulnByEntity(vuln, (&vuln).GetTable())
	}
	return nil
}
func (c *ControllerWrite) addVulnerabilitiesByLanguage(entity *analysis.Analysis) error {
	vulnsByLanguage := c.useCase.ParseAnalysisToVulnerabilitiesByLanguage(entity)
	for _, vuln := range vulnsByLanguage {
		return c.repoDashboard.SaveNewVulnByEntity(vuln, (&vuln).GetTable())
	}
	return nil
}
func (c *ControllerWrite) addVulnerabilitiesByTime(entity *analysis.Analysis) error {
	vulnsByTime := c.useCase.ParseAnalysisToVulnerabilitiesByTime(entity)
	for _, vuln := range vulnsByTime {
		return c.repoDashboard.SaveNewVulnByEntity(vuln, (&vuln).GetTable())
	}
	return nil
}

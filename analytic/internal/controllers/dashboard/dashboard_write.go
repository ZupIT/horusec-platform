package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"

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
	tableName := (&dashboard.VulnerabilitiesByAuthor{}).GetTable()
	vulnsByAuthor := c.useCase.ParseAnalysisToVulnerabilitiesByAuthor(entity)
	for index := range vulnsByAuthor {
		vuln := &vulnsByAuthor[index]
		if index == 0 {
			conditionToUpdate := map[string]interface{}{
				"active": true,
				"repository_id": vuln.RepositoryID,
			}
			if err := c.repoDashboard.Update(map[string]interface{}{"active": false}, conditionToUpdate, tableName); err != nil {
				return err
			}
		}
		if err := c.repoDashboard.Save(vuln, tableName); err != nil {
			return err
		}
	}
	return nil
}
func (c *ControllerWrite) addVulnerabilitiesByRepository(entity *analysis.Analysis) error {
	tableName := (&dashboard.VulnerabilitiesByRepository{}).GetTable()
	vulnsByRepository := c.useCase.ParseAnalysisToVulnerabilitiesByRepository(entity)
	for index := range vulnsByRepository {
		vuln := &vulnsByRepository[index]
		if index == 0 {
			conditionToUpdate := map[string]interface{}{
				"active": true,
				"repository_id": vuln.RepositoryID,
			}
			if err := c.repoDashboard.Update(map[string]interface{}{"active": false}, conditionToUpdate, tableName); err != nil {
				return err
			}
		}
		if err := c.repoDashboard.Save(vuln, tableName); err != nil {
			return err
		}
	}
	return nil
}
func (c *ControllerWrite) addVulnerabilitiesByLanguage(entity *analysis.Analysis) error {
	tableName := (&dashboard.VulnerabilitiesByLanguage{}).GetTable()
	vulnsByLanguage := c.useCase.ParseAnalysisToVulnerabilitiesByLanguage(entity)
	for index := range vulnsByLanguage {
		vuln := &vulnsByLanguage[index]
		if index == 0 {
			conditionToUpdate := map[string]interface{}{
				"active": true,
				"repository_id": vuln.RepositoryID,
			}
			if err := c.repoDashboard.Update(map[string]interface{}{"active": false}, conditionToUpdate, tableName); err != nil {
				return err
			}
		}
		if err := c.repoDashboard.Save(vuln, tableName); err != nil {
			return err
		}
	}
	return nil
}
func (c *ControllerWrite) addVulnerabilitiesByTime(entity *analysis.Analysis) error {
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

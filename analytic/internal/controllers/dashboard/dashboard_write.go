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
	useCase useCaseAnalysis.IUseCase
}

func NewControllerDashboardWrite(repositoryDashboard repoDashboard.IRepoDashboard) IWriteController {
	return &ControllerWrite{
		repoDashboard: repositoryDashboard,
		useCase: useCaseAnalysis.NewUseCaseAnalysis(),
	}
}

func (c *ControllerWrite) AddNewAnalysis(analysis *analysis.Analysis) error {
	//vulnByAuthor := c.useCase.ParseAnalysisToVulnerabilitiesByAuthor(analysis)
	//vulnByRepository := c.useCase.ParseAnalysisToVulnerabilitiesByRepository(analysis)
	//vulnByLanguage := c.useCase.ParseAnalysisToVulnerabilitiesByLanguage(analysis)
	//vulnByTime := c.useCase.ParseAnalysisToVulnerabilitiesByTime(analysis)
	return nil
}

package analysis

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) ParseAnalysisToVulnerabilitiesByAuthor(_ *analysis.Analysis) []dashboard.VulnerabilitiesByAuthor {
	args := m.MethodCalled("ParseAnalysisToVulnerabilitiesByAuthor")
	return args.Get(0).([]dashboard.VulnerabilitiesByAuthor)
}
func (m *Mock) ParseAnalysisToVulnerabilitiesByRepository(_ *analysis.Analysis) []dashboard.VulnerabilitiesByRepository {
	args := m.MethodCalled("ParseAnalysisToVulnerabilitiesByRepository")
	return args.Get(0).([]dashboard.VulnerabilitiesByRepository)
}
func (m *Mock) ParseAnalysisToVulnerabilitiesByLanguage(_ *analysis.Analysis) []dashboard.VulnerabilitiesByLanguage {
	args := m.MethodCalled("ParseAnalysisToVulnerabilitiesByLanguage")
	return args.Get(0).([]dashboard.VulnerabilitiesByLanguage)
}
func (m *Mock) ParseAnalysisToVulnerabilitiesByTime(_ *analysis.Analysis) []dashboard.VulnerabilitiesByTime {
	args := m.MethodCalled("ParseAnalysisToVulnerabilitiesByTime")
	return args.Get(0).([]dashboard.VulnerabilitiesByTime)
}

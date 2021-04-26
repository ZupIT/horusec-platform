package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetAllCharts(_ *dashboard.FilterDashboard) (*dashboard.Response, error) {
	args := m.MethodCalled("GetAllCharts")
	return args.Get(0).(*dashboard.Response), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) AddVulnerabilitiesByAuthor(_ *analysis.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByAuthor")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) AddVulnerabilitiesByRepository(_ *analysis.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByRepository")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) AddVulnerabilitiesByLanguage(_ *analysis.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByLanguage")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) AddVulnerabilitiesByTime(_ *analysis.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByTime")
	return utilsMock.ReturnNilOrError(args, 0)
}

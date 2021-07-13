package dashboard

import (
	"github.com/stretchr/testify/mock"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetAllDashboardChartsWorkspace(_ *dashboard.Filter) (*dashboard.Response, error) {
	args := m.MethodCalled("GetAllDashboardChartsWorkspace")
	return args.Get(0).(*dashboard.Response), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAllDashboardChartsRepository(_ *dashboard.Filter) (*dashboard.Response, error) {
	args := m.MethodCalled("GetAllDashboardChartsRepository")
	return args.Get(0).(*dashboard.Response), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) AddVulnerabilitiesByAuthor(_ *analysisEntities.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByAuthor")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) AddVulnerabilitiesByRepository(_ *analysisEntities.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByRepository")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) AddVulnerabilitiesByLanguage(_ *analysisEntities.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByLanguage")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) AddVulnerabilitiesByTime(_ *analysisEntities.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByTime")
	return utilsMock.ReturnNilOrError(args, 0)
}

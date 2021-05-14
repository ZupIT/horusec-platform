package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/database"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetAllDashboardCharts(_ *database.Filter) (*response.Response, error) {
	args := m.MethodCalled("GetAllDashboardCharts")
	return args.Get(0).(*response.Response), utilsMock.ReturnNilOrError(args, 1)
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

package dashboard

import (
	"github.com/stretchr/testify/mock"

	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetDashboardTotalDevelopers(_ *dashboard.Filter) (int, error) {
	args := m.MethodCalled("GetDashboardTotalDevelopers")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) GetDashboardTotalRepositories(_ *dashboard.Filter) (int, error) {
	args := m.MethodCalled("GetDashboardTotalRepositories")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) GetDashboardVulnBySeverity(_ *dashboard.Filter) (*dashboard.Vulnerability, error) {
	args := m.MethodCalled("GetDashboardVulnBySeverity")
	return args.Get(0).(*dashboard.Vulnerability), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) GetDashboardVulnByAuthor(_ *dashboard.Filter) ([]*dashboard.VulnerabilitiesByAuthor, error) {
	args := m.MethodCalled("GetDashboardVulnByAuthor")
	return args.Get(0).([]*dashboard.VulnerabilitiesByAuthor), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) GetDashboardVulnByRepository(_ *dashboard.Filter) ([]*dashboard.VulnerabilitiesByRepository, error) {
	args := m.MethodCalled("GetDashboardVulnByRepository")
	return args.Get(0).([]*dashboard.VulnerabilitiesByRepository), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) GetDashboardVulnByLanguage(_ *dashboard.Filter) ([]*dashboard.VulnerabilitiesByLanguage, error) {
	args := m.MethodCalled("GetDashboardVulnByLanguage")
	return args.Get(0).([]*dashboard.VulnerabilitiesByLanguage), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) GetDashboardVulnByTime(_ *dashboard.Filter) ([]*dashboard.VulnerabilitiesByTime, error) {
	args := m.MethodCalled("GetDashboardVulnByTime")
	return args.Get(0).([]*dashboard.VulnerabilitiesByTime), utilsMock.ReturnNilOrError(args, 1)
}

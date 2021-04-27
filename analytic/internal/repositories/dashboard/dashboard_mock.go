package dashboard

import (
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Save(_ interface{}, _ string) error {
	args := m.MethodCalled("Save")
	return utilsMock.ReturnNilOrError(args, 0)
}
func (m *Mock) Inactive(_ map[string]interface{}, _ string) error {
	args := m.MethodCalled("Inactive")
	return utilsMock.ReturnNilOrError(args, 0)
}
func (m *Mock) GetDashboardTotalDevelopers(_ *dashboard.FilterDashboard) (int, error) {
	args := m.MethodCalled("GetDashboardTotalDevelopers")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardTotalRepositories(_ *dashboard.FilterDashboard) (int, error) {
	args := m.MethodCalled("GetDashboardTotalRepositories")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnBySeverity(_ *dashboard.FilterDashboard) (*dashboard.Vulnerability, error) {
	args := m.MethodCalled("GetDashboardVulnBySeverity")
	return args.Get(0).(*dashboard.Vulnerability), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByAuthor(_ *dashboard.FilterDashboard) ([]*dashboard.VulnerabilitiesByAuthor, error) {
	args := m.MethodCalled("GetDashboardVulnByAuthor")
	return args.Get(0).([]*dashboard.VulnerabilitiesByAuthor), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByRepository(_ *dashboard.FilterDashboard) ([]*dashboard.VulnerabilitiesByRepository, error) {
	args := m.MethodCalled("GetDashboardVulnByRepository")
	return args.Get(0).([]*dashboard.VulnerabilitiesByRepository), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByLanguage(_ *dashboard.FilterDashboard) ([]*dashboard.VulnerabilitiesByLanguage, error) {
	args := m.MethodCalled("GetDashboardVulnByLanguage")
	return args.Get(0).([]*dashboard.VulnerabilitiesByLanguage), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByTime(_ *dashboard.FilterDashboard) ([]*dashboard.VulnerabilitiesByTime, error) {
	args := m.MethodCalled("GetDashboardVulnByTime")
	return args.Get(0).([]*dashboard.VulnerabilitiesByTime), utilsMock.ReturnNilOrError(args, 1)
}

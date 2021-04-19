package dashboard

import (
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetTotalDevelopers(filter *dashboard.FilterDashboard) (int, error) {
	args := m.MethodCalled("GetTotalDevelopers")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetTotalRepositories(filter *dashboard.FilterDashboard) (int, error) {
	args := m.MethodCalled("GetTotalRepositories")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnBySeverity(filter *dashboard.FilterDashboard) (*dashboard.VulnerabilityBySeverity, error) {
	args := m.MethodCalled("GetVulnBySeverity")
	return args.Get(0).(*dashboard.VulnerabilityBySeverity), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnByDeveloper(filter *dashboard.FilterDashboard) (*[]dashboard.VulnerabilityByDeveloper, error) {
	args := m.MethodCalled("GetVulnByDeveloper")
	return args.Get(0).(*[]dashboard.VulnerabilityByDeveloper), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnByRepository(filter *dashboard.FilterDashboard) (*[]dashboard.VulnerabilityByRepository, error) {
	args := m.MethodCalled("GetVulnByRepository")
	return args.Get(0).(*[]dashboard.VulnerabilityByRepository), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnByLanguage(filter *dashboard.FilterDashboard) (*[]dashboard.VulnerabilityByLanguage, error) {
	args := m.MethodCalled("GetVulnByLanguage")
	return args.Get(0).(*[]dashboard.VulnerabilityByLanguage), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnByTime(filter *dashboard.FilterDashboard) (*[]dashboard.VulnerabilityByTime, error) {
	args := m.MethodCalled("GetVulnByTime")
	return args.Get(0).(*[]dashboard.VulnerabilityByTime), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnDetails(filter *dashboard.FilterDashboard) (*dashboard.VulnerabilityDetails, error) {
	args := m.MethodCalled("GetVulnDetails")
	return args.Get(0).(*dashboard.VulnerabilityDetails), utilsMock.ReturnNilOrError(args, 1)
}

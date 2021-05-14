package dashboard

import (
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/database"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"
	"github.com/stretchr/testify/mock"
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
func (m *Mock) GetDashboardTotalDevelopers(_ *database.Filter) (int, error) {
	args := m.MethodCalled("GetDashboardTotalDevelopers")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardTotalRepositories(_ *database.Filter) (int, error) {
	args := m.MethodCalled("GetDashboardTotalRepositories")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnBySeverity(_ *database.Filter) (*response.Vulnerability, error) {
	args := m.MethodCalled("GetDashboardVulnBySeverity")
	return args.Get(0).(*response.Vulnerability), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByAuthor(_ *database.Filter) ([]*database.VulnerabilitiesByAuthor, error) {
	args := m.MethodCalled("GetDashboardVulnByAuthor")
	return args.Get(0).([]*database.VulnerabilitiesByAuthor), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByRepository(_ *database.Filter) ([]*database.VulnerabilitiesByRepository, error) {
	args := m.MethodCalled("GetDashboardVulnByRepository")
	return args.Get(0).([]*database.VulnerabilitiesByRepository), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByLanguage(_ *database.Filter) ([]*database.VulnerabilitiesByLanguage, error) {
	args := m.MethodCalled("GetDashboardVulnByLanguage")
	return args.Get(0).([]*database.VulnerabilitiesByLanguage), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByTime(_ *database.Filter) ([]*database.VulnerabilitiesByTime, error) {
	args := m.MethodCalled("GetDashboardVulnByTime")
	return args.Get(0).([]*database.VulnerabilitiesByTime), utilsMock.ReturnNilOrError(args, 1)
}

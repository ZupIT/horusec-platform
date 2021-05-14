package dashboard

import (
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/repositories"
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
func (m *Mock) GetDashboardTotalDevelopers(_ *repositories.Filter) (int, error) {
	args := m.MethodCalled("GetDashboardTotalDevelopers")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardTotalRepositories(_ *repositories.Filter) (int, error) {
	args := m.MethodCalled("GetDashboardTotalRepositories")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnBySeverity(_ *repositories.Filter) (*response.Vulnerability, error) {
	args := m.MethodCalled("GetDashboardVulnBySeverity")
	return args.Get(0).(*response.Vulnerability), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByAuthor(_ *repositories.Filter) ([]*repositories.VulnerabilitiesByAuthor, error) {
	args := m.MethodCalled("GetDashboardVulnByAuthor")
	return args.Get(0).([]*repositories.VulnerabilitiesByAuthor), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByRepository(_ *repositories.Filter) ([]*repositories.VulnerabilitiesByRepository, error) {
	args := m.MethodCalled("GetDashboardVulnByRepository")
	return args.Get(0).([]*repositories.VulnerabilitiesByRepository), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByLanguage(_ *repositories.Filter) ([]*repositories.VulnerabilitiesByLanguage, error) {
	args := m.MethodCalled("GetDashboardVulnByLanguage")
	return args.Get(0).([]*repositories.VulnerabilitiesByLanguage), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByTime(_ *repositories.Filter) ([]*repositories.VulnerabilitiesByTime, error) {
	args := m.MethodCalled("GetDashboardVulnByTime")
	return args.Get(0).([]*repositories.VulnerabilitiesByTime), utilsMock.ReturnNilOrError(args, 1)
}

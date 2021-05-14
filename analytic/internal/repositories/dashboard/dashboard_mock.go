package dashboard

import (
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/repository"
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
func (m *Mock) GetDashboardTotalDevelopers(_ *repository.Filter) (int, error) {
	args := m.MethodCalled("GetDashboardTotalDevelopers")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardTotalRepositories(_ *repository.Filter) (int, error) {
	args := m.MethodCalled("GetDashboardTotalRepositories")
	return args.Get(0).(int), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnBySeverity(_ *repository.Filter) (*response.Vulnerability, error) {
	args := m.MethodCalled("GetDashboardVulnBySeverity")
	return args.Get(0).(*response.Vulnerability), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByAuthor(_ *repository.Filter) ([]*repository.VulnerabilitiesByAuthor, error) {
	args := m.MethodCalled("GetDashboardVulnByAuthor")
	return args.Get(0).([]*repository.VulnerabilitiesByAuthor), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByRepository(_ *repository.Filter) ([]*repository.VulnerabilitiesByRepository, error) {
	args := m.MethodCalled("GetDashboardVulnByRepository")
	return args.Get(0).([]*repository.VulnerabilitiesByRepository), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByLanguage(_ *repository.Filter) ([]*repository.VulnerabilitiesByLanguage, error) {
	args := m.MethodCalled("GetDashboardVulnByLanguage")
	return args.Get(0).([]*repository.VulnerabilitiesByLanguage), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetDashboardVulnByTime(_ *repository.Filter) ([]*repository.VulnerabilitiesByTime, error) {
	args := m.MethodCalled("GetDashboardVulnByTime")
	return args.Get(0).([]*repository.VulnerabilitiesByTime), utilsMock.ReturnNilOrError(args, 1)
}

package dashboard

import (
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetTotalDevelopers(_ *dashboard.FilterDashboard) (interface{}, error) {
	args := m.MethodCalled("GetTotalDevelopers")
	return args.Get(0), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetTotalRepositories(_ *dashboard.FilterDashboard) (interface{}, error) {
	args := m.MethodCalled("GetTotalRepositories")
	return args.Get(0), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnBySeverity(_ *dashboard.FilterDashboard) (interface{}, error) {
	args := m.MethodCalled("GetVulnBySeverity")
	return args.Get(0), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnByDeveloper(_ *dashboard.FilterDashboard) (interface{}, error) {
	args := m.MethodCalled("GetVulnByDeveloper")
	return args.Get(0), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnByRepository(_ *dashboard.FilterDashboard) (interface{}, error) {
	args := m.MethodCalled("GetVulnByRepository")
	return args.Get(0), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnByLanguage(_ *dashboard.FilterDashboard) (interface{}, error) {
	args := m.MethodCalled("GetVulnByLanguage")
	return args.Get(0), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnByTime(_ *dashboard.FilterDashboard) (interface{}, error) {
	args := m.MethodCalled("GetVulnByTime")
	return args.Get(0), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) GetVulnDetails(_ *dashboard.FilterDashboard) (interface{}, error) {
	args := m.MethodCalled("GetVulnDetails")
	return args.Get(0), utilsMock.ReturnNilOrError(args, 1)
}

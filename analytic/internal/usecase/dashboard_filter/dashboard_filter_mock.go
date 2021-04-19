package dashboardfilter

import (
	netHTTP "net/http"

	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) ExtractFilterDashboardByWorkspace(r *netHTTP.Request) (*dashboard.FilterDashboard, error) {
	args := m.MethodCalled("ExtractFilterDashboardByWorkspace")
	return args.Get(0).(*dashboard.FilterDashboard), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) ExtractFilterDashboardByRepository(r *netHTTP.Request) (*dashboard.FilterDashboard, error) {
	args := m.MethodCalled("ExtractFilterDashboardByRepository")
	return args.Get(0).(*dashboard.FilterDashboard), utilsMock.ReturnNilOrError(args, 1)
}

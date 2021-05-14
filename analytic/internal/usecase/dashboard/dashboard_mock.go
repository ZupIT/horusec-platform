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

func (m *Mock) ExtractFilterDashboard(r *netHTTP.Request) (*dashboard.Filters, error) {
	args := m.MethodCalled("ExtractFilterDashboard")
	return args.Get(0).(*dashboard.Filters), utilsMock.ReturnNilOrError(args, 1)
}

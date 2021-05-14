package dashboardfilter

import (
	netHTTP "net/http"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/repositories"

	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) ExtractFilterDashboard(r *netHTTP.Request) (*repositories.Filter, error) {
	args := m.MethodCalled("ExtractFilterDashboard")
	return args.Get(0).(*repositories.Filter), utilsMock.ReturnNilOrError(args, 1)
}

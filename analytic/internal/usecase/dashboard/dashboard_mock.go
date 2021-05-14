package dashboardfilter

import (
	netHTTP "net/http"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/database"

	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) ExtractFilterDashboard(r *netHTTP.Request) (*database.Filter, error) {
	args := m.MethodCalled("ExtractFilterDashboard")
	return args.Get(0).(*database.Filter), utilsMock.ReturnNilOrError(args, 1)
}

package dispatcher

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) DispatchRequest(_ *analysis.Analysis) error {
	args := m.MethodCalled("DispatchRequest")
	return utilsMock.ReturnNilOrError(args, 0)
}
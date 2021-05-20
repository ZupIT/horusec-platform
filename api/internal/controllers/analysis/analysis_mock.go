package analysis

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) SaveAnalysis(_ *analysis.Analysis) (uuid.UUID, error) {
	args := m.MethodCalled("SaveAnalysis")
	return args.Get(0).(uuid.UUID), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAnalysis(_ uuid.UUID) (*analysis.Analysis, error) {
	args := m.MethodCalled("GetAnalysis")
	return args.Get(0).(*analysis.Analysis), mockUtils.ReturnNilOrError(args, 1)
}

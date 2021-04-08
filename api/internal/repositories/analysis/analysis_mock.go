package analysis

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) FindAnalysisByID(_ uuid.UUID) response.IResponse {
	args := m.MethodCalled("FindAnalysisByID")
	return args.Get(0).(response.IResponse)
}
func (m *Mock) CreateFullAnalysis(analysisArgument *analysis.Analysis) error {
	m.MethodCalled("CreateFullAnalysisArguments").Get(0).(func(*analysis.Analysis))(analysisArgument)
	args := m.MethodCalled("CreateFullAnalysisResponse")
	return utilsMock.ReturnNilOrError(args, 0)
}

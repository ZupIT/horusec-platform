// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
func (m *Mock) FindAllVulnerabilitiesByHashesAndRepository(_ []string, _ uuid.UUID) response.IResponse {
	args := m.MethodCalled("FindAllVulnerabilitiesByHashesAndRepository")
	return args.Get(0).(response.IResponse)
}

func (m *Mock) SaveTreatCompatibility(_ map[string]uuid.UUID, _ *analysis.Analysis) error {
	args := m.MethodCalled("SaveTreatCompatibility")
	return utilsMock.ReturnNilOrError(args, 0)
}

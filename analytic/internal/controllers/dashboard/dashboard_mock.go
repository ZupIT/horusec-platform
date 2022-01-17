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

package dashboard

import (
	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetAllDashboardChartsWorkspace(_ *dashboard.Filter) (*dashboard.Response, error) {
	args := m.MethodCalled("GetAllDashboardChartsWorkspace")
	return args.Get(0).(*dashboard.Response), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAllDashboardChartsRepository(_ *dashboard.Filter) (*dashboard.Response, error) {
	args := m.MethodCalled("GetAllDashboardChartsRepository")
	return args.Get(0).(*dashboard.Response), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) AddVulnerabilitiesByAuthor(_ *analysisEntities.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByAuthor")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) AddVulnerabilitiesByRepository(_ *analysisEntities.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByRepository")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) AddVulnerabilitiesByLanguage(_ *analysisEntities.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByLanguage")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) AddVulnerabilitiesByTime(_ *analysisEntities.Analysis) error {
	args := m.MethodCalled("AddVulnerabilitiesByTime")
	return utilsMock.ReturnNilOrError(args, 0)
}

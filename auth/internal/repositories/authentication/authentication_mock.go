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

package authentication

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	accountEnums "github.com/ZupIT/horusec-devkit/pkg/enums/account"
	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"

	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetWorkspaceGroups(_ uuid.UUID) (*authEntities.AuthzGroups, error) {
	args := m.MethodCalled("GetWorkspaceGroups")
	return args.Get(0).(*authEntities.AuthzGroups), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetRepositoryGroups(_ uuid.UUID) (*authEntities.AuthzGroups, error) {
	args := m.MethodCalled("GetRepositoryGroups")
	return args.Get(0).(*authEntities.AuthzGroups), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetWorkspaceRole(_, _ uuid.UUID) (accountEnums.Role, error) {
	args := m.MethodCalled("GetWorkspaceRole")
	return args.Get(0).(accountEnums.Role), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetRepositoryRole(_, _ uuid.UUID) (accountEnums.Role, error) {
	args := m.MethodCalled("GetRepositoryRole")
	return args.Get(0).(accountEnums.Role), mockUtils.ReturnNilOrError(args, 1)
}

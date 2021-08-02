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

package workspace

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"

	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) ListWorkspacesAuthTypeHorusec(_ uuid.UUID) (*[]workspaceEntities.Response, error) {
	args := m.MethodCalled("ListWorkspacesAuthTypeHorusec")
	return args.Get(0).(*[]workspaceEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListWorkspacesAuthTypeLdap(_ []string) (*[]workspaceEntities.Response, error) {
	args := m.MethodCalled("ListWorkspacesAuthTypeLdap")
	return args.Get(0).(*[]workspaceEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetWorkspace(_ uuid.UUID) (*workspaceEntities.Workspace, error) {
	args := m.MethodCalled("GetWorkspace")
	return args.Get(0).(*workspaceEntities.Workspace), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAccountWorkspace(_, _ uuid.UUID) (*workspaceEntities.AccountWorkspace, error) {
	args := m.MethodCalled("GetAccountWorkspace")
	return args.Get(0).(*workspaceEntities.AccountWorkspace), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListAllWorkspaceUsers(_ uuid.UUID) (*[]roleEntities.Response, error) {
	args := m.MethodCalled("ListAllWorkspaceUsers")
	return args.Get(0).(*[]roleEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListWorkspacesApplicationAdmin() (*[]workspaceEntities.Response, error) {
	args := m.MethodCalled("ListWorkspacesApplicationAdmin")
	return args.Get(0).(*[]workspaceEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) IsWorkspaceAdmin(_, _ uuid.UUID) bool {
	args := m.MethodCalled("IsWorkspaceAdmin")
	return args.Get(0).(bool)
}

func (m *Mock) ListWorkspaceUsersNoBelong(_, _ uuid.UUID) (*[]roleEntities.Response, error) {
	args := m.MethodCalled("ListWorkspaceUsersNoBelong")
	return args.Get(0).(*[]roleEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetWorkspaceLdap(_ uuid.UUID, _ []string) (*workspaceEntities.Response, error) {
	args := m.MethodCalled("GetWorkspaceLdap")
	return args.Get(0).(*workspaceEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

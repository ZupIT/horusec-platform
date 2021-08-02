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

package repository

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"

	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetRepositoryByName(_ uuid.UUID, _ string) (*repositoryEntities.Repository, error) {
	args := m.MethodCalled("GetRepositoryByName")
	return args.Get(0).(*repositoryEntities.Repository), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetRepository(_ uuid.UUID) (*repositoryEntities.Repository, error) {
	args := m.MethodCalled("GetRepository")
	return args.Get(0).(*repositoryEntities.Repository), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAccountRepository(_, _ uuid.UUID) (*repositoryEntities.AccountRepository, error) {
	args := m.MethodCalled("GetAccountRepository")
	return args.Get(0).(*repositoryEntities.AccountRepository), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListRepositoriesAuthTypeHorusec(_, _ uuid.UUID) (*[]repositoryEntities.Response, error) {
	args := m.MethodCalled("ListRepositoriesAuthTypeHorusec")
	return args.Get(0).(*[]repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListRepositoriesAuthTypeLdap(_ uuid.UUID, _ []string) (*[]repositoryEntities.Response, error) {
	args := m.MethodCalled("ListRepositoriesAuthTypeLdap")
	return args.Get(0).(*[]repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) IsNotMemberOfWorkspace(_, _ uuid.UUID) bool {
	args := m.MethodCalled("IsNotMemberOfWorkspace")
	return args.Get(0).(bool)
}

func (m *Mock) ListAllRepositoryUsers(_ uuid.UUID) (*[]roleEntities.Response, error) {
	args := m.MethodCalled("ListAllRepositoryUsers")
	return args.Get(0).(*[]roleEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetWorkspace(_ uuid.UUID) (*workspaceEntities.Workspace, error) {
	args := m.MethodCalled("GetWorkspace")
	return args.Get(0).(*workspaceEntities.Workspace), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ListRepositoriesWhenApplicationAdmin() (*[]repositoryEntities.Response, error) {
	args := m.MethodCalled("ListRepositoriesWhenApplicationAdmin")
	return args.Get(0).(*[]repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetRepositoryLdap(_ uuid.UUID, _ []string) (*repositoryEntities.Response, error) {
	args := m.MethodCalled("GetRepositoryLdap")
	return args.Get(0).(*repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

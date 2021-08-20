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
	tokenEntities "github.com/ZupIT/horusec-platform/core/internal/entities/token"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Create(_ *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	args := m.MethodCalled("Create")
	return args.Get(0).(*repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Get(_ *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	args := m.MethodCalled("Get")
	return args.Get(0).(*repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Update(_ *repositoryEntities.Data) (*repositoryEntities.Response, error) {
	args := m.MethodCalled("Update")
	return args.Get(0).(*repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Delete(_ uuid.UUID) error {
	args := m.MethodCalled("Delete")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) List(_ *repositoryEntities.Data,
	_ *repositoryEntities.PaginatedContent) (*[]repositoryEntities.Response, error) {
	args := m.MethodCalled("List")
	return args.Get(0).(*[]repositoryEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) UpdateRole(_ *roleEntities.Data) (*roleEntities.Response, error) {
	args := m.MethodCalled("UpdateRole")
	return args.Get(0).(*roleEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) InviteUser(_ *roleEntities.UserData) (*roleEntities.Response, error) {
	args := m.MethodCalled("InviteUser")
	return args.Get(0).(*roleEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetUsers(_ uuid.UUID) (*[]roleEntities.Response, error) {
	args := m.MethodCalled("GetUsers")
	return args.Get(0).(*[]roleEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) RemoveUser(_ *roleEntities.Data) error {
	args := m.MethodCalled("RemoveUser")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) CreateToken(_ *tokenEntities.Data) (string, error) {
	args := m.MethodCalled("CreateToken")
	return args.Get(0).(string), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) DeleteToken(_ *tokenEntities.Data) error {
	args := m.MethodCalled("DeleteToken")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) ListTokens(_ *tokenEntities.Data) (*[]tokenEntities.Response, error) {
	args := m.MethodCalled("ListTokens")
	return args.Get(0).(*[]tokenEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

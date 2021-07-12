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

package account

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"

	accountEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetAccount(_ uuid.UUID) (*accountEntities.Account, error) {
	args := m.MethodCalled("GetAccount")
	return args.Get(0).(*accountEntities.Account), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAccountByEmail(_ string) (*accountEntities.Account, error) {
	args := m.MethodCalled("GetAccountByEmail")
	return args.Get(0).(*accountEntities.Account), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAccountByUsername(_ string) (*accountEntities.Account, error) {
	args := m.MethodCalled("GetAccountByUsername")
	return args.Get(0).(*accountEntities.Account), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) CreateAccount(_ *accountEntities.Account) (*accountEntities.Account, error) {
	args := m.MethodCalled("CreateAccount")
	return args.Get(0).(*accountEntities.Account), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Update(_ *accountEntities.Account) (*accountEntities.Account, error) {
	args := m.MethodCalled("Update")
	return args.Get(0).(*accountEntities.Account), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) Delete(_ uuid.UUID) error {
	args := m.MethodCalled("Delete")
	return mockUtils.ReturnNilOrError(args, 0)
}

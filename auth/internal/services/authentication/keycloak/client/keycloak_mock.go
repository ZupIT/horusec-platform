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

package keycloak

import (
	"github.com/Nerzal/gocloak/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Authenticate(_, _ string) (*gocloak.JWT, error) {
	args := m.MethodCalled("Authenticate")
	return args.Get(0).(*gocloak.JWT), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) IsActiveToken(_ string) (bool, error) {
	args := m.MethodCalled("IsActiveToken")
	return args.Get(0).(bool), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAccountIDByJWTToken(_ string) (uuid.UUID, error) {
	args := m.MethodCalled("GetAccountIDByJWTToken")
	return args.Get(0).(uuid.UUID), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetUserInfo(_ string) (*gocloak.UserInfo, error) {
	args := m.MethodCalled("GetUserInfo")
	return args.Get(0).(*gocloak.UserInfo), mockUtils.ReturnNilOrError(args, 1)
}

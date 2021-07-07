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
	authEntities "github.com/ZupIT/horusec-platform/auth/internal/entities/authentication"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) CreateAccountKeycloak(_ string) (*accountEntities.Response, error) {
	args := m.MethodCalled("CreateAccountKeycloak")
	return args.Get(0).(*accountEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) CreateAccountHorusec(_ *accountEntities.Data) (*accountEntities.Response, error) {
	args := m.MethodCalled("CreateAccountHorusec")
	return args.Get(0).(*accountEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) ValidateAccountEmail(_ uuid.UUID) error {
	args := m.MethodCalled("ValidateAccountEmail")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) SendResetPasswordCode(_ string) error {
	args := m.MethodCalled("SendResetPasswordCode")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) ChangePassword(_ *accountEntities.ChangePasswordData) error {
	args := m.MethodCalled("ChangePassword")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) CheckExistingEmailOrUsername(_ *accountEntities.CheckEmailAndUsername) error {
	args := m.MethodCalled("CheckExistingEmailOrUsername")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) DeleteAccount(_ uuid.UUID) error {
	args := m.MethodCalled("DeleteAccount")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (m *Mock) Logout(_ string) {
	_ = m.MethodCalled("Logout")
}

func (m *Mock) CheckResetPasswordCode(_ *accountEntities.ResetCodeData) (string, error) {
	args := m.MethodCalled("CheckResetPasswordCode")
	return args.Get(0).(string), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) RefreshToken(_ string) (*authEntities.LoginResponse, error) {
	args := m.MethodCalled("RefreshToken")
	return args.Get(0).(*authEntities.LoginResponse), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) GetAccountID(_ string) (uuid.UUID, error) {
	args := m.MethodCalled("GetAccountID")
	return args.Get(0).(uuid.UUID), mockUtils.ReturnNilOrError(args, 1)
}

func (m *Mock) UpdateAccount(_ *accountEntities.UpdateAccount) (*accountEntities.Response, error) {
	args := m.MethodCalled("UpdateAccount")
	return args.Get(0).(*accountEntities.Response), mockUtils.ReturnNilOrError(args, 1)
}

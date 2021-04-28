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

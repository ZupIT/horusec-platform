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

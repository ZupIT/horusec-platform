package administrator

import (
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	entity "github.com/ZupIT/horusec-platform/auth/internal/entities/account"
	"github.com/ZupIT/horusec-platform/auth/test/mocks"
)

func TestUseCase_CreateOrUpdate(t *testing.T) {
	t.Run("SHOULD update existing account WHEN provided email already exists", func(t *testing.T) {
		// expected behavior
		writeMock := new(mocks.IDatabaseWrite)
		writeMock.
			On("Update", mock.Anything, mock.Anything, mock.Anything).
			Return(&response.Response{})
		readMock := new(mocks.IDatabaseRead)
		oldest := account1()
		readMock.
			On("Find", mock.Anything, mock.Anything, mock.Anything).
			Run(updateAccountRef([]*entity.Account{oldest})).
			Return(response.NewResponse(1, nil, ""))

		// state under test
		newest := account2()
		err := NewUseCase(&database.Connection{Read: readMock, Write: writeMock}).
			CreateOrUpdate(newest)

		// assertions
		assert.NoError(t, err)
		writeMock.AssertCalled(t, "Update", mock.Anything, map[string]interface{}{"account_id": oldest.AccountID}, mock.Anything)
	})
	t.Run("SHOULD create a new account WHEN provided email doesn't exist", func(t *testing.T) {
		// expected behavior
		writeMock := new(mocks.IDatabaseWrite)
		writeMock.
			On("Create", mock.Anything, mock.Anything).
			Return(&response.Response{})
		readMock := new(mocks.IDatabaseRead)
		readMock.
			On("Find", mock.Anything, mock.Anything, mock.Anything).
			Return(response.NewResponse(1, nil, ""))

		// state under test
		newest := account1()
		err := NewUseCase(&database.Connection{Read: readMock, Write: writeMock}).
			CreateOrUpdate(newest)

		// assertions
		assert.NoError(t, err)
		writeMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
	})
	t.Run("SHOULD delete old accounts WHEN provided email doesn't exist", func(t *testing.T) {
		// expected behavior
		writeMock := new(mocks.IDatabaseWrite)
		writeMock.
			On("Delete", mock.Anything, mock.Anything).
			Return(&response.Response{})
		writeMock.
			On("Create", mock.Anything, mock.Anything).
			Return(&response.Response{})
		readMock := new(mocks.IDatabaseRead)
		oldest := account1()
		readMock.
			On("Find", mock.Anything, mock.Anything, mock.Anything).
			Run(updateAccountRef([]*entity.Account{oldest})).
			Return(response.NewResponse(1, nil, ""))

		// state under test
		newest := account3()
		err := NewUseCase(&database.Connection{Read: readMock, Write: writeMock}).
			CreateOrUpdate(newest)

		// assertions
		assert.NoError(t, err)
		writeMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
		writeMock.AssertCalled(t, "Delete", map[string]interface{}{"account_id": oldest.AccountID}, mock.Anything)
	})
}

func account1() *entity.Account {
	accountID, err := uuid.Parse("9468a7cb-d378-4630-adb3-ec9b5e705e56")
	if err != nil {
		panic(err)
	}
	return &entity.Account{
		AccountID: accountID,
		Email:     "luiz.ferreira@armyspy.com",
		Username:  "Luiz Correia Ferreira",
	}
}

func account2() *entity.Account {
	accountID, err := uuid.Parse("14aa1b52-961c-4900-999e-a1ef7d051010")
	if err != nil {
		panic(err)
	}
	return &entity.Account{
		AccountID: accountID,
		Email:     "luiz.ferreira@armyspy.com",
		Username:  "Luiz C. Ferreira",
	}
}

func account3() *entity.Account {
	accountID, err := uuid.Parse("7c0bfa48-7346-4a31-a76e-838ae7a233b7")
	if err != nil {
		panic(err)
	}
	return &entity.Account{
		AccountID: accountID,
		Email:     "camila.c.almeida@rhyta.com",
		Username:  "Camila Costa Almeida",
	}
}

func updateAccountRef(accounts []*entity.Account) func(args mock.Arguments) {
	return func(args mock.Arguments) {
		arg0 := args.Get(0).(*[]*entity.Account)
		*arg0 = append(*arg0, accounts...)
	}
}

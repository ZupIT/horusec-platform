package token

import (
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) FindTokenByValue(_ string) response.IResponse {
	args := m.MethodCalled("FindTokenByValue")
	return args.Get(0).(response.IResponse)
}

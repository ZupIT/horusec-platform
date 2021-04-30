package email

import (
	"github.com/stretchr/testify/mock"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) SendEmail(_ *emailEntities.Message) error {
	args := m.MethodCalled("SendEmail")
	return mockUtils.ReturnNilOrError(args, 0)
}

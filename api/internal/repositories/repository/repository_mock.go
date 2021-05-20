package repository

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) FindRepository(_ uuid.UUID, _ string) (uuid.UUID, error) {
	args := m.MethodCalled("FindRepository")
	return args.Get(0).(uuid.UUID), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) CreateRepository(_, _ uuid.UUID, _ string) error {
	args := m.MethodCalled("CreateRepository")
	return utilsMock.ReturnNilOrError(args, 0)
}

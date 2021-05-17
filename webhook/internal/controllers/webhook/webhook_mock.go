package webhook

import (
	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Save(_ *webhook.Webhook) (uuid.UUID, error) {
	args := m.MethodCalled("Save")
	return args.Get(0).(uuid.UUID), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) Update(_ *webhook.Webhook, _ uuid.UUID) error {
	args := m.MethodCalled("Update")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) ListAll(_ uuid.UUID) (*[]webhook.WithRepository, error) {
	args := m.MethodCalled("ListAll")
	return args.Get(0).(*[]webhook.WithRepository), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) Remove(_ uuid.UUID) error {
	args := m.MethodCalled("Remove")
	return utilsMock.ReturnNilOrError(args, 0)
}

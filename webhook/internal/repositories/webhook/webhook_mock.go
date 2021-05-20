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

func (m *Mock) Save(_ *webhook.Webhook) error {
	args := m.MethodCalled("Save")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) Update(_ *webhook.Webhook, _ uuid.UUID) error {
	args := m.MethodCalled("Update")
	return utilsMock.ReturnNilOrError(args, 0)
}

func (m *Mock) ListAll(_ uuid.UUID) (entities *[]webhook.WithRepository, err error) {
	args := m.MethodCalled("ListAll")
	return args.Get(0).(*[]webhook.WithRepository), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) ListOne(_ map[string]interface{}) (entity *webhook.Webhook, err error) {
	args := m.MethodCalled("ListOne")
	return args.Get(0).(*webhook.Webhook), utilsMock.ReturnNilOrError(args, 1)
}

func (m *Mock) Remove(_ uuid.UUID) error {
	args := m.MethodCalled("Remove")
	return utilsMock.ReturnNilOrError(args, 0)
}

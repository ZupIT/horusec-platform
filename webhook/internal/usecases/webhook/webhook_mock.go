package webhook

import (
	netHTTP "net/http"

	utilsMock "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) DecodeWebhookFromIoRead(_ *netHTTP.Request) (*webhook.Webhook, error) {
	args := m.MethodCalled("DecodeWebhookFromIoRead")
	return args.Get(0).(*webhook.Webhook), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) ExtractWebhookIDFromURL(_ *netHTTP.Request) (uuid.UUID, error) {
	args := m.MethodCalled("ExtractWebhookIDFromURL")
	return args.Get(0).(uuid.UUID), utilsMock.ReturnNilOrError(args, 1)
}
func (m *Mock) ExtractWorkspaceIDFromURL(_ *netHTTP.Request) (uuid.UUID, error) {
	args := m.MethodCalled("ExtractWorkspaceIDFromURL")
	return args.Get(0).(uuid.UUID), utilsMock.ReturnNilOrError(args, 1)
}

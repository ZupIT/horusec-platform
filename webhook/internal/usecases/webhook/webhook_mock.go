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

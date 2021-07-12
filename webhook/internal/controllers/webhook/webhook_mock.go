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

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
	"errors"
	"testing"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/webhook/internal/controllers/dispatcher"
)

func TestNewWebhookEvent(t *testing.T) {
	t.Run("Should start consume queues without not panics", func(t *testing.T) {
		message := &amqp.Delivery{}
		entity := packet.NewPacket(message)
		entity.SetBody((&analysis.Analysis{}).ToBytes())
		brokerMock := &broker.Mock{}
		brokerMock.On("Consume")
		brokerMock.On("ConsumeHandlerFunc").Return(entity)
		controllerMock := &dispatcher.Mock{}
		controllerMock.On("DispatchRequest").Return(nil)
		assert.NotPanics(t, func() {
			NewWebhookEvent(brokerMock, controllerMock)
			time.Sleep(5 * time.Second)
			brokerMock.AssertCalled(t, "ConsumeHandlerFunc")
		})
	})
	t.Run("Should return error on parse packet to analysis", func(t *testing.T) {
		event := &Event{}
		assert.NotPanics(t, func() {
			event.handleNewAnalysis(packet.NewPacket(&amqp.Delivery{}))
		})
	})
	t.Run("Should return error on parse packet to analysis", func(t *testing.T) {
		controllerMock := &dispatcher.Mock{}
		controllerMock.On("DispatchRequest").Return(errors.New("unexpected error"))
		event := &Event{
			controller: controllerMock,
		}
		pkg := packet.NewPacket(&amqp.Delivery{})
		pkg.SetBody([]byte("{}"))
		assert.NotPanics(t, func() {
			event.handleNewAnalysis(pkg)
		})
	})
}

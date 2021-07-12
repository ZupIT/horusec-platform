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

package email

import (
	"encoding/json"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerPacket "github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	emailController "github.com/ZupIT/horusec-platform/messages/internal/controllers/email"
	eventsEnums "github.com/ZupIT/horusec-platform/messages/internal/enums/events"
)

type EventHandler struct {
	controller emailController.IController
	broker     broker.IBroker
}

func NewEmailEventHandler(controller emailController.IController, brokerLib broker.IBroker) *EventHandler {
	return &EventHandler{
		controller: controller,
		broker:     brokerLib,
	}
}

func (e *EventHandler) StartConsumers() {
	go e.broker.Consume(queues.HorusecEmail.ToString(), "", "", e.handleEmailPacket)
}

func (e *EventHandler) handleEmailPacket(packet brokerPacket.IPacket) {
	var emailData *emailEntities.Message

	if err := json.Unmarshal(packet.GetBody(), &emailData); err != nil {
		logger.LogError(eventsEnums.MessageFailedToUnmarshalPacket, err)
		_ = packet.Ack()
		return
	}

	if err := e.controller.SendEmail(emailData); err != nil {
		logger.LogError(eventsEnums.MessageFailedToSendEmail, err)
	}

	_ = packet.Ack()
}

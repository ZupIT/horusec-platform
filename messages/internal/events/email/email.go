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

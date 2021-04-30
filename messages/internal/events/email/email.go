package email

import (
	"encoding/json"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	emailController "github.com/ZupIT/horusec-platform/messages/internal/controllers/email"
	eventsEnums "github.com/ZupIT/horusec-platform/messages/internal/enums/events"
)

type Consumer struct {
	controller emailController.IController
}

func NewEmailEventConsumer(controller emailController.IController) *Consumer {
	return &Consumer{
		controller: controller,
	}
}

func (c *Consumer) ConsumeEmailPacket(packet packet.IPacket) {
	var emailData *emailEntities.Message

	if err := json.Unmarshal(packet.GetBody(), &emailData); err != nil {
		logger.LogError(eventsEnums.MessageFailedToUnmarshalPacket, err)
		_ = packet.Ack()
		return
	}

	if err := c.controller.SendEmail(emailData); err != nil {
		logger.LogError(eventsEnums.MessageFailedToSendEmail, err)
	}

	_ = packet.Ack()
}

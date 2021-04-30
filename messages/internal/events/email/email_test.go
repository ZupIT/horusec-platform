package email

import (
	"errors"
	"testing"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"

	emailController "github.com/ZupIT/horusec-platform/messages/internal/controllers/email"
)

func TestNewEmailEventConsumer(t *testing.T) {
	t.Run("should success create a new event consumer", func(t *testing.T) {
		assert.NotNil(t, NewEmailEventConsumer(nil))
	})
}

func TestConsumeEmailPacket(t *testing.T) {
	t.Run("should success consume packet", func(t *testing.T) {
		controllerMock := &emailController.Mock{}
		controllerMock.On("SendEmail").Return(nil)

		consumer := NewEmailEventConsumer(controllerMock)

		emailData := emailEntities.Message{}
		assert.NotPanics(t, func() {
			consumer.ConsumeEmailPacket(packet.NewPacket(&amqp.Delivery{Body: emailData.ToBytes()}))
		})
	})

	t.Run("should log error when failed to send email", func(t *testing.T) {
		controllerMock := &emailController.Mock{}
		controllerMock.On("SendEmail").Return(errors.New("test"))

		consumer := NewEmailEventConsumer(controllerMock)

		emailData := emailEntities.Message{}
		assert.NotPanics(t, func() {
			consumer.ConsumeEmailPacket(packet.NewPacket(&amqp.Delivery{Body: emailData.ToBytes()}))
		})
	})

	t.Run("should log error and ack packet when failed to parse body", func(t *testing.T) {
		controllerMock := &emailController.Mock{}

		consumer := NewEmailEventConsumer(controllerMock)

		assert.NotPanics(t, func() {
			consumer.ConsumeEmailPacket(packet.NewPacket(&amqp.Delivery{}))
		})
	})
}

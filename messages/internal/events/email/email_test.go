package email

import (
	"errors"
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"

	emailEntities "github.com/ZupIT/horusec-devkit/pkg/entities/email"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"

	emailController "github.com/ZupIT/horusec-platform/messages/internal/controllers/email"
)

func TestNewEmailEventHandler(t *testing.T) {
	t.Run("should success create a new event consumer", func(t *testing.T) {
		assert.NotNil(t, NewEmailEventHandler(nil, nil))
	})
}

func TestStartConsumers(t *testing.T) {
	t.Run("should panic when failed to consume", func(t *testing.T) {
		handler := NewEmailEventHandler(nil, nil)

		assert.Panics(t, func() {
			handler.StartConsumers()
		})
	})
}

func TestHandleEmailPacket(t *testing.T) {
	t.Run("should success handle packet", func(t *testing.T) {
		brokerMock := &broker.Mock{}

		controllerMock := &emailController.Mock{}
		controllerMock.On("SendEmail").Return(nil)

		handler := NewEmailEventHandler(controllerMock, brokerMock)

		emailData := emailEntities.Message{}
		assert.NotPanics(t, func() {
			handler.handleEmailPacket(packet.NewPacket(&amqp.Delivery{Body: emailData.ToBytes()}))
		})
	})

	t.Run("should log error when failed to send email", func(t *testing.T) {
		brokerMock := &broker.Mock{}

		controllerMock := &emailController.Mock{}
		controllerMock.On("SendEmail").Return(errors.New("test"))

		handler := NewEmailEventHandler(controllerMock, brokerMock)

		emailData := emailEntities.Message{}
		assert.NotPanics(t, func() {
			handler.handleEmailPacket(packet.NewPacket(&amqp.Delivery{Body: emailData.ToBytes()}))
		})
	})

	t.Run("should log error and ack packet when failed to parse body", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		controllerMock := &emailController.Mock{}

		handler := NewEmailEventHandler(controllerMock, brokerMock)

		assert.NotPanics(t, func() {
			handler.handleEmailPacket(packet.NewPacket(&amqp.Delivery{}))
		})
	})
}

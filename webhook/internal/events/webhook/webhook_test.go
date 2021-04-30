package webhook

import (
	"errors"
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
	"github.com/ZupIT/horusec-platform/webhook/internal/controllers/dispatcher"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
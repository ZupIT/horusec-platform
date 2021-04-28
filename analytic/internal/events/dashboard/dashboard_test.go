package dashboard

import (
	"errors"
	"testing"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"

	controller "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
)

func TestNewDashboardEvent(t *testing.T) {
	t.Run("Should consume queues without panics", func(t *testing.T) {
		message := &amqp.Delivery{}
		entity := packet.NewPacket(message)
		entity.SetBody((&analysis.Analysis{}).ToBytes())
		brokerMock := &broker.Mock{}
		brokerMock.On("ConsumeHandlerFunc").Return(entity)
		brokerMock.On("Consume").Return()
		controllerMock := &controller.Mock{}
		controllerMock.On("AddVulnerabilitiesByAuthor").Return(nil)
		controllerMock.On("AddVulnerabilitiesByRepository").Return(nil)
		controllerMock.On("AddVulnerabilitiesByLanguage").Return(nil)
		controllerMock.On("AddVulnerabilitiesByTime").Return(nil)
		assert.NotPanics(t, func() {
			NewDashboardEvent(brokerMock, controllerMock)
			time.Sleep(5 * time.Second)
			brokerMock.AssertCalled(t, "ConsumeHandlerFunc")
		})
	})
	t.Run("Should Handle new analysis with error on parse packet", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		controllerMock := &controller.Mock{}
		controllerMock.On("AddVulnerabilitiesByAuthor").Return(nil)
		e := &Event{broker: brokerMock, controller: controllerMock}
		message := &amqp.Delivery{}
		entity := packet.NewPacket(message)
		e.handleNewAnalysis(entity, queues.HorusecAnalyticAuthors)
	})
	t.Run("Should Handle new analysis with error on call controller method", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		controllerMock := &controller.Mock{}
		controllerMock.On("AddVulnerabilitiesByAuthor").Return(errors.New("unexpected error"))
		e := &Event{broker: brokerMock, controller: controllerMock}
		message := &amqp.Delivery{}
		entity := packet.NewPacket(message)
		entity.SetBody((&analysis.Analysis{}).ToBytes())
		e.handleNewAnalysis(entity, queues.HorusecAnalyticAuthors)
		controllerMock.AssertCalled(t, "AddVulnerabilitiesByAuthor")
	})
	t.Run("Should Handle new analysis with error because pass wrong queue name", func(t *testing.T) {
		brokerMock := &broker.Mock{}
		controllerMock := &controller.Mock{}
		controllerMock.On("AddVulnerabilitiesByAuthor").Return(errors.New("unexpected error"))
		e := &Event{broker: brokerMock, controller: controllerMock}
		message := &amqp.Delivery{}
		entity := packet.NewPacket(message)
		entity.SetBody((&analysis.Analysis{}).ToBytes())
		assert.Panics(t, func() {
			e.handleNewAnalysis(entity, queues.HorusecEmail)
		})

	})
}

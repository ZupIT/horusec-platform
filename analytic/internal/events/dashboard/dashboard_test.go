package dashboard

import (
	"errors"
	"testing"
	"time"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerPacket "github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"

	dashboardController "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
)

func TestNewDashboardEvent(t *testing.T) {
	t.Run("should start consumers and consume without errors", func(t *testing.T) {
		controllerMock := &dashboardController.Mock{}
		brokerMock := &broker.Mock{}

		delivery := &amqp.Delivery{}
		packet := brokerPacket.NewPacket(delivery)
		packet.SetBody((&analysis.Analysis{}).ToBytes())

		brokerMock.On("ConsumeHandlerFunc").Return(packet)
		brokerMock.On("Consume").Return()

		controllerMock.On("AddVulnerabilitiesByAuthor").Return(nil)
		controllerMock.On("AddVulnerabilitiesByRepository").Return(nil)
		controllerMock.On("AddVulnerabilitiesByLanguage").Return(nil)
		controllerMock.On("AddVulnerabilitiesByTime").Return(nil)

		assert.NotPanics(t, func() {
			NewDashboardEvents(brokerMock, controllerMock)

			time.Sleep(1 * time.Second)

			brokerMock.AssertCalled(t, "ConsumeHandlerFunc")
		})
	})

	t.Run("should return error when failed parse packet", func(t *testing.T) {
		controllerMock := &dashboardController.Mock{}
		brokerMock := &broker.Mock{}

		controllerMock.On("AddVulnerabilitiesByAuthor").Return(nil)

		events := &Events{broker: brokerMock, controller: controllerMock}

		delivery := &amqp.Delivery{}
		packet := brokerPacket.NewPacket(delivery)

		assert.NotPanics(t, func() {
			events.handleNewAnalysis(packet, queues.HorusecAnalyticNewAnalysisByAuthors)
		})
	})

	t.Run("should return error when failed to process packet", func(t *testing.T) {
		controllerMock := &dashboardController.Mock{}
		brokerMock := &broker.Mock{}

		controllerMock.On("AddVulnerabilitiesByAuthor").Return(errors.New("test"))

		events := &Events{broker: brokerMock, controller: controllerMock}

		delivery := &amqp.Delivery{}
		packet := brokerPacket.NewPacket(delivery)
		packet.SetBody((&analysis.Analysis{}).ToBytes())

		assert.NotPanics(t, func() {
			events.handleNewAnalysis(packet, queues.HorusecAnalyticNewAnalysisByAuthors)

			time.Sleep(1 * time.Second)

			controllerMock.AssertCalled(t, "AddVulnerabilitiesByAuthor")
		})
	})

	t.Run("should panic because invalid queue name", func(t *testing.T) {
		controllerMock := &dashboardController.Mock{}
		brokerMock := &broker.Mock{}

		events := &Events{broker: brokerMock, controller: controllerMock}

		delivery := &amqp.Delivery{}
		packet := brokerPacket.NewPacket(delivery)
		packet.SetBody((&analysis.Analysis{}).ToBytes())

		assert.Panics(t, func() {
			events.handleNewAnalysis(packet, queues.HorusecEmail)
		})
	})
}

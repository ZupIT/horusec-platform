package webhook

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/exchange"
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/webhook/internal/controllers/dispatcher"
)

type IEvent interface{}

type Event struct {
	broker     broker.IBroker
	controller dispatcher.IDispatcherController
}

func NewWebhookEvent(iBroker broker.IBroker, controller dispatcher.IDispatcherController) IEvent {
	e := &Event{
		broker:     iBroker,
		controller: controller,
	}
	return e.consumeQueues()
}

func (e *Event) consumeQueues() IEvent {
	go e.broker.Consume(queues.HorusecWebhook.ToString(), exchange.NewAnalysis, exchange.Fanout,
		e.handleNewAnalysis)
	return e
}

func (e *Event) handleNewAnalysis(brokerPacket packet.IPacket) {
	logger.LogInfo("{HORUSEC} Packet received from new analysis")
	entity := analysis.Analysis{}
	if err := parser.ParsePacketToEntity(brokerPacket, &entity); err != nil {
		logger.LogError("{HORUSEC} Read packet error", err)
		return
	}

	if err := e.controller.DispatchRequest(&entity); err != nil {
		logger.LogError("{HORUSEC} Error on dispatch new analysis", err)
		return
	}
	_ = brokerPacket.Ack()
}

package dashboard

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/exchange"
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
)

type IEvent interface {}

type Event struct {
	broker     broker.IBroker
	controller dashboard.IWriteController
}

func NewDashboardEvent(iBroker broker.IBroker, controller dashboard.IWriteController) IEvent {
	e := &Event{
		broker:     iBroker,
		controller: controller,
	}
	return e.consumeQueues()
}

func (e *Event) consumeQueues() IEvent {
	go e.broker.Consume(queues.HorusecNewAnalysis.ToString(), exchange.NewAnalysis.ToString(),
		exchange.Fanout.ToString(), e.handleNewAnalysis)
	return e
}

func (e *Event) handleNewAnalysis(packet packet.IPacket) {
	entity := analysis.Analysis{}
	if err := parser.ParsePacketToEntity(packet, &entity); err != nil {
		logger.LogError("{HORUSEC} Read packet error", err)
		return
	}

	if err := e.controller.AddNewAnalysis(&entity); err != nil {
		logger.LogError("{HORUSEC} Save new analysis", err)
		return
	}
	_ = packet.Ack()
	logger.LogInfo("{HORUSEC} New analysis received and registered on analytic")
}

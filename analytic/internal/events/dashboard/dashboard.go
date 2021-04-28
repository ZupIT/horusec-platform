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

type IEvent interface{}

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
	go e.broker.Consume(
		queues.HorusecAnalyticAuthors.ToString(), exchange.NewAnalysis.ToString(), exchange.Fanout.ToString(),
		func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticAuthors) })
	go e.broker.Consume(
		queues.HorusecAnalyticRepositories.ToString(), exchange.NewAnalysis.ToString(), exchange.Fanout.ToString(),
		func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticRepositories) })
	go e.broker.Consume(
		queues.HorusecAnalyticLanguages.ToString(), exchange.NewAnalysis.ToString(), exchange.Fanout.ToString(),
		func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticLanguages) })
	go e.broker.Consume(
		queues.HorusecAnalyticTimes.ToString(), exchange.NewAnalysis.ToString(), exchange.Fanout.ToString(),
		func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticTimes) })
	return e
}

// nolint:exhaustive // others queues is not necessary
func (e *Event) execControllerByQueueType(queue queues.Queue) func(*analysis.Analysis) error {
	switch queue {
	case queues.HorusecAnalyticAuthors:
		return e.controller.AddVulnerabilitiesByAuthor
	case queues.HorusecAnalyticRepositories:
		return e.controller.AddVulnerabilitiesByRepository
	case queues.HorusecAnalyticLanguages:
		return e.controller.AddVulnerabilitiesByLanguage
	case queues.HorusecAnalyticTimes:
		return e.controller.AddVulnerabilitiesByTime
	}
	return nil
}

func (e *Event) handleNewAnalysis(brokerPacket packet.IPacket, queue queues.Queue) {
	logger.LogInfo("{HORUSEC} Packet received from new analysis")
	entity := analysis.Analysis{}
	if err := parser.ParsePacketToEntity(brokerPacket, &entity); err != nil {
		logger.LogError("{HORUSEC} Read packet error by "+queue.ToString(), err)
		return
	}

	if err := e.execControllerByQueueType(queue)(&entity); err != nil {
		logger.LogError("{HORUSEC} Error on save new analysis by "+queue.ToString(), err)
		return
	}
	_ = brokerPacket.Ack()
}
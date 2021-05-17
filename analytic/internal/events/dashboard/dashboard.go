package dashboard

import (
	"fmt"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/exchange"
	"github.com/ZupIT/horusec-devkit/pkg/enums/queues"
	brokerLib "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	"github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	eventsEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/events"
)

type Events struct {
	broker     brokerLib.IBroker
	controller dashboard.IController
}

func NewDashboardEvents(broker brokerLib.IBroker, controller dashboard.IController) *Events {
	events := &Events{
		broker:     broker,
		controller: controller,
	}

	return events.startConsumers()
}

func (e *Events) startConsumers() *Events {
	go e.broker.Consume(queues.HorusecAnalyticAuthors.ToString(), exchange.NewAnalysis.ToString(), exchange.Fanout.
		ToString(), func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticAuthors) })

	go e.broker.Consume(queues.HorusecAnalyticRepositories.ToString(), exchange.NewAnalysis.ToString(), exchange.Fanout.
		ToString(), func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticRepositories) })

	go e.broker.Consume(queues.HorusecAnalyticLanguages.ToString(), exchange.NewAnalysis.ToString(), exchange.Fanout.
		ToString(), func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticLanguages) })

	go e.broker.Consume(queues.HorusecAnalyticTimes.ToString(), exchange.NewAnalysis.ToString(), exchange.Fanout.
		ToString(), func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticTimes) })

	return e
}

func (e *Events) processAnalysisPacketByQueue(queue queues.Queue) func(*analysisEntities.Analysis) error {
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

func (e *Events) handleNewAnalysis(analysisPacket packet.IPacket, queue queues.Queue) {
	logger.LogInfo(eventsEnums.MessageNewAnalysisReceivedAnalytic)
	analysis := &analysisEntities.Analysis{}

	if err := parser.ParsePacketToEntity(analysisPacket, analysis); err != nil {
		logger.LogError(fmt.Sprintf(eventsEnums.MessageFailedToParsePacket, analysisPacket.GetBody(), queue), err)
		_ = analysisPacket.Ack()
		return
	}

	if err := e.processAnalysisPacketByQueue(queue)(analysis); err != nil {
		logger.LogError(fmt.Sprintf(eventsEnums.MessageFailedToProcessPacket, analysisPacket.GetBody(), queue), err)
		_ = analysisPacket.Ack()
		return
	}

	_ = analysisPacket.Ack()
}

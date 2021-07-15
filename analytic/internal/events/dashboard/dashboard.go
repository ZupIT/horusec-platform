// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	go e.broker.Consume(queues.HorusecAnalyticNewAnalysisByAuthors.ToString(), exchange.NewAnalysis, exchange.Fanout,
		func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticNewAnalysisByAuthors) })

	go e.broker.Consume(queues.HorusecAnalyticNewAnalysisByRepository.ToString(), exchange.NewAnalysis, exchange.Fanout,
		func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticNewAnalysisByRepository) })

	go e.broker.Consume(queues.HorusecAnalyticNewAnalysisByLanguage.ToString(), exchange.NewAnalysis, exchange.Fanout,
		func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticNewAnalysisByLanguage) })

	go e.broker.Consume(queues.HorusecAnalyticNewAnalysisByTime.ToString(), exchange.NewAnalysis, exchange.Fanout,
		func(pack packet.IPacket) { e.handleNewAnalysis(pack, queues.HorusecAnalyticNewAnalysisByTime) })

	return e
}

func (e *Events) handleNewAnalysis(analysisPacket packet.IPacket, queue queues.Queue) {
	logger.LogInfo(eventsEnums.MessageNewAnalysisReceivedAnalytic)
	analysis := &analysisEntities.Analysis{}

	if err := parser.ParsePacketToEntity(analysisPacket, analysis); err != nil {
		logger.LogError(fmt.Sprintf(eventsEnums.MessageFailedToParsePacket, analysisPacket.GetBody(), queue), err)
		_ = analysisPacket.Ack()
		return
	}

	logger.LogError(fmt.Sprintf(eventsEnums.MessageFailedToProcessPacket,
		analysisPacket.GetBody(), queue), e.processNewAnalysisPacketByQueue(queue)(analysis))

	_ = analysisPacket.Ack()
}

//nolint:exhaustive // no need of all constants
func (e *Events) processNewAnalysisPacketByQueue(queue queues.Queue) func(*analysisEntities.Analysis) error {
	switch queue {
	case queues.HorusecAnalyticNewAnalysisByAuthors:
		return e.controller.AddVulnerabilitiesByAuthor
	case queues.HorusecAnalyticNewAnalysisByRepository:
		return e.controller.AddVulnerabilitiesByRepository
	case queues.HorusecAnalyticNewAnalysisByLanguage:
		return e.controller.AddVulnerabilitiesByLanguage
	case queues.HorusecAnalyticNewAnalysisByTime:
		return e.controller.AddVulnerabilitiesByTime
	}

	return nil
}

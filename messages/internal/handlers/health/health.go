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

package health

import (
	"net/http"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	_ "github.com/ZupIT/horusec-devkit/pkg/utils/http/entities" // swagger import

	"github.com/ZupIT/horusec-platform/messages/internal/enums/health"
	"github.com/ZupIT/horusec-platform/messages/internal/services/mailer"
)

type Handler struct {
	mailerService mailer.IService
	broker        broker.IBroker
}

func NewHealthHandler(brokerLib broker.IBroker, mailerService mailer.IService) *Handler {
	return &Handler{
		broker:        brokerLib,
		mailerService: mailerService,
	}
}

func (h *Handler) Options(w http.ResponseWriter, _ *http.Request) {
	httpUtil.StatusNoContent(w)
}

// @Tags Health
// @Description Check if application is healthy
// @ID health
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.Response
// @Failure 200 {object} entities.Response
// @Router /messages/health [get]
// @Security ApiKeyAuth
func (h *Handler) Get(w http.ResponseWriter, _ *http.Request) {
	if !h.mailerService.IsAvailable() {
		httpUtil.StatusInternalServerError(w, health.ErrorUnhealthyMailer)
		return
	}

	if !h.broker.IsAvailable() {
		httpUtil.StatusInternalServerError(w, health.ErrorUnhealthyBroker)
		return
	}

	httpUtil.StatusOK(w, nil)
}

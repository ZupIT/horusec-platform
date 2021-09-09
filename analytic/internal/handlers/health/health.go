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

	brokerLib "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	_ "github.com/ZupIT/horusec-devkit/pkg/utils/http/entities" // [swagger-import]
	httpUtilEnums "github.com/ZupIT/horusec-devkit/pkg/utils/http/enums"
)

type Handler struct {
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	broker        brokerLib.IBroker
}

func NewHealthHandler(databaseConnection *database.Connection, broker brokerLib.IBroker) *Handler {
	return &Handler{
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
		broker:        broker,
	}
}

func (h *Handler) Options(w http.ResponseWriter, _ *http.Request) {
	httpUtil.StatusNoContent(w)
}

// Get
// @Tags Health
// @Description Check if Health of service it's OK!
// @ID health
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.Response{content=string} "OK"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/health [get]
func (h *Handler) Get(w http.ResponseWriter, _ *http.Request) {
	if h.IsDatabaseNotAvailable() {
		httpUtil.StatusInternalServerError(w, httpUtilEnums.ErrorDatabaseIsNotHealth)
		return
	}

	if isAvailable := h.broker.IsAvailable(); !isAvailable {
		httpUtil.StatusInternalServerError(w, httpUtilEnums.ErrorBrokerIsNotHealth)
		return
	}

	httpUtil.StatusOK(w, nil)
}

func (h *Handler) IsDatabaseNotAvailable() bool {
	return !h.databaseRead.IsAvailable() || !h.databaseWrite.IsAvailable()
}

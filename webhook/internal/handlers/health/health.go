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
	"fmt"
	netHTTP "net/http"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"

	"google.golang.org/grpc"

	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/health"
	httpUtilEnums "github.com/ZupIT/horusec-devkit/pkg/utils/http/enums"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"

	_ "github.com/ZupIT/horusec-devkit/pkg/utils/http/entities" // [swagger-import]
)

type Handler struct {
	databaseRead           database.IDatabaseRead
	databaseWrite          database.IDatabaseWrite
	grpcHealthCheckService health.ICheckClient
	broker                 broker.IBroker
}

func NewHealthHandler(databaseConnection *database.Connection, authConGRPC grpc.ClientConnInterface,
	brokerConnection broker.IBroker) *Handler {
	return &Handler{
		databaseRead:           databaseConnection.Read,
		databaseWrite:          databaseConnection.Write,
		broker:                 brokerConnection,
		grpcHealthCheckService: health.NewHealthCheckGrpcClient(authConGRPC.(*grpc.ClientConn)),
	}
}

func (h *Handler) Options(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
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
// @Router /webhook/health [get]
func (h *Handler) Get(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	if h.databaseNotAvailable() {
		httpUtil.StatusInternalServerError(w, httpUtilEnums.ErrorDatabaseIsNotHealth)
		return
	}
	if isAvailable := h.broker.IsAvailable(); !isAvailable {
		httpUtil.StatusInternalServerError(w, httpUtilEnums.ErrorBrokerIsNotHealth)
		return
	}
	if isAvailable, state := h.grpcAvailable(); !isAvailable {
		httpUtil.StatusInternalServerError(w, fmt.Errorf("%e %s", httpUtilEnums.ErrorGrpcIsNotHealth, state))
		return
	}
	httpUtil.StatusOK(w, "service is healthy")
}

func (h *Handler) databaseNotAvailable() bool {
	return !h.databaseRead.IsAvailable() || !h.databaseWrite.IsAvailable()
}
func (h *Handler) grpcAvailable() (bool, string) {
	return h.grpcHealthCheckService.IsAvailable()
}

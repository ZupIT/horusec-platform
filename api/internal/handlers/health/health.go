// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

	"google.golang.org/grpc"

	brokerService "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/health"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	httpUtilEnums "github.com/ZupIT/horusec-devkit/pkg/utils/http/enums"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
)

type Handler struct {
	broker                 brokerService.IBroker
	brokerConfig           brokerConfig.IConfig
	databaseRead           database.IDatabaseRead
	databaseWrite          database.IDatabaseWrite
	grpcHealthCheckService health.ICheckClient
}

func NewHealthHandler(broker brokerService.IBroker, brokerConfiguration brokerConfig.IConfig,
	databaseConnection *database.Connection, grpcCon *grpc.ClientConn) *Handler {
	return &Handler{
		broker:                 broker,
		brokerConfig:           brokerConfiguration,
		databaseRead:           databaseConnection.Read,
		databaseWrite:          databaseConnection.Write,
		grpcHealthCheckService: health.NewHealthCheckGrpcClient(grpcCon),
	}
}

func (h *Handler) Options(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
}

func (h *Handler) Get(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	if h.databaseAvailable() {
		httpUtil.StatusInternalServerError(w, httpUtilEnums.ErrorDatabaseIsNotHealth)
		return
	}
	if h.brokerAvailable() {
		httpUtil.StatusInternalServerError(w, httpUtilEnums.ErrorBrokerIsNotHealth)
		return
	}
	if isAvailable, state := h.grpcAvailable(); !isAvailable {
		httpUtil.StatusInternalServerError(w, fmt.Errorf("%e %s", httpUtilEnums.ErrorGrpcIsNotHealth, state))
		return
	}
	httpUtil.StatusOK(w, "service is healthy")
}

func (h *Handler) databaseAvailable() bool {
	return !h.databaseRead.IsAvailable() || !h.databaseWrite.IsAvailable()
}
func (h *Handler) brokerAvailable() bool {
	return h.brokerConfig.GetEnableBroker() && !h.broker.IsAvailable()
}
func (h *Handler) grpcAvailable() (bool, string) {
	return h.grpcHealthCheckService.IsAvailable()
}

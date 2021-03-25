//+build wireinject

package providers

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	analysisHandler "github.com/ZupIT/horusec-platform/api/internal/handlers/analysis"
	healthHandler "github.com/ZupIT/horusec-platform/api/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/api/internal/router"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-platform/api/config/cors"
	analysisController "github.com/ZupIT/horusec-platform/api/internal/controllers/analysis"
)

var providers = wire.NewSet(
	brokerConfig.NewBrokerConfig,
	broker.NewBroker,
	databaseConfig.NewDatabaseConfig,
	database.NewDatabaseReadAndWrite,
	auth.NewAuthGRPCConnection,
	cors.NewCorsConfig,
	http.NewHTTPRouter,
	router.NewHTTPRouter,
	analysisController.NewAnalysisController,
	analysisHandler.NewAnalysisHandler,
	healthHandler.NewHealthHandler,
)

func Initialize(defaultPort string) (router.IRouter, error) {
	wire.Build(providers)
	return &router.Router{}, nil
}

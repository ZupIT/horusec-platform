//+build wireinject

package providers

import (
	analysisHandler "github.com/ZupIT/horusec-platform/api/internal/handlers/analysis"
	healthHandler "github.com/ZupIT/horusec-platform/api/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/api/internal/middelwares/token"
	"github.com/ZupIT/horusec-platform/api/internal/repositories/analysis"
	"github.com/ZupIT/horusec-platform/api/internal/repositories/repository"
	repositoriesToken "github.com/ZupIT/horusec-platform/api/internal/repositories/token"
	"github.com/ZupIT/horusec-platform/api/internal/router"
	"github.com/google/wire"

	appConfiguration "github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/http"

	"github.com/ZupIT/horusec-platform/api/config/cors"
	analysisController "github.com/ZupIT/horusec-platform/api/internal/controllers/analysis"
)

var providers = wire.NewSet(
	brokerConfig.NewBrokerConfig,
	broker.NewBroker,
	databaseConfig.NewDatabaseConfig,
	database.NewDatabaseReadAndWrite,
	auth.NewAuthGRPCConnection,
	proto.NewAuthServiceClient,
	token.NewTokenAuthz,
	analysis.NewRepositoriesAnalysis,
	repository.NewRepositoriesRepository,
	repositoriesToken.NewRepositoriesToken,
	cors.NewCorsConfig,
	http.NewHTTPRouter,
	appConfiguration.NewAppConfig,
	analysisController.NewAnalysisController,
	analysisHandler.NewAnalysisHandler,
	healthHandler.NewHealthHandler,
	router.NewHTTPRouter,
)

func Initialize(defaultPort string) (router.IRouter, error) {
	wire.Build(providers)
	return &router.Router{}, nil
}

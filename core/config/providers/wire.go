//+build wireinject

package providers

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/http"
	workspaceHandler "github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
	"github.com/ZupIT/horusec-platform/core/internal/router"
	"github.com/google/wire"

	"github.com/ZupIT/horusec-platform/core/config/cors"
	workspaceController "github.com/ZupIT/horusec-platform/core/internal/controllers/workspace"
)

var providers = wire.NewSet(
	brokerConfig.NewBrokerConfig,
	broker.NewBroker,
	databaseConfig.NewDatabaseConfig,
	database.NewDatabaseReadAndWrite,
	cors.NewCorsConfig,
	http.NewHTTPRouter,
	router.NewHTTPRouter,
	workspaceController.NewWorkspaceController,
	workspaceHandler.NewWorkspaceHandler,
)

func Initialize(defaultPort string) (router.IRouter, error) {
	wire.Build(providers)
	return &router.Router{}, nil
}

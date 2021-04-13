//+build wireinject
package providers

import (
	"github.com/google/wire"

	"github.com/ZupIT/horusec-platform/auth/config/grpc"
	"github.com/ZupIT/horusec-platform/auth/internal/router"
)

var devKitProviders = wire.NewSet()

var configProviders = wire.NewSet(
	grpc.NewAuthGRPCServer,
)

var controllerProviders = wire.NewSet()

var handleProviders = wire.NewSet()

var useCasesProviders = wire.NewSet()

var repositoriesProviders = wire.NewSet()

func Initialize(_ string) (router.IRouter, error) {
	wire.Build(devKitProviders, configProviders, controllerProviders, handleProviders,
		useCasesProviders, repositoriesProviders)

	return &router.Router{}, nil
}

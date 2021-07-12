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

package router

import (
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/services/http/router"

	"github.com/stretchr/testify/assert"
	grpcService "google.golang.org/grpc"

	"github.com/ZupIT/horusec-platform/auth/config/cors"
	"github.com/ZupIT/horusec-platform/auth/config/grpc"
	accountHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/account"
	authHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/authentication"
	healthHandler "github.com/ZupIT/horusec-platform/auth/internal/handlers/health"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("should success create a new http router and set routes", func(t *testing.T) {
		routerService := router.NewHTTPRouter(cors.NewCorsConfig(), "9999")
		authGRPCServer := &grpc.AuthGRPCServer{Port: 9998, GRPCServer: grpcService.NewServer()}

		assert.NotPanics(t, func() {
			assert.NotNil(t, NewHTTPRouter(routerService, authGRPCServer,
				&authHandler.Handler{}, &accountHandler.Handler{}, &healthHandler.Handler{}))
		})
	})
}

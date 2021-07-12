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

	"github.com/stretchr/testify/assert"

	httpRouter "github.com/ZupIT/horusec-devkit/pkg/services/http/router"
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/ZupIT/horusec-platform/core/config/cors"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/repository"
	"github.com/ZupIT/horusec-platform/core/internal/handlers/workspace"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("should success create a new http router and set routes", func(t *testing.T) {
		routerService := httpRouter.NewHTTPRouter(cors.NewCorsConfig(), "9999")
		middlewareService := middlewares.NewAuthzMiddleware(nil)
		workspaceHandler := &workspace.Handler{}
		repositoryHandler := &repository.Handler{}
		healthHandler := &health.Handler{}

		assert.NotPanics(t, func() {
			assert.NotNil(t, NewHTTPRouter(routerService, middlewareService, workspaceHandler,
				repositoryHandler, healthHandler))
		})
	})
}

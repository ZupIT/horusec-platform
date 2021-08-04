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

	"github.com/ZupIT/horusec-devkit/pkg/services/tracer"

	"github.com/ZupIT/horusec-devkit/pkg/services/http/router"

	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"

	analysisHandler "github.com/ZupIT/horusec-platform/api/internal/handlers/analysis"
	healthHandler "github.com/ZupIT/horusec-platform/api/internal/handlers/health"
	"github.com/ZupIT/horusec-platform/api/internal/middelwares/token"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("Should add all necessary routes", func(t *testing.T) {
		routerConnection := router.NewHTTPRouter(&cors.Options{}, "8000", tracer.Jaeger{})
		healthMock := &healthHandler.Handler{}
		analysisMock := &analysisHandler.Handler{}
		tokenMiddlewareMock := token.NewTokenAuthz(nil)
		instance := NewHTTPRouter(routerConnection, tokenMiddlewareMock, analysisMock, healthMock)
		assert.NotEmpty(t, instance)
	})
}

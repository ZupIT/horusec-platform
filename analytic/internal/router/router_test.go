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
	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"
	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"

	eventDashboard "github.com/ZupIT/horusec-platform/analytic/internal/events/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/handlers/health"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("should add all necessary routes", func(t *testing.T) {
		routerConn := router.NewHTTPRouter(&cors.Options{}, "8009")
		healthMock := &health.Handler{}
		dashboardHandlerMock := &dashboard.Handler{}
		middlewareMock := &middlewares.AuthzMiddleware{}
		eventMock := &eventDashboard.Events{}
		instance := NewHTTPRouter(routerConn, middlewareMock, healthMock, dashboardHandlerMock, eventMock)
		assert.NotEmpty(t, instance)
	})
}

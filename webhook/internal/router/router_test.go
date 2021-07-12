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

	webhook2 "github.com/ZupIT/horusec-platform/webhook/internal/events/webhook"
	"github.com/ZupIT/horusec-platform/webhook/internal/handlers/webhook"

	"github.com/ZupIT/horusec-platform/webhook/internal/handlers/health"

	"github.com/ZupIT/horusec-devkit/pkg/services/http/router"

	"github.com/ZupIT/horusec-devkit/pkg/services/middlewares"

	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("Should add all necessary routes", func(t *testing.T) {
		routerConn := router.NewHTTPRouter(&cors.Options{}, "8005")
		middlewareMock := &middlewares.AuthzMiddleware{}
		healthMock := &health.Handler{}
		webhookHandlerMock := &webhook.Handler{}
		webhookEventMock := &webhook2.Event{}
		instance := NewHTTPRouter(routerConn, middlewareMock, healthMock, webhookHandlerMock, webhookEventMock)
		assert.NotEmpty(t, instance)
	})
}

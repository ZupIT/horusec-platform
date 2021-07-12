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

package dispatcher

import (
	"errors"
	"net/http"
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/http/request"
	"github.com/ZupIT/horusec-devkit/pkg/services/http/request/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
	repositoryWebhook "github.com/ZupIT/horusec-platform/webhook/internal/repositories/webhook"
)

func TestNewDispatcherController(t *testing.T) {
	assert.NotEmpty(t, NewDispatcherController(&repositoryWebhook.Mock{}))
}

func TestController_DispatchRequest(t *testing.T) {
	t.Run("Should dispatch request without error", func(t *testing.T) {
		httpRequestMock := &request.Mock{}
		httpRequestMock.On("NewHTTPRequest").Return(&http.Request{}, nil)
		httpRequestMock.On("DoRequest").Return(&entities.HTTPResponse{Response: &http.Response{StatusCode: http.StatusOK}}, nil)
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListOne").Return(&webhook.Webhook{WebhookID: uuid.New()}, nil)
		controller := &Controller{
			repository:  repoMock,
			httpRequest: httpRequestMock,
		}
		err := controller.DispatchRequest(&analysis.Analysis{})
		assert.NoError(t, err)
	})
	t.Run("Should NOT dispatch request because not exists webhook", func(t *testing.T) {
		httpRequestMock := &request.Mock{}
		httpRequestMock.On("NewHTTPRequest").Return(&http.Request{}, nil)
		httpRequestMock.On("DoRequest").Return(&entities.HTTPResponse{Response: &http.Response{StatusCode: http.StatusOK}}, nil)
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListOne").Return(&webhook.Webhook{}, nil)
		controller := &Controller{
			repository:  repoMock,
			httpRequest: httpRequestMock,
		}
		err := controller.DispatchRequest(&analysis.Analysis{})
		assert.NoError(t, err)
	})
	t.Run("Should return error because on list return unexpected error", func(t *testing.T) {
		httpRequestMock := &request.Mock{}
		httpRequestMock.On("NewHTTPRequest").Return(&http.Request{}, nil)
		httpRequestMock.On("DoRequest").Return(&entities.HTTPResponse{Response: &http.Response{StatusCode: http.StatusOK}}, nil)
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListOne").Return(&webhook.Webhook{}, errors.New("unexpected error"))
		controller := &Controller{
			repository:  repoMock,
			httpRequest: httpRequestMock,
		}
		err := controller.DispatchRequest(&analysis.Analysis{})
		assert.Error(t, err)
	})
	t.Run("Should return error because on mount request is return unexpected error", func(t *testing.T) {
		httpRequestMock := &request.Mock{}
		httpRequestMock.On("NewHTTPRequest").Return(&http.Request{}, errors.New("unexpected error"))
		httpRequestMock.On("DoRequest").Return(&entities.HTTPResponse{Response: &http.Response{StatusCode: http.StatusOK}}, nil)
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListOne").Return(&webhook.Webhook{WebhookID: uuid.New()}, nil)
		controller := &Controller{
			repository:  repoMock,
			httpRequest: httpRequestMock,
		}
		err := controller.DispatchRequest(&analysis.Analysis{})
		assert.Error(t, err)
	})
	t.Run("Should return error because on do request is return unexpected error", func(t *testing.T) {
		httpRequestMock := &request.Mock{}
		httpRequestMock.On("NewHTTPRequest").Return(&http.Request{}, nil)
		httpRequestMock.On("DoRequest").Return(&entities.HTTPResponse{}, errors.New("unexpected error"))
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListOne").Return(&webhook.Webhook{WebhookID: uuid.New()}, nil)
		controller := &Controller{
			repository:  repoMock,
			httpRequest: httpRequestMock,
		}
		err := controller.DispatchRequest(&analysis.Analysis{})
		assert.Error(t, err)
	})
}

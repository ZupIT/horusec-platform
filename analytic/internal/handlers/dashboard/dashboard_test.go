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

package dashboard

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	controller "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

func TestOptions(t *testing.T) {
	t.Run("should return no content when options", func(t *testing.T) {
		controllerMock := &controller.Mock{}

		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)

		handler := NewDashboardHandler(controllerMock)

		handler.Options(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

func TestGetAllChartsByWorkspace(t *testing.T) {
	layoutDateTime := "2006-01-02T15:04:05Z"
	startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
	endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")

	t.Run("should return 200 when success get all charts", func(t *testing.T) {
		controllerMock := &controller.Mock{}
		controllerMock.On("GetAllDashboardCharts").Return(&dashboard.Response{}, nil)

		handler := NewDashboardHandler(controllerMock)

		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			startTime.Format(layoutDateTime), endTime.Format(layoutDateTime), 18, 18)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.New().String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetAllChartsByWorkspace(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when failed to get charts", func(t *testing.T) {
		controllerMock := &controller.Mock{}
		controllerMock.On("GetAllDashboardCharts").Return(&dashboard.Response{}, errors.New("test"))

		handler := NewDashboardHandler(controllerMock)

		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			startTime.Format(layoutDateTime), endTime.Format(layoutDateTime), 18, 18)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.New().String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetAllChartsByWorkspace(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid filter", func(t *testing.T) {
		controllerMock := &controller.Mock{}

		handler := NewDashboardHandler(controllerMock)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "", nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.New().String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetAllChartsByWorkspace(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGetAllChartsByRepository(t *testing.T) {
	layoutDateTime := "2006-01-02T15:04:05Z"
	startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
	endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")

	t.Run("should return 200 when success get all charts", func(t *testing.T) {
		controllerMock := &controller.Mock{}
		controllerMock.On("GetAllDashboardCharts").Return(&dashboard.Response{}, nil)

		handler := NewDashboardHandler(controllerMock)

		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			startTime.Format(layoutDateTime), endTime.Format(layoutDateTime), 18, 18)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.New().String())
		ctx.URLParams.Add("repositoryID", uuid.New().String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetAllChartsByRepository(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when failed to get charts", func(t *testing.T) {
		controllerMock := &controller.Mock{}
		controllerMock.On("GetAllDashboardCharts").Return(&dashboard.Response{}, errors.New("test"))

		handler := NewDashboardHandler(controllerMock)

		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			startTime.Format(layoutDateTime), endTime.Format(layoutDateTime), 18, 18)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.New().String())
		ctx.URLParams.Add("repositoryID", uuid.New().String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetAllChartsByRepository(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 400 when invalid filter", func(t *testing.T) {
		controllerMock := &controller.Mock{}

		handler := NewDashboardHandler(controllerMock)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "", nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", uuid.New().String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.GetAllChartsByRepository(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

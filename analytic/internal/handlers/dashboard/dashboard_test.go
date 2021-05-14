package dashboard

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/database"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	controller "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/enums"
	dashboardfilter "github.com/ZupIT/horusec-platform/analytic/internal/usecase/dashboard"
)

func TestHandler_Options(t *testing.T) {
	t.Run("Should return no content when call options", func(t *testing.T) {
		controllerMock := &controller.Mock{}
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		NewDashboardHandler(controllerMock).Options(w, r)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

func TestHandler_GetAllCharts(t *testing.T) {
	t.Run("Should return status OK when get all charts by workspace", func(t *testing.T) {
		workspaceID := uuid.New()
		controllerMock := &controller.Mock{}
		controllerMock.On("GetAllDashboardCharts").Return(&response.Response{}, nil)
		useCaseMock := &dashboardfilter.Mock{}
		useCaseMock.On("ExtractFilterDashboard").Return(&database.Filter{}, nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		handler.GetAllChartsByWorkspace(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Should return status OK when get all charts by repository", func(t *testing.T) {
		workspaceID := uuid.New()
		repositoryID := uuid.New()
		controllerMock := &controller.Mock{}
		controllerMock.On("GetAllDashboardCharts").Return(&response.Response{}, nil)
		useCaseMock := &dashboardfilter.Mock{}
		useCaseMock.On("ExtractFilterDashboard").Return(&database.Filter{}, nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		ctx.URLParams.Add("repositoryID", repositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		handler.GetAllChartsByRepository(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Should return status bad request when get all charts and wrong workspaceID", func(t *testing.T) {
		workspaceID := uuid.New()
		controllerMock := &controller.Mock{}
		controllerMock.On("GetAllDashboardCharts").Return(&response.Response{}, nil)
		useCaseMock := &dashboardfilter.Mock{}
		useCaseMock.On("ExtractFilterDashboard").Return(&database.Filter{}, enums.ErrorWrongWorkspaceID)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		handler.GetAllChartsByWorkspace(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("Should return status internal server error when get all charts", func(t *testing.T) {
		workspaceID := uuid.New()
		controllerMock := &controller.Mock{}
		controllerMock.On("GetAllDashboardCharts").Return(&response.Response{}, errors.New("unexpected error"))
		useCaseMock := &dashboardfilter.Mock{}
		useCaseMock.On("ExtractFilterDashboard").Return(&database.Filter{}, nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		handler.GetAllChartsByWorkspace(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

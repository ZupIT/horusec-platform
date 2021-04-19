package dashboardfilter

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/analytic/internal/enums"
)

func TestUseCaseDashboard_ExtractFilterDashboardByRepository(t *testing.T) {
	t.Run("Should get filter of dashboard from workspace without error", func(t *testing.T) {
		workspaceID := uuid.New()
		initialDate := "2006-01-02T15:04:05Z"
		finalDate := "2006-01-02T15:04:05Z"
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%s&size=%s", initialDate, finalDate, "", "")
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		filter, err := NewUseCaseDashboard().ExtractFilterDashboardByWorkspace(r)
		assert.NoError(t, err)
		assert.NotEmpty(t, filter)
	})
	t.Run("Should get filter of dashboard from workspace without error and not get size default", func(t *testing.T) {
		workspaceID := uuid.New()
		initialDate := "2006-01-02T15:04:05Z"
		finalDate := "2006-01-02T15:04:05Z"
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%s&size=%s", initialDate, finalDate, "", "15")
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		filter, err := NewUseCaseDashboard().ExtractFilterDashboardByWorkspace(r)
		assert.NoError(t, err)
		assert.NotEmpty(t, filter)
	})
	t.Run("Should get filter of dashboard from workspace with error from workspace invalid", func(t *testing.T) {
		initialDate := "2006-01-02T15:04:05Z"
		finalDate := "2006-01-02T15:04:05Z"
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%s&size=%s", initialDate, finalDate, "", "")
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "invalid")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		filter, err := NewUseCaseDashboard().ExtractFilterDashboardByWorkspace(r)
		assert.Equal(t, enums.ErrorWrongWorkspaceID, err)
		assert.Empty(t, filter)
	})
	t.Run("Should get filter of dashboard from workspace with error from initialDate invalid", func(t *testing.T) {
		workspaceID := uuid.New()
		initialDate := "invalid"
		finalDate := "2006-01-02T15:04:05Z"
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%s&size=%s", initialDate, finalDate, "", "")
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		filter, err := NewUseCaseDashboard().ExtractFilterDashboardByWorkspace(r)
		assert.Equal(t, enums.ErrorWrongInitialDate, err)
		assert.Empty(t, filter)
	})
	t.Run("Should get filter of dashboard from workspace with error from initialDate invalid", func(t *testing.T) {
		workspaceID := uuid.New()
		initialDate := "2006-01-02T15:04:05Z"
		finalDate := "invalid"
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%s&size=%s", initialDate, finalDate, "", "")
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		filter, err := NewUseCaseDashboard().ExtractFilterDashboardByWorkspace(r)
		assert.Equal(t, enums.ErrorWrongFinalDate, err)
		assert.Empty(t, filter)
	})
	t.Run("Should get filter of dashboard from workspace with error from initialDate blank", func(t *testing.T) {
		workspaceID := uuid.New()
		initialDate := ""
		finalDate := "2006-01-02T15:04:05Z"
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%s&size=%s", initialDate, finalDate, "", "")
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		_, err := NewUseCaseDashboard().ExtractFilterDashboardByWorkspace(r)
		assert.Contains(t, err.Error(), "InitialDate: cannot be blank.")
	})
	t.Run("Should get filter of dashboard from workspace with error from finalDate blank", func(t *testing.T) {
		workspaceID := uuid.New()
		initialDate := "2006-01-02T15:04:05Z"
		finalDate := ""
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%s&size=%s", initialDate, finalDate, "", "")
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		_, err := NewUseCaseDashboard().ExtractFilterDashboardByWorkspace(r)
		assert.Contains(t, err.Error(), "FinalDate: cannot be blank.")
	})
	t.Run("Should get filter of dashboard from repository without error", func(t *testing.T) {
		workspaceID := uuid.New()
		repositoryID := uuid.New()
		initialDate := "2006-01-02T15:04:05Z"
		finalDate := "2006-01-02T15:04:05Z"
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%s&size=%s", initialDate, finalDate, "", "")
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		ctx.URLParams.Add("repositoryID", repositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		filter, err := NewUseCaseDashboard().ExtractFilterDashboardByRepository(r)
		assert.NoError(t, err)
		assert.NotEmpty(t, filter)
	})
	t.Run("Should get filter of dashboard from repository with error when exists workspaceID invalid", func(t *testing.T) {
		repositoryID := uuid.New()
		initialDate := "2006-01-02T15:04:05Z"
		finalDate := "2006-01-02T15:04:05Z"
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%s&size=%s", initialDate, finalDate, "", "")
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")
		ctx.URLParams.Add("repositoryID", repositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		_, err := NewUseCaseDashboard().ExtractFilterDashboardByRepository(r)
		assert.Equal(t, enums.ErrorWrongWorkspaceID, err)
	})
	t.Run("Should get filter of dashboard from repository with error of repositoryID invalid", func(t *testing.T) {
		workspaceID := uuid.New()
		initialDate := "2006-01-02T15:04:05Z"
		finalDate := "2006-01-02T15:04:05Z"
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%s&size=%s", initialDate, finalDate, "", "")
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", workspaceID.String())
		ctx.URLParams.Add("repositoryID", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		filter, err := NewUseCaseDashboard().ExtractFilterDashboardByRepository(r)
		assert.Equal(t, enums.ErrorWrongRepositoryID, err)
		assert.Empty(t, filter)
	})
}

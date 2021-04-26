package dashboardfilter

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/enums"
)

func TestUseCaseDashboard_ExtractFilterDashboard(t *testing.T) {
	layoutDateTime := "2006-01-02T15:04:05Z"
	t.Run("Should extract filter with repository", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			expected.StartTime.Format(layoutDateTime), expected.EndTime.Format(layoutDateTime), expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		filter, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.NoError(t, err)
		assert.Equal(t, expected.RepositoryID, filter.RepositoryID)
		assert.Equal(t, expected.WorkspaceID, filter.WorkspaceID)
		assert.Equal(t, expected.StartTime, filter.StartTime)
		assert.Equal(t, expected.EndTime, filter.EndTime)
		assert.Equal(t, expected.Page, filter.Page)
		assert.Equal(t, expected.Size, filter.Size)
	})
	t.Run("Should extract filter without page size", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s",
			expected.StartTime.Format(layoutDateTime), expected.EndTime.Format(layoutDateTime))
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		filter, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.NoError(t, err)
		assert.Equal(t, expected.RepositoryID, filter.RepositoryID)
		assert.Equal(t, expected.WorkspaceID, filter.WorkspaceID)
		assert.Equal(t, expected.StartTime, filter.StartTime)
		assert.Equal(t, expected.EndTime, filter.EndTime)
		assert.Equal(t, 0, filter.Page)
		assert.Equal(t, 10, filter.Size)
	})
	t.Run("Should return error on extract filter because not send empty startTime", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?page=%v&size=%v",
			expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		_, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "EndTime: cannot be blank; StartTime: cannot be blank.")
	})
	t.Run("Should return error on extract filter because wrong workspaceID", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			expected.StartTime.Format(layoutDateTime), expected.EndTime.Format(layoutDateTime), expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "wrong type")
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		_, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.Error(t, err)
		assert.Equal(t, err, enums.ErrorWrongWorkspaceID)
	})
	t.Run("Should return error on extract filter because wrong repositoryID", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			expected.StartTime.Format(layoutDateTime), expected.EndTime.Format(layoutDateTime), expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", "wrong type")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		_, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.Error(t, err)
		assert.Equal(t, err, enums.ErrorWrongRepositoryID)
	})
	t.Run("Should return error on extract filter because wrong startTime", func(t *testing.T) {
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=wrongStartTime&finalDate=%s&page=%v&size=%v",
			expected.EndTime.Format(layoutDateTime), expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		_, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.Error(t, err)
		assert.Equal(t, err, enums.ErrorWrongInitialDate)
	})
	t.Run("Should return error on extract filter because wrong endTime", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		expected := &dashboard.FilterDashboard{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=wrongStartTime&page=%v&size=%v",
			expected.StartTime.Format(layoutDateTime), expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		_, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.Error(t, err)
		assert.Equal(t, err, enums.ErrorWrongFinalDate)
	})
}

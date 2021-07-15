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
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
)

func TestGetRepositoryFilter(t *testing.T) {
	t.Run("should success get filter", func(t *testing.T) {
		filter := &Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
			Page:         20,
			Size:         20,
		}

		args := filter.GetRepositoryFilter()
		assert.NotEmpty(t, args)
	})
}

func TestGetWorkspaceFilter(t *testing.T) {
	t.Run("should success get filter", func(t *testing.T) {
		filter := &Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
			Page:         20,
			Size:         20,
		}

		args := filter.GetWorkspaceFilter()
		assert.NotEmpty(t, args)
	})
}

func TestGetDateFilter(t *testing.T) {
	t.Run("should success get filter", func(t *testing.T) {
		filter := &Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
			Page:         20,
			Size:         20,
		}

		condition, args := filter.GetDateFilter()
		assert.NotEmpty(t, args)
		assert.Equal(t, "AND created_at >= @startTime AND created_at <= @endTime ", condition)
	})
}

func TestValidate(t *testing.T) {
	t.Run("should return no error when valid filter", func(t *testing.T) {
		filter := &Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    time.Now(),
			EndTime:      time.Now(),
			Page:         20,
			Size:         20,
		}

		assert.NoError(t, filter.Validate())
	})

	t.Run("should return error when invalid filter", func(t *testing.T) {
		filter := &Filter{}

		assert.Error(t, filter.Validate())
	})
}

func TestSetDateRangeAndPagination(t *testing.T) {
	layoutDateTime := "2006-01-02T15:04:05Z"
	startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
	endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")

	t.Run("should success set pagination and date range", func(t *testing.T) {
		filter := &Filter{}

		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			startTime.Format(layoutDateTime), endTime.Format(layoutDateTime), 18, 18)

		ctx := chi.NewRouteContext()
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		assert.NoError(t, filter.SetDateRangeAndPagination(r))
		assert.Equal(t, startTime, filter.StartTime)
		assert.Equal(t, endTime, filter.EndTime)
		assert.Equal(t, 18, filter.Page)
		assert.Equal(t, 18, filter.Size)
	})

	t.Run("should success set pagination and date range when size is lower than min", func(t *testing.T) {
		filter := &Filter{}

		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			startTime.Format(layoutDateTime), endTime.Format(layoutDateTime), 18, 0)

		ctx := chi.NewRouteContext()
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		assert.NoError(t, filter.SetDateRangeAndPagination(r))
		assert.Equal(t, startTime, filter.StartTime)
		assert.Equal(t, endTime, filter.EndTime)
		assert.Equal(t, 18, filter.Page)
		assert.Equal(t, dashboardEnums.DefaultPaginationSize, filter.Size)
	})

	t.Run("should return error when failed to parse initial date", func(t *testing.T) {
		filter := &Filter{}

		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			"test", "test", 0, 0)

		ctx := chi.NewRouteContext()
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		assert.Error(t, filter.SetDateRangeAndPagination(r))
	})

	t.Run("should return error when failed to parse final date", func(t *testing.T) {
		filter := &Filter{}

		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			startTime.Format(layoutDateTime), "test", 0, 0)

		ctx := chi.NewRouteContext()
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		assert.Error(t, filter.SetDateRangeAndPagination(r))
	})

	t.Run("should success set pagination without date", func(t *testing.T) {
		filter := &Filter{}

		url := fmt.Sprintf("/test?&page=%v&size=%v", 18, 18)

		ctx := chi.NewRouteContext()
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		assert.NoError(t, filter.SetDateRangeAndPagination(r))
		assert.Equal(t, 18, filter.Page)
		assert.Equal(t, 18, filter.Size)
	})
}

func TestSetWorkspaceAndRepositoryID(t *testing.T) {
	t.Run("should success workspace and repository id", func(t *testing.T) {
		filter := &Filter{}

		id := uuid.New()

		r, _ := http.NewRequest(http.MethodGet, "/test", nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", id.String())
		ctx.URLParams.Add("repositoryID", id.String())

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		assert.NoError(t, filter.SetWorkspaceAndRepositoryID(r))
		assert.Equal(t, id, filter.WorkspaceID)
		assert.Equal(t, id, filter.RepositoryID)
	})

	t.Run("should return error when invalid repository id", func(t *testing.T) {
		filter := &Filter{}

		id := uuid.New()

		r, _ := http.NewRequest(http.MethodGet, "/test", nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", id.String())
		ctx.URLParams.Add("repositoryID", "test")

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		err := filter.SetWorkspaceAndRepositoryID(r)
		assert.Error(t, err)
		assert.Equal(t, dashboardEnums.ErrorInvalidRepositoryID, err)
	})

	t.Run("should return error when invalid workspaces id", func(t *testing.T) {
		filter := &Filter{}

		r, _ := http.NewRequest(http.MethodGet, "/test", nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "test")

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		err := filter.SetWorkspaceAndRepositoryID(r)
		assert.Error(t, err)
		assert.Equal(t, dashboardEnums.ErrorInvalidWorkspaceID, err)
	})
}

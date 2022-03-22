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
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
)

type Filter struct {
	RepositoryID uuid.UUID
	WorkspaceID  uuid.UUID
	StartTime    time.Time
	EndTime      time.Time
	Page         int
	Size         int
}

func (f *Filter) GetRepositoryFilter() interface{} {
	return sql.Named("repositoryID", f.RepositoryID)
}

func (f *Filter) GetWorkspaceFilter() interface{} {
	return sql.Named("workspaceID", f.WorkspaceID)
}

func (f *Filter) GetDateFilter() (query string, args []interface{}) {
	if !f.StartTime.IsZero() {
		query += "AND created_at >= @startTime "
	}

	if !f.EndTime.IsZero() {
		query += "AND created_at <= @endTime "
	}

	return query, f.getDateFilterArgs()
}

func (f *Filter) getDateFilterArgs() []interface{} {
	return []interface{}{
		sql.Named("startTime", f.StartTime),
		sql.Named("endTime", f.EndTime),
		sql.Named("repositoryID", f.RepositoryID),
		sql.Named("workspaceID", f.WorkspaceID),
	}
}

func (f *Filter) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.WorkspaceID, validation.Required, validation.NotIn(uuid.Nil)),
		validation.Field(&f.StartTime, validation.Required),
		validation.Field(&f.EndTime, validation.Required),
		validation.Field(&f.Page, validation.Min(0)),
		validation.Field(&f.Size, validation.Min(dashboard.DefaultPaginationSize)),
	)
}

func (f *Filter) SetDateRangeAndPagination(request *http.Request) error {
	initialDate, err := f.parseDate(request.URL.Query().Get(dashboard.InitialDateHeader))
	if err != nil {
		return errors.Wrap(err, dashboard.MessageInvalidInitialDate)
	}

	finalDate, err := f.parseDate(request.URL.Query().Get(dashboard.FinalDateHeader))
	if err != nil {
		return errors.Wrap(err, dashboard.MessageInvalidFinalDate)
	}

	f.StartTime = initialDate
	f.EndTime = finalDate
	f.setPageAndSize(request)

	return nil
}

func (f *Filter) parseDate(date string) (time.Time, error) {
	if date != "" {
		return time.Parse("2006-01-02T15:04:05Z", date)
	}

	return time.Time{}, nil
}

func (f *Filter) setPageAndSize(request *http.Request) {
	page, _ := strconv.Atoi(request.URL.Query().Get(dashboard.PageHeader))
	size, _ := strconv.Atoi(request.URL.Query().Get(dashboard.SizeHeader))

	f.Page = page
	f.Size = f.getPaginationMinSize(size)
}

func (f *Filter) getPaginationMinSize(size int) int {
	if size < dashboard.DefaultPaginationSize {
		return dashboard.DefaultPaginationSize
	}

	return size
}

func (f *Filter) SetWorkspaceAndRepositoryID(request *http.Request) (err error) {
	f.WorkspaceID, err = uuid.Parse(chi.URLParam(request, dashboard.WorkspaceID))
	if err != nil {
		return dashboard.ErrorInvalidWorkspaceID
	}

	if chi.URLParam(request, dashboard.RepositoryID) != "" {
		f.RepositoryID, err = uuid.Parse(chi.URLParam(request, dashboard.RepositoryID))
		if err != nil {
			return dashboard.ErrorInvalidRepositoryID
		}
	}

	return nil
}

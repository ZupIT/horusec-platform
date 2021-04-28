package dashboardfilter

import (
	netHTTP "net/http"
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/enums"
)

const DefaultSize = 10

type IUseCaseDashboard interface {
	ExtractFilterDashboard(r *netHTTP.Request) (*dashboard.FilterDashboard, error)
}
type UseCaseDashboard struct{}

func NewUseCaseDashboard() IUseCaseDashboard {
	return &UseCaseDashboard{}
}

func (u *UseCaseDashboard) ExtractFilterDashboard(r *netHTTP.Request) (*dashboard.FilterDashboard, error) {
	filter, err := u.extractFilterDashboardByWorkspace(r)
	if err != nil {
		return nil, err
	}
	if chi.URLParam(r, "repositoryID") != "" {
		filter.RepositoryID, err = uuid.Parse(chi.URLParam(r, "repositoryID"))
		if err != nil {
			return nil, enums.ErrorWrongRepositoryID
		}
	}
	return filter, u.validateFilterDashboard(filter, true)
}

func (u *UseCaseDashboard) extractFilterDashboardByWorkspace(r *netHTTP.Request) (*dashboard.FilterDashboard, error) {
	workspaceID, err := uuid.Parse(chi.URLParam(r, "workspaceID"))
	if err != nil {
		return nil, enums.ErrorWrongWorkspaceID
	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	initialDate, finalDate, err := u.getDateRangeFromRequestQuery(r)
	if err != nil {
		return nil, err
	}
	entity := u.buildFilterDasboard(workspaceID, initialDate, finalDate, page, u.getSizeOrMin(size))
	return entity, u.validateFilterDashboard(entity, false)
}

func (u *UseCaseDashboard) getSizeOrMin(size int) int {
	if size < DefaultSize {
		return DefaultSize
	}
	return size
}

func (u *UseCaseDashboard) getDateRangeFromRequestQuery(r *netHTTP.Request) (time.Time, time.Time, error) {
	initial, err := u.getDateFromRequestQuery(r, "initialDate")
	if err != nil {
		return time.Time{}, time.Time{}, enums.ErrorWrongInitialDate
	}

	final, err := u.getDateFromRequestQuery(r, "finalDate")
	if err != nil {
		return time.Time{}, time.Time{}, enums.ErrorWrongFinalDate
	}

	return initial, final, nil
}

func (u *UseCaseDashboard) getDateFromRequestQuery(r *netHTTP.Request, queryStrKey string) (time.Time, error) {
	date := r.URL.Query().Get(queryStrKey)
	if date != "" {
		return time.Parse("2006-01-02T15:04:05Z", date)
	}

	return time.Time{}, nil
}

func (u *UseCaseDashboard) validateFilterDashboard(
	entity *dashboard.FilterDashboard, isFilterFromRepository bool) error {
	return validation.ValidateStruct(entity,
		validation.Field(&entity.WorkspaceID, validation.Required, validation.NotIn(uuid.Nil)),
		validation.Field(&entity.RepositoryID, validation.When(isFilterFromRepository, validation.NotIn(uuid.Nil))),
		validation.Field(&entity.StartTime, validation.Required),
		validation.Field(&entity.EndTime, validation.Required),
		validation.Field(&entity.Page, validation.Min(0)),
		validation.Field(&entity.Size, validation.Min(DefaultSize)),
	)
}

func (u *UseCaseDashboard) buildFilterDasboard(workspaceID uuid.UUID, initialDate, finalDate time.Time,
	page int, size int) *dashboard.FilterDashboard {
	return &dashboard.FilterDashboard{
		WorkspaceID: workspaceID,
		StartTime:   initialDate,
		EndTime:     finalDate,
		Page:        page,
		Size:        size,
	}
}

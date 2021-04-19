package dashboardworkspace

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	controllerDashboard "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	dashboardfilter "github.com/ZupIT/horusec-platform/analytic/internal/usecase/dashboard_filter"
)

func TestHandler_GetTotalDevelopers(t *testing.T) {
	t.Run("Should return success when call GetTotalDevelopers", func(t *testing.T) {
		controllerMock := &controllerDashboard.Mock{}
		controllerMock.On("GetTotalDevelopers").Return(nil, nil)
		useCaseMock := &dashboardfilter.Mock{}
		useCaseMock.On("ExtractFilterDashboardByWorkspace").Return(&dashboard.FilterDashboard{
			RepositoryID: uuid.Nil,
			WorkspaceID:  uuid.New(),
			InitialDate:  time.Now(),
			FinalDate:    time.Now().Add((24 * time.Hour) * 90),
			Page:         0,
			Size:         10,
		}, nil)
		handler := &Handler{
			controller: controllerMock,
			useCase:    useCaseMock,
		}
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		handler.GetTotalDevelopers(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Should return bad request when call GetTotalDevelopers with wrong workspaceID", func(t *testing.T) {

	})
	t.Run("Should return internal server error when call GetTotalDevelopers", func(t *testing.T) {

	})

}
func TestHandler_GetTotalRepositories(t *testing.T) {
	t.Run("Should return success when call GetTotalRepositories", func(t *testing.T) {

	})
	t.Run("Should return bad request when call GetTotalRepositories with wrong workspaceID", func(t *testing.T) {

	})
	t.Run("Should return internal server error when call GetTotalRepositories", func(t *testing.T) {

	})

}
func TestHandler_GetVulnBySeverity(t *testing.T) {
	t.Run("Should return success when call GetVulnBySeverity", func(t *testing.T) {

	})
	t.Run("Should return bad request when call GetVulnBySeverity with wrong workspaceID", func(t *testing.T) {

	})
	t.Run("Should return internal server error when call GetVulnBySeverity", func(t *testing.T) {

	})

}
func TestHandler_GetVulnByDeveloper(t *testing.T) {
	t.Run("Should return success when call GetVulnByDeveloper", func(t *testing.T) {

	})
	t.Run("Should return bad request when call GetVulnByDeveloper with wrong workspaceID", func(t *testing.T) {

	})
	t.Run("Should return internal server error when call GetVulnByDeveloper", func(t *testing.T) {

	})

}
func TestHandler_GetVulnByRepository(t *testing.T) {
	t.Run("Should return success when call GetVulnByRepository", func(t *testing.T) {

	})
	t.Run("Should return bad request when call GetVulnByRepository with wrong workspaceID", func(t *testing.T) {

	})
	t.Run("Should return internal server error when call GetVulnByRepository", func(t *testing.T) {

	})

}
func TestHandler_GetVulnByLanguage(t *testing.T) {
	t.Run("Should return success when call GetVulnByLanguage", func(t *testing.T) {

	})
	t.Run("Should return bad request when call GetVulnByLanguage with wrong workspaceID", func(t *testing.T) {

	})
	t.Run("Should return internal server error when call GetVulnByLanguage", func(t *testing.T) {

	})

}
func TestHandler_GetVulnByTime(t *testing.T) {
	t.Run("Should return success when call GetVulnByTime", func(t *testing.T) {

	})
	t.Run("Should return bad request when call GetVulnByTime with wrong workspaceID", func(t *testing.T) {

	})
	t.Run("Should return internal server error when call GetVulnByTime", func(t *testing.T) {

	})

}
func TestHandler_GetVulnDetails(t *testing.T) {
	t.Run("Should return success when call GetVulnDetails", func(t *testing.T) {

	})
	t.Run("Should return bad request when call GetVulnDetails with wrong workspaceID", func(t *testing.T) {

	})
	t.Run("Should return internal server error when call GetVulnDetails", func(t *testing.T) {

	})

}

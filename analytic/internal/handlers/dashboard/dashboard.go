package dashboard

import (
	"net/http"

	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"

	controller "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	_ "github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard" // [swagger-usage]
	useCase "github.com/ZupIT/horusec-platform/analytic/internal/usecases/dashboard"
)

type Handler struct {
	controller controller.IController
	useCase    useCase.IUseCases
}

func NewDashboardHandler(dashboardController controller.IController) *Handler {
	return &Handler{
		controller: dashboardController,
		useCase:    useCase.NewUseCaseDashboard(),
	}
}

func (h *Handler) Options(w http.ResponseWriter, _ *http.Request) {
	httpUtil.StatusNoContent(w)
}

// GetAllChartsByWorkspace
// @Tags Dashboard
// @Security ApiKeyAuth
// @Description Get all charts of dashboard screen
// @ID GetAllChartsByWorkspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=dashboard.Response} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID} [get]
func (h *Handler) GetAllChartsByWorkspace(w http.ResponseWriter, r *http.Request) {
	filter, err := h.useCase.FilterFromRequest(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	result, err := h.controller.GetAllDashboardChartsWorkspace(filter)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, result)
}

// GetAllChartsByRepository
// @Tags Dashboard
// @Security ApiKeyAuth
// @Description Get all charts of dashboard screen
// @ID GetAllChartsByRepository
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param repositoryID path string true "repositoryID of the repository"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=dashboard.Response} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/{repositoryID} [get]
func (h *Handler) GetAllChartsByRepository(w http.ResponseWriter, r *http.Request) {
	filter, err := h.useCase.FilterFromRequest(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	result, err := h.controller.GetAllDashboardChartsRepository(filter)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, result)
}

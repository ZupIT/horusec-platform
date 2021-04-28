package dashboard

import (
	netHTTP "net/http"

	_ "github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard" // [swagger-usage]

	controller "github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	useCase "github.com/ZupIT/horusec-platform/analytic/internal/usecase/dashboard_filter"

	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
)

type Handler struct {
	controller controller.IReadController
	useCase    useCase.IUseCaseDashboard
}

func NewDashboardHandler(dashboardController controller.IReadController) *Handler {
	return &Handler{
		controller: dashboardController,
		useCase:    useCase.NewUseCaseDashboard(),
	}
}

func (h *Handler) Options(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
}

// GetAllChartsByWorkspace
// @Tags DashboardByWorkspace
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
// @Router /analytic/dashboard/{workspaceID}/dashboard-charts [get]
func (h *Handler) GetAllChartsByWorkspace(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	h.getAllCharts(w, r)
}

// GetAllChartsByRepository
// @Tags DashboardByWorkspace
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
// @Router /analytic/dashboard/{workspaceID}/{repositoryID}/dashboard-charts [get]
func (h *Handler) GetAllChartsByRepository(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	h.getAllCharts(w, r)
}

func (h *Handler) getAllCharts(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboard(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	result, err := h.controller.GetAllCharts(filter)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusOK(w, result)
}

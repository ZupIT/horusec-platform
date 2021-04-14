package dashboardworkspace

import (
	netHTTP "net/http"

	"github.com/ZupIT/horusec-platform/analytic/internal/controllers/dashboard"
	dashboardfilter "github.com/ZupIT/horusec-platform/analytic/internal/usecase/dashboard_filter"

	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
)

type Handler struct {
	controller dashboard.IController
	useCase    dashboardfilter.IUseCaseDashboard
}

func NewDashboardWorkspaceHandler(controller dashboard.IController) *Handler {
	return &Handler{
		controller: controller,
		useCase:    dashboardfilter.NewUseCaseDashboard(),
	}
}

func (h *Handler) Options(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
}

// GetTotalDevelopers
// @Tags DashboardByWorkspace
// @Security ApiKeyAuth
// @Description Get count total vulnerabilities by range date in workspace
// @ID get-total-developers-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=int} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/total-developers [get]
func (h *Handler) GetTotalDevelopers(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByWorkspace(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	result, err := h.controller.GetTotalDevelopers(filter)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusOK(w, result)
}

// GetTotalRepositories
// @Tags DashboardByWorkspace
// @Security ApiKeyAuth
// @Description Get count total repositories by range date in workspace
// @ID get-total-repositories-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=int} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/total-repositories [get]
func (h *Handler) GetTotalRepositories(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByWorkspace(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	result, err := h.controller.GetTotalRepositories(filter)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusOK(w, result)
}

// GetVulnBySeverity
// @Tags DashboardByWorkspace
// @Security ApiKeyAuth
// @Description Get all vulnerabilities by severities by range date in workspace
// @ID get-vulnerabilities-by-severity-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=object} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/vulnerabilities-by-severities [get]
func (h *Handler) GetVulnBySeverity(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByWorkspace(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	result, err := h.controller.GetVulnBySeverity(filter)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusOK(w, result)
}

// GetVulnByDeveloper
// @Tags DashboardByWorkspace
// @Security ApiKeyAuth
// @Description Get all vulnerabilities by top 5 developer by range date in workspace
// @ID get-vulnerabilities-by-top-5-developers-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=object} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/vulnerabilities-by-developer [get]
func (h *Handler) GetVulnByDeveloper(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByWorkspace(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	result, err := h.controller.GetVulnByDeveloper(filter)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusOK(w, result)
}

// GetVulnByRepository
// @Tags DashboardByWorkspace
// @Security ApiKeyAuth
// @Description Get all vulnerabilities by top 5 repositories by range date in workspace
// @ID get-vulnerabilities-by-top-5-repositories-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=object} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/vulnerabilities-by-repositories [get]
func (h *Handler) GetVulnByRepository(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByWorkspace(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	result, err := h.controller.GetVulnByRepository(filter)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusOK(w, result)
}

// GetVulnByLanguage
// @Tags DashboardByWorkspace
// @Security ApiKeyAuth
// @Description Get all vulnerabilities by top 5 languages by range date in workspace
// @ID get-vulnerabilities-by-top-5-languages-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=object} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/vulnerabilities-by-languages [get]
func (h *Handler) GetVulnByLanguage(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByWorkspace(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	result, err := h.controller.GetVulnByLanguage(filter)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusOK(w, result)
}

// GetVulnByTime
// @Tags DashboardByWorkspace
// @Security ApiKeyAuth
// @Description Get count vulnerabilities by time by range date in workspace
// @ID get-vulnerabilities-by-time-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=object} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/vulnerabilities-by-time [get]
func (h *Handler) GetVulnByTime(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByWorkspace(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	result, err := h.controller.GetVulnByTime(filter)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusOK(w, result)
}

// GetVulnDetails
// @Tags DashboardByWorkspace
// @Security ApiKeyAuth
// @Description Get all vulnerabilities by details by range date in workspace
// @ID get-vulnerabilities-by-details-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Param page query string false "page query string"
// @Param size query string false "size query string"
// @Success 200 {object} entities.Response{content=object} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/details [get]
func (h *Handler) GetVulnDetails(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByWorkspace(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	result, err := h.controller.GetVulnDetails(filter)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusOK(w, result)
}

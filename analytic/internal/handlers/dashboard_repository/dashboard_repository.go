package dashboardrepository

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

func NewDashboardRepositoryHandler(controller dashboard.IController) *Handler {
	return &Handler{
		controller: controller,
		useCase:    dashboardfilter.NewUseCaseDashboard(),
	}
}

func (h *Handler) Options(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
}

// GetTotalDevelopers
// @Tags DashboardByRepository
// @Security ApiKeyAuth
// @Description Get count total vulnerabilities by range date in repository
// @ID get-total-developers-repository
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param repositoryID path string true "repositoryID of the repository"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=int} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/{repositoryID}/total-developers [get]
func (h *Handler) GetTotalDevelopers(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByRepository(r)
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

// GetVulnBySeverity
// @Tags DashboardByRepository
// @Security ApiKeyAuth
// @Description Get all vulnerabilities by severities by range date in repository
// @ID get-vulnerabilities-by-severity-repository
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param repositoryID path string true "repositoryID of the repository"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=object} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/{repositoryID}/vulnerabilities-by-severities [get]
func (h *Handler) GetVulnBySeverity(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByRepository(r)
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
// @Tags DashboardByRepository
// @Security ApiKeyAuth
// @Description Get all vulnerabilities by top 5 developer by range date in repository
// @ID get-vulnerabilities-by-top-5-developers-repository
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param repositoryID path string true "repositoryID of the repository"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=object} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/{repositoryID}/vulnerabilities-by-developer [get]
func (h *Handler) GetVulnByDeveloper(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByRepository(r)
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

// GetVulnByLanguage
// @Tags DashboardByRepository
// @Security ApiKeyAuth
// @Description Get all vulnerabilities by top 5 languages by range date in repository
// @ID get-vulnerabilities-by-top-5-languages-repository
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param repositoryID path string true "repositoryID of the repository"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=object} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/{repositoryID}/vulnerabilities-by-languages [get]
func (h *Handler) GetVulnByLanguage(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByRepository(r)
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
// @Tags DashboardByRepository
// @Security ApiKeyAuth
// @Description Get count vulnerabilities by time by range date in repository
// @ID get-vulnerabilities-by-time-repository
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param repositoryID path string true "repositoryID of the repository"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Success 200 {object} entities.Response{content=object} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/{repositoryID}/vulnerabilities-by-time [get]
func (h *Handler) GetVulnByTime(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByRepository(r)
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
// @Tags DashboardByRepository
// @Security ApiKeyAuth
// @Description Get all vulnerabilities by details by range date in repository
// @ID get-vulnerabilities-by-details-repository
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param repositoryID path string true "repositoryID of the repository"
// @Param initialDate query string false "initialDate query string"
// @Param finalDate query string false "finalDate query string"
// @Param page query string false "page query string"
// @Param size query string false "size query string"
// @Success 200 {object} entities.Response{content=object} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /analytic/dashboard/{workspaceID}/{repositoryID}/details [get]
func (h *Handler) GetVulnDetails(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	filter, err := h.useCase.ExtractFilterDashboardByRepository(r)
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

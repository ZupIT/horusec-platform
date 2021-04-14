package dashboardworkspace

import (
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	netHTTP "net/http"
)

type Handler struct {
}

func NewDashboardWorkspaceHandler() *Handler {
	return &Handler{
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
func (h *Handler) GetTotalDevelopers(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetTotalRepositories(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetVulnBySeverity(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetVulnByDeveloper(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetVulnByRepository(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetVulnByLanguage(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetVulnByTime(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetVulnDetails(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
}
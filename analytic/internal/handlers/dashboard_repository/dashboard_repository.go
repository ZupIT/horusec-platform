package dashboardrepository

import (
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	netHTTP "net/http"
)

type Handler struct {
}

func NewDashboardRepositoryHandler() *Handler {
	return &Handler{
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
func (h *Handler) GetTotalDevelopers(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetVulnBySeverity(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetVulnByDeveloper(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetVulnByLanguage(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetVulnByTime(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
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
func (h *Handler) GetVulnDetails(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
}
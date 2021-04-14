package dashboard

import (
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
)

type IController interface {
	GetTotalDevelopers(filter *dashboard.FilterDashboard) (interface{}, error)
	GetTotalRepositories(filter *dashboard.FilterDashboard) (interface{}, error)
	GetVulnBySeverity(filter *dashboard.FilterDashboard) (interface{}, error)
	GetVulnByDeveloper(filter *dashboard.FilterDashboard) (interface{}, error)
	GetVulnByRepository(filter *dashboard.FilterDashboard) (interface{}, error)
	GetVulnByLanguage(filter *dashboard.FilterDashboard) (interface{}, error)
	GetVulnByTime(filter *dashboard.FilterDashboard) (interface{}, error)
	GetVulnDetails(filter *dashboard.FilterDashboard) (interface{}, error)
}

type Controller struct {
	repoDashboard repoDashboard.IRepoDashboard
}

func NewControllerDashboard(repositoryDashboard repoDashboard.IRepoDashboard) IController {
	return &Controller{
		repoDashboard: repositoryDashboard,
	}
}

func (c *Controller) GetTotalDevelopers(filter *dashboard.FilterDashboard) (interface{}, error) {
	return nil, nil
}
func (c *Controller) GetTotalRepositories(filter *dashboard.FilterDashboard) (interface{}, error) {
	return nil, nil
}
func (c *Controller) GetVulnBySeverity(filter *dashboard.FilterDashboard) (interface{}, error) {
	return nil, nil
}
func (c *Controller) GetVulnByDeveloper(filter *dashboard.FilterDashboard) (interface{}, error) {
	return nil, nil
}
func (c *Controller) GetVulnByRepository(filter *dashboard.FilterDashboard) (interface{}, error) {
	return nil, nil
}
func (c *Controller) GetVulnByLanguage(filter *dashboard.FilterDashboard) (interface{}, error) {
	return nil, nil
}
func (c *Controller) GetVulnByTime(filter *dashboard.FilterDashboard) (interface{}, error) {
	return nil, nil
}
func (c *Controller) GetVulnDetails(filter *dashboard.FilterDashboard) (interface{}, error) {
	return nil, nil
}

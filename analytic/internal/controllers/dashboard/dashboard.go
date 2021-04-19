package dashboard

import (
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	repoDashboard "github.com/ZupIT/horusec-platform/analytic/internal/repositories/dashboard"
)

type IController interface {
	GetTotalDevelopers(filter *dashboard.FilterDashboard) (int, error)
	GetTotalRepositories(filter *dashboard.FilterDashboard) (int, error)
	GetVulnBySeverity(filter *dashboard.FilterDashboard) (*dashboard.VulnerabilityBySeverity, error)
	GetVulnByDeveloper(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByDeveloper, error)
	GetVulnByRepository(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByRepository, error)
	GetVulnByLanguage(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByLanguage, error)
	GetVulnByTime(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByTime, error)
	GetVulnDetails(filter *dashboard.FilterDashboard) (*dashboard.VulnerabilityDetails, error)
}

type Controller struct {
	repoDashboard repoDashboard.IRepoDashboard
}

func NewControllerDashboard(repositoryDashboard repoDashboard.IRepoDashboard) IController {
	return &Controller{
		repoDashboard: repositoryDashboard,
	}
}

func (c *Controller) GetTotalDevelopers(filter *dashboard.FilterDashboard) (int, error) {
	filter.Size = 0
	filter.Page = 0
	return c.repoDashboard.GetTotalDevelopers(filter)
}
func (c *Controller) GetTotalRepositories(filter *dashboard.FilterDashboard) (int, error) {
	filter.Size = 0
	filter.Page = 0
	return c.repoDashboard.GetTotalRepositories(filter)
}
func (c *Controller) GetVulnBySeverity(filter *dashboard.FilterDashboard) (*dashboard.VulnerabilityBySeverity, error) {
	filter.Size = 0
	filter.Page = 0
	return c.repoDashboard.GetVulnBySeverity(filter)
	//if err != nil {
	//	return nil, err
	//}
	//return c.decoratorListDTOToVulnerabilityBySeverity(listDTO), nil
}
func (c *Controller) GetVulnByDeveloper(
	filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByDeveloper, error) {
	filter.Size = 0
	filter.Page = 0
	return c.repoDashboard.GetVulnByDeveloper(filter)
}
func (c *Controller) GetVulnByRepository(
	filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByRepository, error) {
	filter.Size = 0
	filter.Page = 0
	return c.repoDashboard.GetVulnByRepository(filter)
}
func (c *Controller) GetVulnByLanguage(
	filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByLanguage, error) {
	filter.Size = 0
	filter.Page = 0
	return c.repoDashboard.GetVulnByLanguage(filter)
}
func (c *Controller) GetVulnByTime(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByTime, error) {
	filter.Size = 0
	filter.Page = 0
	return c.repoDashboard.GetVulnByTime(filter)
}
func (c *Controller) GetVulnDetails(filter *dashboard.FilterDashboard) (*dashboard.VulnerabilityDetails, error) {
	return c.repoDashboard.GetVulnDetails(filter)
}

//func (c *Controller) decoratorListDTOToVulnerabilityBySeverity(
//	listCount *[]dashboard.DTOVulnerabilityBySeverity) *dashboard.VulnerabilityBySeverity {
//	vuln := &dashboard.VulnerabilityBySeverity{}
//	for _, itemCount := range *listCount {
//		vuln = c.factorySetVulnCountBySeverity(itemCount, vuln)
//	}
//	return vuln
//}
//
//func (c *Controller) getDefaultVulnCount(itemCount dashboard.DTOVulnerabilityBySeverity) dashboard.VulnerabilityCount {
//	return dashboard.VulnerabilityCount{
//		Count: itemCount.Count,
//		Types: dashboard.VulnerabilityTypes{
//			Vulnerability: itemCount.Vulnerability,
//			RiskAccept:    itemCount.RiskAccept,
//			FalsePositive: itemCount.FalsePositive,
//			Corrected:     itemCount.Corrected,
//		},
//	}
//}
//
//// nolint:funlen,gocyclo // factory is necessary all severities conditions
//func (c *Controller) factorySetVulnCountBySeverity(itemCount dashboard.DTOVulnerabilityBySeverity,
//	vuln *dashboard.VulnerabilityBySeverity) *dashboard.VulnerabilityBySeverity {
//	switch itemCount.Severity {
//	case severities.Critical:
//		vuln.Critical = c.getDefaultVulnCount(itemCount)
//	case severities.High:
//		vuln.High = c.getDefaultVulnCount(itemCount)
//	case severities.Medium:
//		vuln.Medium = c.getDefaultVulnCount(itemCount)
//	case severities.Low:
//		vuln.Low = c.getDefaultVulnCount(itemCount)
//	case severities.Info:
//		vuln.Info = c.getDefaultVulnCount(itemCount)
//	case severities.Unknown:
//		vuln.Unknown = c.getDefaultVulnCount(itemCount)
//	}
//	return vuln
//}

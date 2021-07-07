package dashboard

import (
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
)

type IRepoDashboard interface {
	GetDashboardTotalDevelopers(filter *dashboard.Filter) (int, error)
	GetDashboardVulnBySeverity(filter *dashboard.Filter) (*dashboard.Vulnerability, error)
	GetDashboardVulnByAuthor(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByAuthor, error)
	GetDashboardVulnByLanguage(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByLanguage, error)
	GetDashboardVulnByTime(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByTime, error)
}

type RepoDashboard struct {
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
}

func NewRepoDashboard(connection *database.Connection) IRepoDashboard {
	return &RepoDashboard{
		databaseRead:  connection.Read,
		databaseWrite: connection.Write,
	}
}

func (r *RepoDashboard) GetDashboardTotalDevelopers(filter *dashboard.Filter) (count int, err error) {
	query := fmt.Sprintf(r.queryGetDashboardTotalDevelopers(), dashboardEnums.TableVulnerabilitiesByAuthor)

	return count, r.databaseRead.Raw(query, &count, filter.GetRepositoryFilter()).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardTotalDevelopers() string {
	return `
			SELECT COUNT(DISTINCT(author)) 
			FROM %[1]s
			WHERE repository_id = @repositoryID
			AND created_at = (SELECT MAX(created_at) FROM %[1]s WHERE repository_id = @repositoryID)  
	`
}

func (r *RepoDashboard) GetDashboardVulnBySeverity(filter *dashboard.Filter) (*dashboard.Vulnerability, error) {
	vulns := &dashboard.Vulnerability{}

	query := fmt.Sprintf(r.queryGetDashboardVulnBySeverity(), dashboardEnums.TableVulnerabilitiesByTime)

	return vulns, r.databaseRead.Raw(query, vulns, filter.GetRepositoryFilter()).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnBySeverity() string {
	return `
			SELECT *
			FROM %[1]s
			WHERE repository_id = @repositoryID
			AND vulnerability_id = (SELECT vulnerability_id FROM %[1]s ORDER BY created_at DESC LIMIT 1)
	`
}

func (r *RepoDashboard) GetDashboardVulnByAuthor(
	filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByAuthor, err error) {

	query := fmt.Sprintf(r.queryGetDashboardVulnByAuthor(), dashboardEnums.TableVulnerabilitiesByAuthor)

	return vulns, r.databaseRead.Raw(query, &vulns, filter.GetRepositoryFilter()).GetErrorExceptNotFound()
}

//nolint:funlen // need to be bigger than 15
func (r *RepoDashboard) queryGetDashboardVulnByAuthor() string {
	return `
			SELECT DISTINCT ON(author) author, *
			FROM %[1]s
			WHERE repository_id = @repositoryID
			AND created_at = (SELECT MAX(created_at) FROM %[1]s WHERE repository_id = @repositoryID)  
			LIMIT 5
	`
}

func (r *RepoDashboard) GetDashboardVulnByLanguage(
	filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByLanguage, err error) {
	query := fmt.Sprintf(r.queryGetDashboardVulnByLanguage(), dashboardEnums.TableVulnerabilitiesByLanguage)

	return vulns, r.databaseRead.Raw(query, &vulns, filter.GetRepositoryFilter()).GetErrorExceptNotFound()
}

//nolint:funlen // need to be bigger than 15
func (r *RepoDashboard) queryGetDashboardVulnByLanguage() string {
	return `
			SELECT DISTINCT ON(language) language, *
			FROM %[1]s
			WHERE repository_id = @repositoryID
			AND created_at = (SELECT MAX(created_at) FROM %[1]s WHERE repository_id = @repositoryID)  
			LIMIT 5
	`
}

func (r *RepoDashboard) GetDashboardVulnByTime(filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByTime, err error) {
	condition, args := filter.GetDateFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnByTime(), dashboardEnums.TableVulnerabilitiesByTime, condition)

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnByTime() string {
	return `
		SELECT DISTINCT ON(created_at) *
		FROM %[1]s
		WHERE repository_id = @repositoryID
		%[2]s
		AND created_at IN (
			SELECT MAX(created_at)
			FROM %[1]s
			WHERE repository_id = @repositoryID
			%[2]s
			GROUP BY created_at )		
	`
}

package dashboard

import (
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
)

type IRepoDashboard interface {
	GetDashboardTotalDevelopers(filter *dashboard.Filter) (int, error)
	GetDashboardTotalRepositories(filter *dashboard.Filter) (int, error)
	GetDashboardVulnBySeverity(filter *dashboard.Filter) (*dashboard.Vulnerability, error)
	GetDashboardVulnByAuthor(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByAuthor, error)
	GetDashboardVulnByRepository(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByRepository, error)
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
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardTotalDevelopers(), dashboardEnums.TableVulnerabilitiesByAuthor, condition)

	return count, r.databaseRead.Raw(query, &count, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardTotalDevelopers() string {
	return `
		SELECT COUNT(*) 
		FROM (
				SELECT DISTINCT ON(author) author
				FROM %[1]s
				WHERE %[2]s AND created_at 
				IN 
				(
					SELECT MAX(created_at) FROM vulnerabilities_by_language GROUP BY repository_id
				)
				ORDER BY author, created_at DESC
		) AS result
	`
}

func (r *RepoDashboard) GetDashboardTotalRepositories(filter *dashboard.Filter) (count int, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardTotalRepositories(),
		dashboardEnums.TableVulnerabilitiesByRepository, condition)

	return count, r.databaseRead.Raw(query, &count, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardTotalRepositories() string {
	return `
		SELECT COUNT(*) 
		FROM (
				SELECT DISTINCT ON(repository_id) repository_id
				FROM %[1]s
				WHERE %[2]s AND created_at 
				IN 
				(
					SELECT MAX(created_at) FROM vulnerabilities_by_language GROUP BY repository_id
				)
				ORDER BY repository_id, created_at DESC
		) AS result
	`
}

func (r *RepoDashboard) GetDashboardVulnBySeverity(filter *dashboard.Filter) (*dashboard.Vulnerability, error) {
	vulns := &dashboard.Vulnerability{}
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnBySeverity(), r.queryDefaultFields(),
		dashboardEnums.TableVulnerabilitiesByTime, condition)

	return vulns, r.databaseRead.Raw(query, vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnBySeverity() string {
	return `
		SELECT %[1]s
		FROM (
				SELECT DISTINCT ON(repository_id) *
				FROM %[2]s
				WHERE %[3]s AND created_at 
				IN 
				(
					SELECT MAX(created_at) FROM vulnerabilities_by_language GROUP BY repository_id
				)
				ORDER BY repository_id, created_at DESC
		) AS result
	`
}

func (r *RepoDashboard) GetDashboardVulnByAuthor(filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByAuthor, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnByAuthor(), r.queryDefaultFields(),
		dashboardEnums.TableVulnerabilitiesByAuthor, condition)

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnByAuthor() string {
	return `
		SELECT author, %[1]s
		FROM (
				SELECT *
				FROM %[2]s
				WHERE %[3]s AND created_at 
				IN 
				(
					SELECT MAX(created_at) FROM vulnerabilities_by_language GROUP BY repository_id
				)
				ORDER BY repository_id, created_at DESC
				LIMIT 5
		) AS result
		GROUP BY author
	`
}

func (r *RepoDashboard) GetDashboardVulnByRepository(filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByRepository, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnByRepository(),
		r.queryDefaultFields(), dashboardEnums.TableVulnerabilitiesByRepository, condition)

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnByRepository() string {
	return `
		SELECT repository_name, %[1]s
		FROM (
				SELECT DISTINCT ON (repository_id) *
				FROM %[2]s
				WHERE %[3]s AND created_at 
				IN 
				(
					SELECT MAX(created_at) FROM vulnerabilities_by_language GROUP BY repository_id
				)
				ORDER BY repository_id, created_at DESC
		) AS result
		GROUP BY repository_name
		LIMIT 5
	`
}

func (r *RepoDashboard) GetDashboardVulnByLanguage(filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByLanguage, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnByLanguage(), r.queryDefaultFields(),
		dashboardEnums.TableVulnerabilitiesByLanguage, condition)

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnByLanguage() string {
	return `
		SELECT language, %[1]s
		FROM (
				SELECT DISTINCT ON (repository_id) *
				FROM %[2]s
				WHERE %[3]s AND created_at 
				IN 
				(
					SELECT MAX(created_at) FROM vulnerabilities_by_language GROUP BY repository_id
				)
				ORDER BY repository_id, created_at DESC
				LIMIT 5
		) AS result
		GROUP BY language
	`
}

func (r *RepoDashboard) GetDashboardVulnByTime(filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByTime, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnByTime(), r.queryDefaultFields(), dashboardEnums.TableVulnerabilitiesByTime, condition)

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnByTime() string {
	return `
		SELECT DATE(created_at) AS created_at, %[1]s
		FROM %[2]s AS vuln_by_time
		INNER JOIN
		(
			SELECT MAX(created_at) AS max_time 
			FROM vulnerabilities_by_time 
			GROUP BY DATE(created_at), repository_id
		) AS vuln_by_time_sub_query
		ON vuln_by_time.created_at  = vuln_by_time_sub_query.max_time
		WHERE %[3]s  
		GROUP BY DATE(created_at)
	`
}

func (r *RepoDashboard) queryDefaultFields() string {
	return `
		SUM(critical_vulnerability) as critical_vulnerability, SUM(critical_false_positive) as critical_false_positive, 
	    SUM(critical_risk_accepted) as critical_risk_accepted, SUM(critical_corrected) as critical_corrected,
		SUM(high_vulnerability) as high_vulnerability, SUM(high_false_positive) as high_false_positive, 
	    SUM(high_risk_accepted) as high_risk_accepted, SUM(high_corrected) as high_corrected,
		SUM(medium_vulnerability) as medium_vulnerability, SUM(medium_false_positive) as medium_false_positive, 
	    SUM(medium_risk_accepted) as medium_risk_accepted, SUM(medium_corrected) as medium_corrected,
		SUM(low_vulnerability) as low_vulnerability, SUM(low_false_positive) as low_false_positive, 
	    SUM(low_risk_accepted) as low_risk_accepted, SUM(low_corrected) as low_corrected,
		SUM(info_vulnerability) as info_vulnerability, SUM(info_false_positive) as info_false_positive, 
	    SUM(info_risk_accepted) as info_risk_accepted, SUM(info_corrected) as info_corrected,
		SUM(unknown_vulnerability) as unknown_vulnerability, SUM(unknown_false_positive) as unknown_false_positive, 
		SUM(unknown_risk_accepted) as unknown_risk_accepted, SUM(unknown_corrected) as unknown_corrected
	`
}

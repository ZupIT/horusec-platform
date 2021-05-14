package dashboard

import (
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	repositoriesEntities "github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/repository"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"
	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
)

type IRepoDashboard interface {
	GetDashboardTotalDevelopers(filter *repositoriesEntities.Filter) (int, error)
	GetDashboardTotalRepositories(filter *repositoriesEntities.Filter) (int, error)
	GetDashboardVulnBySeverity(filter *repositoriesEntities.Filter) (*response.Vulnerability, error)
	GetDashboardVulnByAuthor(filter *repositoriesEntities.Filter) ([]*repositoriesEntities.VulnerabilitiesByAuthor, error)
	GetDashboardVulnByRepository(filter *repositoriesEntities.Filter) ([]*repositoriesEntities.VulnerabilitiesByRepository, error)
	GetDashboardVulnByLanguage(filter *repositoriesEntities.Filter) ([]*repositoriesEntities.VulnerabilitiesByLanguage, error)
	GetDashboardVulnByTime(filter *repositoriesEntities.Filter) ([]*repositoriesEntities.VulnerabilitiesByTime, error)
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

func (r *RepoDashboard) GetDashboardTotalDevelopers(filter *repositoriesEntities.Filter) (count int, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardTotalDevelopers(), dashboardEnums.TableVulnerabilitiesByAuthor, condition)

	return count, r.databaseRead.Raw(query, &count, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardTotalDevelopers() string {
	return `
		SELECT COUNT( DISTINCT ( author ) ) 
		FROM %[1]s 
		WHERE %[2]s AND created_at = (SELECT MAX(created_at) FROM %[1]s)
	`
}

func (r *RepoDashboard) GetDashboardTotalRepositories(filter *repositoriesEntities.Filter) (count int, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardTotalRepositories(),
		dashboardEnums.TableVulnerabilitiesByRepository, condition)

	return count, r.databaseRead.Raw(query, &count, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardTotalRepositories() string {
	return `
		SELECT COUNT( DISTINCT ( repository_id ) ) 
		FROM %[1]s 
		WHERE %[2]s AND created_at = (SELECT MAX(created_at) FROM %[1]s)
	`
}

func (r *RepoDashboard) GetDashboardVulnBySeverity(filter *repositoriesEntities.Filter) (*response.Vulnerability, error) {
	vulns := &response.Vulnerability{}
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnBySeverity(),
		r.queryDefaultFields(), dashboardEnums.TableVulnerabilitiesByTime, condition, r.orderBySeverity())

	return vulns, r.databaseRead.Raw(query, vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnBySeverity() string {
	return `
		SELECT %[1]s 
		FROM %[2]s 
		WHERE %[3]s AND created_at = (SELECT MAX(created_at) FROM %[2]s) 
		GROUP BY workspace_id
		%[4]s
		LIMIT 1
	`
}

func (r *RepoDashboard) GetDashboardVulnByAuthor(filter *repositoriesEntities.Filter) (vulns []*repositoriesEntities.VulnerabilitiesByAuthor, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnByAuthor(),
		r.queryDefaultFields(), dashboardEnums.TableVulnerabilitiesByAuthor, condition, r.orderBySeverity())

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnByAuthor() string {
	return `
		SELECT author, %[1]s 
		FROM %[2]s 
		WHERE %[3]s AND created_at = (SELECT MAX(created_at) FROM %[2]s)  
		GROUP BY author 
		%[4]s 
		LIMIT 5
	`
}

func (r *RepoDashboard) GetDashboardVulnByRepository(filter *repositoriesEntities.Filter) (vulns []*repositoriesEntities.VulnerabilitiesByRepository, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnByRepository(),
		r.queryDefaultFields(), dashboardEnums.TableVulnerabilitiesByRepository, condition, r.orderBySeverity())

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnByRepository() string {
	return `
		SELECT repository_name, %[1]s 
		FROM %[2]s 
		WHERE %[3]s AND created_at = (SELECT MAX(created_at) FROM %[2]s)   
		GROUP BY repository_name
		%[4]s 
		LIMIT 5
	`
}

func (r *RepoDashboard) GetDashboardVulnByLanguage(filter *repositoriesEntities.Filter) (vulns []*repositoriesEntities.VulnerabilitiesByLanguage, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnByLanguage(),
		r.queryDefaultFields(), dashboardEnums.TableVulnerabilitiesByLanguage, condition, r.orderBySeverity())

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnByLanguage() string {
	return `
		SELECT language, %[1]s 
		FROM %[2]s 
		WHERE %[3]s AND created_at = (SELECT MAX(created_at) FROM %[2]s)  
		GROUP BY language
		%[4]s 
		LIMIT 5
	`
}

func (r *RepoDashboard) GetDashboardVulnByTime(filter *repositoriesEntities.Filter) (vulns []*repositoriesEntities.VulnerabilitiesByTime, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnByTime(),
		dashboardEnums.TableVulnerabilitiesByTime, condition, r.orderBySeverity())

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoDashboard) queryGetDashboardVulnByTime() string {
	return `
		SELECT * 
		FROM %[1]s 
		WHERE %[2]s AND created_at = (SELECT MAX(created_at) FROM %[1]s)   
		%[3]s
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

func (r *RepoDashboard) orderBySeverity() interface{} {
	return `ORDER BY critical_vulnerability DESC,
		high_vulnerability DESC,
		medium_vulnerability DESC,
		low_vulnerability DESC,
		info_vulnerability DESC,
		unknown_vulnerability DESC`
}

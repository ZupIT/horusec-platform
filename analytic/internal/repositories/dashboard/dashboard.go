package dashboard

import (
	"fmt"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"

	database2 "github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/database"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
)

type IRepoDashboard interface {
	GetDashboardTotalDevelopers(filter *database2.Filter) (int, error)
	GetDashboardTotalRepositories(filter *database2.Filter) (int, error)
	GetDashboardVulnBySeverity(filter *database2.Filter) (*response.Vulnerability, error)
	GetDashboardVulnByAuthor(filter *database2.Filter) ([]*database2.VulnerabilitiesByAuthor, error)
	GetDashboardVulnByRepository(filter *database2.Filter) ([]*database2.VulnerabilitiesByRepository, error)
	GetDashboardVulnByLanguage(filter *database2.Filter) ([]*database2.VulnerabilitiesByLanguage, error)
	GetDashboardVulnByTime(filter *database2.Filter) ([]*database2.VulnerabilitiesByTime, error)
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

func (r *RepoDashboard) GetDashboardTotalDevelopers(filter *database2.Filter) (count int, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(`SELECT COUNT( DISTINCT ( author ) ) FROM %s WHERE %s AND active = true `,
		dashboardEnums.TableVulnerabilitiesByAuthor, condition)

	return count, r.databaseRead.Raw(query, &count, args).GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardTotalRepositories(filter *database2.Filter) (count int, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(`SELECT COUNT( DISTINCT ( repository_id ) ) FROM %s WHERE %s AND active = true `,
		dashboardEnums.TableVulnerabilitiesByRepository, condition)

	return count, r.databaseRead.Raw(query, &count, args).GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardVulnBySeverity(filter *database2.Filter) (vulns *response.Vulnerability, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(`SELECT %s FROM %s WHERE %s AND active = true GROUP BY "workspace_id", "active" %s LIMIT 1`,
		r.queryDefaultFields(), dashboardEnums.TableVulnerabilitiesByTime, condition, r.orderBySeverity())

	return vulns, r.databaseRead.Raw(query, vulns, args).GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardVulnByAuthor(filter *database2.Filter) (vulns []*database2.VulnerabilitiesByAuthor, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(`SELECT author, %s FROM %s WHERE %s AND active = true GROUP BY "author", "active" %s LIMIT 5`,
		r.queryDefaultFields(), dashboardEnums.TableVulnerabilitiesByAuthor, condition, r.orderBySeverity())

	return vulns, r.databaseRead.Raw(query, &vulns, args).GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardVulnByRepository(filter *database2.Filter) (vulns []*database2.VulnerabilitiesByRepository, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(`
		SELECT repository_name, %s FROM %s WHERE %s AND active = true GROUP BY "repository_name", "active" %s LIMIT 5`,
		r.queryDefaultFields(), dashboardEnums.TableVulnerabilitiesByRepository, condition, r.orderBySeverity())

	return vulns, r.databaseRead.Raw(query, &vulns, args).GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardVulnByLanguage(filter *database2.Filter) (vulns []*database2.VulnerabilitiesByLanguage, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(`
		SELECT language, %s FROM %s WHERE %s AND active = true GROUP BY "language", "active" %s LIMIT 5`,
		r.queryDefaultFields(), dashboardEnums.TableVulnerabilitiesByLanguage, condition, r.orderBySeverity())

	return vulns, r.databaseRead.Raw(query, &vulns, args).GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardVulnByTime(filter *database2.Filter) (vulns []*database2.VulnerabilitiesByTime, err error) {
	condition, args := filter.GetConditionFilter()

	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s AND active = true %s`,
		dashboardEnums.TableVulnerabilitiesByTime, condition, r.orderBySeverity())

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
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

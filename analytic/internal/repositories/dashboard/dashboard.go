package dashboard

import (
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type IRepoDashboard interface {
	Save(entity interface{}, table string) error
	Update(entity interface{}, condition map[string]interface{}, table string) error

	GetDashboardTotalDevelopers(filter *dashboard.FilterDashboard) (int, error)
	GetDashboardTotalRepositories(filter *dashboard.FilterDashboard) (int, error)
	GetDashboardVulnBySeverity(filter *dashboard.FilterDashboard) (*dashboard.Vulnerability, error)
	GetDashboardVulnByAuthor(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilitiesByAuthor, error)
	GetDashboardVulnByRepository(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilitiesByRepository, error)
	GetDashboardVulnByLanguage(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilitiesByLanguage, error)
	GetDashboardVulnByTime(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilitiesByTime, error)
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

func (r *RepoDashboard) Save(entity interface{}, table string) error {
	return r.databaseWrite.Create(entity, table).GetErrorExceptNotFound()
}

func (r *RepoDashboard) Update(entity interface{}, condition map[string]interface{}, table string) error {
	return r.databaseWrite.Update(entity, condition, table).GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardTotalDevelopers(
	filter *dashboard.FilterDashboard) (count int, err error) {
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT COUNT( DISTINCT ( author ) ) FROM %s WHERE %s AND active = true `,
		(&dashboard.VulnerabilitiesByAuthor{}).GetTable(), condition)
	result := r.databaseRead.Raw(query, &count, args...)
	return count, result.GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardTotalRepositories(
	filter *dashboard.FilterDashboard) (count int, err error) {
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT COUNT( DISTINCT ( repository_id ) ) FROM %s WHERE %s AND active = true `,
		(&dashboard.VulnerabilitiesByRepository{}).GetTable(), condition)
	result := r.databaseRead.Raw(query, &count, args...)
	return count, result.GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardVulnBySeverity(
	filter *dashboard.FilterDashboard) (vuln *dashboard.Vulnerability, err error) {
	vuln = &dashboard.Vulnerability{}
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE %s AND active = true GROUP BY "workspace_id", "active" %s LIMIT 1`,
		r.querySelectFieldsDefault(), (&dashboard.VulnerabilitiesByTime{}).GetTable(), condition, r.orderByDefault())
	result := r.databaseRead.Raw(query, &vuln, args...)
	return vuln, result.GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardVulnByAuthor(
	filter *dashboard.FilterDashboard) (vulns []*dashboard.VulnerabilitiesByAuthor, err error) {
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`
		SELECT author, %s FROM %s WHERE %s AND active = true GROUP BY "author", "active" %s LIMIT 5`,
		r.querySelectFieldsDefault(), (&dashboard.VulnerabilitiesByAuthor{}).GetTable(), condition, r.orderByDefault())
	result := r.databaseRead.Raw(query, &vulns, args...)
	if vulns == nil {
		return []*dashboard.VulnerabilitiesByAuthor{}, result.GetErrorExceptNotFound()
	}
	return vulns, result.GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardVulnByRepository(
	filter *dashboard.FilterDashboard) (vulns []*dashboard.VulnerabilitiesByRepository, err error) {
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`
		SELECT repository_name, %s FROM %s WHERE %s AND active = true GROUP BY "repository_name", "active" %s LIMIT 5`,
		r.querySelectFieldsDefault(), (&dashboard.VulnerabilitiesByRepository{}).GetTable(), condition, r.orderByDefault())
	result := r.databaseRead.Raw(query, &vulns, args...)
	if vulns == nil {
		return []*dashboard.VulnerabilitiesByRepository{}, result.GetErrorExceptNotFound()
	}
	return vulns, result.GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardVulnByLanguage(
	filter *dashboard.FilterDashboard) (vulns []*dashboard.VulnerabilitiesByLanguage, err error) {
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`
		SELECT language, %s FROM %s WHERE %s AND active = true GROUP BY "language", "active" %s LIMIT 5`,
		r.querySelectFieldsDefault(), (&dashboard.VulnerabilitiesByLanguage{}).GetTable(), condition, r.orderByDefault())
	result := r.databaseRead.Raw(query, &vulns, args...)
	if vulns == nil {
		return []*dashboard.VulnerabilitiesByLanguage{}, result.GetErrorExceptNotFound()
	}
	return vulns, result.GetErrorExceptNotFound()
}

func (r *RepoDashboard) GetDashboardVulnByTime(
	filter *dashboard.FilterDashboard) (vulns []*dashboard.VulnerabilitiesByTime, err error) {
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`
		SELECT created_at, %s FROM %s WHERE %s AND active = true group by "created_at", "active" %s LIMIT 5`,
		r.querySelectFieldsDefault(), (&dashboard.VulnerabilitiesByTime{}).GetTable(), condition, r.orderByDefault())
	result := r.databaseRead.Raw(query, &vulns, args...)
	if vulns == nil {
		return []*dashboard.VulnerabilitiesByTime{}, result.GetErrorExceptNotFound()
	}
	return vulns, result.GetErrorExceptNotFound()
}

func (r *RepoDashboard) getConditionFilter(
	filter *dashboard.FilterDashboard) (string, []interface{}) {
	query := "workspace_id = ? "
	args := []interface{}{filter.WorkspaceID}
	query, args = r.addRepositoryIDOnConditionFilter(filter, query, args)
	query, args = r.addInitialDateOnConditionFilter(filter, query, args)
	query, args = r.addFinalDateOnConditionFilter(filter, query, args)
	return query, args
}

func (r *RepoDashboard) addRepositoryIDOnConditionFilter(
	filter *dashboard.FilterDashboard, query string, args []interface{}) (string, []interface{}) {
	if filter.RepositoryID != uuid.Nil {
		query += "AND repository_id = ? "
		args = append(args, filter.RepositoryID)
	}
	return query, args
}

func (r *RepoDashboard) addInitialDateOnConditionFilter(
	filter *dashboard.FilterDashboard, query string, args []interface{}) (string, []interface{}) {
	if !filter.StartTime.IsZero() {
		query += "AND created_at >= ? "
		args = append(args, filter.StartTime)
	}
	return query, args
}

func (r *RepoDashboard) addFinalDateOnConditionFilter(
	filter *dashboard.FilterDashboard, query string, args []interface{}) (string, []interface{}) {
	if !filter.EndTime.IsZero() {
		query += "AND created_at <= ? "
		args = append(args, filter.EndTime)
	}
	return query, args
}

// nolint:lll // query need all severities in one line
func (r *RepoDashboard) querySelectFieldsDefault() string {
	return `active,
		SUM(critical_vulnerability) as critical_vulnerability, SUM(critical_false_positive) as critical_false_positive, SUM(critical_risk_accepted) as critical_risk_accepted, SUM(critical_corrected) as critical_corrected,
		SUM(high_vulnerability) as high_vulnerability, SUM(high_false_positive) as high_false_positive, SUM(high_risk_accepted) as high_risk_accepted, SUM(high_corrected) as high_corrected,
		SUM(medium_vulnerability) as medium_vulnerability, SUM(medium_false_positive) as medium_false_positive, SUM(medium_risk_accepted) as medium_risk_accepted, SUM(medium_corrected) as medium_corrected,
		SUM(low_vulnerability) as low_vulnerability, SUM(low_false_positive) as low_false_positive, SUM(low_risk_accepted) as low_risk_accepted, SUM(low_corrected) as low_corrected,
		SUM(info_vulnerability) as info_vulnerability, SUM(info_false_positive) as info_false_positive, SUM(info_risk_accepted) as info_risk_accepted, SUM(info_corrected) as info_corrected,
		SUM(unknown_vulnerability) as unknown_vulnerability, SUM(unknown_false_positive) as unknown_false_positive, SUM(unknown_risk_accepted) as unknown_risk_accepted, SUM(unknown_corrected) as unknown_corrected
	`
}

func (r *RepoDashboard) orderByDefault() interface{} {
	return `ORDER BY critical_vulnerability DESC,
		high_vulnerability DESC,
		medium_vulnerability DESC,
		low_vulnerability DESC,
		info_vulnerability DESC,
		unknown_vulnerability DESC`
}

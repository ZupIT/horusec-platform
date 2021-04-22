package dashboard

import (
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type IRepoDashboard interface {
	Save(entity interface{}, table string) error
	Update(entity interface{}, condition map[string]interface{}, table string) error

	GetDashboardTotalDevelopers(filter *dashboard.FilterDashboard) response.IResponse
	GetDashboardTotalRepositories(filter *dashboard.FilterDashboard) response.IResponse
	GetDashboardVulnBySeverity(filter *dashboard.FilterDashboard) response.IResponse
	GetDashboardVulnByAuthor(filter *dashboard.FilterDashboard) response.IResponse
	GetDashboardVulnByRepository(filter *dashboard.FilterDashboard) response.IResponse
	GetDashboardVulnByLanguage(filter *dashboard.FilterDashboard) response.IResponse
	GetDashboardVulnByTime(filter *dashboard.FilterDashboard) response.IResponse
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

func (r *RepoDashboard) GetDashboardTotalDevelopers(filter *dashboard.FilterDashboard) response.IResponse {
	count := 0
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT COUNT( DISTINCT ( author ) ) FROM %s WHERE %s AND active = true `,
		(&dashboard.VulnerabilitiesByAuthor{}).GetTable(), condition)
	result := r.databaseRead.Raw(query, &count, args...)
	return response.NewResponse(int64(result.GetRowsAffected()), result.GetErrorExceptNotFound(), count)
}

func (r *RepoDashboard) GetDashboardTotalRepositories(filter *dashboard.FilterDashboard) response.IResponse {
	count := 0
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT COUNT( DISTINCT ( repository_id ) ) FROM %s WHERE %s AND active = true `,
		(&dashboard.VulnerabilitiesByRepository{}).GetTable(), condition)
	result := r.databaseRead.Raw(query, &count, args...)
	return response.NewResponse(int64(result.GetRowsAffected()), result.GetErrorExceptNotFound(), count)
}

func (r *RepoDashboard) GetDashboardVulnBySeverity(filter *dashboard.FilterDashboard) response.IResponse {
	vuln := &dashboard.VulnerabilitiesByTime{}
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s AND active = true ORDER BY "created_at" ASC LIMIT 1`,
		(&dashboard.VulnerabilitiesByTime{}).GetTable(), condition)
	result := r.databaseRead.Raw(query, &vuln, args...)
	return response.NewResponse(int64(result.GetRowsAffected()), result.GetErrorExceptNotFound(), vuln)
}

func (r *RepoDashboard) GetDashboardVulnByAuthor(filter *dashboard.FilterDashboard) response.IResponse {
	vulns := []*dashboard.VulnerabilitiesByAuthor{}
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s AND active = true ORDER BY "critical_vulnerability" DESC LIMIT 5`,
		(&dashboard.VulnerabilitiesByAuthor{}).GetTable(), condition)
	result := r.databaseRead.Raw(query, &vulns, args...)
	return response.NewResponse(int64(result.GetRowsAffected()), result.GetErrorExceptNotFound(), vulns)
}

func (r *RepoDashboard) GetDashboardVulnByRepository(filter *dashboard.FilterDashboard) response.IResponse {
	vulns := []*dashboard.VulnerabilitiesByRepository{}
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s AND active = true ORDER BY "critical_vulnerability" DESC LIMIT 5`,
		(&dashboard.VulnerabilitiesByRepository{}).GetTable(), condition)
	result := r.databaseRead.Raw(query, &vulns, args...)
	return response.NewResponse(int64(result.GetRowsAffected()), result.GetErrorExceptNotFound(), vulns)
}

func (r *RepoDashboard) GetDashboardVulnByLanguage(filter *dashboard.FilterDashboard) response.IResponse {
	vulns := []*dashboard.VulnerabilitiesByLanguage{}
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s AND active = true ORDER BY "critical_vulnerability" DESC`,
		(&dashboard.VulnerabilitiesByLanguage{}).GetTable(), condition)
	result := r.databaseRead.Raw(query, &vulns, args...)
	return response.NewResponse(int64(result.GetRowsAffected()), result.GetErrorExceptNotFound(), vulns)
}

func (r *RepoDashboard) GetDashboardVulnByTime(filter *dashboard.FilterDashboard) response.IResponse {
	vulns := []*dashboard.VulnerabilitiesByTime{}
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s AND active = true ORDER BY "created_at" ASC`,
		(&dashboard.VulnerabilitiesByTime{}).GetTable(), condition)
	result := r.databaseRead.Raw(query, &vulns, args...)
	return response.NewResponse(int64(result.GetRowsAffected()), result.GetErrorExceptNotFound(), vulns)
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

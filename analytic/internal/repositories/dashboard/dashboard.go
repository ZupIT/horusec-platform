package dashboard

import (
	"fmt"
	vulnerability2 "github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/pagination"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type IRepoDashboard interface {
	GetTotalDevelopers(filter *dashboard.FilterDashboard) (int, error)
	GetTotalRepositories(filter *dashboard.FilterDashboard) (int, error)
	GetVulnBySeverity(filter *dashboard.FilterDashboard) (*dashboard.VulnerabilityBySeverity, error)
	GetVulnByDeveloper(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByDeveloper, error)
	GetVulnByRepository(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByRepository, error)
	GetVulnByLanguage(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByLanguage, error)
	GetVulnByTime(filter *dashboard.FilterDashboard) ([]*dashboard.VulnerabilityByTime, error)
	GetVulnDetails(filter *dashboard.FilterDashboard) (*dashboard.VulnerabilityDetails, error)
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
func (r *RepoDashboard) GetTotalDevelopers(filter *dashboard.FilterDashboard) (count int, err error) {
	condition, args := r.getConditionFilter(filter)
	result := r.databaseRead.Raw(`
		SELECT COUNT( DISTINCT ( vulnerabilities.commit_email ) )
		FROM analysis
		JOIN analysis_vulnerabilities ON analysis.analysis_id = analysis_vulnerabilities.analysis_id
		JOIN vulnerabilities ON analysis_vulnerabilities.vulnerability_id = vulnerabilities.vulnerability_id
		WHERE `+condition, &count, args...)
	return count, result.GetError()
}
func (r *RepoDashboard) GetTotalRepositories(filter *dashboard.FilterDashboard) (count int, err error) {
	condition, args := r.getConditionFilter(filter)
	result := r.databaseRead.Raw(`
		SELECT COUNT( DISTINCT ( analysis.repository_id ) )
		FROM analysis
		WHERE `+condition, &count, args...)
	return count, result.GetError()
}

func (r *RepoDashboard) GetVulnBySeverity(filter *dashboard.FilterDashboard) (*dashboard.VulnerabilityBySeverity, error) {
	vulnBySeverity := dashboard.NewVulnerabilityBySeverity()
	for _, severity := range severities.Values() {
		countSeverity, err := r.getCountBySeverity(filter, severity)
		if err != nil {
			return nil, err
		}
		vulnBySeverity, err = r.getCountByVulnType(filter, severity, countSeverity, vulnBySeverity)
		if err != nil {
			return nil, err
		}
	}
	return vulnBySeverity, nil
}
func (r *RepoDashboard) GetVulnByDeveloper(filter *dashboard.FilterDashboard) (vulnByDeveloper []*dashboard.VulnerabilityByDeveloper, err error) {
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT vulnerabilities.commit_email, COUNT( DISTINCT (vulnerabilities.vulnerability_id) )
		FROM analysis
		INNER JOIN analysis_vulnerabilities ON analysis.analysis_id = analysis_vulnerabilities.analysis_id
		INNER JOIN vulnerabilities ON analysis_vulnerabilities.vulnerability_id = vulnerabilities.vulnerability_id
		WHERE %s
		GROUP BY vulnerabilities.commit_email
		order by "count" desc
		limit 5`, condition)
	result := r.databaseRead.Raw(query, &vulnByDeveloper, args...)
	if result.GetError() != nil && result.GetError() != enums.ErrorNotFoundRecords {
		return nil, result.GetError()
	}
	for index := range vulnByDeveloper {
		filter.AdditionalQuery.Query = "AND vulnerabilities.commit_email = ?"
		filter.AdditionalQuery.Args = []interface{}{vulnByDeveloper[index].DeveloperEmail}
		vulnBySeverity, err := r.GetVulnBySeverity(filter)
		if err != nil {
			return nil, err
		}
		vulnByDeveloper[index] = &dashboard.VulnerabilityByDeveloper{
			DeveloperEmail:          vulnByDeveloper[index].DeveloperEmail,
			VulnerabilityBySeverity: *vulnBySeverity,
		}
	}
	return vulnByDeveloper, nil
}
func (r *RepoDashboard) GetVulnByRepository(filter *dashboard.FilterDashboard) (vulnByRepository []*dashboard.VulnerabilityByRepository, err error) {
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT analysis.repository_id, analysis.repository_name,
		COUNT( DISTINCT (vulnerabilities.vulnerability_id) ) FROM analysis
		INNER JOIN analysis_vulnerabilities ON analysis.analysis_id = analysis_vulnerabilities.analysis_id
		INNER JOIN vulnerabilities ON analysis_vulnerabilities.vulnerability_id = vulnerabilities.vulnerability_id
		WHERE %s
		GROUP BY analysis.repository_id, analysis.repository_name
		order by "count" desc
		limit 5`, condition)
	result := r.databaseRead.Raw(query, &vulnByRepository, args...)
	if result.GetError() != nil && result.GetError() != enums.ErrorNotFoundRecords {
		return nil, result.GetError()
	}
	for index := range vulnByRepository {
		filter.AdditionalQuery.Query = "AND analysis.repository_id = ?"
		filter.AdditionalQuery.Args = []interface{}{vulnByRepository[index].RepositoryID}
		vulnBySeverity, err := r.GetVulnBySeverity(filter)
		if err != nil {
			return nil, err
		}
		vulnByRepository[index] = &dashboard.VulnerabilityByRepository{
			RepositoryID:            vulnByRepository[index].RepositoryID,
			RepositoryName:          vulnByRepository[index].RepositoryName,
			VulnerabilityBySeverity: *vulnBySeverity,
		}
	}
	return vulnByRepository, nil
}
func (r *RepoDashboard) GetVulnByLanguage(filter *dashboard.FilterDashboard) (vulnByLanguage []*dashboard.VulnerabilityByLanguage, err error) {
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT vulnerabilities."language", COUNT( DISTINCT (vulnerabilities.vulnerability_id) )
		FROM analysis
		INNER JOIN analysis_vulnerabilities ON analysis.analysis_id = analysis_vulnerabilities.analysis_id
		INNER JOIN vulnerabilities ON analysis_vulnerabilities.vulnerability_id = vulnerabilities.vulnerability_id
		WHERE %s
		GROUP BY vulnerabilities."language"`, condition)
	result := r.databaseRead.Raw(query, &vulnByLanguage, args...)
	if result.GetError() != nil && result.GetError() != enums.ErrorNotFoundRecords {
		return nil, result.GetError()
	}
	for index := range vulnByLanguage {
		filter.AdditionalQuery.Query = "AND vulnerabilities.language = ?"
		filter.AdditionalQuery.Args = []interface{}{vulnByLanguage[index].Language}
		vulnBySeverity, err := r.GetVulnBySeverity(filter)
		if err != nil {
			return nil, err
		}
		vulnByLanguage[index] = &dashboard.VulnerabilityByLanguage{
			Language:                vulnByLanguage[index].Language,
			VulnerabilityBySeverity: *vulnBySeverity,
		}
	}
	return vulnByLanguage, nil
}
func (r *RepoDashboard) GetVulnByTime(filter *dashboard.FilterDashboard) (vulnByTime []*dashboard.VulnerabilityByTime, err error) {
	condition, args := r.getConditionFilter(filter)
	query := fmt.Sprintf(`SELECT analysis_vulnerabilities.created_at, COUNT( DISTINCT (vulnerabilities.vulnerability_id) )
		FROM analysis
		INNER JOIN analysis_vulnerabilities ON analysis.analysis_id = analysis_vulnerabilities.analysis_id
		INNER JOIN vulnerabilities ON analysis_vulnerabilities.vulnerability_id = vulnerabilities.vulnerability_id
		WHERE %s
		GROUP BY analysis_vulnerabilities.created_at`, condition)
	result := r.databaseRead.Raw(query, &vulnByTime, args...)
	if result.GetError() != nil && result.GetError() != enums.ErrorNotFoundRecords {
		return nil, result.GetError()
	}
	for index := range vulnByTime {
		filter.AdditionalQuery.Query = "AND analysis_vulnerabilities.created_at = ?"
		filter.AdditionalQuery.Args = []interface{}{vulnByTime[index].Time}
		vulnBySeverity, err := r.GetVulnBySeverity(filter)
		if err != nil {
			return nil, err
		}
		vulnByTime[index] = &dashboard.VulnerabilityByTime{
			Time:                vulnByTime[index].Time,
			VulnerabilityBySeverity: *vulnBySeverity,
		}
	}
	return vulnByTime, nil
}
func (r *RepoDashboard) GetVulnDetails(filter *dashboard.FilterDashboard) (*dashboard.VulnerabilityDetails, error) {
	totalItems, err := r.getTotalVulnerabilitiesDetails(filter)
	if err != nil {
		return nil, err
	}
	allVulnerabilities, err := r.getAllVulnerabilitiesDetails(filter)
	if err != nil {
		return nil, err
	}
	return &dashboard.VulnerabilityDetails{
		TotalItems:      totalItems,
		Vulnerabilities: allVulnerabilities,
	}, nil
}

func (r *RepoDashboard) getConditionFilter(
	filter *dashboard.FilterDashboard) (string, []interface{}) {
	query := "analysis.workspace_id = ? "
	args := []interface{}{filter.WorkspaceID}
	query, args = r.addRepositoryIDOnConditionFilter(filter, query, args)
	query, args = r.addInitialDateOnConditionFilter(filter, query, args)
	query, args = r.addFinalDateOnConditionFilter(filter, query, args)
	query, args = r.addAdditionalQueryOnConditionFilter(filter, query, args)
	return query, args
}

func (r *RepoDashboard) addRepositoryIDOnConditionFilter(
	filter *dashboard.FilterDashboard, query string, args []interface{}) (string, []interface{}) {
	if filter.RepositoryID != uuid.Nil {
		query += "AND analysis.repository_id = ? "
		args = append(args, filter.RepositoryID)
	}
	return query, args
}

func (r *RepoDashboard) addInitialDateOnConditionFilter(
	filter *dashboard.FilterDashboard, query string, args []interface{}) (string, []interface{}) {
	if !filter.InitialDate.IsZero() {
		query += "AND analysis.created_at >= ? "
		args = append(args, filter.InitialDate)
	}
	return query, args
}

func (r *RepoDashboard) addFinalDateOnConditionFilter(
	filter *dashboard.FilterDashboard, query string, args []interface{}) (string, []interface{}) {
	if !filter.FinalDate.IsZero() {
		query += "AND analysis.created_at <= ? "
		args = append(args, filter.FinalDate)
	}
	return query, args
}

func (r *RepoDashboard) addAdditionalQueryOnConditionFilter(
	filter *dashboard.FilterDashboard, query string, args []interface{}) (string, []interface{}) {
	if filter.AdditionalQuery.Query != "" {
		query += filter.AdditionalQuery.Query
		args = append(args, filter.AdditionalQuery.Args...)
	}
	return query, args
}

func (r *RepoDashboard) getCountBySeverity(filter *dashboard.FilterDashboard, severity severities.Severity) (int, error) {
	condition, argsCondition := r.getConditionFilter(filter)
	countSeverity := 0
	query := fmt.Sprintf(`select COUNT( DISTINCT (vulnerabilities.vulnerability_id) )
			FROM analysis
			INNER JOIN analysis_vulnerabilities ON analysis.analysis_id = analysis_vulnerabilities.analysis_id
			INNER JOIN vulnerabilities ON analysis_vulnerabilities.vulnerability_id = vulnerabilities.vulnerability_id
			WHERE vulnerabilities.severity = ? AND %s
			GROUP BY vulnerabilities.severity`, condition)
	args := []interface{}{severity}
	args = append(args, argsCondition...)
	if err := r.databaseRead.Raw(query, &countSeverity, args...).GetError(); err != nil && err != enums.ErrorNotFoundRecords {
		return 0, err
	}
	return countSeverity, nil
}

func (r *RepoDashboard) getCountByVulnType(filter *dashboard.FilterDashboard, severity severities.Severity,
	countSeverity int, vulnBySeverity *dashboard.VulnerabilityBySeverity) (*dashboard.VulnerabilityBySeverity, error) {
	newVulnBySeverity := vulnBySeverity
	condition, argsCondition := r.getConditionFilter(filter)
	for _, vulnType := range vulnerability.Values() {
		countType := 0
		args := []interface{}{severity, vulnType}
		queryType := fmt.Sprintf(`SELECT COUNT( DISTINCT (vulnerabilities.vulnerability_id) )
				FROM analysis
				INNER JOIN analysis_vulnerabilities ON analysis.analysis_id = analysis_vulnerabilities.analysis_id
				INNER JOIN vulnerabilities ON analysis_vulnerabilities.vulnerability_id = vulnerabilities.vulnerability_id
				WHERE vulnerabilities.severity = ? AND vulnerabilities."type" = ? AND %s 
				GROUP BY vulnerabilities."severity"`, condition)
		args = append(args, argsCondition...)
		if err := r.databaseRead.Raw(queryType, &countType, args...).GetError(); err != nil && err != enums.ErrorNotFoundRecords {
			return nil, err
		}
		newVulnBySeverity = newVulnBySeverity.SetCountBySeverityAndCountType(countSeverity, countType, severity, vulnType)
	}
	return newVulnBySeverity, nil
}

func (r *RepoDashboard) getTotalVulnerabilitiesDetails(filter *dashboard.FilterDashboard) (count int, err error) {
	condition, args := r.getConditionFilter(filter)
	result := r.databaseRead.Raw(`SELECT COUNT( DISTINCT ( vulnerabilities.vulnerability_id ) )
		FROM analysis
		JOIN analysis_vulnerabilities ON analysis.analysis_id = analysis_vulnerabilities.analysis_id
		JOIN vulnerabilities ON analysis_vulnerabilities.vulnerability_id = vulnerabilities.vulnerability_id
		WHERE `+ condition, &count, args...)
	if result.GetError() != nil && result.GetError() != enums.ErrorNotFoundRecords {
		return 0, err
	}
	return count, nil
}

func (r *RepoDashboard) getAllVulnerabilitiesDetails(filter *dashboard.FilterDashboard) (allVuln []vulnerability2.Vulnerability, err error) {
	condition, argsCondition := r.getConditionFilter(filter)
	args := []interface{}{}
	args = append(args, argsCondition...)
	args = append(args, filter.Size)
	args = append(args, pagination.GetSkip(int64(filter.Page), int64(filter.Size)))
	result := r.databaseRead.Raw(`SELECT vulnerabilities.*
		FROM analysis
		JOIN analysis_vulnerabilities ON analysis.analysis_id = analysis_vulnerabilities.analysis_id
		JOIN vulnerabilities ON analysis_vulnerabilities.vulnerability_id = vulnerabilities.vulnerability_id
		WHERE `+ condition + `GROUP BY vulnerabilities.vulnerability_id
		ORDER BY CASE vulnerabilities.severity
		WHEN 'CRITICAL' THEN 1 WHEN 'HIGH' THEN 2 WHEN 'MEDIUM' THEN 3 WHEN 'LOW' THEN 4
		WHEN 'UNKNOWN' THEN 5 WHEN 'INFO' THEN 6 END, vulnerabilities.type DESC
		LIMIT ? OFFSET ?`, &allVuln, args...)
	if result.GetError() != nil && result.GetError() != enums.ErrorNotFoundRecords {
		return nil, err
	}
	return allVuln, nil
}

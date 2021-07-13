package dashboard

import (
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
)

type IRepoRepository interface {
	GetDashboardTotalDevelopers(filter *dashboard.Filter) (int, error)
	GetDashboardVulnBySeverity(filter *dashboard.Filter) (*dashboard.Vulnerability, error)
	GetDashboardVulnByAuthor(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByAuthor, error)
	GetDashboardVulnByLanguage(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByLanguage, error)
	GetDashboardVulnByTime(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByTime, error)
}

type RepoRepository struct {
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
}

func NewRepoDashboard(connection *database.Connection) IRepoRepository {
	return &RepoRepository{
		databaseRead:  connection.Read,
		databaseWrite: connection.Write,
	}
}

func (r *RepoRepository) GetDashboardTotalDevelopers(filter *dashboard.Filter) (count int, err error) {
	query := fmt.Sprintf(r.queryGetDashboardTotalDevelopers(), dashboardEnums.TableVulnerabilitiesByAuthor)

	return count, r.databaseRead.Raw(query, &count, filter.GetRepositoryFilter()).GetErrorExceptNotFound()
}

func (r *RepoRepository) queryGetDashboardTotalDevelopers() string {
	return `
			SELECT COUNT(DISTINCT(author)) 
			FROM %[1]s
			WHERE repository_id = @repositoryID
			AND created_at = (SELECT MAX(created_at) FROM %[1]s WHERE repository_id = @repositoryID)  
	`
}

func (r *RepoRepository) GetDashboardVulnBySeverity(filter *dashboard.Filter) (*dashboard.Vulnerability, error) {
	vulns := &dashboard.Vulnerability{}

	query := fmt.Sprintf(r.queryGetDashboardVulnBySeverity(), dashboardEnums.TableVulnerabilitiesByTime)

	return vulns, r.databaseRead.Raw(query, vulns, filter.GetRepositoryFilter()).GetErrorExceptNotFound()
}

func (r *RepoRepository) queryGetDashboardVulnBySeverity() string {
	return `
			SELECT *
			FROM %[1]s
			WHERE repository_id = @repositoryID
			AND vulnerability_id = (
										SELECT vulnerability_id 
										FROM %[1]s 
										WHERE repository_id = @repositoryID 
										ORDER BY created_at DESC LIMIT 1
									)
	`
}

func (r *RepoRepository) GetDashboardVulnByAuthor(
	filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByAuthor, err error) {
	query := fmt.Sprintf(r.queryGetDashboardVulnByAuthor(), dashboardEnums.TableVulnerabilitiesByAuthor)

	return vulns, r.databaseRead.Raw(query, &vulns, filter.GetRepositoryFilter()).GetErrorExceptNotFound()
}

//nolint:funlen // need to be bigger than 15
func (r *RepoRepository) queryGetDashboardVulnByAuthor() string {
	return `
	SELECT *
	FROM 
	(
		SELECT  DISTINCT ON(vba.author) vulnSum.total, vba.*
		FROM %[1]s as vba
		INNER JOIN 
		(
			SELECT vulnerability_id,
			(
					critical_vulnerability + critical_false_positive + critical_risk_accepted + critical_corrected + 
					high_vulnerability + high_false_positive + high_risk_accepted + high_corrected +
					medium_vulnerability + medium_false_positive + medium_risk_accepted + medium_corrected +
					low_vulnerability + low_false_positive + low_risk_accepted + low_corrected +info_vulnerability + 
					info_false_positive + info_risk_accepted + info_corrected +unknown_vulnerability + 
					unknown_false_positive + unknown_risk_accepted + unknown_corrected
			) AS total
			FROM %[1]s
			WHERE repository_id = @repositoryID
			AND created_at = (SELECT MAX(created_at) FROM %[1]s WHERE repository_id = @repositoryID)
		) AS vulnSum 
		ON vba.vulnerability_id = vulnSum.vulnerability_id
		WHERE repository_id = @repositoryID
		AND created_at = (SELECT MAX(created_at) FROM %[1]s WHERE repository_id = @repositoryID) 
	) AS vulnsResult
	ORDER BY (vulnsResult.total) DESC
	LIMIT 5
	`
}

func (r *RepoRepository) GetDashboardVulnByLanguage(
	filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByLanguage, err error) {
	query := fmt.Sprintf(r.queryGetDashboardVulnByLanguage(), dashboardEnums.TableVulnerabilitiesByLanguage)

	return vulns, r.databaseRead.Raw(query, &vulns, filter.GetRepositoryFilter()).GetErrorExceptNotFound()
}

func (r *RepoRepository) queryGetDashboardVulnByLanguage() string {
	return `
			SELECT DISTINCT ON(language) language, *
			FROM %[1]s
			WHERE repository_id = @repositoryID
			AND created_at = (SELECT MAX(created_at) FROM %[1]s WHERE repository_id = @repositoryID)  
			LIMIT 5
	`
}

func (r *RepoRepository) GetDashboardVulnByTime(
	filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByTime, err error) {
	condition, args := filter.GetDateFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnByTime(), dashboardEnums.TableVulnerabilitiesByTime, condition)

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
}

func (r *RepoRepository) queryGetDashboardVulnByTime() string {
	return `
		SELECT DISTINCT ON(DATE(created_at)) *
		FROM %[1]s
		WHERE repository_id = @repositoryID
		%[2]s
		AND created_at IN (
			SELECT MAX(created_at)
			FROM %[1]s
			WHERE repository_id = @repositoryID
			%[2]s
			GROUP BY DATE(created_at) )		
	`
}

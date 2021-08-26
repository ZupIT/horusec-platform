// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dashboard

import (
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	dashboardEnums "github.com/ZupIT/horusec-platform/analytic/internal/enums/dashboard"
)

type IWorkspaceRepository interface {
	GetDashboardTotalDevelopers(filter *dashboard.Filter) (int, error)
	GetDashboardTotalRepositories(filter *dashboard.Filter) (int, error)
	GetDashboardVulnBySeverity(filter *dashboard.Filter) (*dashboard.Vulnerability, error)
	GetDashboardVulnByAuthor(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByAuthor, error)
	GetDashboardVulnByRepository(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByRepository, error)
	GetDashboardVulnByLanguage(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByLanguage, error)
	GetDashboardVulnByTime(filter *dashboard.Filter) ([]*dashboard.VulnerabilitiesByTime, error)
}

type WorkspaceRepository struct {
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
}

func NewWorkspaceDashboard(connection *database.Connection) IWorkspaceRepository {
	return &WorkspaceRepository{
		databaseRead:  connection.Read,
		databaseWrite: connection.Write,
	}
}

func (r *WorkspaceRepository) GetDashboardTotalDevelopers(filter *dashboard.Filter) (count int, err error) {
	query := fmt.Sprintf(r.queryGetDashboardTotalDevelopers(), dashboardEnums.TableVulnerabilitiesByAuthor)

	return count, r.databaseRead.Raw(query, &count, filter.GetWorkspaceFilter()).GetErrorExceptNotFound()
}

func (r *WorkspaceRepository) queryGetDashboardTotalDevelopers() string {
	return `
		SELECT COUNT(DISTINCT(vulns.author))
		FROM %[1]s AS vulns
		INNER JOIN 
		(
			SELECT MAX(created_at) max_time, repository_id
			FROM %[1]s
			WHERE workspace_id = @workspaceID
			GROUP BY(repository_id)
		) AS last_analysis
		ON vulns.created_at = last_analysis.max_time 
		AND vulns.repository_id = last_analysis.repository_id
		WHERE workspace_id = @workspaceID AND vulns.author != ''
	`
}

func (r *WorkspaceRepository) GetDashboardTotalRepositories(filter *dashboard.Filter) (count int, err error) {
	query := fmt.Sprintf(r.queryGetDashboardTotalRepositories(), dashboardEnums.TableVulnerabilitiesByRepository)

	return count, r.databaseRead.Raw(query, &count, filter.GetWorkspaceFilter()).GetErrorExceptNotFound()
}

func (r *WorkspaceRepository) queryGetDashboardTotalRepositories() string {
	return `
		SELECT COUNT(DISTINCT(repository_id)) 
		FROM %[1]s
		WHERE workspace_id = @workspaceID
	`
}

func (r *WorkspaceRepository) GetDashboardVulnBySeverity(filter *dashboard.Filter) (*dashboard.Vulnerability, error) {
	vulns := &dashboard.Vulnerability{}

	query := fmt.Sprintf(r.queryGetDashboardVulnBySeverity(), r.queryDefaultFields(),
		dashboardEnums.TableVulnerabilitiesByTime)

	return vulns, r.databaseRead.Raw(query, vulns, filter.GetWorkspaceFilter()).GetErrorExceptNotFound()
}

//nolint:funlen // need to be bigger than 15
func (r *WorkspaceRepository) queryGetDashboardVulnBySeverity() string {
	return `
		SELECT %[1]s 
		FROM 
		(
			SELECT vulns.*
			FROM %[2]s AS vulns
			INNER JOIN 
			(
				SELECT MAX(created_at) max_time, repository_id
				FROM %[2]s
				WHERE workspace_id = @workspaceID
				GROUP BY(repository_id)
			) AS last_analysis
			ON vulns.created_at = last_analysis.max_time 
			AND vulns.repository_id = last_analysis.repository_id
			WHERE workspace_id = @workspaceID
		) AS result
	`
}

func (r *WorkspaceRepository) GetDashboardVulnByAuthor(
	filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByAuthor, err error) {
	query := fmt.Sprintf(r.queryGetDashboardVulnByAuthor(), r.queryDefaultFields(),
		dashboardEnums.TableVulnerabilitiesByAuthor)

	return vulns, r.databaseRead.Raw(query, &vulns, filter.GetWorkspaceFilter()).GetErrorExceptNotFound()
}

//nolint:funlen // need to be bigger than 15
func (r *WorkspaceRepository) queryGetDashboardVulnByAuthor() string {
	return `
		SELECT %[1]s, author
		FROM 
		(
			SELECT vulns.*
			FROM %[2]s AS vulns
			INNER JOIN 
			(
				SELECT MAX(created_at) max_time, repository_id
				FROM %[2]s
				WHERE workspace_id = @workspaceID
				GROUP BY(repository_id)
			) AS last_analysis
			ON vulns.created_at = last_analysis.max_time 
			AND vulns.repository_id = last_analysis.repository_id
			WHERE workspace_id = @workspaceID
		) AS result
		GROUP BY(author)
		LIMIT 5
	`
}

func (r *WorkspaceRepository) GetDashboardVulnByRepository(
	filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByRepository, err error) {
	query := fmt.Sprintf(r.queryGetDashboardVulnByRepository(), r.queryDefaultFields(),
		dashboardEnums.TableVulnerabilitiesByRepository)

	return vulns, r.databaseRead.Raw(query, &vulns, filter.GetWorkspaceFilter()).GetErrorExceptNotFound()
}

//nolint:funlen // need to be bigger than 15
func (r *WorkspaceRepository) queryGetDashboardVulnByRepository() string {
	return `
		SELECT %[1]s, repository_name, repository_id
		FROM 
		(
			SELECT vulns.*
			FROM %[2]s AS vulns
			INNER JOIN 
			(
				SELECT MAX(created_at) max_time, repository_id
				FROM %[2]s
				WHERE workspace_id = @workspaceID
				GROUP BY(repository_id)
			) AS last_analysis
			ON vulns.created_at = last_analysis.max_time 
			AND vulns.repository_id = last_analysis.repository_id
			WHERE workspace_id = @workspaceID
		) AS result
		GROUP BY(repository_id, repository_name)
		LIMIT 5
	`
}

func (r *WorkspaceRepository) GetDashboardVulnByLanguage(
	filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByLanguage, err error) {
	query := fmt.Sprintf(r.queryGetDashboardVulnByLanguage(), r.queryDefaultFields(),
		dashboardEnums.TableVulnerabilitiesByLanguage)

	return vulns, r.databaseRead.Raw(query, &vulns, filter.GetWorkspaceFilter()).GetErrorExceptNotFound()
}

//nolint:funlen // need to be bigger than 15
func (r *WorkspaceRepository) queryGetDashboardVulnByLanguage() string {
	return `
		SELECT %[1]s, language
		FROM 
		(
			SELECT vulns.*
			FROM %[2]s AS vulns
			INNER JOIN 
			(
				SELECT MAX(created_at) max_time, repository_id
				FROM %[2]s
				WHERE workspace_id = @workspaceID
				GROUP BY(repository_id)
			) AS last_analysis
			ON vulns.created_at = last_analysis.max_time 
			AND vulns.repository_id = last_analysis.repository_id
			WHERE workspace_id = @workspaceID
		) AS result
		GROUP BY(language)
		LIMIT 5
	`
}

func (r *WorkspaceRepository) GetDashboardVulnByTime(
	filter *dashboard.Filter) (vulns []*dashboard.VulnerabilitiesByTime, err error) {
	condition, args := filter.GetDateFilter()

	query := fmt.Sprintf(r.queryGetDashboardVulnByTime(),
		r.queryDefaultFields(), dashboardEnums.TableVulnerabilitiesByTime, condition)

	return vulns, r.databaseRead.Raw(query, &vulns, args...).GetErrorExceptNotFound()
}

//nolint:funlen // need to be bigger than 15
func (r *WorkspaceRepository) queryGetDashboardVulnByTime() string {
	return `
		SELECT %[1]s, DATE(created_at) AS created_at
		FROM 
		(
			SELECT vulns.*
			FROM %[2]s AS vulns
			INNER JOIN 
			(
				SELECT MAX(created_at) max_time, repository_id
				FROM %[2]s
				WHERE workspace_id = @workspaceID
				GROUP BY(DATE(created_at), repository_id)
			) AS last_analysis
			ON vulns.created_at = last_analysis.max_time 
			AND vulns.repository_id = last_analysis.repository_id
			WHERE workspace_id = @workspaceID
		) AS result
		WHERE workspace_id = @workspaceID
		%[3]s
		GROUP BY(DATE(created_at))
	`
}

func (r *WorkspaceRepository) queryDefaultFields() string {
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

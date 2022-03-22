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
	"net/http"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type IUseCases interface {
	FilterFromRequest(request *http.Request) (*dashboard.Filter, error)
	ParseAnalysisToVulnerabilitiesByAuthor(entity *analysis.Analysis) []*dashboard.VulnerabilitiesByAuthor
	ParseAnalysisToVulnerabilitiesByRepository(entity *analysis.Analysis) []*dashboard.VulnerabilitiesByRepository
	ParseAnalysisToVulnerabilitiesByLanguage(entity *analysis.Analysis) []*dashboard.VulnerabilitiesByLanguage
	ParseAnalysisToVulnerabilitiesByTime(entity *analysis.Analysis) *dashboard.VulnerabilitiesByTime
}
type UseCases struct{}

func NewUseCaseDashboard() IUseCases {
	return &UseCases{}
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByAuthor(
	entity *analysis.Analysis) []*dashboard.VulnerabilitiesByAuthor {
	if len(entity.AnalysisVulnerabilities) == 0 {
		return u.emptyAnalysisResponseByAuthor(entity)
	}

	return u.processVulnerabilitiesByAuthor(entity)
}

// nolint:funlen // TODO: Check if is necessary broken method
func (u *UseCases) processVulnerabilitiesByAuthor(
	entity *analysis.Analysis) []*dashboard.VulnerabilitiesByAuthor {
	mapVulnByAuthor := map[string]*dashboard.VulnerabilitiesByAuthor{}

	for index := range entity.AnalysisVulnerabilities {
		if vulnByAuthor, ok := mapVulnByAuthor[entity.AnalysisVulnerabilities[index].Vulnerability.CommitEmail]; ok {
			vulnByAuthor.AddCountVulnerabilityBySeverity(entity.AnalysisVulnerabilities[index].Vulnerability.Severity,
				entity.AnalysisVulnerabilities[index].Vulnerability.Type)

			continue
		}

		// nolint:lll // attribution is necessary
		mapVulnByAuthor[entity.AnalysisVulnerabilities[index].Vulnerability.CommitEmail] = u.newVulnerabilitiesByAuthor(entity, index)
	}

	return u.mapVulnByAuthorToSlice(mapVulnByAuthor)
}

func (u *UseCases) emptyAnalysisResponseByAuthor(
	entity *analysis.Analysis) []*dashboard.VulnerabilitiesByAuthor {
	return []*dashboard.VulnerabilitiesByAuthor{
		{
			Author:        "",
			Vulnerability: u.newVulnerabilityFromAnalysis(entity),
		},
	}
}

func (u *UseCases) newVulnerabilitiesByAuthor(entity *analysis.Analysis,
	index int) *dashboard.VulnerabilitiesByAuthor {
	vulnsByAuthor := &dashboard.VulnerabilitiesByAuthor{
		Author:        entity.AnalysisVulnerabilities[index].Vulnerability.CommitEmail,
		Vulnerability: u.newVulnerabilityFromAnalysis(entity),
	}

	vulnsByAuthor.AddCountVulnerabilityBySeverity(entity.AnalysisVulnerabilities[index].Vulnerability.Severity,
		entity.AnalysisVulnerabilities[index].Vulnerability.Type)

	return vulnsByAuthor
}

func (u *UseCases) mapVulnByAuthorToSlice(mapVulnByAuthor map[string]*dashboard.VulnerabilitiesByAuthor) (
	sliceVulnsByAuthor []*dashboard.VulnerabilitiesByAuthor) {
	for _, vulnsByAuthor := range mapVulnByAuthor {
		sliceVulnsByAuthor = append(sliceVulnsByAuthor, vulnsByAuthor)
	}

	return sliceVulnsByAuthor
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByRepository(
	entity *analysis.Analysis) []*dashboard.VulnerabilitiesByRepository {
	if len(entity.AnalysisVulnerabilities) == 0 {
		return u.emptyAnalysisResponseByRepository(entity)
	}

	return u.processVulnerabilitiesByRepository(entity)
}

func (u *UseCases) processVulnerabilitiesByRepository(
	entity *analysis.Analysis) []*dashboard.VulnerabilitiesByRepository {
	mapVulnByRepository := map[string]*dashboard.VulnerabilitiesByRepository{}

	for index := range entity.AnalysisVulnerabilities {
		if vulnByRepository, ok := mapVulnByRepository[entity.RepositoryName]; ok {
			vulnByRepository.AddCountVulnerabilityBySeverity(entity.AnalysisVulnerabilities[index].Vulnerability.Severity,
				entity.AnalysisVulnerabilities[index].Vulnerability.Type)

			continue
		}

		mapVulnByRepository[entity.RepositoryName] = u.newVulnerabilitiesByRepository(entity, index)
	}

	return u.mapVulnByRepositoryToSlice(mapVulnByRepository)
}

func (u *UseCases) emptyAnalysisResponseByRepository(
	entity *analysis.Analysis) []*dashboard.VulnerabilitiesByRepository {
	return []*dashboard.VulnerabilitiesByRepository{
		{
			RepositoryName: entity.RepositoryName,
			Vulnerability:  u.newVulnerabilityFromAnalysis(entity),
		},
	}
}

func (u *UseCases) newVulnerabilitiesByRepository(entity *analysis.Analysis,
	index int) *dashboard.VulnerabilitiesByRepository {
	vulnByRepository := &dashboard.VulnerabilitiesByRepository{
		RepositoryName: entity.RepositoryName,
		Vulnerability:  u.newVulnerabilityFromAnalysis(entity),
	}

	vulnByRepository.AddCountVulnerabilityBySeverity(entity.AnalysisVulnerabilities[index].Vulnerability.Severity,
		entity.AnalysisVulnerabilities[index].Vulnerability.Type)

	return vulnByRepository
}

func (u *UseCases) mapVulnByRepositoryToSlice(mapVulnByRepository map[string]*dashboard.VulnerabilitiesByRepository) (
	sliceVulnsByRepository []*dashboard.VulnerabilitiesByRepository) {
	for _, vulnsByRepository := range mapVulnByRepository {
		sliceVulnsByRepository = append(sliceVulnsByRepository, vulnsByRepository)
	}

	return sliceVulnsByRepository
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByLanguage(
	entity *analysis.Analysis) []*dashboard.VulnerabilitiesByLanguage {
	if len(entity.AnalysisVulnerabilities) == 0 {
		return u.emptyAnalysisResponseByLanguage(entity)
	}

	return u.processVulnerabilitiesByLanguage(entity)
}

// nolint:funlen // TODO: Check if is necessary broken method
func (u *UseCases) processVulnerabilitiesByLanguage(
	entity *analysis.Analysis) []*dashboard.VulnerabilitiesByLanguage {
	mapVulnByLanguage := map[languages.Language]*dashboard.VulnerabilitiesByLanguage{}

	for index := range entity.AnalysisVulnerabilities {
		if vulnByRepository, ok := mapVulnByLanguage[entity.AnalysisVulnerabilities[index].Vulnerability.Language]; ok {
			vulnByRepository.AddCountVulnerabilityBySeverity(
				entity.AnalysisVulnerabilities[index].Vulnerability.Severity,
				entity.AnalysisVulnerabilities[index].Vulnerability.Type)

			continue
		}

		// nolint:lll // attribution is necessary
		mapVulnByLanguage[entity.AnalysisVulnerabilities[index].Vulnerability.Language] = u.newVulnerabilitiesByLanguage(entity, index)
	}

	return u.mapVulnByLanguageToSlice(mapVulnByLanguage)
}

func (u *UseCases) newVulnerabilitiesByLanguage(entity *analysis.Analysis,
	index int) *dashboard.VulnerabilitiesByLanguage {
	vulnByLanguage := &dashboard.VulnerabilitiesByLanguage{
		Language:      entity.AnalysisVulnerabilities[index].Vulnerability.Language,
		Vulnerability: u.newVulnerabilityFromAnalysis(entity),
	}

	vulnByLanguage.AddCountVulnerabilityBySeverity(entity.AnalysisVulnerabilities[index].Vulnerability.Severity,
		entity.AnalysisVulnerabilities[index].Vulnerability.Type)

	return vulnByLanguage
}

func (u *UseCases) mapVulnByLanguageToSlice(
	mapVulnByLanguage map[languages.Language]*dashboard.VulnerabilitiesByLanguage) (
	sliceVulnsByLanguage []*dashboard.VulnerabilitiesByLanguage) {
	for _, vulnsByLanguage := range mapVulnByLanguage {
		sliceVulnsByLanguage = append(sliceVulnsByLanguage, vulnsByLanguage)
	}

	return sliceVulnsByLanguage
}

func (u *UseCases) emptyAnalysisResponseByLanguage(
	entity *analysis.Analysis) []*dashboard.VulnerabilitiesByLanguage {
	return []*dashboard.VulnerabilitiesByLanguage{
		{
			Language:      "",
			Vulnerability: u.newVulnerabilityFromAnalysis(entity),
		},
	}
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByTime(
	entity *analysis.Analysis) *dashboard.VulnerabilitiesByTime {
	vulnsByTime := &dashboard.VulnerabilitiesByTime{
		Vulnerability: u.newVulnerabilityFromAnalysis(entity),
	}

	for index := range entity.AnalysisVulnerabilities {
		vulnsByTime.AddCountVulnerabilityBySeverity(entity.AnalysisVulnerabilities[index].Vulnerability.Severity,
			entity.AnalysisVulnerabilities[index].Vulnerability.Type)
	}

	return vulnsByTime
}

func (u *UseCases) newVulnerabilityFromAnalysis(entity *analysis.Analysis) dashboard.Vulnerability {
	return dashboard.Vulnerability{
		VulnerabilityID: uuid.New(),
		CreatedAt:       entity.CreatedAt,
		WorkspaceID:     entity.WorkspaceID,
		RepositoryID:    entity.RepositoryID,
	}
}

func (u *UseCases) FilterFromRequest(request *http.Request) (*dashboard.Filter, error) {
	filter := &dashboard.Filter{}

	if err := filter.SetWorkspaceAndRepositoryID(request); err != nil {
		return nil, err
	}

	if err := filter.SetDateRangeAndPagination(request); err != nil {
		return nil, err
	}

	return filter, filter.Validate()
}

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

	"github.com/google/uuid"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
)

type IUseCases interface {
	FilterFromRequest(request *http.Request) (*dashboard.Filter, error)
	ParseAnalysisToVulnerabilitiesByAuthor(analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByAuthor
	ParseAnalysisToVulnerabilitiesByRepository(
		analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByRepository
	ParseAnalysisToVulnerabilitiesByLanguage(analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByLanguage
	ParseAnalysisToVulnerabilitiesByTime(analysis *analysisEntities.Analysis) *dashboard.VulnerabilitiesByTime
}
type UseCases struct{}

func NewUseCaseDashboard() IUseCases {
	return &UseCases{}
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByAuthor(
	analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByAuthor {
	if len(analysis.AnalysisVulnerabilities) == 0 {
		return u.emptyAnalysisResponseByAuthor(analysis)
	}

	return u.processVulnerabilitiesByAuthor(analysis)
}

func (u *UseCases) processVulnerabilitiesByAuthor(
	analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByAuthor {
	mapVulnByAuthor := map[string]*dashboard.VulnerabilitiesByAuthor{}

	for index := range analysis.AnalysisVulnerabilities {
		if vulnByAuthor, ok := mapVulnByAuthor[analysis.AnalysisVulnerabilities[index].Vulnerability.CommitEmail]; ok {
			vulnByAuthor.AddCountVulnerabilityBySeverity(analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
				analysis.AnalysisVulnerabilities[index].Vulnerability.Type)
			continue
		}

		mapVulnByAuthor[analysis.AnalysisVulnerabilities[index].Vulnerability.CommitEmail] =
			u.newVulnerabilitiesByAuthor(analysis, index)
	}

	return u.mapVulnByAuthorToSlice(mapVulnByAuthor)
}

func (u *UseCases) emptyAnalysisResponseByAuthor(
	analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByAuthor {
	return []*dashboard.VulnerabilitiesByAuthor{
		{
			Author:        "-",
			Vulnerability: u.newVulnerabilityFromAnalysis(analysis),
		},
	}
}

func (u *UseCases) newVulnerabilitiesByAuthor(analysis *analysisEntities.Analysis,
	index int) *dashboard.VulnerabilitiesByAuthor {
	vulnsByAuthor := &dashboard.VulnerabilitiesByAuthor{
		Author:        analysis.AnalysisVulnerabilities[index].Vulnerability.CommitEmail,
		Vulnerability: u.newVulnerabilityFromAnalysis(analysis),
	}

	vulnsByAuthor.AddCountVulnerabilityBySeverity(analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
		analysis.AnalysisVulnerabilities[index].Vulnerability.Type)

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
	analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByRepository {
	if len(analysis.AnalysisVulnerabilities) == 0 {
		return u.emptyAnalysisResponseByRepository(analysis)
	}

	return u.processVulnerabilitiesByRepository(analysis)
}

func (u *UseCases) processVulnerabilitiesByRepository(
	analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByRepository {
	mapVulnByRepository := map[string]*dashboard.VulnerabilitiesByRepository{}

	for index := range analysis.AnalysisVulnerabilities {
		if vulnByRepository, ok := mapVulnByRepository[analysis.RepositoryName]; ok {
			vulnByRepository.AddCountVulnerabilityBySeverity(analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
				analysis.AnalysisVulnerabilities[index].Vulnerability.Type)
			continue
		}

		mapVulnByRepository[analysis.RepositoryName] =
			u.newVulnerabilitiesByRepository(analysis, index)
	}

	return u.mapVulnByRepositoryToSlice(mapVulnByRepository)
}

func (u *UseCases) emptyAnalysisResponseByRepository(
	analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByRepository {
	return []*dashboard.VulnerabilitiesByRepository{
		{
			RepositoryName: analysis.RepositoryName,
			Vulnerability:  u.newVulnerabilityFromAnalysis(analysis),
		},
	}
}

func (u *UseCases) newVulnerabilitiesByRepository(analysis *analysisEntities.Analysis,
	index int) *dashboard.VulnerabilitiesByRepository {
	vulnByRepository := &dashboard.VulnerabilitiesByRepository{
		RepositoryName: analysis.RepositoryName,
		Vulnerability:  u.newVulnerabilityFromAnalysis(analysis),
	}

	vulnByRepository.AddCountVulnerabilityBySeverity(analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
		analysis.AnalysisVulnerabilities[index].Vulnerability.Type)

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
	analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByLanguage {
	if len(analysis.AnalysisVulnerabilities) == 0 {
		return u.emptyAnalysisResponseByLanguage(analysis)
	}

	return u.processVulnerabilitiesByLanguage(analysis)
}

func (u *UseCases) processVulnerabilitiesByLanguage(
	analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByLanguage {
	mapVulnByLanguage := map[languages.Language]*dashboard.VulnerabilitiesByLanguage{}

	for index := range analysis.AnalysisVulnerabilities {
		if vulnByRepository, ok := mapVulnByLanguage[analysis.AnalysisVulnerabilities[index].Vulnerability.Language]; ok {
			vulnByRepository.AddCountVulnerabilityBySeverity(analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
				analysis.AnalysisVulnerabilities[index].Vulnerability.Type)
			continue
		}

		mapVulnByLanguage[analysis.AnalysisVulnerabilities[index].Vulnerability.Language] =
			u.newVulnerabilitiesByLanguage(analysis, index)
	}

	return u.mapVulnByLanguageToSlice(mapVulnByLanguage)
}

func (u *UseCases) newVulnerabilitiesByLanguage(analysis *analysisEntities.Analysis,
	index int) *dashboard.VulnerabilitiesByLanguage {
	vulnByLanguage := &dashboard.VulnerabilitiesByLanguage{
		Language:      analysis.AnalysisVulnerabilities[index].Vulnerability.Language,
		Vulnerability: u.newVulnerabilityFromAnalysis(analysis),
	}

	vulnByLanguage.AddCountVulnerabilityBySeverity(analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
		analysis.AnalysisVulnerabilities[index].Vulnerability.Type)

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
	analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByLanguage {
	return []*dashboard.VulnerabilitiesByLanguage{
		{
			Language:      "-",
			Vulnerability: u.newVulnerabilityFromAnalysis(analysis),
		},
	}
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByTime(
	analysis *analysisEntities.Analysis) *dashboard.VulnerabilitiesByTime {
	vulnsByTime := &dashboard.VulnerabilitiesByTime{
		Vulnerability: u.newVulnerabilityFromAnalysis(analysis),
	}

	for index := range analysis.AnalysisVulnerabilities {
		vulnsByTime.AddCountVulnerabilityBySeverity(analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
			analysis.AnalysisVulnerabilities[index].Vulnerability.Type)
	}

	return vulnsByTime
}

func (u *UseCases) newVulnerabilityFromAnalysis(analysis *analysisEntities.Analysis) dashboard.Vulnerability {
	return dashboard.Vulnerability{
		VulnerabilityID: uuid.New(),
		CreatedAt:       analysis.CreatedAt,
		WorkspaceID:     analysis.WorkspaceID,
		RepositoryID:    analysis.RepositoryID,
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

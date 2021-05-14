package dashboardfilter

import (
	"net/http"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/repository"

	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"

	"github.com/google/uuid"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
)

type IUseCases interface {
	FilterFromRequest(request *http.Request) (*repository.Filter, error)
	ParseAnalysisToVulnerabilitiesByAuthor(analysis *analysisEntities.Analysis) []*repository.VulnerabilitiesByAuthor
	ParseAnalysisToVulnerabilitiesByRepository(analysis *analysisEntities.Analysis) []*repository.VulnerabilitiesByRepository
	ParseAnalysisToVulnerabilitiesByLanguage(analysis *analysisEntities.Analysis) []*repository.VulnerabilitiesByLanguage
	ParseAnalysisToVulnerabilitiesByTime(analysis *analysisEntities.Analysis) *repository.VulnerabilitiesByTime
}
type UseCases struct{}

func NewUseCaseDashboard() IUseCases {
	return &UseCases{}
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByAuthor(
	analysis *analysisEntities.Analysis) []*repository.VulnerabilitiesByAuthor {
	mapVulnByAuthor := map[string]*repository.VulnerabilitiesByAuthor{}

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

func (u *UseCases) newVulnerabilitiesByAuthor(analysis *analysisEntities.Analysis,
	index int) *repository.VulnerabilitiesByAuthor {
	vulnsByAuthor := &repository.VulnerabilitiesByAuthor{
		Author: analysis.AnalysisVulnerabilities[index].Vulnerability.CommitEmail,
		Vulnerability: response.Vulnerability{
			VulnerabilityID: uuid.New(),
			CreatedAt:       analysis.CreatedAt,
			WorkspaceID:     analysis.WorkspaceID,
			RepositoryID:    analysis.RepositoryID,
		},
	}

	vulnsByAuthor.AddCountVulnerabilityBySeverity(analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
		analysis.AnalysisVulnerabilities[index].Vulnerability.Type)

	return vulnsByAuthor
}

func (u *UseCases) mapVulnByAuthorToSlice(mapVulnByAuthor map[string]*repository.VulnerabilitiesByAuthor) (
	sliceVulnsByAuthor []*repository.VulnerabilitiesByAuthor) {
	for _, vulnsByAuthor := range mapVulnByAuthor {
		sliceVulnsByAuthor = append(sliceVulnsByAuthor, vulnsByAuthor)
	}

	return sliceVulnsByAuthor
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByRepository(
	analysis *analysisEntities.Analysis) []*repository.VulnerabilitiesByRepository {
	mapVulnByRepository := map[string]*repository.VulnerabilitiesByRepository{}

	for index := range analysis.AnalysisVulnerabilities {
		if vulnByRepository, ok := mapVulnByRepository[analysis.RepositoryName]; ok {
			vulnByRepository.AddCountVulnerabilityBySeverity(
				analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
				analysis.AnalysisVulnerabilities[index].Vulnerability.Type)
			continue
		}

		mapVulnByRepository[analysis.RepositoryName] =
			u.newVulnerabilitiesByRepository(analysis, index)
	}

	return u.mapVulnByRepositoryToSlice(mapVulnByRepository)
}

func (u *UseCases) newVulnerabilitiesByRepository(analysis *analysisEntities.Analysis,
	index int) *repository.VulnerabilitiesByRepository {
	vulnByRepository := &repository.VulnerabilitiesByRepository{
		RepositoryName: analysis.RepositoryName,
		Vulnerability: response.Vulnerability{
			VulnerabilityID: uuid.New(),
			CreatedAt:       analysis.CreatedAt,
			WorkspaceID:     analysis.WorkspaceID,
			RepositoryID:    analysis.RepositoryID,
		},
	}

	vulnByRepository.AddCountVulnerabilityBySeverity(analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
		analysis.AnalysisVulnerabilities[index].Vulnerability.Type)

	return vulnByRepository
}

func (u *UseCases) mapVulnByRepositoryToSlice(mapVulnByRepository map[string]*repository.VulnerabilitiesByRepository) (
	sliceVulnsByRepository []*repository.VulnerabilitiesByRepository) {

	for _, vulnsByRepository := range mapVulnByRepository {
		sliceVulnsByRepository = append(sliceVulnsByRepository, vulnsByRepository)
	}

	return sliceVulnsByRepository
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByLanguage(
	analysis *analysisEntities.Analysis) []*repository.VulnerabilitiesByLanguage {
	mapVulnByLanguage := map[languages.Language]*repository.VulnerabilitiesByLanguage{}

	for index := range analysis.AnalysisVulnerabilities {
		if vulnByRepository, ok := mapVulnByLanguage[analysis.AnalysisVulnerabilities[index].Vulnerability.Language]; ok {
			vulnByRepository.AddCountVulnerabilityBySeverity(
				analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
				analysis.AnalysisVulnerabilities[index].Vulnerability.Type)
			continue
		}

		mapVulnByLanguage[analysis.AnalysisVulnerabilities[index].Vulnerability.Language] =
			u.newVulnerabilitiesByLanguage(analysis, index)
	}

	return u.mapVulnByLanguageToSlice(mapVulnByLanguage)
}

func (u *UseCases) newVulnerabilitiesByLanguage(analysis *analysisEntities.Analysis,
	index int) *repository.VulnerabilitiesByLanguage {
	vulnByLanguage := &repository.VulnerabilitiesByLanguage{
		Language: analysis.AnalysisVulnerabilities[index].Vulnerability.Language,
		Vulnerability: response.Vulnerability{
			VulnerabilityID: uuid.New(),
			CreatedAt:       analysis.CreatedAt,
			WorkspaceID:     analysis.WorkspaceID,
			RepositoryID:    analysis.RepositoryID,
		},
	}

	vulnByLanguage.AddCountVulnerabilityBySeverity(analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
		analysis.AnalysisVulnerabilities[index].Vulnerability.Type)

	return vulnByLanguage
}

func (u *UseCases) mapVulnByLanguageToSlice(mapVulnByLanguage map[languages.Language]*repository.VulnerabilitiesByLanguage) (
	sliceVulnsByLanguage []*repository.VulnerabilitiesByLanguage) {
	for _, vulnsByLanguage := range mapVulnByLanguage {
		sliceVulnsByLanguage = append(sliceVulnsByLanguage, vulnsByLanguage)
	}

	return sliceVulnsByLanguage
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByTime(
	analysis *analysisEntities.Analysis) *repository.VulnerabilitiesByTime {
	vulnsByTime := &repository.VulnerabilitiesByTime{
		Vulnerability: response.Vulnerability{
			VulnerabilityID: uuid.New(),
			CreatedAt:       analysis.CreatedAt,
			WorkspaceID:     analysis.WorkspaceID,
			RepositoryID:    analysis.RepositoryID,
		},
	}

	for index := range analysis.AnalysisVulnerabilities {
		vulnsByTime.AddCountVulnerabilityBySeverity(analysis.AnalysisVulnerabilities[index].Vulnerability.Severity,
			analysis.AnalysisVulnerabilities[index].Vulnerability.Type)
	}

	return vulnsByTime
}

func (u *UseCases) FilterFromRequest(request *http.Request) (*repository.Filter, error) {
	filter := &repository.Filter{}

	if err := filter.SetWorkspaceAndRepositoryID(request); err != nil {
		return nil, err
	}

	if err := filter.SetDateRangeAndPagination(request); err != nil {
		return nil, err
	}

	return filter, filter.Validate()
}

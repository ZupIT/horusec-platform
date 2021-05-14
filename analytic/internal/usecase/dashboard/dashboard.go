package dashboardfilter

import (
	netHTTP "net/http"
	"strconv"
	"time"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/response"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/repositories"

	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"

	"github.com/ZupIT/horusec-platform/analytic/internal/enums"
)

const DefaultSize = 10

type IUseCases interface {
	ExtractFilterDashboard(r *netHTTP.Request) (*repositories.Filter, error)
	ParseAnalysisToVulnerabilitiesByAuthor(analysis *analysisEntities.Analysis) []*repositories.VulnerabilitiesByAuthor
	ParseAnalysisToVulnerabilitiesByRepository(analysis *analysisEntities.Analysis) []*repositories.VulnerabilitiesByRepository
	ParseAnalysisToVulnerabilitiesByLanguage(analysis *analysisEntities.Analysis) []*repositories.VulnerabilitiesByLanguage
	ParseAnalysisToVulnerabilitiesByTime(analysis *analysisEntities.Analysis) *repositories.VulnerabilitiesByTime
}
type UseCases struct{}

func NewUseCaseDashboard() IUseCases {
	return &UseCases{}
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByAuthor(
	analysis *analysisEntities.Analysis) []*repositories.VulnerabilitiesByAuthor {
	mapVulnByAuthor := map[string]*repositories.VulnerabilitiesByAuthor{}

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
	index int) *repositories.VulnerabilitiesByAuthor {
	vulnsByAuthor := &repositories.VulnerabilitiesByAuthor{
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

func (u *UseCases) mapVulnByAuthorToSlice(mapVulnByAuthor map[string]*repositories.VulnerabilitiesByAuthor) (
	sliceVulnsByAuthor []*repositories.VulnerabilitiesByAuthor) {
	for _, vulnsByAuthor := range mapVulnByAuthor {
		sliceVulnsByAuthor = append(sliceVulnsByAuthor, vulnsByAuthor)
	}

	return sliceVulnsByAuthor
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByRepository(
	analysis *analysisEntities.Analysis) []*repositories.VulnerabilitiesByRepository {
	mapVulnByRepository := map[string]*repositories.VulnerabilitiesByRepository{}

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
	index int) *repositories.VulnerabilitiesByRepository {
	vulnByRepository := &repositories.VulnerabilitiesByRepository{
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

func (u *UseCases) mapVulnByRepositoryToSlice(mapVulnByRepository map[string]*repositories.VulnerabilitiesByRepository) (
	sliceVulnsByRepository []*repositories.VulnerabilitiesByRepository) {

	for _, vulnsByRepository := range mapVulnByRepository {
		sliceVulnsByRepository = append(sliceVulnsByRepository, vulnsByRepository)
	}

	return sliceVulnsByRepository
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByLanguage(
	analysis *analysisEntities.Analysis) []*repositories.VulnerabilitiesByLanguage {
	mapVulnByLanguage := map[languages.Language]*repositories.VulnerabilitiesByLanguage{}

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
	index int) *repositories.VulnerabilitiesByLanguage {
	vulnByLanguage := &repositories.VulnerabilitiesByLanguage{
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

func (u *UseCases) mapVulnByLanguageToSlice(mapVulnByLanguage map[languages.Language]*repositories.VulnerabilitiesByLanguage) (
	sliceVulnsByLanguage []*repositories.VulnerabilitiesByLanguage) {
	for _, vulnsByLanguage := range mapVulnByLanguage {
		sliceVulnsByLanguage = append(sliceVulnsByLanguage, vulnsByLanguage)
	}

	return sliceVulnsByLanguage
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByTime(
	analysis *analysisEntities.Analysis) *repositories.VulnerabilitiesByTime {
	vulnsByTime := &repositories.VulnerabilitiesByTime{
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

func (u *UseCases) ExtractFilterDashboard(r *netHTTP.Request) (*repositories.Filter, error) {
	filter, err := u.extractFilterDashboardByWorkspace(r)
	if err != nil {
		return nil, err
	}
	if chi.URLParam(r, "repositoryID") != "" {
		filter.RepositoryID, err = uuid.Parse(chi.URLParam(r, "repositoryID"))
		if err != nil {
			return nil, enums.ErrorWrongRepositoryID
		}
	}
	return filter, u.validateFilterDashboard(filter, true)
}

func (u *UseCases) extractFilterDashboardByWorkspace(r *netHTTP.Request) (*repositories.Filter, error) {
	workspaceID, err := uuid.Parse(chi.URLParam(r, "workspaceID"))
	if err != nil {
		return nil, enums.ErrorWrongWorkspaceID
	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	initialDate, finalDate, err := u.getDateRangeFromRequestQuery(r)
	if err != nil {
		return nil, err
	}
	entity := u.buildFilterDasboard(workspaceID, initialDate, finalDate, page, u.getSizeOrMin(size))
	return entity, u.validateFilterDashboard(entity, false)
}

func (u *UseCases) getSizeOrMin(size int) int {
	if size < DefaultSize {
		return DefaultSize
	}
	return size
}

func (u *UseCases) getDateRangeFromRequestQuery(r *netHTTP.Request) (time.Time, time.Time, error) {
	initial, err := u.getDateFromRequestQuery(r, "initialDate")
	if err != nil {
		return time.Time{}, time.Time{}, enums.ErrorWrongInitialDate
	}

	final, err := u.getDateFromRequestQuery(r, "finalDate")
	if err != nil {
		return time.Time{}, time.Time{}, enums.ErrorWrongFinalDate
	}

	return initial, final, nil
}

func (u *UseCases) getDateFromRequestQuery(r *netHTTP.Request, queryStrKey string) (time.Time, error) {
	date := r.URL.Query().Get(queryStrKey)
	if date != "" {
		return time.Parse("2006-01-02T15:04:05Z", date)
	}

	return time.Time{}, nil
}

func (u *UseCases) validateFilterDashboard(
	entity *repositories.Filter, isFilterFromRepository bool) error {
	return validation.ValidateStruct(entity,
		validation.Field(&entity.WorkspaceID, validation.Required, validation.NotIn(uuid.Nil)),
		validation.Field(&entity.RepositoryID, validation.When(isFilterFromRepository, validation.NotIn(uuid.Nil))),
		validation.Field(&entity.StartTime, validation.Required),
		validation.Field(&entity.EndTime, validation.Required),
		validation.Field(&entity.Page, validation.Min(0)),
		validation.Field(&entity.Size, validation.Min(DefaultSize)),
	)
}

func (u *UseCases) buildFilterDasboard(workspaceID uuid.UUID, initialDate, finalDate time.Time,
	page int, size int) *repositories.Filter {
	return &repositories.Filter{
		WorkspaceID: workspaceID,
		StartTime:   initialDate,
		EndTime:     finalDate,
		Page:        page,
		Size:        size,
	}
}

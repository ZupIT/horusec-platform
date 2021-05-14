package dashboardfilter

import (
	netHTTP "net/http"
	"strconv"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard"
	"github.com/ZupIT/horusec-platform/analytic/internal/enums"
)

const DefaultSize = 10

type IUseCases interface {
	ExtractFilterDashboard(r *netHTTP.Request) (*dashboard.Filters, error)
	ParseAnalysisToVulnerabilitiesByAuthor(analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByAuthor
	ParseAnalysisToVulnerabilitiesByRepository(analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByRepository
	ParseAnalysisToVulnerabilitiesByLanguage(analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByLanguage
	ParseAnalysisToVulnerabilitiesByTime(analysis *analysisEntities.Analysis) *dashboard.VulnerabilitiesByTime
}
type UseCases struct{}

func NewUseCaseDashboard() IUseCases {
	return &UseCases{}
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByAuthor(
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

func (u *UseCases) newVulnerabilitiesByAuthor(analysis *analysisEntities.Analysis,
	index int) *dashboard.VulnerabilitiesByAuthor {
	vulnsByAuthor := &dashboard.VulnerabilitiesByAuthor{
		Author: analysis.AnalysisVulnerabilities[index].Vulnerability.CommitEmail,
		Vulnerability: dashboard.Vulnerability{
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

func (u *UseCases) mapVulnByAuthorToSlice(mapVulnByAuthor map[string]*dashboard.VulnerabilitiesByAuthor) (
	sliceVulnsByAuthor []*dashboard.VulnerabilitiesByAuthor) {
	for _, vulnsByAuthor := range mapVulnByAuthor {
		sliceVulnsByAuthor = append(sliceVulnsByAuthor, vulnsByAuthor)
	}

	return sliceVulnsByAuthor
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByRepository(
	analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByRepository {
	mapVulnByRepository := map[string]*dashboard.VulnerabilitiesByRepository{}

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
	index int) *dashboard.VulnerabilitiesByRepository {
	vulnByRepository := &dashboard.VulnerabilitiesByRepository{
		RepositoryName: analysis.RepositoryName,
		Vulnerability: dashboard.Vulnerability{
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

func (u *UseCases) mapVulnByRepositoryToSlice(mapVulnByRepository map[string]*dashboard.VulnerabilitiesByRepository) (
	sliceVulnsByRepository []*dashboard.VulnerabilitiesByRepository) {

	for _, vulnsByRepository := range mapVulnByRepository {
		sliceVulnsByRepository = append(sliceVulnsByRepository, vulnsByRepository)
	}

	return sliceVulnsByRepository
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByLanguage(
	analysis *analysisEntities.Analysis) []*dashboard.VulnerabilitiesByLanguage {
	mapVulnByLanguage := map[languages.Language]*dashboard.VulnerabilitiesByLanguage{}

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
	index int) *dashboard.VulnerabilitiesByLanguage {
	vulnByLanguage := &dashboard.VulnerabilitiesByLanguage{
		Language: analysis.AnalysisVulnerabilities[index].Vulnerability.Language,
		Vulnerability: dashboard.Vulnerability{
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

func (u *UseCases) mapVulnByLanguageToSlice(mapVulnByLanguage map[languages.Language]*dashboard.VulnerabilitiesByLanguage) (
	sliceVulnsByLanguage []*dashboard.VulnerabilitiesByLanguage) {
	for _, vulnsByLanguage := range mapVulnByLanguage {
		sliceVulnsByLanguage = append(sliceVulnsByLanguage, vulnsByLanguage)
	}

	return sliceVulnsByLanguage
}

func (u *UseCases) ParseAnalysisToVulnerabilitiesByTime(
	analysis *analysisEntities.Analysis) *dashboard.VulnerabilitiesByTime {
	vulnsByTime := &dashboard.VulnerabilitiesByTime{
		Vulnerability: dashboard.Vulnerability{
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

func (u *UseCases) ExtractFilterDashboard(r *netHTTP.Request) (*dashboard.Filters, error) {
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

func (u *UseCases) extractFilterDashboardByWorkspace(r *netHTTP.Request) (*dashboard.Filters, error) {
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
	entity *dashboard.Filters, isFilterFromRepository bool) error {
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
	page int, size int) *dashboard.Filters {
	return &dashboard.Filters{
		WorkspaceID: workspaceID,
		StartTime:   initialDate,
		EndTime:     finalDate,
		Page:        page,
		Size:        size,
	}
}

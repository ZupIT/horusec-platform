package dashboard

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	analysisEnum "github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/confidence"
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"github.com/ZupIT/horusec-devkit/pkg/enums/tools"
	vulnerabilityEnum "github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"
)

func getAnalysisMock() *analysisEntities.Analysis {
	analysisID := uuid.New()
	vulnerabilityID1 := uuid.New()
	vulnerabilityID2 := uuid.New()

	return &analysisEntities.Analysis{
		ID:             analysisID,
		RepositoryID:   uuid.New(),
		RepositoryName: "my-repository",
		WorkspaceID:    uuid.New(),
		WorkspaceName:  "my-workspace",
		Status:         analysisEnum.Success,
		Errors:         "",
		CreatedAt:      time.Now(),
		FinishedAt:     time.Now(),
		AnalysisVulnerabilities: []analysisEntities.AnalysisVulnerabilities{
			{
				VulnerabilityID: vulnerabilityID1,
				AnalysisID:      analysisID,
				CreatedAt:       time.Now(),
				Vulnerability: vulnerability.Vulnerability{
					VulnerabilityID: vulnerabilityID1,
					Line:            "1",
					Column:          "1",
					Confidence:      confidence.High,
					File:            "/deployments/cert.pem",
					Code:            "-----BEGIN CERTIFICATE-----",
					Details:         "Asymmetric Private Key \n Found SSH and/or x.509 ...",
					SecurityTool:    tools.HorusecEngine,
					Language:        languages.Leaks,
					Severity:        severities.Critical,
					VulnHash:        "1234567890",
					Type:            vulnerabilityEnum.Vulnerability,
					CommitAuthor:    "Horusec",
					CommitEmail:     "horusec@zup.com.br",
					CommitHash:      "9876543210",
					CommitMessage:   "Initial Commit",
					CommitDate:      "2021-03-31T10:58:42Z",
				},
			},
			{
				VulnerabilityID: vulnerabilityID2,
				AnalysisID:      analysisID,
				CreatedAt:       time.Now(),
				Vulnerability: vulnerability.Vulnerability{
					VulnerabilityID: vulnerabilityID2,
					Line:            "1",
					Column:          "1",
					Confidence:      confidence.High,
					File:            "/deployments/key.pem",
					Code:            "-----BEGIN OPENSSH PRIVATE KEY-----",
					Details:         "Asymmetric Private Key \n Found SSH and/or x.509 ...",
					SecurityTool:    tools.HorusecEngine,
					Language:        languages.Leaks,
					Severity:        severities.Critical,
					VulnHash:        "0987654321",
					Type:            vulnerabilityEnum.Vulnerability,
					CommitAuthor:    "Horusec",
					CommitEmail:     "horusec@zup.com.br",
					CommitHash:      "1234567890",
					CommitMessage:   "Initial Commit",
					CommitDate:      "2021-03-31T10:58:42Z",
				},
			},
		},
	}
}

func TestParseAnalysisToVulnerabilitiesByAuthor(t *testing.T) {
	t.Run("should success parse without errors", func(t *testing.T) {
		useCases := NewUseCaseDashboard()

		assert.Len(t, useCases.ParseAnalysisToVulnerabilitiesByAuthor(getAnalysisMock()), 1)
	})
}

func TestParseAnalysisToVulnerabilitiesByRepository(t *testing.T) {
	t.Run("should success parse without errors", func(t *testing.T) {
		useCases := NewUseCaseDashboard()

		assert.Len(t, useCases.ParseAnalysisToVulnerabilitiesByRepository(getAnalysisMock()), 1)
	})
}

func TestParseAnalysisToVulnerabilitiesByLanguage(t *testing.T) {
	t.Run("should success parse without errors", func(t *testing.T) {
		useCases := NewUseCaseDashboard()

		assert.Len(t, useCases.ParseAnalysisToVulnerabilitiesByLanguage(getAnalysisMock()), 1)
	})
}

func TestParseAnalysisToVulnerabilitiesByTime(t *testing.T) {
	t.Run("should success parse without errors", func(t *testing.T) {
		useCases := NewUseCaseDashboard()

		assert.NotNil(t, useCases.ParseAnalysisToVulnerabilitiesByTime(getAnalysisMock()))
	})
}

func TestFilterFromRequest(t *testing.T) {
	layoutDateTime := "2006-01-02T15:04:05Z"
	startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
	endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")

	t.Run("should success create a new filter from request data", func(t *testing.T) {
		useCases := NewUseCaseDashboard()

		id := uuid.New()

		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			startTime.Format(layoutDateTime), endTime.Format(layoutDateTime), 18, 18)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", id.String())
		ctx.URLParams.Add("repositoryID", id.String())
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		filter, err := useCases.FilterFromRequest(r)
		assert.NoError(t, err)
		assert.Equal(t, startTime, filter.StartTime)
		assert.Equal(t, endTime, filter.EndTime)
		assert.Equal(t, 18, filter.Page)
		assert.Equal(t, 18, filter.Size)
		assert.Equal(t, id, filter.WorkspaceID)
		assert.Equal(t, id, filter.RepositoryID)
	})

	t.Run("should return error when failed to set pagination", func(t *testing.T) {
		useCases := NewUseCaseDashboard()

		id := uuid.New()

		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			startTime.Format(layoutDateTime), "test", 18, 18)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", id.String())
		ctx.URLParams.Add("repositoryID", id.String())
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		filter, err := useCases.FilterFromRequest(r)
		assert.Error(t, err)
		assert.Nil(t, filter)
	})

	t.Run("should return error when failed to workspace or repository id", func(t *testing.T) {
		useCases := NewUseCaseDashboard()

		ctx := chi.NewRouteContext()
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		filter, err := useCases.FilterFromRequest(r)
		assert.Error(t, err)
		assert.Nil(t, filter)
	})
}

package dashboardfilter

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/ZupIT/horusec-platform/analytic/internal/entities/dashboard/repository"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	analysisEnum "github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/confidence"
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"github.com/ZupIT/horusec-devkit/pkg/enums/tools"
	vulnerabilityEnum "github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"

	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/analytic/internal/enums"
)

func TestUseCaseDashboard_ExtractFilterDashboard(t *testing.T) {
	layoutDateTime := "2006-01-02T15:04:05Z"
	t.Run("Should extract filter with repository", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &repository.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			expected.StartTime.Format(layoutDateTime), expected.EndTime.Format(layoutDateTime), expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		filter, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.NoError(t, err)
		assert.Equal(t, expected.RepositoryID, filter.RepositoryID)
		assert.Equal(t, expected.WorkspaceID, filter.WorkspaceID)
		assert.Equal(t, expected.StartTime, filter.StartTime)
		assert.Equal(t, expected.EndTime, filter.EndTime)
		assert.Equal(t, expected.Page, filter.Page)
		assert.Equal(t, expected.Size, filter.Size)
	})
	t.Run("Should extract filter without page size", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &repository.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s",
			expected.StartTime.Format(layoutDateTime), expected.EndTime.Format(layoutDateTime))
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		filter, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.NoError(t, err)
		assert.Equal(t, expected.RepositoryID, filter.RepositoryID)
		assert.Equal(t, expected.WorkspaceID, filter.WorkspaceID)
		assert.Equal(t, expected.StartTime, filter.StartTime)
		assert.Equal(t, expected.EndTime, filter.EndTime)
		assert.Equal(t, 0, filter.Page)
		assert.Equal(t, 10, filter.Size)
	})
	t.Run("Should return error on extract filter because not send empty startTime", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &repository.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?page=%v&size=%v",
			expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		_, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "EndTime: cannot be blank; StartTime: cannot be blank.")
	})
	t.Run("Should return error on extract filter because wrong workspaceID", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &repository.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			expected.StartTime.Format(layoutDateTime), expected.EndTime.Format(layoutDateTime), expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", "wrong type")
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		_, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.Error(t, err)
		assert.Equal(t, err, enums.ErrorWrongWorkspaceID)
	})
	t.Run("Should return error on extract filter because wrong repositoryID", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &repository.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=%s&page=%v&size=%v",
			expected.StartTime.Format(layoutDateTime), expected.EndTime.Format(layoutDateTime), expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", "wrong type")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		_, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.Error(t, err)
		assert.Equal(t, err, enums.ErrorWrongRepositoryID)
	})
	t.Run("Should return error on extract filter because wrong startTime", func(t *testing.T) {
		endTime, _ := time.Parse(layoutDateTime, "2022-01-01T00:00:00Z")
		expected := &repository.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			EndTime:      endTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=wrongStartTime&finalDate=%s&page=%v&size=%v",
			expected.EndTime.Format(layoutDateTime), expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		_, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.Error(t, err)
		assert.Equal(t, err, enums.ErrorWrongInitialDate)
	})
	t.Run("Should return error on extract filter because wrong endTime", func(t *testing.T) {
		startTime, _ := time.Parse(layoutDateTime, "2020-01-01T00:00:00Z")
		expected := &repository.Filter{
			RepositoryID: uuid.New(),
			WorkspaceID:  uuid.New(),
			StartTime:    startTime,
			Page:         0,
			Size:         10,
		}
		url := fmt.Sprintf("/test?initialDate=%s&finalDate=wrongStartTime&page=%v&size=%v",
			expected.StartTime.Format(layoutDateTime), expected.Page, expected.Size)
		r, _ := http.NewRequest(http.MethodGet, url, nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("workspaceID", expected.WorkspaceID.String())
		ctx.URLParams.Add("repositoryID", expected.RepositoryID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		_, err := NewUseCaseDashboard().ExtractFilterDashboard(r)
		assert.Error(t, err)
		assert.Equal(t, err, enums.ErrorWrongFinalDate)
	})
}

func getAnalysisMock() *analysis.Analysis {
	input := &analysis.Analysis{
		ID:                      uuid.New(),
		RepositoryID:            uuid.New(),
		RepositoryName:          "my-repository",
		WorkspaceID:             uuid.New(),
		WorkspaceName:           "my-workspace",
		Status:                  analysisEnum.Success,
		Errors:                  "",
		CreatedAt:               time.Now(),
		FinishedAt:              time.Now(),
		AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{},
	}
	VulnerabilityID1 := uuid.New()
	VulnerabilityID2 := uuid.New()
	input.AnalysisVulnerabilities = append(input.AnalysisVulnerabilities, analysis.AnalysisVulnerabilities{
		VulnerabilityID: VulnerabilityID1,
		AnalysisID:      input.ID,
		CreatedAt:       time.Now(),
		Vulnerability: vulnerability.Vulnerability{
			VulnerabilityID: VulnerabilityID1,
			Line:            "1",
			Column:          "1",
			Confidence:      confidence.High,
			File:            "/deployments/cert.pem",
			Code:            "-----BEGIN CERTIFICATE-----",
			Details:         "Asymmetric Private Key \n Found SSH and/or x.509 Cerficates among the files of your project, make sure you want this kind of information inside your Git repo, since it can be missused by someone with access to any kind of copy.  For more information checkout the CWE-312 (https://cwe.mitre.org/data/definitions/312.html) advisory.",
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
	})
	input.AnalysisVulnerabilities = append(input.AnalysisVulnerabilities, analysis.AnalysisVulnerabilities{
		VulnerabilityID: VulnerabilityID2,
		AnalysisID:      input.ID,
		CreatedAt:       time.Now(),
		Vulnerability: vulnerability.Vulnerability{
			VulnerabilityID: VulnerabilityID2,
			Line:            "1",
			Column:          "1",
			Confidence:      confidence.High,
			File:            "/deployments/key.pem",
			Code:            "-----BEGIN OPENSSH PRIVATE KEY-----",
			Details:         "Asymmetric Private Key \n Found SSH and/or x.509 Cerficates among the files of your project, make sure you want this kind of information inside your Git repo, since it can be missused by someone with access to any kind of copy.  For more information checkout the CWE-312 (https://cwe.mitre.org/data/definitions/312.html) advisory.",
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
	})
	return input
}

func TestUseCase_ParseAnalysisToVulnerabilitiesByAuthor(t *testing.T) {
	t.Run("Should parse analytic to vuln by author with success", func(t *testing.T) {
		input := getAnalysisMock()

		output := NewUseCaseAnalysis().ParseAnalysisToVulnerabilitiesByAuthor(input)

		assert.Len(t, output, 1)
		assert.Equal(t, output[0].Vulnerability.CriticalVulnerability, 2)
		assert.Equal(t, output[0].Vulnerability.WorkspaceID, input.WorkspaceID)
		assert.Equal(t, output[0].Vulnerability.RepositoryID, input.RepositoryID)
		assert.Equal(t, output[0].Vulnerability.Active, true)
		assert.Equal(t, output[0].Author, "horusec@zup.com.br")
	})
}

func TestUseCase_ParseAnalysisToVulnerabilitiesByLanguage(t *testing.T) {
	t.Run("Should parse analytic to vuln by language with success", func(t *testing.T) {
		input := getAnalysisMock()

		output := NewUseCaseAnalysis().ParseAnalysisToVulnerabilitiesByLanguage(input)

		assert.Len(t, output, 1)
		assert.Equal(t, output[0].Vulnerability.CriticalVulnerability, 2)
		assert.Equal(t, output[0].Vulnerability.WorkspaceID, input.WorkspaceID)
		assert.Equal(t, output[0].Vulnerability.RepositoryID, input.RepositoryID)
		assert.Equal(t, output[0].Vulnerability.Active, true)
		assert.Equal(t, output[0].Language, languages.Leaks)
	})
}

func TestUseCase_ParseAnalysisToVulnerabilitiesByRepository(t *testing.T) {
	t.Run("Should parse analytic to vuln by repository with success", func(t *testing.T) {
		input := getAnalysisMock()

		output := NewUseCaseAnalysis().ParseAnalysisToVulnerabilitiesByRepository(input)

		assert.Len(t, output, 1)
		assert.Equal(t, output[0].Vulnerability.CriticalVulnerability, 2)
		assert.Equal(t, output[0].Vulnerability.WorkspaceID, input.WorkspaceID)
		assert.Equal(t, output[0].Vulnerability.RepositoryID, input.RepositoryID)
		assert.Equal(t, output[0].Vulnerability.Active, true)
		assert.Equal(t, output[0].RepositoryName, input.RepositoryName)
	})
}

func TestUseCase_ParseAnalysisToVulnerabilitiesByTime(t *testing.T) {
	t.Run("Should parse analytic to vuln by time with success", func(t *testing.T) {
		input := getAnalysisMock()

		output := NewUseCaseAnalysis().ParseAnalysisToVulnerabilitiesByTime(input)

		assert.Len(t, output, 1)
		assert.Equal(t, output[0].Vulnerability.CriticalVulnerability, 2)
		assert.Equal(t, output[0].Vulnerability.WorkspaceID, input.WorkspaceID)
		assert.Equal(t, output[0].Vulnerability.RepositoryID, input.RepositoryID)
		assert.Equal(t, output[0].Vulnerability.Active, true)
		assert.Equal(t, output[0].CreatedAt, input.CreatedAt)
	})
}

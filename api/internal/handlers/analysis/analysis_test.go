package analysis

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/enums/confidence"

	tokensEnums "github.com/ZupIT/horusec-platform/api/internal/middelwares/token/enums"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	analysisController "github.com/ZupIT/horusec-platform/api/internal/controllers/analysis"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/cli"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	analysisEnum "github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"github.com/ZupIT/horusec-devkit/pkg/enums/tools"
	vulnerabilityEnum "github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
)

func TestHandler_Options(t *testing.T) {
	t.Run("should return 204 when options", func(t *testing.T) {
		handler := NewAnalysisHandler(&analysisController.Mock{})
		r, _ := http.NewRequest(http.MethodOptions, "/test", nil)
		w := httptest.NewRecorder()

		handler.Options(w, r)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

func TestHandler_Get(t *testing.T) {
	t.Run("should return 200 with analysis created", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("GetAnalysis").Return(&analysis.Analysis{}, nil)
		handler := NewAnalysisHandler(controllerMock)
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("analysisID", "85d08ec1-7786-4c2d-bf4e-5fee3a010315")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("should return 400 when not exists analysisID", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("GetAnalysis").Return(&analysis.Analysis{}, nil)
		handler := NewAnalysisHandler(controllerMock)
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		handler.Get(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("should return 404 when not exists analysis", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("GetAnalysis").Return(&analysis.Analysis{}, enums.ErrorNotFoundRecords)
		handler := NewAnalysisHandler(controllerMock)
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("analysisID", "85d08ec1-7786-4c2d-bf4e-5fee3a010315")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("should return 500 when return error unexpected", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("GetAnalysis").Return(&analysis.Analysis{}, errors.New("unexpected error"))
		handler := NewAnalysisHandler(controllerMock)
		r, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("analysisID", "85d08ec1-7786-4c2d-bf4e-5fee3a010315")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		handler.Get(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestHandler_Post(t *testing.T) {
	VulnerabilityID := uuid.New()
	AnalysisID := uuid.New()
	analysisDataMock := &cli.AnalysisData{
		Analysis: &analysis.Analysis{
			ID:         AnalysisID,
			Status:     analysisEnum.Running,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
			AnalysisVulnerabilities: []analysis.AnalysisVulnerabilities{
				{
					VulnerabilityID: VulnerabilityID,
					AnalysisID:      AnalysisID,
					CreatedAt:       time.Now(),
					Vulnerability: vulnerability.Vulnerability{
						VulnerabilityID: VulnerabilityID,
						Line:            "1",
						Column:          "1",
						Confidence:      confidence.High,
						File:            "/cmd/app/cert.pem",
						Code:            "=========PRIVATE KEY",
						Details:         "",
						SecurityTool:    tools.HorusecEngine,
						Language:        languages.Leaks,
						Severity:        severities.Critical,
						VulnHash:        "124356789",
						Type:            vulnerabilityEnum.Vulnerability,
					},
				},
			},
		},
		RepositoryName: "",
	}

	t.Run("should return 201 when analysis was created with success using token of repository", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.New(), nil)
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(analysisDataMock.ToBytes()))
		ctx := r.Context()
		ctx = context.WithValue(ctx, tokensEnums.RepositoryID, uuid.New())
		ctx = context.WithValue(ctx, tokensEnums.RepositoryName, uuid.New().String())
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceID, uuid.New())
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceName, uuid.New().String())
		r = r.WithContext(ctx)
		r.Header.Set("X-Horusec-Authorization", uuid.New().String())

		handler.Post(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return 201 when analysis was created with success using token of workspace", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.New(), nil)
		analysisWithRepositoryName := &cli.AnalysisData{
			Analysis:       analysisDataMock.Analysis,
			RepositoryName: uuid.New().String(),
		}
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(analysisWithRepositoryName.ToBytes()))
		ctx := r.Context()
		ctx = context.WithValue(ctx, tokensEnums.RepositoryID, uuid.Nil)
		ctx = context.WithValue(ctx, tokensEnums.RepositoryName, "")
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceID, uuid.New())
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceName, uuid.New().String())
		r = r.WithContext(ctx)
		r.Header.Set("X-Horusec-Authorization", uuid.New().String())

		handler.Post(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return 400 when analysis not exists status and other fields required", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.Nil, nil)
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader((&cli.AnalysisData{
			Analysis:       &analysis.Analysis{},
			RepositoryName: "",
		}).ToBytes()))
		ctx := r.Context()
		ctx = context.WithValue(ctx, tokensEnums.RepositoryID, uuid.Nil)
		ctx = context.WithValue(ctx, tokensEnums.RepositoryName, "")
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceID, uuid.New())
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceName, uuid.New().String())
		r = r.WithContext(ctx)

		handler.Post(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when analysis send wrong data", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.Nil, nil)
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader([]byte("[]")))
		ctx := r.Context()
		r = r.WithContext(ctx)

		handler.Post(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when not exists company on context", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(analysisDataMock.ToBytes()))
		ctx := r.Context()
		ctx = context.WithValue(ctx, tokensEnums.RepositoryID, uuid.Nil)
		ctx = context.WithValue(ctx, tokensEnums.RepositoryName, "")
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceID, uuid.Nil)
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceName, "")
		r = r.WithContext(ctx)

		handler.Post(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when not exists repository on context and not exists repository name on body", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(analysisDataMock.ToBytes()))
		ctx := r.Context()
		ctx = context.WithValue(ctx, tokensEnums.RepositoryID, uuid.Nil)
		ctx = context.WithValue(ctx, tokensEnums.RepositoryName, "")
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceID, uuid.New())
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceName, uuid.New().String())
		r = r.WithContext(ctx)

		handler.Post(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 500 when save analysis and return unknown error", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.Nil, errors.New("some unexpected error"))
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(analysisDataMock.ToBytes()))
		ctx := r.Context()
		ctx = context.WithValue(ctx, tokensEnums.RepositoryID, uuid.New())
		ctx = context.WithValue(ctx, tokensEnums.RepositoryName, uuid.New().String())
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceID, uuid.New())
		ctx = context.WithValue(ctx, tokensEnums.WorkspaceName, uuid.New().String())
		r = r.WithContext(ctx)
		r.Header.Set("X-Horusec-Authorization", uuid.New().String())

		handler.Post(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

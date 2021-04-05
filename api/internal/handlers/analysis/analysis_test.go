package analysis

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	analysisController "github.com/ZupIT/horusec-platform/api/internal/controllers/analysis"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/cli"
	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	analysisEnum "github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-devkit/pkg/enums/severities"
	"github.com/ZupIT/horusec-devkit/pkg/enums/tools"
	vulnerabilityEnum "github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	middlewaresEnums "github.com/ZupIT/horusec-devkit/pkg/services/middlewares/enums"
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
	var VulnerabilityID = uuid.New()
	var AnalysisID = uuid.New()
	var analysisDataMock = &cli.AnalysisData{
		Analysis: &analysis.Analysis{
			ID:         AnalysisID,
			Status:     analysisEnum.Running,
			Errors:     "",
			CreatedAt:  time.Now(),
			FinishedAt: time.Now(),
			AnalysisVulnerabilities: []analysis.RelationshipAnalysisVuln{
				{
					VulnerabilityID: VulnerabilityID,
					AnalysisID:      AnalysisID,
					CreatedAt:       time.Now(),
					Vulnerability: vulnerability.Vulnerability{
						VulnerabilityID: VulnerabilityID,
						Line:            "1",
						Column:          "1",
						Confidence:      "",
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
		RepositoryName: "repository",
	}

	t.Run("should return 201 when analysis was created with success", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.New(), nil)
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(analysisDataMock.ToBytes()))
		ctx := r.Context()
		ctx = context.WithValue(ctx, middlewaresEnums.RepositoryID, uuid.New())
		ctx = context.WithValue(ctx, middlewaresEnums.CompanyID, uuid.New())
		r = r.WithContext(ctx)
		r.Header.Set("X-Horusec-Authorization", uuid.New().String())

		handler.Post(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return 400 when analysis was created", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.Nil, nil)
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader((&cli.AnalysisData{
			Analysis:       &analysis.Analysis{},
			RepositoryName: "",
		}).ToBytes()))
		ctx := r.Context()
		ctx = context.WithValue(ctx, middlewaresEnums.RepositoryID, uuid.New())
		ctx = context.WithValue(ctx, middlewaresEnums.CompanyID, uuid.New())
		r = r.WithContext(ctx)
		r.Header.Set("X-Horusec-Authorization", uuid.New().String())

		handler.Post(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when analysis was created", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.Nil, nil)
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader([]byte("[]")))
		ctx := r.Context()
		ctx = context.WithValue(ctx, middlewaresEnums.RepositoryID, uuid.New())
		ctx = context.WithValue(ctx, middlewaresEnums.CompanyID, uuid.New())
		r = r.WithContext(ctx)
		r.Header.Set("X-Horusec-Authorization", uuid.New().String())

		handler.Post(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return 400 when not exists company on context", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.New(), nil)
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(analysisDataMock.ToBytes()))
		ctx := r.Context()
		r = r.WithContext(ctx)
		r.Header.Set("X-Horusec-Authorization", uuid.New().String())

		handler.Post(w, r)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 201 when analysis was created with success without repositoryID", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.New(), nil)
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(analysisDataMock.ToBytes()))
		ctx := r.Context()
		ctx = context.WithValue(ctx, middlewaresEnums.CompanyID, uuid.New())
		r = r.WithContext(ctx)
		r.Header.Set("X-Horusec-Authorization", uuid.New().String())

		handler.Post(w, r)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return 404 when analysis was created with success", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.Nil, enums.ErrorNotFoundRecords)
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(analysisDataMock.ToBytes()))
		ctx := r.Context()
		ctx = context.WithValue(ctx, middlewaresEnums.RepositoryID, uuid.New())
		ctx = context.WithValue(ctx, middlewaresEnums.CompanyID, uuid.New())
		r = r.WithContext(ctx)
		r.Header.Set("X-Horusec-Authorization", uuid.New().String())

		handler.Post(w, r)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should return 500 when analysis was created with success", func(t *testing.T) {
		controllerMock := &analysisController.Mock{}
		controllerMock.On("SaveAnalysis").Return(uuid.Nil, errors.New("some unexpected error"))
		handler := NewAnalysisHandler(controllerMock)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(analysisDataMock.ToBytes()))
		ctx := r.Context()
		ctx = context.WithValue(ctx, middlewaresEnums.RepositoryID, uuid.New())
		ctx = context.WithValue(ctx, middlewaresEnums.CompanyID, uuid.New())
		r = r.WithContext(ctx)
		r.Header.Set("X-Horusec-Authorization", uuid.New().String())

		handler.Post(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

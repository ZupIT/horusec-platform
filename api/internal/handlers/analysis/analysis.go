// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package analysis

import (
	"context"
	netHTTP "net/http"

	"github.com/ZupIT/horusec-devkit/pkg/services/tracer"

	"github.com/opentracing/opentracing-go"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	analysisController "github.com/ZupIT/horusec-platform/api/internal/controllers/analysis"
	handlersEnums "github.com/ZupIT/horusec-platform/api/internal/handlers/analysis/enums"
	tokenMiddlewareEnum "github.com/ZupIT/horusec-platform/api/internal/middelwares/token/enums"
	analysisUseCases "github.com/ZupIT/horusec-platform/api/internal/usecases/analysis"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"

	analysisEntities "github.com/ZupIT/horusec-devkit/pkg/entities/analysis"

	_ "github.com/ZupIT/horusec-devkit/pkg/entities/cli"        // [swagger-import]
	_ "github.com/ZupIT/horusec-devkit/pkg/utils/http/entities" // [swagger-import]
)

type Handler struct {
	controller analysisController.IController
	useCases   analysisUseCases.Interface
}

func NewAnalysisHandler(controller analysisController.IController) *Handler {
	return &Handler{
		controller: controller,
		useCases:   analysisUseCases.NewAnalysisUseCases(),
	}
}

func (h *Handler) Options(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
}

// Post
// @Tags Analysis
// @Security ApiKeyAuth
// @Description Start new analysis
// @ID start-new-analysis
// @Accept  json
// @Produce  json
// @Param SendNewAnalysis body cli.AnalysisData true "send new analysis info"
// @Success 201 {object} entities.Response{content=string} "CREATED"
// @Success 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Success 404 {object} entities.Response{content=string} "NOT FOUND"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /api/analysis [post]
func (h *Handler) Post(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "POST ANALYSIS")
	defer span.Finish()
	analysisEntity, err := h.validateAnalysis(ctx, r)
	if err != nil {
		tracer.SetSpanError(span, err)
		httpUtil.StatusBadRequest(w, err)
		return
	}
	h.saveAnalysis(r.Context(), w, analysisEntity)
}

func (h *Handler) validateAnalysis(ctx context.Context, r *netHTTP.Request) (*analysisEntities.Analysis, error) {
	analysisData, err := h.useCases.DecodeAnalysisDataFromIoRead(r.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	analysisEntity := h.decoratorAnalysisFromContext(analysisData.Analysis, r.WithContext(ctx))
	analysisEntity, err = h.decoratorAnalysisToRepositoryName(r.Context(), analysisEntity, analysisData.RepositoryName)
	if err != nil {
		return nil, err
	}
	return analysisEntity, nil
}

func (h *Handler) decoratorAnalysisFromContext(
	analysisEntity *analysisEntities.Analysis, r *netHTTP.Request) *analysisEntities.Analysis {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "decoratorAnalysisFromContext")
	defer span.Finish()
	r = r.WithContext(ctx)
	analysisEntity.WorkspaceID = r.Context().Value(tokenMiddlewareEnum.WorkspaceID).(uuid.UUID)
	analysisEntity.WorkspaceName = r.Context().Value(tokenMiddlewareEnum.WorkspaceName).(string)
	analysisEntity.RepositoryID = r.Context().Value(tokenMiddlewareEnum.RepositoryID).(uuid.UUID)
	analysisEntity.RepositoryName = r.Context().Value(tokenMiddlewareEnum.RepositoryName).(string)
	return analysisEntity
}

func (h *Handler) decoratorAnalysisToRepositoryName(ctx context.Context,
	analysisEntity *analysisEntities.Analysis, repositoryName string) (*analysisEntities.Analysis, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "decoratorAnalysisToRepositoryName")
	defer span.Finish()
	if h.isInvalidWorkspaceToCreateAnalysis(analysisEntity) {
		return nil, handlersEnums.ErrorWorkspaceNotSelected
	}
	if h.isValidRepositoryToCreateAnalysis(analysisEntity, repositoryName) {
		return nil, handlersEnums.ErrorRepositoryNotSelected
	}
	if h.isToCreateNewRepository(analysisEntity) {
		analysisEntity.RepositoryName = repositoryName
	}
	return analysisEntity, nil
}

func (h *Handler) isInvalidWorkspaceToCreateAnalysis(analysisEntity *analysisEntities.Analysis) bool {
	return analysisEntity.WorkspaceName == "" || analysisEntity.WorkspaceID == uuid.Nil
}

func (h *Handler) isValidRepositoryToCreateAnalysis(
	analysisEntity *analysisEntities.Analysis, repositoryName string) bool {
	return repositoryName == "" && analysisEntity.RepositoryName == ""
}

func (h *Handler) isToCreateNewRepository(analysisEntity *analysisEntities.Analysis) bool {
	return analysisEntity.RepositoryName == "" && analysisEntity.RepositoryID == uuid.Nil
}

func (h *Handler) saveAnalysis(ctx context.Context, w netHTTP.ResponseWriter,
	analysisEntity *analysisEntities.Analysis) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "decoratorAnalysisFromContext")
	defer span.Finish()
	analysisID, err := h.controller.SaveAnalysis(ctx, analysisEntity)
	if err != nil {
		tracer.SetSpanError(span, err)
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusCreated(w, analysisID)
}

// Get
// @Tags Analysis
// @Security ApiKeyAuth
// @Description Get analysis on database
// @ID get-one-analysis
// @Accept  json
// @Produce  json
// @Param analysisID path string true "analysisID of the analysis"
// @Success 200 {object} entities.Response{content=analysisEntities.Analysis} "OK"
// @Success 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Success 404 {object} entities.Response{content=string} "NOT FOUND"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /api/analysis/{analysisID} [get]
func (h *Handler) Get(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "GetAnalysis")
	defer span.Finish()
	analysisID, err := uuid.Parse(chi.URLParam(r, "analysisID"))
	if err != nil || analysisID == uuid.Nil {
		tracer.SetSpanError(span, err)
		httpUtil.StatusBadRequest(w, err)
		return
	}
	response, err := h.controller.GetAnalysis(ctx, analysisID)
	treatGetAnalysisResponse(w, err, span, response)
}

func treatGetAnalysisResponse(w netHTTP.ResponseWriter, err error, span opentracing.Span,
	response *analysisEntities.Analysis) {
	if err != nil {
		tracer.SetSpanError(span, err)
		if err == enums.ErrorNotFoundRecords {
			httpUtil.StatusNotFound(w, err)
		} else {
			httpUtil.StatusInternalServerError(w, err)
		}
	} else {
		httpUtil.StatusOK(w, response)
	}
}

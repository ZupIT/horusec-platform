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
	netHTTP "net/http"

	handlersEnums "github.com/ZupIT/horusec-platform/api/internal/handlers/analysis/enums"
	tokenMiddlewareEnum "github.com/ZupIT/horusec-platform/api/internal/middelwares/token/enums"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"

	analysisController "github.com/ZupIT/horusec-platform/api/internal/controllers/analysis"
	analysisUseCases "github.com/ZupIT/horusec-platform/api/internal/usecases/analysis"
	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
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

func (h *Handler) Post(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	analysisData, err := h.useCases.DecodeAnalysisDataFromIoRead(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	analysisEntity := h.decoratorAnalysisFromContext(analysisData.Analysis, r)
	analysisEntity, err = h.decoratorAnalysisToRepositoryName(analysisEntity, analysisData.RepositoryName)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	h.saveAnalysis(w, analysisEntity)
}

func (h *Handler) decoratorAnalysisFromContext(
	analysisEntity *analysis.Analysis, r *netHTTP.Request) *analysis.Analysis {
	analysisEntity.WorkspaceID = r.Context().Value(tokenMiddlewareEnum.WorkspaceID).(uuid.UUID)
	analysisEntity.WorkspaceName = r.Context().Value(tokenMiddlewareEnum.WorkspaceName).(string)
	analysisEntity.RepositoryID = r.Context().Value(tokenMiddlewareEnum.RepositoryID).(uuid.UUID)
	analysisEntity.RepositoryName = r.Context().Value(tokenMiddlewareEnum.RepositoryName).(string)
	return analysisEntity
}

func (h *Handler) decoratorAnalysisToRepositoryName(
	analysisEntity *analysis.Analysis, repositoryName string) (*analysis.Analysis, error) {
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

func (h *Handler) isInvalidWorkspaceToCreateAnalysis(analysisEntity *analysis.Analysis) bool {
	return analysisEntity.WorkspaceName == "" || analysisEntity.WorkspaceID == uuid.Nil
}

func (h *Handler) isValidRepositoryToCreateAnalysis(analysisEntity *analysis.Analysis, repositoryName string) bool {
	return repositoryName == "" && analysisEntity.RepositoryName == ""
}

func (h *Handler) isToCreateNewRepository(analysisEntity *analysis.Analysis) bool {
	return analysisEntity.RepositoryName == "" && analysisEntity.RepositoryID == uuid.Nil
}

func (h *Handler) saveAnalysis(w netHTTP.ResponseWriter, analysisEntity *analysis.Analysis) {
	analysisID, err := h.controller.SaveAnalysis(analysisEntity)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusCreated(w, analysisID)
}

func (h *Handler) Get(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	analysisID, err := uuid.Parse(chi.URLParam(r, "analysisID"))
	if err != nil || analysisID == uuid.Nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	response, err := h.controller.GetAnalysis(analysisID)
	if err != nil {
		if err == enums.ErrorNotFoundRecords {
			httpUtil.StatusNotFound(w, err)
		} else {
			httpUtil.StatusInternalServerError(w, err)
		}
	} else {
		httpUtil.StatusOK(w, response)
	}
}

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

	analysisController "github.com/ZupIT/horusec-platform/api/internal/controllers/analysis"
	analysisUseCases "github.com/ZupIT/horusec-platform/api/internal/usecases/analysis"
	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/entities/cli"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	middlewaresEnums "github.com/ZupIT/horusec-devkit/pkg/services/middlewares/enums"
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
	companyID, repositoryID, err := h.getCompanyIDAndRepositoryIDInCxt(r)
	if err != nil {
		httpUtil.StatusUnauthorized(w, err)
		return
	}
	analysisData = h.setCompanyIDRepositoryIDInAnalysis(analysisData, companyID, repositoryID)
	h.saveAnalysis(w, analysisData)
}

func (h *Handler) setCompanyIDRepositoryIDInAnalysis(
	analysisData *cli.AnalysisData, companyID uuid.UUID, repositoryID uuid.UUID) *cli.AnalysisData {
	analysisData.Analysis.CompanyID = companyID
	analysisData.Analysis.RepositoryID = repositoryID
	return analysisData
}

func (h *Handler) saveAnalysis(w netHTTP.ResponseWriter, analysisData *cli.AnalysisData) {
	analysisID, err := h.controller.SaveAnalysis(analysisData)
	if err != nil {
		if err == enums.ErrorNotFoundRecords {
			httpUtil.StatusNotFound(w, err)
			return
		}
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusCreated(w, analysisID)
}

func (h *Handler) getCompanyIDAndRepositoryIDInCxt(r *netHTTP.Request) (uuid.UUID, uuid.UUID, error) {
	companyIDCtx := r.Context().Value(middlewaresEnums.CompanyID)
	if companyIDCtx == nil {
		return uuid.Nil, uuid.Nil, middlewaresEnums.ErrorUnauthorized
	}
	repositoryIDCtx := r.Context().Value(middlewaresEnums.RepositoryID)
	if repositoryIDCtx == nil {
		return companyIDCtx.(uuid.UUID), uuid.Nil, nil
	}
	return companyIDCtx.(uuid.UUID), repositoryIDCtx.(uuid.UUID), nil
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

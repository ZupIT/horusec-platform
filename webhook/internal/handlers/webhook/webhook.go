// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package webhook

import (
	netHTTP "net/http"

	enumsWebhook "github.com/ZupIT/horusec-platform/webhook/internal/enums"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"

	controllerWebhook "github.com/ZupIT/horusec-platform/webhook/internal/controllers/webhook"
	useCase "github.com/ZupIT/horusec-platform/webhook/internal/usecases/webhook"
)

type Handler struct {
	controller controllerWebhook.IWebhookController
	useCase    useCase.IUseCaseWebhook
}

func NewWebhookHandler(controller controllerWebhook.IWebhookController) *Handler {
	return &Handler{
		controller: controller,
		useCase:    useCase.NewUseCaseWebhook(),
	}
}

func (h *Handler) Options(w netHTTP.ResponseWriter, _ *netHTTP.Request) {
	httpUtil.StatusNoContent(w)
}

// ListAll
// @Tags Webhook
// @Security ApiKeyAuth
// @Description Get all webhooks by workspace
// @ID GetAllWebhooksByWorkspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Success 200 {object} entities.Response{content=[]webhook.Webhook} "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /webhook/webhook/{workspaceID} [get]
func (h *Handler) ListAll(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	workspaceID, err := h.useCase.ExtractWorkspaceIDFromURL(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	webhooks, err := h.controller.ListAll(workspaceID)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusOK(w, webhooks)
}

// Remove
// @Tags Webhook
// @Security ApiKeyAuth
// @Description Remove webhook by id
// @ID RemoveWebhookByID
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param webhookID path string true "webhookID of the webhook"
// @Success 204    "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /webhook/webhook/{workspaceID}/{webhookID} [delete]
func (h *Handler) Remove(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	webhookID, err := h.useCase.ExtractWebhookIDFromURL(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	if err := h.controller.Remove(webhookID); err != nil {
		if err == enums.ErrorNotFoundRecords {
			httpUtil.StatusNotFound(w, err)
		} else {
			httpUtil.StatusInternalServerError(w, err)
		}
	} else {
		httpUtil.StatusNoContent(w)
	}
}

// Update
// @Tags Webhook
// @Security ApiKeyAuth
// @Description Update webhook by id
// @ID UpdateWebhookByID
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param webhookID path string true "webhookID of the webhook"
// @Param webhookToUpdate body webhook.Webhook true "update webhook content info"
// @Success 204    "OK"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /webhook/webhook/{workspaceID}/{webhookID} [put]
func (h *Handler) Update(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	webhookID, err := h.useCase.ExtractWebhookIDFromURL(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	body, err := h.useCase.DecodeWebhookFromIoRead(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	h.updateWebhook(w, body, webhookID)
}

func (h *Handler) updateWebhook(w netHTTP.ResponseWriter, body *webhook.Webhook, webhookID uuid.UUID) {
	if err := h.controller.Update(body, webhookID); err != nil {
		if err == enums.ErrorNotFoundRecords {
			httpUtil.StatusNotFound(w, err)
			return
		}
		httpUtil.StatusInternalServerError(w, err)
		return
	}
	httpUtil.StatusNoContent(w)
}

// Save
// @Tags Webhook
// @Security ApiKeyAuth
// @Description Save webhook by id
// @ID SaveWebhook
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "workspaceID of the workspace"
// @Param webhookToSave body webhook.Webhook true "update webhook content info"
// @Success 200 {object} entities.Response{content=string} "NO CONTENT"
// @Failure 400 {object} entities.Response{content=string} "BAD REQUEST"
// @Failure 500 {object} entities.Response{content=string} "INTERNAL SERVER ERROR"
// @Router /webhook/webhook/{workspaceID} [post]
func (h *Handler) Save(w netHTTP.ResponseWriter, r *netHTTP.Request) {
	body, err := h.useCase.DecodeWebhookFromIoRead(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}
	webhookID, err := h.controller.Save(body)
	if err != nil {
		if err == enumsWebhook.ErrorWebhookDuplicate {
			httpUtil.StatusConflict(w, err)
		} else {
			httpUtil.StatusInternalServerError(w, err)
		}
	} else {
		httpUtil.StatusOK(w, webhookID)
	}
}

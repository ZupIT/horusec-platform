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

	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
	"github.com/ZupIT/horusec-platform/webhook/internal/enums"
)

type IUseCaseWebhook interface {
	DecodeWebhookFromIoRead(r *netHTTP.Request) (*webhook.Webhook, error)
	ExtractWebhookIDFromURL(r *netHTTP.Request) (uuid.UUID, error)
	ExtractWorkspaceIDFromURL(r *netHTTP.Request) (uuid.UUID, error)
}

type UseCaseWebhook struct{}

func NewUseCaseWebhook() IUseCaseWebhook {
	return &UseCaseWebhook{}
}

func (uc *UseCaseWebhook) DecodeWebhookFromIoRead(r *netHTTP.Request) (entity *webhook.Webhook, err error) {
	if err := parser.ParseBodyToEntity(r.Body, &entity); err != nil {
		return nil, err
	}
	return entity, uc.validateWebhook(entity)
}

func (uc *UseCaseWebhook) ExtractWebhookIDFromURL(r *netHTTP.Request) (uuid.UUID, error) {
	ID, err := uuid.Parse(chi.URLParam(r, "webhookID"))
	if err != nil || ID == uuid.Nil {
		return uuid.Nil, enums.ErrorWrongWebhookID
	}
	return ID, nil
}

func (uc *UseCaseWebhook) ExtractWorkspaceIDFromURL(r *netHTTP.Request) (uuid.UUID, error) {
	ID, err := uuid.Parse(chi.URLParam(r, "workspaceID"))
	if err != nil || ID == uuid.Nil {
		return uuid.Nil, enums.ErrorWrongWorkspaceID
	}
	return ID, nil
}

func (uc *UseCaseWebhook) validateWebhook(entity *webhook.Webhook) error {
	return validation.ValidateStruct(entity,
		validation.Field(&entity.URL, validation.Required, is.URL),
		validation.Field(&entity.Method, validation.Required, validation.In(netHTTP.MethodPost)),
		validation.Field(&entity.RepositoryID, validation.Required, is.UUID),
		validation.Field(&entity.WorkspaceID, validation.Required, is.UUID),
	)
}

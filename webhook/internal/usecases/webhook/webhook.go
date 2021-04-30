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
	if err != nil {
		return uuid.Nil, enums.ErrorWrongWebhookID
	}
	return ID, nil
}

func (uc *UseCaseWebhook) ExtractWorkspaceIDFromURL(r *netHTTP.Request) (uuid.UUID, error) {
	vulnerabilityID, err := uuid.Parse(chi.URLParam(r, "workspaceID"))
	if err != nil {
		return uuid.Nil, enums.ErrorWrongWorkspaceID
	}
	return vulnerabilityID, nil
}

func (uc *UseCaseWebhook) validateWebhook(entity *webhook.Webhook) error {
	return validation.ValidateStruct(entity,
		validation.Field(&entity.URL, validation.Required, is.URL),
		validation.Field(&entity.Method, validation.Required, validation.In(netHTTP.MethodPost)),
		validation.Field(&entity.RepositoryID, validation.Required, is.UUID),
		validation.Field(&entity.WorkspaceID, validation.Required, is.UUID),
	)
}

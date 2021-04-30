package dispatcher

import (
	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/services/http/request"

	webhookEntity "github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
	"github.com/ZupIT/horusec-platform/webhook/internal/repositories/webhook"
)

type IDispatcherController interface {
	DispatchRequest(entity *analysis.Analysis) error
}

type Controller struct {
	repository  webhook.IWebhookRepository
	httpRequest request.IRequest
}

func NewDispatcherController(repository webhook.IWebhookRepository) IDispatcherController {
	const DefaultTimeoutOnRequests = 10
	return &Controller{
		repository:  repository,
		httpRequest: request.NewHTTPRequestService(DefaultTimeoutOnRequests),
	}
}

func (c *Controller) DispatchRequest(entity *analysis.Analysis) error {
	webhookFound, err := c.repository.ListOne(map[string]interface{}{"repository_id": entity.RepositoryID})
	if err != nil {
		return err
	}
	return c.sendHTTPRequest(webhookFound, entity)
}

func (c *Controller) sendHTTPRequest(webhookFound *webhookEntity.Webhook, entity *analysis.Analysis) error {
	req, err := c.httpRequest.NewHTTPRequest(webhookFound.Method, webhookFound.URL, entity,
		webhookFound.Headers.GetMapHeaders())
	if err != nil {
		return err
	}
	res, err := c.httpRequest.DoRequest(req, nil)
	if err != nil {
		return err
	}
	defer res.CloseBody()
	return res.ErrorByStatusCode()
}

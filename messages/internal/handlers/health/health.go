package health

import (
	"net/http"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	_ "github.com/ZupIT/horusec-devkit/pkg/utils/http/entities" // swagger import

	"github.com/ZupIT/horusec-platform/messages/internal/enums/health"
	"github.com/ZupIT/horusec-platform/messages/internal/services/mailer"
)

type Handler struct {
	mailerService mailer.IService
	broker        broker.IBroker
}

func NewHealthHandler(brokerLib broker.IBroker, mailerService mailer.IService) *Handler {
	return &Handler{
		broker:        brokerLib,
		mailerService: mailerService,
	}
}

func (h *Handler) Options(w http.ResponseWriter, _ *http.Request) {
	httpUtil.StatusNoContent(w)
}

// @Tags Health
// @Description Check if application is healthy
// @ID health
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.Response
// @Failure 200 {object} entities.Response
// @Router /messages/health [get]
// @Security ApiKeyAuth
func (h *Handler) Get(w http.ResponseWriter, _ *http.Request) {
	if !h.mailerService.IsAvailable() {
		httpUtil.StatusInternalServerError(w, health.ErrorUnhealthyMailer)
		return
	}

	if !h.broker.IsAvailable() {
		httpUtil.StatusInternalServerError(w, health.ErrorUnhealthyBroker)
		return
	}

	httpUtil.StatusOK(w, nil)
}

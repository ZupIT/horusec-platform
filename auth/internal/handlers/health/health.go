package health

import (
	"net/http"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	_ "github.com/ZupIT/horusec-devkit/pkg/utils/http/entities" // swagger import

	"github.com/ZupIT/horusec-platform/auth/internal/enums/health"
)

type Handler struct {
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	broker        broker.IBroker
}

func NewHealthHandler(connection *database.Connection, brokerLib broker.IBroker) *Handler {
	return &Handler{
		databaseWrite: connection.Write,
		databaseRead:  connection.Read,
		broker:        brokerLib,
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
// @Router /auth/health [get]
// @Security ApiKeyAuth
func (h *Handler) Get(w http.ResponseWriter, _ *http.Request) {
	if !h.databaseWrite.IsAvailable() || !h.databaseRead.IsAvailable() {
		httpUtil.StatusInternalServerError(w, health.ErrorUnhealthyDatabase)
		return
	}

	if !h.broker.IsAvailable() {
		httpUtil.StatusInternalServerError(w, health.ErrorUnhealthyBroker)
		return
	}

	httpUtil.StatusOK(w, nil)
}

package analysis

import (
	httpUtils "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	"github.com/ZupIT/horusec-platform/api/internal/controllers/analysis"
	"net/http"
)

type Handler struct {
	controller analysis.IController
}

func NewAnalysisHandler(controller analysis.IController) *Handler {
	return &Handler{
		controller: controller,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	h.controller.Test()

	httpUtils.StatusOK(w, "service is healthy")
}

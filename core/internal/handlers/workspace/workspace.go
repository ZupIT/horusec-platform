package workspace

import (
	httpUtils "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	"github.com/ZupIT/horusec-platform/core/internal/controllers/workspace"
	"net/http"
)

type Handler struct {
	controller workspace.IController
}

func NewWorkspaceHandler(controller workspace.IController) *Handler {
	return &Handler{
		controller: controller,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	h.controller.Test()

	httpUtils.StatusOK(w, "service is healthy")
}

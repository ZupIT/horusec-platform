package workspace

import (
	"context"
	"net/http"

	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt/enums"
	"github.com/google/uuid"

	workspaceController "github.com/ZupIT/horusec-platform/core/internal/controllers/workspace"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

type Handler struct {
	controller workspaceController.IController
	useCases   workspaceUseCases.IUseCases
	authGRPC   proto.AuthServiceClient
	context    context.Context
}

func NewWorkspaceHandler(controller workspaceController.IController, useCases workspaceUseCases.IUseCases,
	authGRPC proto.AuthServiceClient) *Handler {
	return &Handler{
		controller: controller,
		useCases:   useCases,
		authGRPC:   authGRPC,
		context:    context.Background(),
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	workspaceData, err := h.useCases.GetCreateWorkspaceData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	response, err := h.authGRPC.GetAccountInfo(h.context,
		&proto.GetAccountData{Token: r.Header.Get(enums.HorusecJWTHeader)})
	if err != nil {
		return
	}

	accountID, _ := uuid.Parse(response.AccountID)
	workspaceData.AccountID = accountID
	workspaceData.Permissions = response.Permissions

	workspace, err := h.controller.Create(workspaceData)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusCreated(w, workspace)
}

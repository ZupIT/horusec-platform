package workspace

import (
	"context"
	"net/http"

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt/enums"

	workspaceController "github.com/ZupIT/horusec-platform/core/internal/controllers/workspace"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

type Handler struct {
	controller workspaceController.IController
	useCases   workspaceUseCases.IUseCases
	authGRPC   proto.AuthServiceClient
	context    context.Context
	appConfig  app.IConfig
}

func NewWorkspaceHandler(controller workspaceController.IController, useCases workspaceUseCases.IUseCases,
	authGRPC proto.AuthServiceClient, appConfig app.IConfig) *Handler {
	return &Handler{
		controller: controller,
		useCases:   useCases,
		authGRPC:   authGRPC,
		context:    context.Background(),
		appConfig:  appConfig,
	}
}

// @Tags Workspace
// @Description Create a new workspace
// @ID create-workspace
// @Accept  json
// @Produce  json
// @Param Workspace body workspaceEntities.CreateWorkspaceData true "create workspace data"
// @Router /core/workspaces [post]
// @Security ApiKeyAuth
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	workspaceData, err := h.getCreateData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	workspace, err := h.controller.Create(workspaceData)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusCreated(w, workspace)
}

func (h *Handler) getCreateData(r *http.Request) (*workspaceEntities.CreateWorkspaceData, error) {
	workspaceData, err := h.useCases.GetCreateWorkspaceData(r.Body)
	if err != nil {
		return nil, err
	}

	accountData, err := h.getAccountData(r)
	if err != nil {
		return nil, err
	}

	workspaceData.SetAccountData(accountData)
	return workspaceData, workspaceData.CheckLdapGroups(h.appConfig.GetAuthorizationType())
}

func (h *Handler) getAccountData(r *http.Request) (*proto.GetAccountDataResponse, error) {
	return h.authGRPC.GetAccountInfo(h.context, &proto.GetAccountData{Token: r.Header.Get(enums.HorusecJWTHeader)})
}

package repository

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser"

	repositoryController "github.com/ZupIT/horusec-platform/core/internal/controllers/repository"
	repositoryEntities "github.com/ZupIT/horusec-platform/core/internal/entities/repository"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	workspaceEnums "github.com/ZupIT/horusec-platform/core/internal/enums/workspace"
	repositoryUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
)

type Handler struct {
	useCases   repositoryUseCases.IUseCases
	controller repositoryController.IController
	appConfig  app.IConfig
	authGRPC   proto.AuthServiceClient
	context    context.Context
}

func NewRepositoryHandler(useCases repositoryUseCases.IUseCases, controller repositoryController.IController,
	appConfig app.IConfig, authGRPC proto.AuthServiceClient) *Handler {
	return &Handler{
		useCases:   useCases,
		controller: controller,
		appConfig:  appConfig,
		authGRPC:   authGRPC,
		context:    context.Background(),
	}
}

func (h *Handler) getAccountData(r *http.Request) (*proto.GetAccountDataResponse, error) {
	return h.authGRPC.GetAccountInfo(h.context, &proto.GetAccountData{Token: r.Header.Get(enums.HorusecJWTHeader)})
}

// @Tags Repository
// @Description Create a new repository
// @ID create-repository
// @Accept  json
// @Produce  json
// @Param Repository body repositoryEntities.Data true "create repository data"
// @Success 201 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories [post]
// @Security ApiKeyAuth
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	workspaceData, err := h.getCreateData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	workspace, err := h.controller.Create(workspaceData)
	if err != nil {
		h.checkCreateRepositoryErrors(w, err)
		return
	}

	httpUtil.StatusCreated(w, workspace)
}

func (h *Handler) getCreateData(r *http.Request) (*repositoryEntities.Data, error) {
	accountData, err := h.getAccountData(r)
	if err != nil {
		return nil, err
	}

	data, err := h.useCases.RepositoryDataFromIOReadCloser(r.Body)
	if err != nil {
		return nil, err
	}

	return data.SetWorkspaceIDAndAccountData(parser.ParseStringToUUID(chi.URLParam(r, workspaceEnums.ID)), accountData),
		data.CheckLdapGroups(h.appConfig.GetAuthorizationType())
}

func (h *Handler) checkCreateRepositoryErrors(w http.ResponseWriter, err error) {
	if err == repositoryEnums.ErrorRepositoryNameAlreadyInUse {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

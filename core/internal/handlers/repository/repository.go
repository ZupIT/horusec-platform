package repository

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"
	_ "github.com/ZupIT/horusec-devkit/pkg/utils/http/entities" // swagger import
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
// @Param workspaceID path string true "ID of the workspace"
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

	return data.SetWorkspaceIDAndAccountData(parser.ParseStringToUUID(
		chi.URLParam(r, workspaceEnums.ID)), accountData), data.CheckLdapGroups(h.appConfig.GetAuthorizationType())
}

func (h *Handler) checkCreateRepositoryErrors(w http.ResponseWriter, err error) {
	if err == repositoryEnums.ErrorRepositoryNameAlreadyInUse {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

// @Tags Repository
// @Description Search for a existing repository by id
// @ID get-repository
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param repositoryID path string true "ID of the repository"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories/{repositoryID} [get]
// @Security ApiKeyAuth
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	data, err := h.getByIDData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	repository, err := h.controller.Get(data)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, repository)
}

func (h *Handler) getByIDData(r *http.Request) (*repositoryEntities.Data, error) {
	repositoryID, err := uuid.Parse(chi.URLParam(r, repositoryEnums.ID))
	if err != nil {
		return nil, err
	}

	accountData, err := h.getAccountData(r)
	if err != nil {
		return nil, err
	}

	return h.useCases.NewRepositoryData(repositoryID, parser.ParseStringToUUID(
		chi.URLParam(r, workspaceEnums.ID)), accountData), nil
}

// @Tags Repository
// @Description Updates a existing repository by id
// @ID update-repository
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param repositoryID path string true "ID of the repository"
// @Param Repository body repositoryEntities.Data true "update repository data"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories/{repositoryID} [patch]
// @Security ApiKeyAuth
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	data, err := h.getUpdateData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	repository, err := h.controller.Update(data)
	if err != nil {
		h.checkUpdateRepositoryErrors(w, err)
		return
	}

	httpUtil.StatusOK(w, repository)
}

func (h *Handler) getUpdateData(r *http.Request) (*repositoryEntities.Data, error) {
	repositoryID, err := uuid.Parse(chi.URLParam(r, repositoryEnums.ID))
	if err != nil {
		return nil, err
	}

	data, err := h.getCreateData(r)
	if err != nil {
		return nil, err
	}

	return data.SetWorkspaceAndRepositoryID(parser.ParseStringToUUID(chi.URLParam(r, workspaceEnums.ID)),
		repositoryID), data.CheckLdapGroups(h.appConfig.GetAuthorizationType())
}

func (h *Handler) checkUpdateRepositoryErrors(w http.ResponseWriter, err error) {
	if err == repositoryEnums.ErrorRepositoryNameAlreadyInUse {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

// @Tags Repository
// @Description Delete a repository by id
// @ID delete-repository
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param repositoryID path string true "ID of the repository"
// @Success 204 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories/{repositoryID} [delete]
// @Security ApiKeyAuth
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	repositoryID, err := uuid.Parse(chi.URLParam(r, repositoryEnums.ID))
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	if err = h.controller.Delete(repositoryID); err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusNoContent(w)
}

// @Tags Repository
// @Description List all repositories of an account in a workspace
// @ID list-repositories
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories [get]
// @Security ApiKeyAuth
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	data, err := h.getListData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	repositories, err := h.controller.List(data)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, repositories)
}

func (h *Handler) getListData(r *http.Request) (*repositoryEntities.Data, error) {
	accountData, err := h.getAccountData(r)
	if err != nil {
		return nil, err
	}

	return h.useCases.NewRepositoryData(uuid.Nil, parser.ParseStringToUUID(
		chi.URLParam(r, workspaceEnums.ID)), accountData), nil
}

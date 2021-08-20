// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	tokenEntities "github.com/ZupIT/horusec-platform/core/internal/entities/token"
	repositoryEnums "github.com/ZupIT/horusec-platform/core/internal/enums/repository"
	roleEnums "github.com/ZupIT/horusec-platform/core/internal/enums/role"
	tokenEnums "github.com/ZupIT/horusec-platform/core/internal/enums/token"
	workspaceEnums "github.com/ZupIT/horusec-platform/core/internal/enums/workspace"
	repositoryUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/repository"
	roleUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/role"
	tokenUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/token"
)

type Handler struct {
	useCases      repositoryUseCases.IUseCases
	controller    repositoryController.IController
	appConfig     app.IConfig
	authGRPC      proto.AuthServiceClient
	context       context.Context
	roleUseCases  roleUseCases.IUseCases
	tokenUseCases tokenUseCases.IUseCases
}

func NewRepositoryHandler(useCases repositoryUseCases.IUseCases, controller repositoryController.IController,
	appConfig app.IConfig, authGRPC proto.AuthServiceClient, useCasesRole roleUseCases.IUseCases,
	useCasesToken tokenUseCases.IUseCases) *Handler {
	return &Handler{
		useCases:      useCases,
		controller:    controller,
		appConfig:     appConfig,
		authGRPC:      authGRPC,
		context:       context.Background(),
		roleUseCases:  useCasesRole,
		tokenUseCases: useCasesToken,
	}
}

func (h *Handler) Options(w http.ResponseWriter, _ *http.Request) {
	httpUtil.StatusNoContent(w)
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
		chi.URLParam(r, workspaceEnums.ID)), accountData), data.CheckLdapGroups(h.appConfig.GetAuthenticationType())
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
		repositoryID), data.CheckLdapGroups(h.appConfig.GetAuthenticationType())
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
	repositoryData, paginatedData, err := h.getListData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	repositories, err := h.controller.List(repositoryData, paginatedData)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, repositories)
}

func (h *Handler) getListData(r *http.Request) (*repositoryEntities.Data, *repositoryEntities.PaginatedContent, error) {
	accountData, err := h.getAccountData(r)
	if err != nil {
		return nil, nil, err
	}

	paginatedData := &repositoryEntities.PaginatedContent{}

	repositoryData := h.useCases.NewRepositoryData(uuid.Nil, parser.ParseStringToUUID(
		chi.URLParam(r, workspaceEnums.ID)), accountData)
	paginatedData = paginatedData.SetEnable(r.URL.Query().Get(repositoryEnums.Page) != "").
		SetPage(r.URL.Query().Get(repositoryEnums.Page)).
		SetSize(r.URL.Query().Get(repositoryEnums.Size)).
		SetSearch(r.URL.Query().Get(repositoryEnums.Search))
	return repositoryData, paginatedData, nil
}

// @Tags Repository
// @Description Update an account role of a repository
// @ID update-repository-role
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param accountID path string true "ID of the account"
// @Param repositoryID path string true "ID of the repository"
// @Param Role body roleEntities.Data true "update role of a account in a specific workspace"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories/{repositoryID}/roles/{accountID} [patch]
// @Security ApiKeyAuth
func (h *Handler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	data, err := h.getUpdateRoleData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	role, err := h.controller.UpdateRole(data)
	if err != nil {
		h.checkUpdateRoleErrors(w, err)
		return
	}

	httpUtil.StatusOK(w, role)
}

func (h *Handler) getUpdateRoleData(r *http.Request) (*roleEntities.Data, error) {
	data, err := h.roleUseCases.RoleDataFromIOReadCloser(r.Body)
	if err != nil {
		return nil, err
	}

	accountID, err := uuid.Parse(chi.URLParam(r, roleEnums.AccountID))
	if err != nil {
		return nil, err
	}

	return data.SetDataIDs(accountID, chi.URLParam(r, workspaceEnums.ID), chi.URLParam(r, repositoryEnums.ID)), nil
}

func (h *Handler) checkUpdateRoleErrors(w http.ResponseWriter, err error) {
	if err == repositoryEnums.ErrorUserDoesNotBelongToWorkspace {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

// @Tags Repository
// @Description Invite a user to a repository
// @ID invite-user-repository
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param repositoryID path string true "ID of the repository"
// @Param User Data body roleEntities.UserData true "user account data"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories/{repositoryID}/roles [post]
// @Security ApiKeyAuth
func (h *Handler) InviteUser(w http.ResponseWriter, r *http.Request) {
	data, err := h.getInviteUserData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	role, err := h.controller.InviteUser(data)
	if err != nil {
		h.checkInviteUserErrors(w, err)
		return
	}

	httpUtil.StatusOK(w, role)
}

func (h *Handler) getInviteUserData(r *http.Request) (*roleEntities.UserData, error) {
	data, err := h.roleUseCases.InviteUserDataFromIOReadCloser(r.Body)
	if err != nil {
		return nil, err
	}

	return data.SetIDs(chi.URLParam(r, workspaceEnums.ID), chi.URLParam(r, repositoryEnums.ID)), nil
}

func (h *Handler) checkInviteUserErrors(w http.ResponseWriter, err error) {
	if err == repositoryEnums.ErrorUserDoesNotBelongToWorkspace {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

// @Tags Repository
// @Description Get all users of a repository
// @ID get-repository-users
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param repositoryID path string true "ID of the repository"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories/{repositoryID}/roles [get]
// @Security ApiKeyAuth
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	repositoryID, err := uuid.Parse(chi.URLParam(r, repositoryEnums.ID))
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	users, err := h.controller.GetUsers(repositoryID)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, users)
}

// @Tags Repository
// @Description Remove a user from a repository
// @ID remove-repository-user
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param repositoryID path string true "ID of the repository"
// @Param accountID path string true "ID of the account"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories/{repositoryID}/roles/{accountID} [delete]
// @Security ApiKeyAuth
func (h *Handler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	data, err := h.getRemoveUserData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	if err := h.controller.RemoveUser(data); err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusNoContent(w)
}

func (h *Handler) getRemoveUserData(r *http.Request) (*roleEntities.Data, error) {
	accountID, err := uuid.Parse(chi.URLParam(r, roleEnums.AccountID))
	if err != nil {
		return nil, err
	}

	return h.roleUseCases.NewRoleData(accountID, parser.ParseStringToUUID(chi.URLParam(r, workspaceEnums.ID)),
		parser.ParseStringToUUID(chi.URLParam(r, repositoryEnums.ID))), nil
}

// @Tags Repository
// @Description Create a new repository token
// @ID create-repository-token
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param repositoryID path string true "ID of the repository"
// @Param Token body tokenEntities.Data true "create repository token data"
// @Success 201 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories/{repositoryID}/tokens [post]
// @Security ApiKeyAuth
func (h *Handler) CreateToken(w http.ResponseWriter, r *http.Request) {
	data, err := h.getCreateTokenData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	token, err := h.controller.CreateToken(data)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusCreated(w, token)
}

func (h *Handler) getCreateTokenData(r *http.Request) (*tokenEntities.Data, error) {
	data, err := h.tokenUseCases.TokenDataFromIOReadCloser(r.Body)
	if err != nil {
		return nil, err
	}

	return data.SetIDsString(uuid.Nil, chi.URLParam(r, workspaceEnums.ID), chi.URLParam(r, repositoryEnums.ID)), nil
}

// @Tags Repository
// @Description Delete a repository token
// @ID delete-repository-token
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param repositoryID path string true "ID of the repository"
// @Param tokenID path string true "ID of the token"
// @Success 204 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories/{repositoryID}/tokens/{tokenID} [delete]
// @Security ApiKeyAuth
func (h *Handler) DeleteToken(w http.ResponseWriter, r *http.Request) {
	data, err := h.getDeleteTokenData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	if err := h.controller.DeleteToken(data); err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusNoContent(w)
}

func (h *Handler) getDeleteTokenData(r *http.Request) (*tokenEntities.Data, error) {
	tokenID, err := uuid.Parse(chi.URLParam(r, tokenEnums.ID))
	if err != nil {
		return nil, err
	}

	return h.tokenUseCases.NewTokenData(tokenID, chi.URLParam(r, workspaceEnums.ID),
		chi.URLParam(r, repositoryEnums.ID)), nil
}

// @Tags Repository
// @Description List all repository tokens
// @ID list-repository-tokens
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param repositoryID path string true "ID of the repository"
// @Success 201 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/repositories/{repositoryID}/tokens [get]
// @Security ApiKeyAuth
func (h *Handler) ListTokens(w http.ResponseWriter, r *http.Request) {
	data, err := h.getListTokensData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	tokens, err := h.controller.ListTokens(data)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, tokens)
}

func (h *Handler) getListTokensData(r *http.Request) (*tokenEntities.Data, error) {
	data := &tokenEntities.Data{}

	workspaceID, err := uuid.Parse(chi.URLParam(r, workspaceEnums.ID))
	if err != nil {
		return nil, err
	}

	repositoryID, err := uuid.Parse(chi.URLParam(r, repositoryEnums.ID))
	if err != nil {
		return nil, err
	}

	return data.SetIDs(workspaceID, repositoryID, uuid.Nil), nil
}

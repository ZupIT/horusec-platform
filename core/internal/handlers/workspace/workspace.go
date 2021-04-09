package workspace

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

	workspaceController "github.com/ZupIT/horusec-platform/core/internal/controllers/workspace"
	roleEntities "github.com/ZupIT/horusec-platform/core/internal/entities/role"
	workspaceEntities "github.com/ZupIT/horusec-platform/core/internal/entities/workspace"
	roleEnums "github.com/ZupIT/horusec-platform/core/internal/enums/role"
	workspaceEnums "github.com/ZupIT/horusec-platform/core/internal/enums/workspace"
	roleUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/role"
	workspaceUseCases "github.com/ZupIT/horusec-platform/core/internal/usecases/workspace"
)

type Handler struct {
	controller   workspaceController.IController
	useCases     workspaceUseCases.IUseCases
	roleUseCases roleUseCases.IUseCases
	authGRPC     proto.AuthServiceClient
	context      context.Context
	appConfig    app.IConfig
}

func NewWorkspaceHandler(controller workspaceController.IController, useCases workspaceUseCases.IUseCases,
	authGRPC proto.AuthServiceClient, appConfig app.IConfig, useCasesRole roleUseCases.IUseCases) *Handler {
	return &Handler{
		controller:   controller,
		useCases:     useCases,
		authGRPC:     authGRPC,
		context:      context.Background(),
		appConfig:    appConfig,
		roleUseCases: useCasesRole,
	}
}

func (h *Handler) getAccountData(r *http.Request) (*proto.GetAccountDataResponse, error) {
	return h.authGRPC.GetAccountInfo(h.context, &proto.GetAccountData{Token: r.Header.Get(enums.HorusecJWTHeader)})
}

// @Tags Workspace
// @Description Create a new workspace
// @ID create-workspace
// @Accept  json
// @Produce  json
// @Param Workspace body workspaceEntities.Data true "create workspace data"
// @Success 201 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 500 {object} entities.Response
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

func (h *Handler) getCreateData(r *http.Request) (*workspaceEntities.Data, error) {
	accountData, err := h.getAccountData(r)
	if err != nil {
		return nil, err
	}

	workspaceData, err := h.useCases.WorkspaceDataFromIOReadCloser(r.Body)
	if err != nil {
		return nil, err
	}

	return workspaceData.SetAccountData(accountData), workspaceData.CheckLdapGroups(h.appConfig.GetAuthorizationType())
}

// @Tags Workspace
// @Description Search for a existing workspace by id
// @ID get-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID} [get]
// @Security ApiKeyAuth
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	data, err := h.getByIDData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	workspace, err := h.controller.Get(data)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, workspace)
}

func (h *Handler) getByIDData(r *http.Request) (*workspaceEntities.Data, error) {
	workspaceID, err := uuid.Parse(chi.URLParam(r, workspaceEnums.ID))
	if err != nil {
		return nil, err
	}

	accountData, err := h.getAccountData(r)
	if err != nil {
		return nil, err
	}

	return h.useCases.NewWorkspaceData(workspaceID, accountData), nil
}

// @Tags Workspace
// @Description Updates a existing workspace by id
// @ID update-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param Workspace body workspaceEntities.Data true "update workspace data"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID} [patch]
// @Security ApiKeyAuth
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	data, err := h.getUpdateData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	workspace, err := h.controller.Update(data)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, workspace)
}

func (h *Handler) getUpdateData(r *http.Request) (*workspaceEntities.Data, error) {
	workspaceID, err := uuid.Parse(chi.URLParam(r, workspaceEnums.ID))
	if err != nil {
		return nil, err
	}

	data, err := h.getCreateData(r)
	if err != nil {
		return nil, err
	}

	return data.SetWorkspaceID(workspaceID), data.CheckLdapGroups(h.appConfig.GetAuthorizationType())
}

// @Tags Workspace
// @Description Delete a workspace by id
// @ID delete-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Success 204 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID} [delete]
// @Security ApiKeyAuth
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	workspaceID, err := uuid.Parse(chi.URLParam(r, workspaceEnums.ID))
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	if err = h.controller.Delete(workspaceID); err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusNoContent(w)
}

// @Tags Workspace
// @Description List all workspaces of an account
// @ID list-workspace
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces [get]
// @Security ApiKeyAuth
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	data, err := h.getListData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	workspaces, err := h.controller.List(data)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, workspaces)
}

func (h *Handler) getListData(r *http.Request) (*workspaceEntities.Data, error) {
	accountData, err := h.getAccountData(r)
	if err != nil {
		return nil, err
	}

	return h.useCases.NewWorkspaceData(uuid.Nil, accountData), nil
}

// @Tags Workspace
// @Description Update an account role of a workspace
// @ID update-workspace-role
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param accountID path string true "ID of the account"
// @Param Workspace body roleEntities.Data true "update role of a account in a specific workspace"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/roles/{accountID} [patch]
// @Security ApiKeyAuth
func (h *Handler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	data, err := h.getUpdateRoleData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	role, err := h.controller.UpdateRole(data)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
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

	return data.SetDataIDs(accountID, chi.URLParam(r, workspaceEnums.ID), ""), nil
}

// @Tags Workspace
// @Description Invite a user to a workspace
// @ID invite-user-workspace
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param Workspace body roleEntities.UserData true "user account data"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/roles [post]
// @Security ApiKeyAuth
func (h *Handler) InviteUser(w http.ResponseWriter, r *http.Request) {
	data, err := h.getInviteUserData(r)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	role, err := h.controller.InviteUser(data)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, role)
}

func (h *Handler) getInviteUserData(r *http.Request) (*roleEntities.UserData, error) {
	data, err := h.roleUseCases.InviteUserDataFromIOReadCloser(r.Body)
	if err != nil {
		return nil, err
	}

	return data.SetIDs(chi.URLParam(r, workspaceEnums.ID), ""), nil
}

// @Tags Workspace
// @Description Get all users of a workspace
// @ID get-users
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/roles [get]
// @Security ApiKeyAuth
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	workspaceID, err := uuid.Parse(chi.URLParam(r, workspaceEnums.ID))
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	users, err := h.controller.GetUsers(workspaceID)
	if err != nil {
		httpUtil.StatusInternalServerError(w, err)
		return
	}

	httpUtil.StatusOK(w, users)
}

// @Tags Workspace
// @Description Remove a user from a workspace
// @ID remove-user
// @Accept  json
// @Produce  json
// @Param workspaceID path string true "ID of the workspace"
// @Param accountID path string true "ID of the account"
// @Success 200 {object} entities.Response
// @Failure 400 {object} entities.Response
// @Failure 401 {object} entities.Response
// @Failure 404 {object} entities.Response
// @Failure 500 {object} entities.Response
// @Router /core/workspaces/{workspaceID}/roles/{accountID} [delete]
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
	workspaceID, err := uuid.Parse(chi.URLParam(r, workspaceEnums.ID))
	if err != nil {
		return nil, err
	}

	accountID, err := uuid.Parse(chi.URLParam(r, roleEnums.AccountID))
	if err != nil {
		return nil, err
	}

	return h.roleUseCases.NewRoleData(accountID, workspaceID, uuid.Nil), nil
}

package account

import (
	"net/http"

	httpUtil "github.com/ZupIT/horusec-devkit/pkg/utils/http"

	accountController "github.com/ZupIT/horusec-platform/auth/internal/controllers/account"
	accountEnums "github.com/ZupIT/horusec-platform/auth/internal/enums/account"
	accountUseCases "github.com/ZupIT/horusec-platform/auth/internal/usecases/account"
)

type Handler struct {
	useCases   accountUseCases.IUseCases
	controller accountController.IController
}

func NewAccountHandler(useCases accountUseCases.IUseCases, controller accountController.IController) *Handler {
	return &Handler{
		useCases:   useCases,
		controller: controller,
	}
}

func (h *Handler) CreateAccountKeycloak(w http.ResponseWriter, r *http.Request) {
	keyCloakToken, err := h.useCases.AccessTokenFromIOReadCloser(r.Body)
	if err != nil {
		httpUtil.StatusBadRequest(w, err)
		return
	}

	response, err := h.controller.CreateAccountKeycloak(keyCloakToken)
	if err != nil {
		h.checkCreateAccountKeycloakErrors(w, err, response)
		return
	}

	httpUtil.StatusOK(w, response)
}

func (h *Handler) checkCreateAccountKeycloakErrors(w http.ResponseWriter, err error, response interface{}) {
	if err == accountEnums.ErrorEmailAlreadyInUse || err == accountEnums.ErrorUsernameAlreadyInUse {
		httpUtil.StatusOK(w, response)
		return
	}

	httpUtil.StatusInternalServerError(w, err)
}

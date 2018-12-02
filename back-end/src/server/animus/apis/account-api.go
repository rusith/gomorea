package apis

import (
	"net/http"
	"server/app/apis"
	"server/app/services"
	"server/app/utils"
)

type AccountApi struct {
	Api
	accountService services.IAccountService
	apiUtils utils.IApiUtils
}

func NewAccountApi(service services.IAccountService, apiUtils utils.IApiUtils) apis.IAccountApi {
	return &AccountApi{ accountService: service, apiUtils: apiUtils }
}

func (a *AccountApi) SignIn(w http.ResponseWriter, r *http.Request) {
	body := a.ReadJson(r, w, &SignInModel{}).(*SignInModel)
	if body == nil { return }
	if body.Username == "" {
		a.apiUtils.SendJson(http.StatusBadRequest, "Username is required", w)
		return
	}
	if body.Password == "" {
		a.apiUtils.SendJson(http.StatusBadRequest, "Password is required", w)
		return
	}

	_, err := a.accountService.ValidateUser(body.Username, body.Password)
	if err != nil {
		a.apiUtils.SendJsonError(http.StatusBadRequest, err, w)
		return
	}
}


type SignInModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
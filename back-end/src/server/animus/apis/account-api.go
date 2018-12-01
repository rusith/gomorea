package apis

import (
	"net/http"
	"server/app/apis"
	"server/app/services"
)

type AccountApi struct {
  accountService services.IAccountService
}

func NewAccountApi(service services.IAccountService) apis.IAccountApi {
	return &AccountApi{ accountService: service }
}

func (a *AccountApi) SignIn(w http.ResponseWriter, r *http.Request) {
	a.accountService.ValidateUser("", "")
}

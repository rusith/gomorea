package apis

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"server/app/apis"
	"server/app/config"
	"server/app/models"
	"server/app/services"
	"server/app/utils"
)

type AccountApi struct {
	Api
	accountService services.IAccountService
	apiUtils utils.IApiUtils
	config config.IAppConfig
}

func NewAccountApi(service services.IAccountService, apiUtils utils.IApiUtils, config config.IAppConfig) apis.IAccountApi {
	return &AccountApi{ accountService: service, apiUtils: apiUtils, config: config }
}

func (a *AccountApi) SignIn(w http.ResponseWriter, r *http.Request) {
	body := a.ReadJson(r, w, &models.User{}).(*models.User)
	if body == nil { return }
	if body.Username == "" {
		a.apiUtils.SendJson(http.StatusBadRequest, "Username is required", w)
		return
	}
	if body.Password == "" {
		a.apiUtils.SendJson(http.StatusBadRequest, "Password is required", w)
		return
	}

	user, err := a.accountService.ValidateUser(body.Username, body.Password)
	if err != nil {
		a.apiUtils.SendJsonError(http.StatusBadRequest, err, w)
		return
	}

	tk := &Token{UserId: user.ID.String()}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv(a.config.TokenPassword())))

	a.apiUtils.SendJson(http.StatusOK, tokenString, w)
}

func (a *AccountApi) SignUp(w http.ResponseWriter, r *http.Request) {
	body := a.ReadJson(r, w, &models.User{}).(*models.User)
	if body == nil { return }
	if body.FirstName == "" {
		a.apiUtils.SendJson(http.StatusBadRequest, "First name is required", w)
		return
	}
	if body.LastName == "" {
		a.apiUtils.SendJson(http.StatusBadRequest, "Last name is required", w)
		return
	}
	if body.Username == "" {
		a.apiUtils.SendJson(http.StatusBadRequest, "Username is required", w)
		return
	}
	if body.Password == "" {
		a.apiUtils.SendJson(http.StatusBadRequest, "Password is required", w)
		return
	}

	created, err := a.accountService.CreateUser(body)
	if err != nil {
		a.apiUtils.SendJsonError(http.StatusInternalServerError, err, w)
	}
	created.Password = ""

	a.apiUtils.SendJson(http.StatusOK, created, w)
}


type Token struct {
	UserId string
	jwt.StandardClaims
}
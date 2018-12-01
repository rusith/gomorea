package apis

import "net/http"

type IAccountApi interface {
	SignIn(w http.ResponseWriter, r *http.Request)
}

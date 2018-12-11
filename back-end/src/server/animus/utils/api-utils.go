package utils

import (
	"encoding/json"
	"net/http"
	"server/app/utils"
)

type ApiUtils struct {
}

func NewApiUtils() utils.IApiUtils {
	return &ApiUtils{}
}

func (a *ApiUtils) SendJson(statusCode int, data interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (a *ApiUtils) SendJsonError(statusCode int, err error, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err.Error())
}

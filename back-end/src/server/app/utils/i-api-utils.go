package utils

import "net/http"

type IApiUtils interface {
	SendJson(statusCode int, data interface{}, w http.ResponseWriter)
}

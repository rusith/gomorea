package apis

import (
	"encoding/json"
	"net/http"
)

type Api struct {

}


func (Api) ReadJson(r *http.Request, w http.ResponseWriter, t interface{}) interface{} {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("Invalid JSON"))
	}
	return t
}

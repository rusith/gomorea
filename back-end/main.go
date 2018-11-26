package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.Handle("/token", getTokenHandler).Methods("GET")
	router.Handle("/get", jwtMiddleware.Handler(notImplemented)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

var notImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

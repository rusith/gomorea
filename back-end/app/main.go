package main

import (
	"fmt"
	"github.com/gorilla/context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/token", getTokenHandler).Methods("GET")
	router.Handle("/get", jwtMiddleware.Handler(notImplemented)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

var notImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println(context.GetAll(r))
	w.Write([]byte("Not Implemented"))
})

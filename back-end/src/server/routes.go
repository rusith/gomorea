package main

import "github.com/gorilla/mux"

func (s *Server) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/accounts", s.accountApi.SignIn).Methods("GET")
	router.HandleFunc("/accounts", s.accountApi.SignUp).Methods("POST")
}

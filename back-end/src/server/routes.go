package main

import "github.com/gorilla/mux"

func (s *Server) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/accounts/token", s.accountApi.SignIn).Methods("POST")
	router.HandleFunc("/accounts", s.accountApi.SignUp).Methods("POST")
}

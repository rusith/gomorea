package main

import "github.com/gorilla/mux"

func (s *Server) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/account/login", s.accountApi.SignIn)
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"server/app"
	"server/models/account"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)
	router.Handle("/users/login", account.Login)
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	fmt.Println(port)
	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}

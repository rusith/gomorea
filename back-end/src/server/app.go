package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"server/app"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)
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

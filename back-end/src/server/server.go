package main

//noinspection SpellCheckingInspection
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	iapi "server/app/apis"
	aconfig "server/app/config"
	autils "server/app/utils"
)


type Server struct {
	accountApi iapi.IAccountApi
	apiUtils autils.IApiUtils
	appConfig aconfig.IAppConfig
}

func NewServer(accountApi iapi.IAccountApi, apiUtils autils.IApiUtils, appConfig aconfig.IAppConfig) *Server {
	return &Server{
		accountApi,
		apiUtils,
		appConfig,
	}
}


func (s *Server) Run() {
	router := mux.NewRouter()
	router.Use(s.GetJwtMiddleware)
	s.RegisterRoutes(router)

	port := s.appConfig.Port()
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}


func main() {

	container := BuildContainer()

	err := container.Invoke(func(server *Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}

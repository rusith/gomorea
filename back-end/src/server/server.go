package main

//noinspection SpellCheckingInspection
import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
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

	headersOk := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"})

	err := http.ListenAndServe(":" + port, handlers.CORS(originsOk, headersOk, methodsOk)(router))
	if err != nil {
		fmt.Print(err)
	}
}


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	container := BuildContainer()

	err = container.Invoke(func(server *Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}

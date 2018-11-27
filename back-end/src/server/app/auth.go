package app

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	models "server/models/auth"
	u "server/utils"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {

		/* Paths that don't need authentication */
		notAuth := []string{ "/user/login", "/user/new" }
		requestPath := r.URL.Path

		/* If path does not need authentication, just return */
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		response := make(map[string] interface{})


		/* If token header not present, then return */
		if tokenHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			u.Respond(w, u.Message(false, "Authentication Failed"))
			return
		}

		/* if auth token doesn't have to parts, its invalid */
		splitted := strings.Split(tokenHeader, "")
		if len(splitted) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			u.Respond(w, u.Message(false, "Auth token is invalid"))
			return
		}


		/* Parse the token */
		tokenPart := splitted[1]
		tk := models.Token{}
		_, err := jwt.ParseWithClaims(tokenPart, tk, func( token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = u.Message(false, "Auth token is invalid")
			u.Respond(w, response)
			return
		}

		/* Continue with context */
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
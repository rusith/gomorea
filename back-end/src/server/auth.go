package main

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}


func (s *Server) GetJwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {

		/* Paths that don't need authentication */
		notAuth := []string{ "/account/login", "/account/new" }
		requestPath := r.URL.Path

		/* If path does not need authentication, just return */
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		/* If token header not present, then return */
		if tokenHeader == "" {
			s.apiUtils.SendJson(http.StatusUnauthorized, "Authentication Failed", w)
			return
		}

		/* if auth token doesn't have to parts, its invalid */
		sp := strings.Split(tokenHeader, "")
		if len(sp) != 2 {
			s.apiUtils.SendJson(http.StatusUnauthorized, "Auth token is invalid", w)
			return
		}


		/* Parse the token */
		tokenPart := sp[1]
		tk := Token{}
		_, err := jwt.ParseWithClaims(tokenPart, tk, func( token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			s.apiUtils.SendJson(http.StatusUnauthorized, "Auth token is invalid", w)
			return
		}

		/* Continue with context */
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

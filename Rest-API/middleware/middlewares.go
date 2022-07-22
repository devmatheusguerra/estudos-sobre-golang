package middleware

import (
	"devmatheusguerra/rest/utils"
	"encoding/json"
	"net/http"
)

func ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/auth" {
			next.ServeHTTP(w, r)

		} else {

			token := r.Header.Get("Authorization")
			if token == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				token = token[7:]
				if utils.Validar(token) {
					if utils.VerificarSeTokenEstaExpirado(token) {
						novoToken := utils.RenovarJWT(token)
						w.WriteHeader(http.StatusUnauthorized)
						json.NewEncoder(w).Encode(novoToken)
					} else {
						next.ServeHTTP(w, r)
					}

				} else {
					json.NewEncoder(w).Encode(map[string]string{
						"Error": "Invalid Token",
					})
				}
			}

		}

	})
}

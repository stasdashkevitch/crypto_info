package middleware

import (
	"net/http"

	"github.com/stasdashkevitch/crypto_info/internal/usecase/service/auth"
)

func AuthMiddleware(auth auth.Auth, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		token = token[len("Bearer "):]

		_, err := auth.ValidateToken(token)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

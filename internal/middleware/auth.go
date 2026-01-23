package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/sujeet-crossml/GoLang_Backend_Project/internal/utils"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteError(w, http.StatusUnauthorized, "Authorization header misssing")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteError(w, http.StatusUnauthorized, "Invalid token format")
			return
		}

		userID, err := utils.ValidateToken(parts[1])
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))

	}
}

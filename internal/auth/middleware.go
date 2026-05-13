package auth

import (
	"context"
	"net/http"
	"strings"
)

// internal/auth/middleware.go
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization") // "Bearer <token>"
		if token != "" {
			userID, err := ParseToken(strings.TrimPrefix(token, "Bearer "))
			if err == nil {
				ctx := context.WithValue(r.Context(), UserIDKey, userID)
				r = r.WithContext(ctx)
			}
		}
		next.ServeHTTP(w, r)
	})
}

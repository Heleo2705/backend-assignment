package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"example.com/backend-assignment/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed token"))
		} else {
			claims, err := utils.VerifyToken(authHeader[1])
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(fmt.Sprintf("error while verifying token %v", err)))
				return
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, "userID", claims.UserID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

		}
	})
}

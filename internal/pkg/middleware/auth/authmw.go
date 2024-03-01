package middleware

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func NewAuthMiddleware(uc usecase.Auth) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionID, err := r.Cookie("session_id")
			if errors.Is(err, http.ErrNoCookie) {
				http.Error(w, "no session", http.StatusUnauthorized)
				return
			}
			userID, err := uc.GetUserIDBySessionID(sessionID.Value)
			ctx := context.WithValue(context.Background(), "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

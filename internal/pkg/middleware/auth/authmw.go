package middleware

import (
	"context"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"net/http"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type UserKey string

const ContextUserKey UserKey = "user_id"

func NewAuthMiddleware(uc usecase.Auth) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionID, err := r.Cookie("session_id")
			if errors.Is(err, http.ErrNoCookie) {
				http.Error(w, "no session", http.StatusUnauthorized)
				return
			}
			userID, err := uc.GetUserIDBySessionID(sessionID.Value)
			if errors.Is(err, models.ErrNoSession) {
				http.Error(w, "no session", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(context.Background(), ContextUserKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

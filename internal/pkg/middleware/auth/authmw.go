package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func NewAuthMiddleware(uc usecase.Auth) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionID, err := r.Cookie("session_id")
			if errors.Is(err, http.ErrNoCookie) {
				err := helper.JSONResponse(w, 200, models.ErrResponse{
					Status: 401,
					Msg:    "no session",
					MsgRus: "авторизация отсутствует",
				})
				if err != nil {
					log.Printf("marshall error: %v", err)
				}
				return
			}
			userID, err := uc.GetUserIDBySessionID(sessionID.Value)
			if errors.Is(err, models.ErrNoSession) {
				err := helper.JSONResponse(w, 200, models.ErrResponse{
					Status: 401,
					Msg:    "no session",
					MsgRus: "авторизация отсутствует",
				})
				if err != nil {
					log.Printf("marshall error: %v", err)
				}
				return
			}
			ctx := context.WithValue(context.Background(), models.ContextUserKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

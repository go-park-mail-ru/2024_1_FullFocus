package middleware

import (
	"context"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

func NewAuthMiddleware(uc usecase.Auth) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			sessionID, err := r.Cookie("session_id")
			if errors.Is(err, http.ErrNoCookie) {
				helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
					Status: 401,
					Msg:    "no session",
					MsgRus: "авторизация отсутствует",
				})
				return
			}
			userID, err := uc.GetUserIDBySessionID(r.Context(), sessionID.Value)
			if errors.Is(err, models.ErrNoSession) {
				helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
					Status: 401,
					Msg:    "no session",
					MsgRus: "авторизация отсутствует",
				})
				return
			}
			ctx = context.WithValue(ctx, helper.UserID{}, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

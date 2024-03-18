package middleware

import (
	"context"
	"fmt"
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
			l := helper.GetLoggerFromContext(ctx)

			sessionID, err := r.Cookie("session_id")
			if errors.Is(err, http.ErrNoCookie) {
				err := helper.JSONResponse(w, 200, models.ErrResponse{
					Status: 401,
					Msg:    "no session",
					MsgRus: "авторизация отсутствует",
				})
				if err != nil {
					l.Error(fmt.Sprintf("marshall error: %v", err))
				}
				return
			}
			userID, err := uc.GetUserIDBySessionID(r.Context(), sessionID.Value)
			if errors.Is(err, models.ErrNoSession) {
				err := helper.JSONResponse(w, 200, models.ErrResponse{
					Status: 401,
					Msg:    "no session",
					MsgRus: "авторизация отсутствует",
				})
				if err != nil {
					l.Error(fmt.Sprintf("marshall error: %v", err))
				}
				return
			}
			ctx = context.WithValue(ctx, helper.UserId{}, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

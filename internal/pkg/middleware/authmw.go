package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	auth "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/auth/grpc"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
)

func NewAuthMiddleware(c *auth.Client) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if !strings.Contains(r.URL.Path, "public") {
				sessionID, err := r.Cookie("session_id")
				if errors.Is(err, http.ErrNoCookie) {
					helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
						Status: 401,
						Msg:    "no session",
						MsgRus: "авторизация отсутствует",
					})
					return
				}
				userID, err := c.GetUserIDBySessionID(ctx, sessionID.Value)
				if errors.Is(err, models.ErrNoSession) {
					helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
						Status: 401,
						Msg:    "no session",
						MsgRus: "авторизация отсутствует",
					})
					return
				}
				ctx = context.WithValue(ctx, helper.UserID{}, userID)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

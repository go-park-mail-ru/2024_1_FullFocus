package middleware

import (
	"fmt"
	csrf "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func CSRFMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			tokens, _ := csrf.NewJwtToken("qsRY2e4hcM5T7X984E9WQ5uZ8Nty7fxB")
			if r.Method == "POST" || r.Method == "PUT" || r.Method == "DELETE" || r.Method == "PATCH" {
				session, err := r.Cookie("session_id")
				sID := session.Value
				CSRFToken := r.FormValue("X-Csrf-Token")
				_, err = tokens.Check(sID, CSRFToken)
				if err != nil {
					w.Write([]byte("{}"))
					logger.Debug(ctx, fmt.Sprintf("csrf token validation error: %v", err))
					return
				}
			} else if r.Method == "GET" {
				session, err := r.Cookie("session_id")
				sID := session.Value
				token, err := tokens.Create(sID, time.Now().Add(1*time.Hour).Unix())
				if err != nil {
					logger.Debug(ctx, fmt.Sprintf("csrf token creation error: %v", err))
					return
				}
				w.Header().Set("X-Csrf-Token", token)
			}
			next.ServeHTTP(w, r)
		})
	}
}

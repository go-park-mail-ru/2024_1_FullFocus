package middleware

import (
	"fmt"
	"slices"
	"time"

	csrf "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/gorilla/mux"

	"net/http"
)

func CSRFMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			httpMethods := []string{http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch}
			found := slices.Contains(httpMethods, r.Method)
			if found {
				err := CheckSCRFToken(r)
				if err != nil {
					logger.Debug(ctx, fmt.Sprintf("csrf token creation error: %v", err))
					return
				}
			}
			if r.Method == http.MethodGet {
				err := SetSCRFToken(w, r)
				if err != nil {
					logger.Debug(ctx, fmt.Sprintf("csrf token creation error: %v", err))
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

func SetSCRFToken(w http.ResponseWriter, r *http.Request) error {
	tokens, _ := csrf.NewJwtToken("qsRY2e4hcM5T7X984E9WQ5uZ8Nty7fxB")
	session, err := r.Cookie("session_id")
	if err != nil {
		return err
	}
	sID := session.Value
	token, err := tokens.Create(sID, time.Now().Add(1*time.Hour).Unix())
	if err != nil {
		return err
	}
	w.Header().Set("X-Csrf-Token", token)
	return nil
}

func CheckSCRFToken(r *http.Request) error {
	tokens, _ := csrf.NewJwtToken("qsRY2e4hcM5T7X984E9WQ5uZ8Nty7fxB")
	session, err := r.Cookie("session_id")
	if err != nil {
		return err
	}
	sID := session.Value
	csrfToken := r.FormValue("X-Csrf-Token")
	_, err = tokens.Check(sID, csrfToken)
	if err != nil {
		return err
	}
	return nil
}

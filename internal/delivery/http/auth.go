package delivery

import (
	"context"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const (
	SessionTTL = time.Hour * 24
)

type AuthHandler struct {
	srv     *http.Server
	router  *mux.Router
	usecase usecase.Auth
}

func NewAuthHandler(s *http.Server, uc usecase.Auth) *AuthHandler {
	return &AuthHandler{
		srv:     s,
		router:  mux.NewRouter(),
		usecase: uc,
	}
}

func (h *AuthHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/auth").Subrouter()
	{
		h.router.Handle("/login", http.HandlerFunc(h.Login)).Methods("GET", "POST", "OPTIONS")
		h.router.Handle("/signup", http.HandlerFunc(h.Signup)).Methods("GET", "POST", "OPTIONS")
		h.router.Handle("/logout", http.HandlerFunc(h.Logout)).Methods("POST", "OPTIONS")
	}
}

func (h *AuthHandler) Run() error {
	return h.srv.ListenAndServe()
}

func (h *AuthHandler) Stop() error {
	return h.srv.Shutdown(context.Background())
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	sID, err := h.usecase.Login(login, password)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sID,
		Secure:   true,
		HttpOnly: true,
		Expires:  time.Now().Add(SessionTTL),
	}
	http.SetCookie(w, cookie)
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	sID, _, err := h.usecase.Signup(login, password)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sID,
		Secure:   true,
		HttpOnly: true,
		Expires:  time.Now().Add(SessionTTL),
	}
	http.SetCookie(w, cookie)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		http.Error(w, `no session`, http.StatusUnauthorized)
		return
	}
	err = h.usecase.Logout(session.Value)
	if errors.Is(err, models.ErrNoSession) {
		http.Error(w, `no session`, http.StatusUnauthorized)
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

package delivery

import (
	"context"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/gorilla/mux"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
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

// Auth godoc
// @Tags Auth
// @Summary Init router
// @Description init new router
// @Param r body object true "*mux.Router"
// @Router /auth [get]
func (h *AuthHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/auth").Subrouter()
	{
		h.router.Handle("/login", http.HandlerFunc(h.Login)).Methods("GET", "POST", "OPTIONS")
		h.router.Handle("/signup", http.HandlerFunc(h.Signup)).Methods("GET", "POST", "OPTIONS")
		h.router.Handle("/logout", http.HandlerFunc(h.Logout)).Methods("POST", "OPTIONS")
	}
}

// Auth godoc
// @Tags Auth
// @Summary Run server
func (h *AuthHandler) Run() error {
	return h.srv.ListenAndServe()
}

// Auth godoc
// @Tags Auth
// @Summary Stop server
func (h *AuthHandler) Stop() error {
	return h.srv.Shutdown(context.Background())
}

// Auth godoc
// @Tags Auth
// @Summary Login
// @Description let login
// @Param w body object true "ResponseWriter"
// @Param r body object true "Request"
// @Router /auth/login [post]
// @Router /auth/login [get]
// @Router /auth/login [options]
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

// Auth godoc
// @Tags Auth
// @Summary Signup
// @Description let signup
// @Param w body object true "ResponseWriter"
// @Param r body object true "Request"
// @Router /auth/signup [post]
// @Router /auth/signup [get]
// @Router /auth/singnup [options]
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

// Auth godoc
// @Tags Auth
// @Summary Logout
// @Description let logout
// @Param w body object true "ResponseWriter"
// @Param r body object true "Request"
// @Router /auth/logout [post]
// @Router /auth/logout [options]
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

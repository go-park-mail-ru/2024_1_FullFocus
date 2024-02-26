package delivery

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

const (
	SessionTTL = time.Hour * 24
)

type AuthHandler struct {
	usecase usecase.Auth
}

func NewAuthHandler(uc usecase.Auth) *AuthHandler {
	return &AuthHandler{
		usecase: uc,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	sID, err := h.usecase.Login(login, password)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sID,
		Expires: time.Now().Add(SessionTTL),
	}
	http.SetCookie(w, cookie)
	w.Write([]byte(sID))
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	sID, uID, err := h.usecase.Signup(login, password)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sID,
		Expires: time.Now().Add(SessionTTL),
	}
	http.SetCookie(w, cookie)
	w.Write([]byte(uID))
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		http.Error(w, `no session`, 401)
		return
	}
	if !h.usecase.IsLoggedIn(session.Value) {
		http.Error(w, `no session`, 401)
		return
	}
	_ = h.usecase.Logout(session.Value)
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

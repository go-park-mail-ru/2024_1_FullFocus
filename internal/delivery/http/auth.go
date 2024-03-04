package delivery

import (
	"context"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/gorilla/mux"

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
	switch err {
	case models.ErrNoUser:
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    "wrong login",
			MsgRus: "логин введен неправильно или учетная запись не существует",
		})
		return
	case models.ErrWrongPassword:
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    "wrong password",
			MsgRus: "введен неверный пароль",
		})
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

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	sID, _, err := h.usecase.Signup(login, password)
	switch err {
	case models.ErrUserAlreadyExists:
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    "user already exists",
			MsgRus: "невозможно создать пользователя с таким логином, такой уже существует",
		})
		return
	case models.ErrShortUsername:
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    "too short username",
			MsgRus: "логин должен содержать больше 5 символов",
		})
		return
	case models.ErrWeakPassword:
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    "too weak password",
			MsgRus: "пароль слишком простой",
		})
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
	session, errSession := r.Cookie("session_id")
	if errors.Is(errSession, http.ErrNoCookie) {
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    "no session",
			MsgRus: "авторизация отсутствует",
		})
		return
	}
	errUsecase := h.usecase.Logout(session.Value)
	if errors.Is(errUsecase, models.ErrNoSession) {
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    "no session",
			MsgRus: "авторизация отсутствует",
		})
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

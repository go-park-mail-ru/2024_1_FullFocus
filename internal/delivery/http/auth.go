package delivery

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
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

type LoginData struct {
	login    string `json:"login"`
	password string `json:"password"`
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
		h.router.Handle("/check", http.HandlerFunc(h.CheckAuth)).Methods("GET", "OPTIONS")
	}
}

func (h *AuthHandler) Run() error {
	return h.srv.ListenAndServe()
}

func (h *AuthHandler) Stop() error {
	return h.srv.Shutdown(context.Background())
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var logindata LoginData
	err := json.NewDecoder(r.Body).Decode(&logindata)
	if err != nil {
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}

	sID, err := h.usecase.Login(logindata.login, logindata.password)
	if err != nil {
		if validationError := new(models.ValidationError); errors.As(err, &validationError) {
			helper.JSONResponse(w, 200, validationError.WithCode(400))
			return
		}
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Неверный логин или пароль",
		})
		return
	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sID,
		HttpOnly: true,
		Expires:  time.Now().Add(SessionTTL),
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	})
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var logindata LoginData
	err := json.NewDecoder(r.Body).Decode(&logindata)
	if err != nil {
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}

	sID, _, err := h.usecase.Signup(logindata.login, logindata.password)
	if err != nil {
		if validationError := new(models.ValidationError); errors.As(err, &validationError) {
			helper.JSONResponse(w, 200, validationError.WithCode(400))
			return
		}
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Пользователь уже существует",
		})
		return
	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sID,
		HttpOnly: true,
		Expires:  time.Now().Add(SessionTTL),
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    err.Error(),
			MsgRus: "Авторизация отсутствует",
		})
		return
	}
	err = h.usecase.Logout(session.Value)
	if errors.Is(err, models.ErrNoSession) {
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    err.Error(),
			MsgRus: "Авторизация отсутствует",
		})
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	})
}

func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    "no session",
			MsgRus: "авторизация отсутствует",
		})
		return
	}
	if !h.usecase.IsLoggedIn(session.Value) {
		helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    "no session",
			MsgRus: "авторизация отсутствует",
		})
		return
	}
	helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	})
}

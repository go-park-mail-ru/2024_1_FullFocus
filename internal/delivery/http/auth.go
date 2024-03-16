package delivery

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

const (
	SessionTTL = time.Hour * 24
)

type AuthHandler struct {
	router  *mux.Router
	usecase usecase.Auth
}

func NewAuthHandler(uc usecase.Auth) *AuthHandler {
	return &AuthHandler{
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

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := helper.GetLoggerFromContext(ctx)

	login := r.FormValue("login")
	password := r.FormValue("password")

	sID, err := h.usecase.Login(ctx, login, password)
	if err != nil {
		if validationError := new(models.ValidationError); errors.As(err, &validationError) {
			if err := helper.JSONResponse(w, 200, validationError.WithCode(400)); err != nil {
				l.Error(fmt.Sprintf("marshall error: %v", err))
			}
			return
		}
		err := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Неверный логин или пароль",
		})
		if err != nil {
			l.Error(fmt.Sprintf("marshall error: %v", err))
		}

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
	err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	})
	if err != nil {
		l.Error(fmt.Sprintf("marshall error: %v", err))
	}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := helper.GetLoggerFromContext(ctx)

	login := r.FormValue("login")
	password := r.FormValue("password")

	sID, _, err := h.usecase.Signup(ctx, login, password)
	if err != nil {
		if validationError := new(models.ValidationError); errors.As(err, &validationError) {
			if err := helper.JSONResponse(w, 200, validationError.WithCode(400)); err != nil {
				l.Error(fmt.Sprintf("marshall error: %v", err))
			}
			return
		}
		err := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Пользователь уже существует",
		})
		if err != nil {
			l.Error(fmt.Sprintf("marshall error: %v", err))
		}
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
	err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	})
	if err != nil {
		l.Error(fmt.Sprintf("marshall error: %v", err))
	}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := helper.GetLoggerFromContext(ctx)

	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		err := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    err.Error(),
			MsgRus: "Авторизация отсутствует",
		})
		if err != nil {
			l.Error(fmt.Sprintf("marshall error: %v", err))
		}
		return
	}
	err = h.usecase.Logout(ctx, session.Value)
	if errors.Is(err, models.ErrNoSession) {
		err := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    err.Error(),
			MsgRus: "Авторизация отсутствует",
		})
		if err != nil {
			l.Error(fmt.Sprintf("marshall error: %v", err))
		}
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	})
	if err != nil {
		l.Error(fmt.Sprintf("marshall error: %v", err))
	}
}

func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := helper.GetLoggerFromContext(ctx)

	session, err := r.Cookie("session_id")
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
	if !h.usecase.IsLoggedIn(ctx, session.Value) {
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
	err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	})
	if err != nil {
		l.Error(fmt.Sprintf("marshall error: %v", err))
	}
}

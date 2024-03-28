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

type AuthHandler struct {
	router     *mux.Router
	usecase    usecase.Auth
	sessionTTL time.Duration
}

func NewAuthHandler(uc usecase.Auth, sessTTL time.Duration) *AuthHandler {
	return &AuthHandler{
		router:     mux.NewRouter(),
		usecase:    uc,
		sessionTTL: sessTTL,
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
			if jsonErr := helper.JSONResponse(w, 200, validationError.WithCode(400)); jsonErr != nil {
				l.Error(fmt.Sprintf("marshall error: %v", jsonErr))
			}
			return
		}
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Неверный логин или пароль",
		}); jsonErr != nil {
			l.Error(fmt.Sprintf("marshall error: %v", jsonErr))
		}

		return
	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sID,
		HttpOnly: true,
		Expires:  time.Now().Add(h.sessionTTL),
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	if err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	}); err != nil {
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
			if jsonErr := helper.JSONResponse(w, 200, validationError.WithCode(400)); jsonErr != nil {
				l.Error(fmt.Sprintf("marshall error: %v", jsonErr))
			}
			return
		}
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Пользователь уже существует",
		}); jsonErr != nil {
			l.Error(fmt.Sprintf("marshall error: %v", jsonErr))
		}
		return
	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sID,
		HttpOnly: true,
		Expires:  time.Now().Add(h.sessionTTL),
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	if err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	}); err != nil {
		l.Error(fmt.Sprintf("marshall error: %v", err))
	}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := helper.GetLoggerFromContext(ctx)

	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    err.Error(),
			MsgRus: "Авторизация отсутствует",
		}); jsonErr != nil {
			l.Error(fmt.Sprintf("marshall error: %v", jsonErr))
		}
		return
	}
	err = h.usecase.Logout(ctx, session.Value)
	if errors.Is(err, models.ErrNoSession) {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    err.Error(),
			MsgRus: "Авторизация отсутствует",
		}); jsonErr != nil {
			l.Error(fmt.Sprintf("marshall error: %v", jsonErr))
		}
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	if err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	}); err != nil {
		l.Error(fmt.Sprintf("marshall error: %v", err))
	}
}

func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := helper.GetLoggerFromContext(ctx)

	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    "no session",
			MsgRus: "авторизация отсутствует",
		}); jsonErr != nil {
			l.Error(fmt.Sprintf("marshall error: %v", jsonErr))
		}
		return
	}
	if !h.usecase.IsLoggedIn(ctx, session.Value) {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    "no session",
			MsgRus: "авторизация отсутствует",
		}); err != nil {
			l.Error(fmt.Sprintf("marshall error: %v", jsonErr))
		}
		return
	}
	if err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	}); err != nil {
		l.Error(fmt.Sprintf("marshall error: %v", err))
	}
}

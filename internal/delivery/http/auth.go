package delivery

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
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
	var loginData dto.LoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		}); jsonErr != nil {
			logger.Error(ctx, "marshall error: "+jsonErr.Error())
		}
		return
	}
	sID, err := h.usecase.Login(ctx, loginData.Login, loginData.Password)
	if err != nil {
		if validationError := new(models.ValidationError); errors.As(err, &validationError) {
			if jsonErr := helper.JSONResponse(w, 200, validationError.WithCode(400)); jsonErr != nil {
				logger.Error(ctx, "marshall error: "+jsonErr.Error())
			}
			return
		}
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Неверный логин или пароль",
		}); jsonErr != nil {
			logger.Error(ctx, "marshall error: "+jsonErr.Error())
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
		logger.Error(ctx, "marshall error: "+err.Error())
	}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var loginData dto.LoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		}); jsonErr != nil {
			logger.Error(ctx, "marshall error: "+jsonErr.Error())
		}
		return
	}
	sID, _, err := h.usecase.Signup(ctx, loginData.Login, loginData.Password)
	if err != nil {
		if validationError := new(models.ValidationError); errors.As(err, &validationError) {
			if jsonErr := helper.JSONResponse(w, 200, validationError.WithCode(400)); jsonErr != nil {
				logger.Error(ctx, "marshall error: "+jsonErr.Error())
			}
			return
		}
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Пользователь уже существует",
		}); jsonErr != nil {
			logger.Error(ctx, "marshall error: "+jsonErr.Error())
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
		logger.Error(ctx, "marshall error: "+err.Error())
	}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    err.Error(),
			MsgRus: "Авторизация отсутствует",
		}); jsonErr != nil {
			logger.Error(ctx, "marshall error: "+jsonErr.Error())
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
			logger.Error(ctx, "marshall error: "+jsonErr.Error())
		}
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	if err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	}); err != nil {
		logger.Error(ctx, "marshall error: "+err.Error())
	}
}

func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    "no session",
			MsgRus: "авторизация отсутствует",
		}); jsonErr != nil {
			logger.Error(ctx, "marshall error: "+jsonErr.Error())
		}
		return
	}
	if !h.usecase.IsLoggedIn(ctx, session.Value) {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    "no session",
			MsgRus: "авторизация отсутствует",
		}); jsonErr != nil {
			logger.Error(ctx, "marshall error: "+jsonErr.Error())
		}
		return
	}
	if err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	}); err != nil {
		logger.Error(ctx, "marshall error: "+err.Error())
	}
}

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
		h.router.Handle("/public/v1/login", http.HandlerFunc(h.Login)).Methods("POST", "OPTIONS")
		h.router.Handle("/public/v1/signup", http.HandlerFunc(h.Signup)).Methods("POST", "OPTIONS")
		h.router.Handle("/v1/logout", http.HandlerFunc(h.Logout)).Methods("POST", "OPTIONS")
		h.router.Handle("/public/v1/check", http.HandlerFunc(h.CheckAuth)).Methods("GET", "OPTIONS")
		h.router.Handle("/v1/password", http.HandlerFunc(h.UpdatePassword)).Methods("POST", "OPTIONS")
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginData, err := helper.GetLoginData(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	sID, err := h.usecase.Login(ctx, loginData.Login, loginData.Password)
	if err != nil {
		if validationError := new(helper.ValidationError); errors.As(err, &validationError) {
			helper.JSONResponse(ctx, w, 200, validationError.WithCode(400))
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
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
		Expires:  time.Now().Add(h.sessionTTL),
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginData, err := helper.GetLoginData(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	sID, err := h.usecase.Signup(ctx, loginData.Login, loginData.Password)
	if err != nil {
		if validationError := new(helper.ValidationError); errors.As(err, &validationError) {
			helper.JSONResponse(ctx, w, 200, validationError.WithCode(400))
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
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
		Expires:  time.Now().Add(h.sessionTTL),
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 401,
			Msg:    err.Error(),
			MsgRus: "Авторизация отсутствует",
		})
		return
	}
	if err = h.usecase.Logout(ctx, session.Value); errors.Is(err, models.ErrNoSession) {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 401,
			Msg:    err.Error(),
			MsgRus: "Авторизация отсутствует",
		})
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 401,
			Msg:    "no session",
			MsgRus: "авторизация отсутствует",
		})
		return
	}
	if !h.usecase.IsLoggedIn(ctx, session.Value) {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 401,
			Msg:    "no session",
			MsgRus: "авторизация отсутствует",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

func (h *AuthHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uID, err := helper.GetUserIDFromContext(ctx)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 403,
			Msg:    err.Error(),
			MsgRus: "Пользователь не авторизован",
		})
		return
	}
	var updateInput dto.UpdatePasswordInput
	if err = json.NewDecoder(r.Body).Decode(&updateInput); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	if err = h.usecase.UpdatePassword(ctx, uID, updateInput.Password, updateInput.NewPassword); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Неверный пароль",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

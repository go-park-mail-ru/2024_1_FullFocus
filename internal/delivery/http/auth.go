package delivery

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

type JwtToken struct {
	Secret []byte
}

func NewJwtToken(secret string) (*JwtToken, error) {
	return &JwtToken{Secret: []byte(secret)}, nil
}

type JwtCsrfClaims struct {
	SessionID string `json:"sid"`
	jwt.StandardClaims
}

func (tk *JwtToken) Create(sID string, tokenExpTime int64) (string, error) {
	data := JwtCsrfClaims{
		SessionID: sID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpTime,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(tk.Secret)
}

func (tk *JwtToken) parseSecretGetter(token *jwt.Token) (interface{}, error) {
	method, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok || method.Alg() != "HS256" {
		return nil, fmt.Errorf("bad sign method")
	}
	return tk.Secret, nil
}

func (tk *JwtToken) Check(sID string, inputToken string) (bool, error) {
	payload := &JwtCsrfClaims{}
	_, err := jwt.ParseWithClaims(inputToken, payload, tk.parseSecretGetter)
	if err != nil {
		return false, fmt.Errorf("cant parse jwt token: %v", err)
	}
	if payload.Valid() != nil {
		return false, fmt.Errorf("invalid jwt token: %v", err)
	}
	return payload.SessionID == sID, nil
}

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

func CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "DELETE" || r.Method == "PATCH" {
			tokens, _ := NewJwtToken("qsRY2e4hcM5T7X984E9WQ5uZ8Nty7fxB")
			session, err := r.Cookie("session_id")
			sID := session.Value
			CSRFToken := r.FormValue("X-Csrf-Token")
			_, err = tokens.Check(sID, CSRFToken)
			if err != nil {
				w.Write([]byte("{}"))
				return
			}
		} else if r.Method == "GET" {
			tokens, _ := NewJwtToken("qsRY2e4hcM5T7X984E9WQ5uZ8Nty7fxB")
			session, err := r.Cookie("session_id")
			sID := session.Value
			token, err := tokens.Create(sID, time.Now().Add(1*time.Hour).Unix())
			if err != nil {
				//logger.Error(ctx, fmt.Sprintf("csrf token creation error: %v", err))
				return
			}
			w.Header().Set("X-Csrf-Token", token)
		}
		next.ServeHTTP(w, r)
	})
}

func (h *AuthHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/auth").Subrouter()
	{
		h.router.Use(CSRFMiddleware)
		h.router.Handle("/login", http.HandlerFunc(h.Login)).Methods("GET", "POST", "OPTIONS")
		h.router.Handle("/signup", http.HandlerFunc(h.Signup)).Methods("GET", "POST", "OPTIONS")
		h.router.Handle("/logout", http.HandlerFunc(h.Logout)).Methods("POST", "OPTIONS")
		h.router.Handle("/check", http.HandlerFunc(h.CheckAuth)).Methods("GET", "OPTIONS")
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	login := r.FormValue("login")
	password := r.FormValue("password")
	sID, err := h.usecase.Login(ctx, login, password)
	//uID, err := h.usecase.GetUserIDBySessionID(ctx, sID)

	if err != nil {
		if validationError := new(models.ValidationError); errors.As(err, &validationError) {
			if jsonErr := helper.JSONResponse(w, 200, validationError.WithCode(400)); jsonErr != nil {
				logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
			}
			return
		}
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Неверный логин или пароль",
		}); jsonErr != nil {
			logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
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
	/*
		tokens, _ := NewHMACHashToken("golangcourse")
		token, err := tokens.Create(sID, uID, time.Now().Add(24*time.Hour).Unix())
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("csrf token creation error: %v", err))
			return
		}
		w.Header().Set("X-Csrf-Token", token)

	*/

	if err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	}); err != nil {
		logger.Error(ctx, fmt.Sprintf("marshall error: %v", err))
	}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	login := r.FormValue("login")
	password := r.FormValue("password")

	sID, _, err := h.usecase.Signup(ctx, login, password)
	//uID, err := h.usecase.GetUserIDBySessionID(ctx, sID)

	if err != nil {
		if validationError := new(models.ValidationError); errors.As(err, &validationError) {
			if jsonErr := helper.JSONResponse(w, 200, validationError.WithCode(400)); jsonErr != nil {
				logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
			}
			return
		}
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Пользователь уже существует",
		}); jsonErr != nil {
			logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
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
	/*tokens, _ := NewHMACHashToken("golangcourse")
	token, err := tokens.Create(sID, uID, time.Now().Add(24*time.Hour).Unix())
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("csrf token creation error: %v", err))
		return
	}
	w.Header().Set("X-Csrf-Token", token)

	*/

	if err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	}); err != nil {
		logger.Error(ctx, fmt.Sprintf("marshall error: %v", err))
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
			logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
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
			logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
		}
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	if err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	}); err != nil {
		logger.Error(ctx, fmt.Sprintf("marshall error: %v", err))
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
			logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
		}
		return
	}
	if !h.usecase.IsLoggedIn(ctx, session.Value) {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 401,
			Msg:    "no session",
			MsgRus: "авторизация отсутствует",
		}); err != nil {
			logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
		}
		return
	}
	if err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	}); err != nil {
		logger.Error(ctx, fmt.Sprintf("marshall error: %v", err))
	}
}

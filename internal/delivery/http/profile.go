package delivery

import (
	"errors"
	"fmt"
	model "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
	"net/http"
)

type ProfileHandler struct {
	router  *mux.Router
	usecase usecase.Profiles
}

func NewProfileHandler(u usecase.Profiles) *ProfileHandler {
	return &ProfileHandler{
		router:  mux.NewRouter(),
		usecase: u,
	}
}

func (h *ProfileHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/profile").Subrouter()
	{
		h.router.Handle("/update", http.HandlerFunc(h.UpdateProfile)).Methods("POST", "GET")
		h.router.Handle("/get", http.HandlerFunc(h.GetProfile)).Methods("GET")
	}
}

func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	newUsername := r.FormValue("newUsername")
	newPassword := r.FormValue("newPassword")

	// Надо вытащить username и и аватарку
	username := r.FormValue("username")
	// Через UploadAvatar?
	newProfile := model.Profile{
		User: model.User{
			Username: newUsername,
			Password: newPassword,
		},
		// Image: ?
	}

	err := h.usecase.UpdateProfile(ctx, username, newProfile)
	if err != nil {
		if validationError := new(model.ValidationError); errors.As(err, &validationError) {
			if jsonErr := helper.JSONResponse(w, 200, validationError.WithCode(400)); jsonErr != nil {
				logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
			}
			return
		} else {
			if jsonErr := helper.JSONResponse(w, 200, model.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Пользователя не существует",
			}); jsonErr != nil {
				logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
			}
			return
		}
	}

	if err = helper.JSONResponse(w, 200, model.SuccessResponse{
		Status: 200,
	}); err != nil {
		logger.Error(ctx, fmt.Sprintf("marshall error: %v", err))
	}
}
func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// надо вытащить username?
	username := r.FormValue("username")

	profile, err := h.usecase.GetProfile(ctx, username)
	if err != nil {
		if jsonErr := helper.JSONResponse(w, 200, model.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Пользователя не существует",
		}); jsonErr != nil {
			logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
		}
		return
	}
	if jsonErr := helper.JSONResponse(w, 200, profile); jsonErr != nil {
		logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
	}
}

package delivery

import (
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
	"github.com/gorilla/mux"
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
	h.router = r.PathPrefix("/v1/profile").Subrouter()
	{
		h.router.Handle("/get", http.HandlerFunc(h.GetProfile)).Methods("GET", "OPTIONS")
		h.router.Handle("/update", http.HandlerFunc(h.UpdateProfile)).Methods("GET", "POST", "OPTIONS")
	}
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
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
	profile, err := h.usecase.GetProfile(ctx, uID)
	if err != nil {
		if errors.Is(err, models.ErrNoProfile) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Пользователя не существует",
			})
			return
		}
		logger.Error(ctx, err.Error())
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    "Internal error",
			MsgRus: "Неизвестная ошибка",
		})
		return
	}
	data := dto.ConvertProfileDataToProfile(profile)
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   data,
	})
}

func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
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
	profileData, err := helper.GetProfileData(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	updateProfileInput := dto.ConvertProfileToProfileData(profileData)
	if err = h.usecase.UpdateProfile(ctx, uID, updateProfileInput); err != nil {
		if validationError := new(helper.ValidationError); errors.As(err, &validationError) {
			helper.JSONResponse(ctx, w, 200, validationError.WithCode(400))
			return
		}
		if errors.Is(err, models.ErrNoProfile) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Пользователя не существует",
			})
			return
		}
		logger.Error(ctx, err.Error())
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    "Internal error",
			MsgRus: "Неизвестная ошибка",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

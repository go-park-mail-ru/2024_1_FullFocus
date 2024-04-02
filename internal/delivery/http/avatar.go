package delivery

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

type AvatarHandler struct {
	router  *mux.Router
	usecase usecase.Avatars
}

func NewAvatarHandler(u usecase.Avatars) *AvatarHandler {
	return &AvatarHandler{
		router:  mux.NewRouter(),
		usecase: u,
	}
}

func (h *AvatarHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/avatar").Subrouter()
	{
		h.router.Handle("/", http.HandlerFunc(h.UploadAvatar)).Methods("POST", "OPTIONS")
		h.router.Handle("/", http.HandlerFunc(h.DeleteAvatar)).Methods("POST", "OPTIONS")
	}
}

func (h *AvatarHandler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	src, hdr, err := r.FormFile("avatar")
	if err != nil {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Файл не загружен",
		}); jsonErr != nil {
			logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
		}
		return
	}
	img := dto.Image{
		Payload:     src,
		PayloadSize: hdr.Size,
	}
	if err = h.usecase.UploadAvatar(ctx, img); err != nil {
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка загрузки фото",
		}); jsonErr != nil {
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

func (h *AvatarHandler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := h.usecase.DeleteAvatar(ctx); err != nil {
		if errors.Is(err, models.ErrNoAvatar) {
			if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Аватар не найден",
			}); jsonErr != nil {
				logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
			}
			return
		}
		if jsonErr := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка удаления фото",
		}); jsonErr != nil {
			logger.Error(ctx, fmt.Sprintf("marshall error: %v", jsonErr))
		}
	}
	if err := helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	}); err != nil {
		logger.Error(ctx, fmt.Sprintf("marshall error: %v", err))
	}
}

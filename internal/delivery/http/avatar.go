package delivery

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/pkg/errors"
	"net/http"

	"github.com/gorilla/mux"

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
	l := helper.GetLoggerFromContext(ctx)

	src, hdr, err := r.FormFile("avatar")
	if err != nil {
		err := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Файл не загружен",
		})
		if err != nil {
			l.Error(fmt.Sprintf("marshall error: %v", err))
		}
		return
	}
	img := dto.Image{
		Payload:     src,
		PayloadSize: hdr.Size,
	}
	if err = h.usecase.UploadAvatar(ctx, img); err != nil {
		err := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка загрузки фото",
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

func (h *AvatarHandler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := helper.GetLoggerFromContext(ctx)

	if err := h.usecase.DeleteAvatar(ctx); err != nil {
		if errors.Is(err, models.ErrNoAvatar) {
			err := helper.JSONResponse(w, 200, models.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Аватар не найден",
			})
			if err != nil {
				l.Error(fmt.Sprintf("marshall error: %v", err))
			}
			return
		}
		err := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка удаления фото",
		})
		if err != nil {
			l.Error(fmt.Sprintf("marshall error: %v", err))
		}
	}
	err := helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
	})
	if err != nil {
		l.Error(fmt.Sprintf("marshall error: %v", err))
	}
}

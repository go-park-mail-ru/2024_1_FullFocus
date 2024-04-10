package delivery

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
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
		h.router.Handle("/upload", http.HandlerFunc(h.Upload)).Methods("POST", "OPTIONS")
		h.router.Handle("/delete", http.HandlerFunc(h.Delete)).Methods("POST", "OPTIONS")
		// h.router.Handle("/get/{filename}", http.HandlerFunc(h.Get)).Methods("GET", "OPTIONS")
	}
}

func (h *AvatarHandler) Upload(w http.ResponseWriter, r *http.Request) {
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
	src, hdr, err := r.FormFile("avatar")
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Файл не загружен",
		})
		return
	}
	img := dto.Image{
		Payload:     src,
		PayloadSize: hdr.Size,
	}
	if err = h.usecase.UploadAvatar(ctx, img, uID); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка загрузки фото",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

func (h *AvatarHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	if err = h.usecase.DeleteAvatar(ctx, uID); err != nil {
		if errors.Is(err, models.ErrNoAvatar) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Аватар не найден",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка удаления фото",
		})
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

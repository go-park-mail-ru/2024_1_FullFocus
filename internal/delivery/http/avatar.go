package delivery

import (
	"fmt"
	"net/http"
	"time"

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
	h.router = r.PathPrefix("/v1/avatar").Subrouter()
	{
		h.router.Handle("/public/{name}", http.HandlerFunc(h.GetAvatarByName)).Methods("GET", "OPTIONS")
		h.router.Handle("/upload", http.HandlerFunc(h.UploadAvatar)).Methods("POST", "OPTIONS")
		h.router.Handle("/delete", http.HandlerFunc(h.DeleteAvatar)).Methods("POST", "OPTIONS")
	}
}

func (h *AvatarHandler) GetAvatarByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fileName, ok := mux.Vars(r)["name"]
	if !ok {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid name value",
			MsgRus: "Невалидное имя файла",
		})
		return
	}
	avatar, err := h.usecase.GetAvatar(ctx, fileName)
	if err != nil {
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
			Msg:    "internal error",
			MsgRus: "Неизвестная ошибка",
		})
		return
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fileName))
	http.ServeContent(w, r, fileName, time.Now(), avatar.Payload)
}

func (h *AvatarHandler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
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
			MsgRus: "Ошибка загрузки",
		})
		return
	}
	img := models.Avatar{
		Payload:     src,
		PayloadSize: hdr.Size,
	}
	if err = h.usecase.UploadAvatar(ctx, uID, img); err != nil {
		if validationError := new(helper.ValidationError); errors.As(err, &validationError) {
			helper.JSONResponse(ctx, w, 200, validationError.WithCode(400))
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка загрузки фото",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

func (h *AvatarHandler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
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
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Аватар не найден",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

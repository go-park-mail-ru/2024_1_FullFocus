package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

type NotificationHandler struct {
	router  *mux.Router
	usecase usecase.Notifications
}

func NewNotificationHandler(uc usecase.Notifications) *NotificationHandler {
	return &NotificationHandler{
		router:  mux.NewRouter(),
		usecase: uc,
	}
}

func (h *NotificationHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/v1/notifications").Subrouter()
	{
		h.router.Handle("", http.HandlerFunc(h.GetAllNotifications)).Methods("GET", "OPTIONS")
		h.router.Handle("/read", http.HandlerFunc(h.MarkNotificationAsRead)).Methods("POST", "OPTIONS")
		h.router.Handle("/connect", http.HandlerFunc(h.ConnectToNotifications)).Methods("POST", "OPTIONS")
	}
}

func (h *NotificationHandler) GetAllNotifications(w http.ResponseWriter, r *http.Request) {
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
	notifications, err := h.usecase.GetAllNotifications(ctx, uID)
	if err != nil {
		if errors.Is(err, models.ErrNoNotifications) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Нет уведомлений",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка получения уведомлений",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertNotifications(notifications),
	})
}

func (h *NotificationHandler) MarkNotificationAsRead(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if _, err := helper.GetUserIDFromContext(ctx); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 403,
			Msg:    err.Error(),
			MsgRus: "Пользователь не авторизован",
		})
		return
	}
	var input dto.ReadNotificationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	if err := h.usecase.MarkNotificationRead(ctx, input.NotificationID); err != nil {
		if errors.Is(err, models.ErrNoNotifications) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Уведомление не найдено",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка сервера",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

func (h *NotificationHandler) ConnectToNotifications(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uID, err := helper.GetUserIDFromContext(ctx)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("{\n  \"disconnect\": {\n    \"code\": 4501,\n    \"reason\": \"unauthorized\"\n  }\n}"))
		return
	}
	uIDstr := strconv.FormatUint(uint64(uID), 10)
	json.NewEncoder(w).Encode(dto.CentrifugoConnectResponse{
		Result: dto.CentrifugoConnectResponseData{
			User:     uIDstr,
			ExpireAt: int(time.Now().Add(time.Hour * 24).Unix()),
			Channels: []string{uIDstr},
		},
	})
}

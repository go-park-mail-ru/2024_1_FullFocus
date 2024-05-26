package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

type OrderHandler struct {
	router  *mux.Router
	usecase usecase.Orders
}

func NewOrderHandler(uc usecase.Orders) *OrderHandler {
	return &OrderHandler{
		router:  mux.NewRouter(),
		usecase: uc,
	}
}

func (h *OrderHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/v1/order").Subrouter()
	{
		h.router.Handle("/create", http.HandlerFunc(h.Create)).Methods("POST", "OPTIONS")
		h.router.Handle("/{id:[0-9]+}", http.HandlerFunc(h.GetOrder)).Methods("GET", "OPTIONS")
		h.router.Handle("/all", http.HandlerFunc(h.GetAllOrders)).Methods("GET", "OPTIONS")
		h.router.Handle("/public/update", http.HandlerFunc(h.UpdateStatus)).Methods("POST", "OPTIONS")
		h.router.Handle("/cancel", http.HandlerFunc(h.Delete)).Methods("POST", "OPTIONS")
	}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	var createOrderInput dto.CreateOrderInput
	if err = json.NewDecoder(r.Body).Decode(&createOrderInput); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	createInput := dto.ConvertCreateOrderInputToModel(uID, createOrderInput)
	orderInfo, err := h.usecase.Create(ctx, createInput)
	if err != nil {
		if errors.Is(err, models.ErrInvalidPromocode) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Недостаточная сумма заказа для активации промокода",
			})
			return
		} else if errors.Is(err, models.ErrNoPromocode) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Промокод не найден",
			})
			return
		} else if errors.Is(err, models.ErrPromocodeExpired) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Промокод просрочен",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка создания заказа",
		})
		return
	}
	data := dto.ConvertCreatePayload(orderInfo)
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   data,
	})
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
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
	orderID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid orderID",
			MsgRus: "Невалидный параметр",
		})
		return
	}
	orderInfo, err := h.usecase.GetOrderByID(ctx, uID, uint(orderID))
	if err != nil {
		if errors.Is(err, models.ErrNoAccess) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 403,
				Msg:    err.Error(),
				MsgRus: "Ошибка доступа",
			})
			return
		}
		if errors.Is(err, models.ErrNoRowsFound) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    "not found",
				MsgRus: "Заказ не найден",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка получения информации о заказе",
		})
		return
	}
	data := dto.GetOrderPayload{
		Products:   dto.ConvertOrderProductsToDTO(orderInfo.Products),
		Sum:        orderInfo.Sum,
		Status:     orderInfo.Status,
		ItemsCount: orderInfo.ItemsCount,
		CreatedAt:  orderInfo.CreatedAt,
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   data,
	})
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
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
	orders, err := h.usecase.GetAllOrders(ctx, uID)
	if err != nil {
		if errors.Is(err, models.ErrNoRowsFound) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    "not found",
				MsgRus: "Заказы не найдены",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка получения истории заказов",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertOrdersToDTO(orders),
	})
}

func (h *OrderHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input dto.UpdateOrderStatusInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	if err := h.usecase.UpdateStatus(ctx, dto.ConvertUpdateOrderStatusInput(input)); err != nil {
		if errors.Is(err, models.ErrInvalidField) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Неверный статус заказа",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка обновления статуса",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

func (h *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	var cancelOrderInput dto.CancelOrderInput
	if err = json.NewDecoder(r.Body).Decode(&cancelOrderInput); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	if err = h.usecase.Delete(ctx, uID, cancelOrderInput.OrderID); err != nil {
		if errors.Is(err, models.ErrNoAccess) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 403,
				Msg:    err.Error(),
				MsgRus: "Ошибка доступа",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка отмены заказа",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

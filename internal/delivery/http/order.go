package delivery

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"

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
	h.router = r.PathPrefix("/order").Subrouter()
	{
		h.router.Handle("/{id:[0-9]+}", http.HandlerFunc(h.Create)).Methods("POST", "OPTIONS")
		h.router.Handle("/{id:[0-9]+}", http.HandlerFunc(h.Create)).Methods("POST", "OPTIONS")
		h.router.Handle("/{id:[0-9]+}", http.HandlerFunc(h.Delete)).Methods("POST", "OPTIONS")
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
	createInput := models.CreateOrderInput{
		UserID:   uID,
		FromCart: createOrderInput.FromCart,
	}
	for _, item := range createOrderInput.Items {
		createInput.Items = append(createInput.Items, models.OrderItem{
			ProductID: item.ProductID,
			Count:     item.Count,
		})
	}
	orderID, err := h.usecase.Create(ctx, createInput)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка создания заказа",
		})
	}
	helper.JSONResponse(ctx, w, 200, dto.CreateOrderPayload{
		OrderID: orderID,
	})
}

func (h *OrderHandler) GetOrderItems(w http.ResponseWriter, r *http.Request) {
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
	var getOrderProductsInput dto.GetOrderProductsInput
	if err = json.NewDecoder(r.Body).Decode(&getOrderProductsInput); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	products, err := h.usecase.GetOrderProducts(ctx, uID, getOrderProductsInput.OrderID)
	if err != nil {
		if errors.Is(err, models.ErrNoAccess) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 403,
				Msg:    err.Error(),
				MsgRus: "Ошибка доступа",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка получения информации о заказе",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertProductsToDTO(products),
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
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка отмены заказа",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

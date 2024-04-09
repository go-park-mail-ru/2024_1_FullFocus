package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type CartHandler struct {
	router  *mux.Router
	usecase usecase.Carts
}

func NewCartHandler(uc usecase.Carts) *CartHandler {
	return &CartHandler{
		router:  mux.NewRouter(),
		usecase: uc,
	}
}

func (h *CartHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/cart").Subrouter()
	{
		h.router.Handle("/v1", http.HandlerFunc(h.GetAllCartItems)).Methods("GET", "OPTIONS")
		h.router.Handle("/v1/add", http.HandlerFunc(h.UpdateCartItem)).Methods("POST", "OPTIONS")
		h.router.Handle("/v1/delete", http.HandlerFunc(h.DeleteCartItem)).Methods("POST", "OPTIONS")
		h.router.Handle("/v1/clear", http.HandlerFunc(h.DeleteAllCartItems)).Methods("POST", "OPTIONS")
	}
}

func (h *CartHandler) GetAllCartItems(w http.ResponseWriter, r *http.Request) {
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

	cartContent, err := h.usecase.GetAllCartItems(ctx, uID)
	if errors.Is(err, models.ErrEmptyCart) {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 404,
			Msg:    err.Error(),
			MsgRus: "Товары в корзине отсутствуют",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertContentToDto(cartContent),
	})
}

func (h *CartHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
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
	cartItem, err := helper.GetCartItemData(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}

	newCount, err := h.usecase.UpdateCartItem(ctx, uID, cartItem.ProductID)
	if errors.Is(err, models.ErrNoProduct) {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 404,
			Msg:    err.Error(),
			MsgRus: "Такой товар в корзине не найден",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data: dto.UpdateCartItemPayload{
			Count: newCount,
		},
	})
}

func (h *CartHandler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
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

	cartItem, err := helper.GetCartItemData(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}

	newCount, err := h.usecase.DeleteCartItem(ctx, uID, cartItem.ProductID)
	if errors.Is(err, models.ErrNoProduct) {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 404,
			Msg:    err.Error(),
			MsgRus: "Такой товар в корзине не найден",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data: dto.UpdateCartItemPayload{
			Count: newCount,
		},
	})
}

func (h *CartHandler) DeleteAllCartItems(w http.ResponseWriter, r *http.Request) {
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

	if err = h.usecase.DeleteAllCartItems(ctx, uID); errors.Is(err, models.ErrEmptyCart) {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 404,
			Msg:    err.Error(),
			MsgRus: "Товары в корзине отсутствуют",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

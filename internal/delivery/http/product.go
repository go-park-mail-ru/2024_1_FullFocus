package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

type ProductHandler struct {
	router  *mux.Router
	usecase usecase.Products
}

func NewProductHandler(u usecase.Products) *ProductHandler {
	return &ProductHandler{
		router:  mux.NewRouter(),
		usecase: u,
	}
}

func (h *ProductHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/public/products").Subrouter()
	{
		h.router.Handle("", http.HandlerFunc(h.GetProducts)).Methods("GET", "OPTIONS")
		h.router.Handle("/{id:[0-9]+}", http.HandlerFunc(h.GetProductByID)).Methods("GET", "OPTIONS")
	}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pageNum, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, models.ErrResponse{
			Status: 400,
			Msg:    "invalid page value",
			MsgRus: "Невалидный параметр",
		})
		return
	}
	pageSize, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, models.ErrResponse{
			Status: 400,
			Msg:    "invalid limit value",
			MsgRus: "Невалидный параметр",
		})
		return
	}
	products, err := h.usecase.GetAllProductCards(ctx, uint(pageNum), uint(pageSize))
	if errors.Is(err, models.ErrNoProduct) {
		helper.JSONResponse(ctx, w, 200, models.ErrResponse{
			Status: 404,
			Msg:    "not found",
			MsgRus: "по данному запросу товары не найдены",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, models.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertProductCardsToDTO(products),
	})
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	productID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, models.ErrResponse{
			Status: 400,
			Msg:    "invalid id value",
			MsgRus: "Невалидный параметр",
		})
		return
	}
	product, err := h.usecase.GetProductById(ctx, uint(productID))
	if err != nil {
		if errors.Is(err, models.ErrNoProduct) {
			helper.JSONResponse(ctx, w, 200, models.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Товар не найден",
			})
		}
		helper.JSONResponse(ctx, w, 200, models.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка поиска товара",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, models.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertProductToDTO(product),
	})
}

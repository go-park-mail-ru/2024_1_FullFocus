package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/gorilla/mux"

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
	h.router = r.PathPrefix("/v1/public/products").Subrouter()
	{
		h.router.Handle("", http.HandlerFunc(h.GetProducts)).Methods("GET", "OPTIONS")
		h.router.Handle("/search", http.HandlerFunc(h.GetProductsByQuery)).Methods("GET", "OPTIONS")
		h.router.Handle("/{id:[1-9]+[0-9]*}", http.HandlerFunc(h.GetProductByID)).Methods("GET", "OPTIONS")
		h.router.Handle("/category/{id:[1-9]+[0-9]*}", http.HandlerFunc(h.GetProductsByCategoryID)).Methods("GET", "OPTIONS")
		h.router.Handle("/sorting", http.HandlerFunc(h.GetProductSortingTypes)).Methods("GET", "OPTIONS")
	}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uID, _ := helper.GetUserIDFromContext(ctx)
	pageNum, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid page value",
			MsgRus: "Невалидные параметры",
		})
		return
	}
	pageSize, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid limit value",
			MsgRus: "Невалидные параметры",
		})
		return
	}
	sortingData, err := helper.GetSortParams(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid sortID value",
			MsgRus: "Невалидный параметр сортировки",
		})
		return
	}
	getProductsInput := models.GetAllProductsInput{
		ProfileID: uID,
		PageNum:   uint(pageNum),
		PageSize:  uint(pageSize),
		Sorting:   sortingData,
	}
	products, err := h.usecase.GetAllProductCards(ctx, getProductsInput)
	if err != nil {
		if errors.Is(err, models.ErrNoRowsFound) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 404,
				Msg:    "not found",
				MsgRus: "Товары не найдены",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка поиска товаров",
		})
		return
	}
	data := dto.GetAllProductsPayload{
		ProductCards: dto.ConvertProductCardsToDTO(products),
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   data,
	})
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uID, _ := helper.GetUserIDFromContext(ctx)
	productID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid id value",
			MsgRus: "Невалидный параметр",
		})
		return
	}
	product, err := h.usecase.GetProductByID(ctx, uID, uint(productID))
	if err != nil {
		if errors.Is(err, models.ErrNoRowsFound) || errors.Is(err, models.ErrNoProduct) || errors.Is(err, models.ErrProductNotFound) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Товар не найден",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка поиска товара",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertPromoProductToDTO(product),
	})
}

func (h *ProductHandler) GetProductsByCategoryID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uID, _ := helper.GetUserIDFromContext(ctx)
	categoryID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid id value",
			MsgRus: "Невалидные параметры",
		})
		return
	}
	pageNum, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid page value",
			MsgRus: "Невалидные параметры",
		})
		return
	}
	pageSize, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid limit value",
			MsgRus: "Невалидные параметры",
		})
		return
	}
	sortingData, err := helper.GetSortParams(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid sortID value",
			MsgRus: "Невалидный параметр сортировки",
		})
		return
	}
	getProductsInput := models.GetProductsByCategoryIDInput{
		CategoryID: uint(categoryID),
		ProfileID:  uID,
		PageNum:    uint(pageNum),
		PageSize:   uint(pageSize),
		Sorting:    sortingData,
	}
	products, err := h.usecase.GetProductsByCategoryID(ctx, getProductsInput)
	if err != nil {
		if errors.Is(err, models.ErrNoProduct) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Товары не найдены",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка поиска товаров",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertProductsByCategoryIDPayload(products),
	})
}

func (h *ProductHandler) GetProductsByQuery(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uID, _ := helper.GetUserIDFromContext(ctx)
	query := r.URL.Query().Get("query")
	pageNum, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid page value",
			MsgRus: "Невалидные параметры",
		})
		return
	}
	pageSize, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid limit value",
			MsgRus: "Невалидные параметры",
		})
		return
	}
	sortingData, err := helper.GetSortParams(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid sortID value",
			MsgRus: "Невалидный параметр сортировки",
		})
		return
	}
	getProductsInput := models.GetProductsByQueryInput{
		Query:     query,
		ProfileID: uID,
		PageNum:   uint(pageNum),
		PageSize:  uint(pageSize),
		Sorting:   sortingData,
	}
	products, err := h.usecase.GetProductsByQuery(ctx, getProductsInput)
	if err != nil {
		if errors.Is(err, models.ErrNoProduct) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Товары не найдены",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
			MsgRus: "Ошибка поиска товаров",
		})
		return
	}
	data := dto.GetProductsByQueryPayload{
		Query:    query,
		Products: dto.ConvertProductCardsToDTO(products),
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   data,
	})
}

func (h *ProductHandler) GetProductSortingTypes(w http.ResponseWriter, r *http.Request) {
	helper.JSONResponse(r.Context(), w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertSortTypesToDTO(helper.GetProductSortTypes()),
	})
}

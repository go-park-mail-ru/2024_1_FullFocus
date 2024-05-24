package delivery

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type PromotionHandler struct {
	router           *mux.Router
	promotionUsecase usecase.Promotion
}

func NewPromotionHandler(u usecase.Promotion) *PromotionHandler {
	return &PromotionHandler{
		router:           mux.NewRouter(),
		promotionUsecase: u,
	}
}

func (h *PromotionHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/promo").Subrouter()
	{
		h.router.Handle("/admin/v1/add", http.HandlerFunc(h.AddPromoProduct)).Methods("POST", "OPTIONS")
		h.router.Handle("/public/v1", http.HandlerFunc(h.GetPromoProducts)).Methods("GET", "OPTIONS")
		h.router.Handle("/public/v1/{id:[1-9]+[0-9]*}", http.HandlerFunc(h.GetPromoProduct)).Methods("GET", "OPTIONS")
		h.router.Handle("/admin/v1/delete", http.HandlerFunc(h.DeletePromoProduct)).Methods("POST", "OPTIONS")
	}
}

func (h *PromotionHandler) GetPromoProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uID, _ := helper.GetUserIDFromContext(ctx)
	amountStr := r.URL.Query().Get("amount")
	var amount uint
	if amountStr == "" {
		amount = 0
	} else {
		amnt, err := strconv.ParseUint(amountStr, 10, 32)
		if err != nil {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    "invalid amount value",
				MsgRus: "Невалидный параметр количества",
			})
			return
		}
		amount = uint(amnt)
	}
	products, err := h.promotionUsecase.GetPromoProductCards(ctx, amount, uID)
	if err != nil {
		if errors.Is(err, models.ErrNoProduct) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 404,
				Msg:    "no promo products",
				MsgRus: "Акционные товары отсутсвуют",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    "Internal server error",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Data: dto.ConvertPromoProductCardsToDTOs(products),
	})
}

func (h *PromotionHandler) GetPromoProduct(w http.ResponseWriter, r *http.Request) {
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
	promoProduct, err := h.promotionUsecase.GetPromoProductInfoByID(ctx, uint(productID), uID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoProduct):
			fallthrough
		case errors.Is(err, models.ErrProductNotFound):
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 404,
				Msg:    "not found",
				MsgRus: "Товар не найден",
			})
			return
		default:
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 500,
				Msg:    "Internal server error",
			})
			return
		}
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertPromoProductToDTO(promoProduct),
	})
}

func (h *PromotionHandler) AddPromoProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	promoData, err := helper.GetPromoDataInput(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	if err := h.promotionUsecase.CreatePromoProduct(ctx, dto.ConvertPromoDataToModel(promoData)); err != nil {
		if validationError := new(helper.ValidationError); errors.As(err, &validationError) {
			helper.JSONResponse(ctx, w, 200, validationError.WithCode(400))
			return
		}
		switch {
		case errors.Is(err, models.ErrInvalidBenefitValue):
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Некорректное значение скидки",
			})
			return
		case errors.Is(err, models.ErrNoProduct):
			fallthrough
		case errors.Is(err, models.ErrProductNotFound):
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 404,
				Msg:    "not found",
				MsgRus: "Товар не найден",
			})
			return
		case errors.Is(err, models.ErrPromoProductAlreadyExists):
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Этот товар уже участвует в распродаже",
			})
			return
		default:
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 500,
				Msg:    "Internal server error",
			})
			return
		}
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

func (h *PromotionHandler) DeletePromoProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	productID, err := helper.GetDeletePromoProductInput(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	if err := h.promotionUsecase.DeletePromoProduct(ctx, productID); err != nil {
		switch {
		case errors.Is(err, models.ErrNoProduct):
			fallthrough
		case errors.Is(err, models.ErrProductNotFound):
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 404,
				Msg:    "not found",
				MsgRus: "Товар не найден",
			})
			return
		default:
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 500,
				Msg:    "Internal server error",
			})
			return
		}
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

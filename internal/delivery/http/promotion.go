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
		h.router.Handle("/public/v1", http.HandlerFunc(h.GetPromoProducts)).Methods("POST", "OPTIONS")
	}
}

func (h *PromotionHandler) GetPromoProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	amount := 3
	products, err := h.promotionUsecase.GetPromoProducts(ctx, uint(amount))
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
		Data: dto.ConvertPromoProductsToDTOs(products),
	})
}

package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/gorilla/mux"
)

type SortingHandler struct {
	router *mux.Router
}

func NewSortingHandler() *SortingHandler {
	return &SortingHandler{
		router: mux.NewRouter(),
	}
}

func (h *SortingHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/public/sorting").Subrouter()
	{
		h.router.Handle("/v1/products", http.HandlerFunc(h.GetProductSorting)).Methods("GET", "OPTIONS")
		h.router.Handle("/v1/reviews", http.HandlerFunc(h.GetReviewSorting)).Methods("GET", "OPTIONS")
	}
}

func (h *SortingHandler) GetProductSorting(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertSortTypesToDTO(helper.GetProductSortTypes()),
	})
}

func (h *SortingHandler) GetReviewSorting(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertSortTypesToDTO(helper.GetReviewSortTypes()),
	})
}

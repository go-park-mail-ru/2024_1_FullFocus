package delivery

import (
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	router  *mux.Router
	usecase usecase.Categories
}

func NewCategoryHandler(uc usecase.Categories) *CategoryHandler {
	return &CategoryHandler{
		router:  mux.NewRouter(),
		usecase: uc,
	}
}

func (h *CategoryHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/category/public/").Subrouter()
	{
		h.router.Handle("/v1", http.HandlerFunc(h.GetAllCategories)).Methods("GET", "OPTIONS")
	}
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	categories, err := h.usecase.GetAllCategories(ctx)
	if errors.Is(err, models.ErrInternal) {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertCategoriesToDto(categories),
	})
}

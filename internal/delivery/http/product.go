package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
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
	h.router = r.PathPrefix("/products").Subrouter()
	{
		h.router.Handle("/", http.HandlerFunc(h.GetProducts)).Methods("GET", "OPTIONS")
	}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var lastID, limit = 1, 10
	qID, ok := r.URL.Query()["lastid"]
	if ok {
		intID, err := strconv.Atoi(qID[0])
		if err == nil && intID > 0 {
			lastID = intID
		}
	}
	qlim, ok := r.URL.Query()["limit"]
	if ok {
		intLim, err := strconv.Atoi(qlim[0])
		if err == nil && intLim > 0 && intLim < 20 {
			limit = intLim
		}
	}
	prods, err := h.usecase.GetProducts(ctx, lastID, limit)
	if errors.Is(err, models.ErrNoProduct) {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 404,
			Msg:    "not found",
			MsgRus: "по данному запросу товары не найдены",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   prods,
	})
}

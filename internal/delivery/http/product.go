package delivery

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	srv     *http.Server
	router  *mux.Router
	usecase usecase.Products
}

func NewProductHandler(srv *http.Server, u usecase.Products) *ProductHandler {
	return &ProductHandler{
		srv:     srv,
		router:  mux.NewRouter(),
		usecase: u,
	}
}

func (h *ProductHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/products").Subrouter()
	h.router.Handle("/", http.HandlerFunc(h.GetProducts)).Methods("GET", "OPTIONS")
}

func (h *ProductHandler) Run() error {
	return h.srv.ListenAndServe()
}

func (h *ProductHandler) Stop() error {
	return h.srv.Shutdown(context.Background())
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := logger.LoggerFromContext(ctx)

	var lastID, limit int = 1, 10
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
		err := helper.JSONResponse(w, 200, models.ErrResponse{
			Status: 404,
			Msg:    "not found",
			MsgRus: "по данному запросу товары не найдены"})
		if err != nil {
			l.Error(fmt.Sprintf("marshall error: %v", err))
		}
		return
	}
	err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
		Data:   prods,
	})
	if err != nil {
		l.Error(fmt.Sprintf("marshall error: %v", err))
	}
}

package delivery

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
)

type ProductsHandler struct {
	srv     *http.Server
	router  *mux.Router
	usecase usecase.Products
}

func NewProductsHandler(srv *http.Server, u usecase.Products) *ProductsHandler {
	return &ProductsHandler{
		srv:     srv,
		router:  mux.NewRouter(),
		usecase: u,
	}
}

func (h *ProductsHandler) InitRouter(r *mux.Router) {
	h.router = r
	h.router.Handle("/", http.HandlerFunc(h.GetProducts)).Methods("GET", "OPTIONS")
}

func (h *ProductsHandler) Run() error {
	return h.srv.ListenAndServe()
}

func (h *ProductsHandler) Stop() error {
	return h.srv.Shutdown(context.Background())
}

func (h *ProductsHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	var lastID, limit int = 0, 10
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
			lastID = intLim
		}
	}
	prods, err := h.usecase.GetProducts(lastID, limit)
	if err != nil {
		return
	}
	err = helper.JSONResponse(w, 200, models.SuccessResponse{
		Status: 200,
		Data:   prods,
	})
	// TODO
	if err != nil {
		_ = err
	}
}

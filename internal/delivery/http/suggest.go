package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
)

type SuggestHandler struct {
	router  *mux.Router
	usecase usecase.Suggests
}

func NewSuggestHandler(uc usecase.Suggests) *SuggestHandler {
	return &SuggestHandler{
		router:  mux.NewRouter(),
		usecase: uc,
	}
}

func (h *SuggestHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/public/v1/suggest").Subrouter()
	{
		h.router.Handle("", http.HandlerFunc(h.GetSuggests)).Methods("GET", "OPTIONS")
	}
}

func (h *SuggestHandler) GetSuggests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query().Get("query")
	suggestions, err := h.usecase.GetSuggestions(ctx, query)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    "search error",
			MsgRus: "Ошибка поиска",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertSuggestionToDTO(suggestions),
	})
}

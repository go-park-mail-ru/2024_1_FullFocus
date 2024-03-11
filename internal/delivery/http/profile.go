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

type ProfileHandler struct {
	srv     *http.Server
	router  *mux.Router
	usecase usecase.Profile
}

// TODO добавить инициализацию в main
func NewProfileHandler(s *http.Server, uc usecase.Profile) *ProfileHandler {
	return &ProfileHandler{
		srv:     s,
		router:  mux.NewRouter(),
		usecase: uc,
	}
}

func (h *ProfileHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/profile").Subrouter()
	{
		h.router.Handle("/{id:^([1-9]+[0-9]*)$}", http.HandlerFunc(h.Profile)).Methods("GET", "POST", "OPTIONS")
	}
}

func (h *ProfileHandler) Run() error {
	return h.srv.ListenAndServe()
}

func (h *ProfileHandler) Stop() error {
	return h.srv.Shutdown(context.Background())
}

func (h *ProfileHandler) Profile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prID, _ := strconv.Atoi(vars["id"])
	if r.Method == "GET" {
		profile, err := h.usecase.GetProfile(uint(prID))
		if err != nil {
			helper.JSONResponse(w, 200, models.ErrResponse{
				Status: 404,
				Msg:    err.Error(),
				MsgRus: "Профиль с таким индентефикатором отсутствует",
			})
			return
		}
		helper.JSONResponse(w, 200, models.SuccessResponse{
			Status: 200,
			Data:   profile,
		})
	}

}

package delivery

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

type CsatHandler struct {
	router  *mux.Router
	usecase usecase.CsatUsecase
}

func NewCsatHandler(u usecase.Avatars) *AvatarHandler {
	return &AvatarHandler{
		router:  mux.NewRouter(),
		usecase: u,
	}
}

func (h *CsatHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/v1/csat").Subrouter()
	{
		h.router.Handle("/{id:[1-9]+[0-9]*}", http.HandlerFunc(h.CreatePollRate)).Methods("POST", "OPTIONS")
		h.router.Handle("/all", http.HandlerFunc(h.GetPolls)).Methods("GET", "OPTIONS")
		h.router.Handle("/{id:[1-9]+[0-9]*", http.HandlerFunc(h.GetPollStats)).Methods("GET", "OPTIONS")
	}
}

func (h *CsatHandler) CreatePollRate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uID, err := helper.GetUserIDFromContext(ctx)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 403,
			Msg:    err.Error(),
			MsgRus: "Пользователь не авторизован",
		})
		return
	}
	orderID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Невалидный параметр",
		})
		return
	}
	createInput := models.CreatePollRateInput{
		ProfileID: uID,
		PollID: ,
	}
	if err = h.usecase.CreatePollRate(ctx, )
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка загрузки фото",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

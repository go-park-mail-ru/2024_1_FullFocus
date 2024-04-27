package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

type CsatHandler struct {
	router  *mux.Router
	usecase *usecase.CsatUsecase
}

func NewCsatHandler(u *usecase.CsatUsecase) *CsatHandler {
	return &CsatHandler{
		router:  mux.NewRouter(),
		usecase: u,
	}
}

func (h *CsatHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/v1/csat").Subrouter()
	{
		h.router.Handle("/vote", http.HandlerFunc(h.CreatePollRate)).Methods("POST", "OPTIONS")
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
	var createPollRateInput dto.CreatePollRateInput
	if err = json.NewDecoder(r.Body).Decode(&createPollRateInput); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}
	createInput := models.CreatePollRateInput{
		ProfileID: uID,
		PollID:    createPollRateInput.PollID,
		Rate:      createPollRateInput.Rate,
	}
	if err = h.usecase.CreatePollRate(ctx, createInput); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка создания оценки",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
	})
}

func (h *CsatHandler) GetPolls(w http.ResponseWriter, r *http.Request) {
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
	polls, err := h.usecase.GetPolls(ctx, uID)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Доступных опросов не найдено",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertPolls(polls),
	})
}

func (h *CsatHandler) GetPollStats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pollID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid id slug value",
			MsgRus: "Невалидный параметр",
		})
		return
	}
	stats, err := h.usecase.GetPollStats(ctx, uint(pollID))
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Статистика не найдена",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertStatsData(stats),
	})
}

package delivery

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

type PromocodeHandler struct {
	router  *mux.Router
	usecase usecase.Promocodes
}

func NewPromocodeHandler(u usecase.Promocodes) *PromocodeHandler {
	return &PromocodeHandler{
		router:  mux.NewRouter(),
		usecase: u,
	}
}

func (h *PromocodeHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/v1/promocode").Subrouter()
	{
		h.router.Handle("/all", http.HandlerFunc(h.GetAllPromocodes)).Methods("GET", "OPTIONS")
		h.router.Handle("/info/{code}", http.HandlerFunc(h.GetPromocodeActivationTerms)).Methods("GET", "OPTIONS")
		h.router.Handle("/{id:[1-9]+[0-9]*}", http.HandlerFunc(h.GetPromocodeByID)).Methods("GET", "OPTIONS")
	}
}

func (h *PromocodeHandler) GetPromocodeActivationTerms(w http.ResponseWriter, r *http.Request) {
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
	code := mux.Vars(r)["code"]
	if code == "" {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid code value",
			MsgRus: "Невалидный параметр",
		})
		return
	}
	promo, err := h.usecase.GetPromocodeItemByActivationCode(ctx, uID, code)
	if err != nil {
		if errors.Is(err, models.ErrPromocodeExpired) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Срок действия промокода истек",
			})
			return
		} else if errors.Is(err, models.ErrNoPromocode) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Промокод не найден",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    "Internal error",
			MsgRus: "Неизвестная ошибка",
		})
		return
	}
	data := dto.ConvertTerms(promo)
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   data,
	})
}

func (h *PromocodeHandler) GetPromocodeByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    "invalid code value",
			MsgRus: "Невалидный параметр",
		})
		return
	}
	promo, err := h.usecase.GetPromocodeByID(ctx, uint(id))
	if err != nil {
		if errors.Is(err, models.ErrNoPromocode) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Промокод не найден",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    "Internal error",
			MsgRus: "Неизвестная ошибка",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertPromocode(promo),
	})
}

func (h *PromocodeHandler) GetAllPromocodes(w http.ResponseWriter, r *http.Request) {
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
	promos, err := h.usecase.GetAvailablePromocodes(ctx, uID)
	if err != nil {
		if errors.Is(err, models.ErrNoPromocode) {
			helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
				Status: 400,
				Msg:    err.Error(),
				MsgRus: "Промокоды не найдены",
			})
			return
		}
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    "Internal error",
			MsgRus: "Неизвестная ошибка",
		})
		return
	}
	data := dto.ConvertPromocodes(promos)
	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   data,
	})
}

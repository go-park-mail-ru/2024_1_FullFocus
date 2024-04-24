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

type ReviewHandler struct {
	router        *mux.Router
	reviewUsecase usecase.Reviews
}

func NewReviewHandler(u usecase.Reviews) *ReviewHandler {
	return &ReviewHandler{
		router:        mux.NewRouter(),
		reviewUsecase: u,
	}
}

func (h *ReviewHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/reviews").Subrouter()
	{
		h.router.Handle("/public/v1/{productID:[1-9]+[0-9]*}", http.HandlerFunc(h.GetProductReviews)).Methods("GET", "OPTIONS")
		h.router.Handle("/v1/new", http.HandlerFunc(h.CreateProductReview)).Methods("POST", "OPTIONS")
	}
}

func (h *ReviewHandler) GetProductReviews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	inputData, err := helper.GetReviewsData(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}

	input := dto.ConvertGetReviewInputToModel(inputData)
	reviews, err := h.reviewUsecase.GetProductReviews(ctx, input)
	switch {
	case errors.Is(err, models.ErrInternal):
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 500,
			Msg:    err.Error(),
		})
		return
	case errors.Is(err, models.ErrNoReviews):
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 404,
			Msg:    err.Error(),
			MsgRus: "Отзывы не найдены",
		})
		return
	}

	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 200,
		Data:   dto.ConvertReviewsToDto(reviews),
	})
}

func (h *ReviewHandler) CreateProductReview(w http.ResponseWriter, r *http.Request) {
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

	inputData, err := helper.GetCreateReviewData(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}

	input := dto.ConvertCreateReviewInputToModel(inputData)
	if err = h.reviewUsecase.CreateProductReview(ctx, uID, input); err != nil {
		helper.JSONResponse(ctx, w, 200, dto.ErrResponse{
			Status: 404,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}

	helper.JSONResponse(ctx, w, 200, dto.SuccessResponse{
		Status: 201,
	})
}

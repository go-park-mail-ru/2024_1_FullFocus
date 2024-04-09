package delivery

import (
	"errors"
	"net/http"

	model "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/gorilla/mux"
)

type ProfileHandler struct {
	router  *mux.Router
	usecase usecase.Profiles
}

func NewProfileHandler(u usecase.Profiles) *ProfileHandler {
	return &ProfileHandler{
		router:  mux.NewRouter(),
		usecase: u,
	}
}

func (h *ProfileHandler) InitRouter(r *mux.Router) {
	h.router = r.PathPrefix("/profile").Subrouter()
	{
		h.router.Handle("/update", http.HandlerFunc(h.UpdateProfile)).Methods("POST", "GET")
		h.router.Handle("/get", http.HandlerFunc(h.GetProfile)).Methods("GET")
	}
}

func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) { // Тест в постмане с фикс ID прошел
	ctx := r.Context()
	uID, err := helper.GetUserIDFromContext(ctx)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, model.ErrResponse{
			Status: 400,
			Msg:    "error with userID ",
			MsgRus: "Проблема с UserID",
		})
		return
	}
	profileData, err := helper.GetProfileData(r)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, model.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Ошибка обработки данных",
		})
		return
	}

	/* uID := uint(10000)

	profileData := dto.ProfileData{
		ID:          uID,
		FullName:    "test",
		Email:       "test@mail.com",
		PhoneNumber: "70000000000",
		ImgSrc:      "test",
	}
	*/

	err = h.usecase.UpdateProfile(ctx, uID, profileData)

	if err != nil {
		if validationError := new(model.ValidationError); errors.As(err, &validationError) {
			helper.JSONResponse(ctx, w, 200, validationError.WithCode(400))
			return
		}
		helper.JSONResponse(ctx, w, 200, model.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Пользователя не существует",
		})
		return
	}

	helper.JSONResponse(ctx, w, 200, model.SuccessResponse{
		Status: 200,
	})
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) { // Тест в постмане с фикс ID и данными прошел
	ctx := r.Context()
	uID, err := helper.GetUserIDFromContext(ctx)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, model.ErrResponse{
			Status: 400,
			Msg:    "error with userID ",
			MsgRus: "Проблема с UserID",
		})
		return
	}
	profile, err := h.usecase.GetProfile(ctx, uID)
	if err != nil {
		helper.JSONResponse(ctx, w, 200, model.ErrResponse{
			Status: 400,
			Msg:    err.Error(),
			MsgRus: "Пользователя не существует",
		})
		return
	}
	helper.JSONResponse(ctx, w, 200, model.SuccessResponse{
		Status: 200,
		Data:   profile, // Что-то не так
	})
}

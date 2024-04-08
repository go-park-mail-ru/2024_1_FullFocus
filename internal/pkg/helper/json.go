package helper

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
)

// JSONResponse - принимает ResponseWriter, значение ответа сервера и соответствующее сообщение.
// Записывает в ResponseWriter json строку, возвращает ошибку или nil.
// Warning: Все ответы от бека, в нашем коде будут 200 – Status OK, чтобы отличать результаты
// от нашего кода от результатов nginx и прочих proxy серверов.
func JSONResponse(ctx context.Context, w http.ResponseWriter, statusCode int, message interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		logger.Error(ctx, "marshall error: "+err.Error())
	}
}

func GetLoginData(r *http.Request) (dto.LoginData, error) {
	var data dto.LoginData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return dto.LoginData{}, err
	}
	return data, nil
}

func GetProfileData(r *http.Request) (dto.ProfileData, error) {
	var data dto.ProfileData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return dto.ProfileData{}, err
	}
	return data, nil
}

package helper

import (
	"context"
	"encoding/json"
	"net/http"

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

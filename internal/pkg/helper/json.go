package helper

import (
	"encoding/json"
	"net/http"
)

// JSONResponse - принимает ResponseWriter, значение ответа сервера и соответствующее сообщение.
// Записывает в ResponseWriter json строку, возвращает ошибку или nil.
// Warning: Все ответы от бека, в нашем коде будут 200 – Status OK, чтобы отличать результаты
// от нашего кода от результатов nginx и прочих proxy серверов.
func JSONResponse(w http.ResponseWriter, statusCode int, message interface{}) error {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		return err
	}
	return nil
}

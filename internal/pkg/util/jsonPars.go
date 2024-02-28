package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ParsingJson - принимает ResponseWriter, значение ответа сервера и соответствующее сообщение.
// Записывает в ResponseWriter json строку, возвращает ошибку или nil.
// Warning: Все ответы от бека, в нашем коде будут 200 – Status OK, чтобы отличать результаты
// от нашего кода от результатов nginx и прочих proxy серверов
func ParsingJson(w http.ResponseWriter, statusCode int, message interface{}) error {
	var result string
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if statusCode == 200 {
		result = fmt.Sprintf(`{"status":%d,"data":%s}`, statusCode, string(bytes))
	} else {
		result = fmt.Sprintf(`{"status":%d,"msg":"error","msgRus":%s}`, statusCode, string(bytes))
	}

	_, err = w.Write([]byte(result))
	if err != nil {
		return err
	}

	return nil
}

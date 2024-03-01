package helper

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
)

type SuccessResponse struct {
	Status int
	Data   interface{}
}

type ErrResponse struct {
	Status int
	Msg    string
	MsgRus string
}

type Test struct {
	StatusCode int
	Message    interface{}
}

func TestParsingJson(t *testing.T) {
	cases := []Test{
		{200, ErrResponse{404, "not valid data", "Проверь данные"}},
		{200, ErrResponse{505, "server err", "Сервак лежит"}},
		{200, ErrResponse{404, "Auth err", "Ошибка, надо перезайти"}},
		{200, SuccessResponse{200, map[string]string{"abc": "abc", "bcd": "bcd"}}},
		{400, SuccessResponse{203, map[string]interface{}{"key": "123", "key2": 123}}},
	}

	for _, item := range cases {
		w := httptest.NewRecorder()
		expect, _ := json.Marshal(item.Message)
		expect = append(expect, byte(10))
		err := JSONResponse(w, item.StatusCode, item.Message)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if w.Body.String() != string(expect) {
			t.Errorf("%v != %v", string(expect), w.Body.String())
		}
		fmt.Println(string(expect))
	}
}

func TestParsingJson2(t *testing.T) {
	data := make(map[string]interface{})
	data["key"] = data
	statusCode := 200
	w := httptest.NewRecorder()
	err := JSONResponse(w, statusCode, data)
	if err.Error() != "json: unsupported value: encountered a cycle via map[string]interface {}" {
		t.Fatalf("функция не обработала корректно ошибку %s", err)
	}
}

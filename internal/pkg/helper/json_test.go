package helper_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	model "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
)

type Test struct {
	StatusCode int
	Message    interface{}
}

func TestJSONResponse(t *testing.T) {
	cases := []Test{
		{200, model.ErrResponse{Status: 404, Msg: "not valid data", MsgRus: "Проверь данные"}},
		{200, model.ErrResponse{Status: 505, Msg: "server err", MsgRus: "Сервак лежит"}},
		{200, model.ErrResponse{Status: 404, Msg: "Auth err", MsgRus: "Ошибка, надо перезайти"}},
		{200, model.SuccessResponse{Status: 200, Data: map[string]string{"abc": "abc", "bcd": "bcd"}}},
		{400, model.SuccessResponse{Status: 203, Data: map[string]interface{}{"key": "123", "key2": 123}}},
	}

	for _, item := range cases {
		w := httptest.NewRecorder()
		expect, _ := json.Marshal(item.Message)
		expect = append(expect, byte(10))
		err := helper.JSONResponse(w, item.StatusCode, item.Message)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if w.Body.String() != string(expect) {
			t.Errorf("%v != %v", string(expect), w.Body.String())
		}
	}
}

func TestJSONResponseErr(t *testing.T) {
	data := make(map[string]interface{})
	data["key"] = data
	statusCode := 200
	w := httptest.NewRecorder()
	err := helper.JSONResponse(w, statusCode, data)
	if err.Error() != "json: unsupported value: encountered a cycle via map[string]interface {}" {
		t.Fatalf("функция не обработала корректно ошибку %s", err)
	}
}

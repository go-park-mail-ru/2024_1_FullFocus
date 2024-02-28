package helper

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
)

type Person struct {
	Age    int
	Weight int
	Height int
}
type Test struct {
	StatusCode int
	Message    interface{}
}

type integer struct {
	Int int
}

func TestParsingJson(t *testing.T) {
	cases := []Test{
		{200, Person{Age: 12, Weight: 50, Height: 180}},
		{200, "Error!"},
		{200, Person{Age: 18, Weight: 80, Height: 200}},
		{200, Person{Age: 18, Weight: 80, Height: 200}},
		{200, 405},
		{200, integer{405}},
		{200, "ошибка, перезайди"},
		{200, map[string]string{"abc": "abc", "bcd": "bcd"}},
		{200, map[string]interface{}{"key": "123", "key2": 123}},

		{400, map[string]interface{}{"key": "123", "key2": 123}},
		{500, map[string]interface{}{"key": "123", "key2": 123}},
	}

	var expect string
	for i := 0; i < len(cases); i++ {
		w := httptest.NewRecorder()

		if cases[i].StatusCode/100 == 2 {
			jsonTest, _ := json.Marshal(cases[i].Message)
			expect = fmt.Sprintf(`{"status":%d,"data":%v}`, cases[i].StatusCode, string(jsonTest))
		} else {
			jsonTest, _ := json.Marshal(cases[i].Message)
			expect = fmt.Sprintf(`{"status":%d,"msg":"error","msgRus":%s}`, cases[i].StatusCode, string(jsonTest))
		}

		err := ParsingJson(w, cases[i].StatusCode, cases[i].Message)
		if err != nil {
			t.Errorf("ParsingJson return error: %v", err)
		}

		if w.Body.String() != expect {
			t.Errorf("Unexpected JSON, expect: %s, get: %s", expect, w.Body.String())
		}
	}
}

func TestParsingJson2(t *testing.T) {
	data := make(map[string]interface{})
	data["key"] = data
	statusCode := 200
	w := httptest.NewRecorder()
	err := ParsingJson(w, statusCode, data)
	if err.Error() != "json: unsupported value: encountered a cycle via map[string]interface {}" {
		t.Fatalf("функция не обработала корректно ошибку %s", err)
	}
}

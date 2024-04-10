package models_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

func TestCreateCSRF(t *testing.T) {
	tokens, _ := models.NewJwtToken("test")
	uID := strconv.Itoa(1)
	_, err := tokens.Create(uID, time.Now().Add(1*time.Hour).Unix())
	if err != nil {
		t.Fatalf("err with creation")
	}
}

func TestCheckCSRF(t *testing.T) {
	tokens, _ := models.NewJwtToken("test")
	uID := strconv.Itoa(1)
	token, err := tokens.Create(uID, time.Now().Add(1*time.Hour).Unix())
	if err != nil {
		t.Fatalf("err with creation 1 token")
	}
	_, err = tokens.Check(uID, token)
	if err != nil {
		t.Fatalf("err with check token")
	}
}

func TestCheckFailCSRF(t *testing.T) {
	tokens, _ := models.NewJwtToken("test")
	uID := strconv.Itoa(1)
	token, err := tokens.Create(uID, time.Now().Add(1*time.Second).Unix())
	time.Sleep(3 * time.Second)
	if err != nil {
		t.Fatalf("err with creation 1 token")
	}
	_, err = tokens.Check(uID, token)
	if err != nil {
		t.Log("success")
	}
}

package models

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	"time"
)

type JwtToken struct {
	Secret []byte
}

func NewJwtToken(secret string) (*JwtToken, error) {
	return &JwtToken{Secret: []byte(secret)}, nil
}

type JwtCsrfClaims struct {
	SessionID string `json:"sid"`
	jwt.StandardClaims
}

func (tk *JwtToken) Create(sID string, tokenExpTime int64) (string, error) {
	data := JwtCsrfClaims{
		SessionID: sID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpTime,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(tk.Secret)
}

func (tk *JwtToken) parseSecretGetter(token *jwt.Token) (interface{}, error) {
	method, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok || method.Alg() != "HS256" {
		return nil, errors.New("bad sign method")
	}
	return tk.Secret, nil
}

func (tk *JwtToken) Check(sID string, inputToken string) (bool, error) {
	payload := &JwtCsrfClaims{}
	_, err := jwt.ParseWithClaims(inputToken, payload, tk.parseSecretGetter)
	if err != nil {
		return false, fmt.Errorf("cant parse jwt token: %w", err)
	}
	if payload.Valid() != nil {
		return false, fmt.Errorf("invalid jwt token: %w", err)
	}
	return payload.SessionID == sID, nil
}

package helper

import (
	"bytes"
	"crypto/rand"
	"golang.org/x/crypto/argon2"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
)

const (
	_countBytes  = 8
	_countMemory = 65536
	_countTreads = 4
	_countKeyLen = 32
)

func hashPass(salt []byte, plainPassword string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), []byte(salt), 1, _countMemory, _countTreads, _countKeyLen)
	return append(salt, hashedPass...)
}

func CheckPass(passHash []byte, plainPassword string) bool {
	salt := passHash[0:8]
	userPassHash := hashPass(salt, plainPassword)
	return bytes.Equal(userPassHash, passHash)
}

func MakeNewPassHash(password string) (string, error) {
	salt := make([]byte, _countBytes)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	passwordHash := hashPass(salt, password)
	charset := "latin1"
	e, err := ianaindex.MIME.Encoding(charset)
	if err != nil {
		log.Fatal(err)
	}
	r := transform.NewReader(bytes.NewBufferString(string(passwordHash)), e.NewDecoder())
	result, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(result), nil
}

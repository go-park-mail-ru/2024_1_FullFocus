package helper

import (
	"crypto/rand"
	"golang.org/x/crypto/argon2"
)

const (
	_countBytes  = 8
	_countMemory = 65536
	_countTreads = 4
	_countKeyLen = 32
)

func PasswordArgon2(plainPassword []byte, salt []byte) []byte {
	return argon2.IDKey(plainPassword, salt, 1, _countMemory, _countTreads, _countKeyLen)
}
func MakeHash(password string, salt []byte) string {
	passwordHash := PasswordArgon2([]byte(password), salt)
	saltWithPasswordHash := string(salt) + string(passwordHash)
	return saltWithPasswordHash
}

func MakeNewHash(password string) (string, error) {
	salt := make([]byte, _countBytes)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	passwordHash := MakeHash(password, salt)
	return passwordHash, nil
}

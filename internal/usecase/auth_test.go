package usecase

import (
	"io"
	"log"
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/stretchr/testify/require"
)

func TestNewAuthUsecase(t *testing.T) {
	t.Run("Check auth usecase creation", func(t *testing.T) {
		uc := NewAuthUsecase(repository.NewUserRepo(), repository.NewSessionRepo())
		require.NotEmpty(t, uc, "auth usecase not created")
	})
}

func TestSignup(t *testing.T) {
	log.SetOutput(io.Discard)
	testUser := models.User{
		Username: "test",
		Password: "test",
	}
	t.Run("Check valid user signup", func(t *testing.T) {
		uc := NewAuthUsecase(repository.NewUserRepo(), repository.NewSessionRepo())
		_, _, err := uc.Signup(testUser.Username, testUser.Password)
		require.Equal(t, nil, err, "user not created")
	})
	t.Run("Check duplicate user signup", func(t *testing.T) {
		uc := NewAuthUsecase(repository.NewUserRepo(), repository.NewSessionRepo())
		_, _, err := uc.Signup(testUser.Username, testUser.Password)
		require.Equal(t, nil, err, "valid user not created")
		_, _, err = uc.Signup(testUser.Username, testUser.Password)
		require.Equal(t, models.ErrUserAlreadyExists, err, "duplicate user created")
	})
}

func TestLogin(t *testing.T) {
	testUser := models.User{
		Username: "test",
		Password: "test",
	}
	t.Run("Check valid login", func(t *testing.T) {
		uc := NewAuthUsecase(repository.NewUserRepo(), repository.NewSessionRepo())
		sID, _, err := uc.Signup(testUser.Username, testUser.Password)
		require.Equal(t, nil, err, "valid user not created")
		require.Equal(t, nil, uc.Logout(sID), "user not logged out")
		_, err = uc.Login(testUser.Username, testUser.Password)
		require.Equal(t, nil, err, "valid user not logged in")
	})
	t.Run("Check invalid password", func(t *testing.T) {
		uc := NewAuthUsecase(repository.NewUserRepo(), repository.NewSessionRepo())
		sID, _, err := uc.Signup(testUser.Username, testUser.Password)
		require.Equal(t, nil, err, "valid user not created")
		require.Equal(t, nil, uc.Logout(sID), "user not logged out")
		_, err = uc.Login(testUser.Username, "")
		require.Equal(t, models.ErrWrongPassword, err, "user logged in by wrong password")
	})
	t.Run("Check invalid login", func(t *testing.T) {
		uc := NewAuthUsecase(repository.NewUserRepo(), repository.NewSessionRepo())
		_, err := uc.Login("abc", "")
		require.Equal(t, models.ErrNoUser, err, "invalid user logged in")
	})
}

func TestIsLoggedIn(t *testing.T) {
	testUser := models.User{
		Username: "test",
		Password: "test",
	}
	t.Run("Check valid login", func(t *testing.T) {
		uc := NewAuthUsecase(repository.NewUserRepo(), repository.NewSessionRepo())
		sID, _, _ := uc.Signup(testUser.Username, testUser.Password)
		require.Equal(t, true, uc.IsLoggedIn(sID), "valid user not logged in")
	})
	t.Run("Check invalid login", func(t *testing.T) {
		uc := NewAuthUsecase(repository.NewUserRepo(), repository.NewSessionRepo())
		require.Equal(t, false, uc.IsLoggedIn("123"), "invalid user logged in")
	})
}

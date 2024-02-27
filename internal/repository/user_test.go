package repository

import (
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

func TestNewUserRepo(t *testing.T) {
	t.Run("Check UserRepo creation", func(t *testing.T) {
		ur := NewUserRepo()
		if ur == nil {
			t.Errorf("UserRepo not created")
		}
	})
}

func TestCreateUser(t *testing.T) {
	ur := NewUserRepo()
	testUser1 := models.User{
		Username: "test",
		Password: "test",
	}
	t.Run("Check new user creation", func(t *testing.T) {
		if _, err := ur.CreateUser(&testUser1); err != nil {
			t.Errorf("User not created")
		}
	})
	t.Run("Check duplicate username creation", func(t *testing.T) {
		if _, err := ur.CreateUser(&testUser1); err != ErrUserAlreadyExists {
			t.Errorf("Duplicate username user created")
		}
	})
}

func TestGetUser(t *testing.T) {
	ur := NewUserRepo()
	testUser1 := models.User{
		Username: "test",
		Password: "test",
	}
	t.Run("Check existing user error", func(t *testing.T) {
		ur.CreateUser(&testUser1)
		if _, err := ur.GetUser(testUser1.Username); err != nil {
			t.Errorf("Got incorrect result")
		}
	})
	testUser2 := models.User{
		Username: "test1",
		Password: "test1",
	}
	t.Run("Check existing user error", func(t *testing.T) {
		uid, _ := ur.CreateUser(&testUser2)
		if gotUser, _ := ur.GetUser(testUser2.Username); gotUser.ID != uid {
			t.Errorf("Got incorrect result")
		}
	})
	t.Run("Check empty user", func(t *testing.T) {
		if _, err := ur.GetUser("abc"); err != ErrNoUser {
			t.Errorf("Got incorrect result")
		}
	})
}

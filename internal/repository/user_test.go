package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

func TestNewUserRepo(t *testing.T) {
	t.Run("Check UserRepo creation", func(t *testing.T) {
		ur := repository.NewUserRepo()
		require.NotEmpty(t, ur, "userrepo not created")
	})
}

func TestCreateUser(t *testing.T) {
	ur := repository.NewUserRepo()
	testUser1 := models.User{
		Username: "test",
		Password: "test",
	}
	t.Run("Check new user creation", func(t *testing.T) {
		_, err := ur.CreateUser(context.Background(), testUser1)
		require.NoError(t, err, "user not created")
	})
	t.Run("Check duplicate username creation", func(t *testing.T) {
		_, err := ur.CreateUser(context.Background(), testUser1)
		require.Equal(t, models.ErrUserAlreadyExists, err, "duplicate username user created")
	})
}

func TestGetUser(t *testing.T) {
	ur := repository.NewUserRepo()
	testUser1 := models.User{
		Username: "test",
		Password: "test",
	}
	t.Run("Check existing user error", func(t *testing.T) {
		_, _ = ur.CreateUser(context.Background(), testUser1)
		_, err := ur.GetUser(context.Background(), testUser1.Username)
		require.NoError(t, err, "existing user not found")
	})
	t.Run("Check empty user", func(t *testing.T) {
		_, err := ur.GetUser(context.Background(), "abc")
		require.Equal(t, models.ErrNoUser, err, "found invalid user")
	})
}

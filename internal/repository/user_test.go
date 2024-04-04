package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	mock_database "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database/mocks"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

func TestNewUserRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock_database.NewMockDatabase(ctrl)
	defer ctrl.Finish()
	t.Run("Check UserRepo creation", func(t *testing.T) {
		ur := repository.NewUserRepo(db)
		require.NotEmpty(t, ur, "userrepo not created")
	})
}

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		name          string
		user          models.User
		mockBehavior  func(*mock_database.MockDatabase, string, database.UserTable)
		expectedID    uint
		expectedError error
	}{
		{
			name: "Test successful user creation",
			user: models.User{
				ID:           1,
				Username:     "test",
				PasswordHash: "test",
			},
			mockBehavior: func(d *mock_database.MockDatabase, q string, u database.UserTable) {
				d.EXPECT().Exec(context.Background(), q, u).Return(mock_database.MockSqlResult{}, nil)
			},
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Test duplicate user creation",
			user: models.User{
				ID:           1,
				Username:     "test",
				PasswordHash: "test",
			},
			mockBehavior: func(d *mock_database.MockDatabase, q string, u database.UserTable) {
				d.EXPECT().Exec(context.Background(), q, u).Return(mock_database.MockSqlResult{}, sql.ErrNoRows)
			},
			expectedID:    0,
			expectedError: models.ErrUserAlreadyExists,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mock_database.NewMockDatabase(ctrl)
			defer ctrl.Finish()
			testCase.mockBehavior(db, "INSERT INTO default_user (id, user_login, password_hash) VALUES ($1, $2, $3);", database.ConvertUserToTable(testCase.user))
			ur := repository.NewUserRepo(db)

			uID, err := ur.CreateUser(context.Background(), testCase.user)
			require.Equal(t, testCase.expectedID, uID)
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

func TestGetUser(t *testing.T) {
	testCases := []struct {
		name          string
		username      string
		mockBehavior  func(*mock_database.MockDatabase, *database.UserTable, string, string)
		expectedError error
	}{
		{
			name:     "Test successful get",
			username: "test",
			mockBehavior: func(d *mock_database.MockDatabase, u *database.UserTable, q string, username string) {
				d.EXPECT().Get(context.Background(), u, q, username).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:     "Test not existing get",
			username: "test",
			mockBehavior: func(d *mock_database.MockDatabase, u *database.UserTable, q string, username string) {
				d.EXPECT().Get(context.Background(), u, q, username).Return(sql.ErrNoRows)
			},
			expectedError: models.ErrNoUser,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mock_database.NewMockDatabase(ctrl)
			defer ctrl.Finish()
			testCase.mockBehavior(db, &database.UserTable{}, "SELECT id FROM default_user WHERE user_login = $1;", testCase.username)
			ur := repository.NewUserRepo(db)
			_, err := ur.GetUser(context.Background(), testCase.username)
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

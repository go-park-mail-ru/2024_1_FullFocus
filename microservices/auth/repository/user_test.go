package repository_test

// import (
// 	"context"
// 	"database/sql"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"

// 	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
// 	mock_database "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database/mocks"
// 	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
// 	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
// )

// func TestNewUserRepo(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	db := mock_database.NewMockDatabase(ctrl)
// 	defer ctrl.Finish()
// 	t.Run("Check UserRepo creation", func(t *testing.T) {
// 		ur := repository.NewUserRepo(db)
// 		require.NotEmpty(t, ur, "userrepo not created")
// 	})
// }

// func TestCreateUser(t *testing.T) {
// 	testCases := []struct {
// 		name         string
// 		user         models.User
// 		mockBehavior func(*mock_database.MockDatabase, *dao.UserTable, string, string, string)
// 		// TODO expectedID    uint
// 		expectedError error
// 	}{
// 		{
// 			name: "Test successful user creation",
// 			user: models.User{
// 				ID:           1,
// 				Username:     "test",
// 				PasswordHash: "test",
// 			},
// 			mockBehavior: func(d *mock_database.MockDatabase, t *dao.UserTable, q string, l, p string) {
// 				d.EXPECT().Get(context.Background(), t, q, l, p).Return(nil)
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name: "Test duplicate user creation",
// 			user: models.User{
// 				ID:           1,
// 				Username:     "test",
// 				PasswordHash: "test",
// 			},
// 			mockBehavior: func(d *mock_database.MockDatabase, t *dao.UserTable, q string, l, p string) {
// 				d.EXPECT().Get(context.Background(), t, q, l, p).Return(sql.ErrNoRows)
// 			},
// 			expectedError: models.ErrUserAlreadyExists,
// 		},
// 	}
// 	for _, testCase := range testCases {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			db := mock_database.NewMockDatabase(ctrl)
// 			defer ctrl.Finish()

// 			q := `INSERT INTO default_user (user_login, password_hash) VALUES ($1, $2) returning id;`
// 			tmpRow := &dao.UserTable{}
// 			testCase.mockBehavior(db, tmpRow, q, testCase.user.Username, testCase.user.PasswordHash)
// 			ur := repository.NewUserRepo(db)

// 			_, err := ur.CreateUser(context.Background(), testCase.user)

// 			require.ErrorIs(t, err, testCase.expectedError)
// 		})
// 	}
// }

// func TestGetUser(t *testing.T) {
// 	testCases := []struct {
// 		name         string
// 		username     string
// 		mockBehavior func(*mock_database.MockDatabase, *dao.UserTable, string, string)
// 		// TODO expectedUser models.User
// 		expectedError error
// 	}{
// 		{
// 			name:     "Test successful get",
// 			username: "test",
// 			mockBehavior: func(d *mock_database.MockDatabase, u *dao.UserTable, q string, username string) {
// 				d.EXPECT().Get(context.Background(), u, q, username).Return(nil)
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name:     "Test not existing get",
// 			username: "test",
// 			mockBehavior: func(d *mock_database.MockDatabase, u *dao.UserTable, q string, username string) {
// 				d.EXPECT().Get(context.Background(), u, q, username).Return(sql.ErrNoRows)
// 			},
// 			expectedError: models.ErrNoUser,
// 		},
// 	}
// 	for _, testCase := range testCases {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			db := mock_database.NewMockDatabase(ctrl)
// 			defer ctrl.Finish()

// 			q := `SELECT id, password_hash FROM default_user WHERE user_login = $1;`
// 			testCase.mockBehavior(db, &dao.UserTable{}, q, testCase.username)
// 			ur := repository.NewUserRepo(db)

// 			_, err := ur.GetUser(context.Background(), testCase.username)
// 			require.ErrorIs(t, err, testCase.expectedError)
// 		})
// 	}
// }

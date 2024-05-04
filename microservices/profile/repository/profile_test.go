package repository_test

// import (
// 	"context"
// 	"database/sql"
// 	"testing"

// 	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
// 	mock_database "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database/mocks"
// 	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
// 	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// func TestNewProfileRepo(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	db := mock_database.NewMockDatabase(ctrl)
// 	defer ctrl.Finish()
// 	t.Run("Check Repo creation", func(t *testing.T) {
// 		pr := repository.NewProfileRepo(db)
// 		require.NotEmpty(t, pr, "profile repo not created")
// 	})
// }

// func TestCreateProfile(t *testing.T) {
// 	testCases := []struct {
// 		name          string
// 		profile       models.Profile
// 		mockBehavior  func(*mock_database.MockDatabase, string, models.Profile)
// 		expectedID    uint
// 		expectedError error
// 	}{
// 		{
// 			name: "Test successful profile creation",
// 			profile: models.Profile{
// 				ID:          1,
// 				FullName:    "test",
// 				Email:       "test@mail.ru",
// 				PhoneNumber: "70000000000",
// 				ImgSrc:      "aaa",
// 			},
// 			mockBehavior: func(d *mock_database.MockDatabase, q string, u models.Profile) {
// 				d.EXPECT().Exec(context.Background(), q, u.ID, u.FullName, u.Email, u.PhoneNumber, u.ImgSrc).Return(mock_database.MockSQLResult{}, nil)
// 			},
// 			expectedID:    1,
// 			expectedError: nil,
// 		},
// 		{
// 			name: "Test duplicate profile creation",
// 			profile: models.Profile{
// 				ID:          1,
// 				FullName:    "test",
// 				Email:       "test@mail.ru",
// 				PhoneNumber: "70000000000",
// 				ImgSrc:      "aaa",
// 			},
// 			mockBehavior: func(d *mock_database.MockDatabase, q string, u models.Profile) {
// 				d.EXPECT().Exec(context.Background(), q, u.ID, u.FullName, u.Email, u.PhoneNumber, u.ImgSrc).Return(mock_database.MockSQLResult{}, sql.ErrNoRows)
// 			},
// 			expectedID:    0,
// 			expectedError: models.ErrProfileAlreadyExists,
// 		},
// 	}
// 	for _, testCase := range testCases {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			db := mock_database.NewMockDatabase(ctrl)
// 			defer ctrl.Finish()
// 			testCase.mockBehavior(db, "INSERT INTO ozon.user_profile (id, full_name, email, phone_number, imgsrc) VALUES ($1, $2, $3, $4, $5);", testCase.profile)
// 			pr := repository.NewProfileRepo(db)

// 			uID, err := pr.CreateProfile(context.Background(), testCase.profile)
// 			require.Equal(t, testCase.expectedID, uID)
// 			require.ErrorIs(t, err, testCase.expectedError)
// 		})
// 	}
// }

// func TestGetProfile(t *testing.T) {
// 	testCases := []struct {
// 		name          string
// 		id            uint
// 		mockBehavior  func(*mock_database.MockDatabase, *dao.ProfileTable, string, uint)
// 		expectedError error
// 	}{
// 		{
// 			name: "Test successful get",
// 			id:   1,
// 			mockBehavior: func(d *mock_database.MockDatabase, u *dao.ProfileTable, q string, id uint) {
// 				d.EXPECT().Get(context.Background(), u, q, id).Return(nil)
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name: "Test not existing get",
// 			id:   1,
// 			mockBehavior: func(d *mock_database.MockDatabase, u *dao.ProfileTable, q string, id uint) {
// 				d.EXPECT().Get(context.Background(), u, q, id).Return(sql.ErrNoRows)
// 			},
// 			expectedError: models.ErrNoProfile,
// 		},
// 	}
// 	for _, testCase := range testCases {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			db := mock_database.NewMockDatabase(ctrl)
// 			defer ctrl.Finish()
// 			testCase.mockBehavior(db, &dao.ProfileTable{}, "SELECT id, full_name, email, phone_number, imgsrc FROM ozon.user_profile WHERE id = $1;", testCase.id)
// 			pr := repository.NewProfileRepo(db)
// 			_, err := pr.GetProfile(context.Background(), testCase.id)
// 			require.ErrorIs(t, err, testCase.expectedError)
// 		})
// 	}
// }

// func TestUpdateProfile(t *testing.T) {
// 	testCases := []struct {
// 		name          string
// 		profile       models.Profile
// 		mockBehavior  func(d *mock_database.MockDatabase, u *dao.ProfileTable, q string, name string, email string, number string, img string, id uint)
// 		expectedError error
// 	}{
// 		{
// 			name: "Test successful get",
// 			profile: models.Profile{
// 				ID:          1,
// 				FullName:    "test",
// 				Email:       "test@mail.ru",
// 				PhoneNumber: "70000000000",
// 				ImgSrc:      "aaa",
// 			},
// 			mockBehavior: func(d *mock_database.MockDatabase, u *dao.ProfileTable, q string, name string, email string, number string, img string, id uint) {
// 				d.EXPECT().Get(context.Background(), u, q, name, email, number, img, id).Return(nil)
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name: "Test fail get",
// 			profile: models.Profile{
// 				ID:          1,
// 				FullName:    "test",
// 				Email:       "test@mail.ru",
// 				PhoneNumber: "70000000000",
// 				ImgSrc:      "aaa",
// 			},
// 			mockBehavior: func(d *mock_database.MockDatabase, u *dao.ProfileTable, q string, name string, email string, number string, img string, id uint) {
// 				d.EXPECT().Get(context.Background(), u, q, name, email, number, img, id).Return(sql.ErrNoRows)
// 			},
// 			expectedError: models.ErrNoProfile,
// 		},
// 	}
// 	for _, testCase := range testCases {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			db := mock_database.NewMockDatabase(ctrl)
// 			defer ctrl.Finish()
// 			testCase.mockBehavior(db, &dao.ProfileTable{},
// 				"UPDATE ozon.user_profile SET full_name=$1, email=$2, phone_number=$3, imgsrc=$4 WHERE id = $5 RETURNING id;",
// 				testCase.profile.FullName, testCase.profile.Email, testCase.profile.PhoneNumber, testCase.profile.ImgSrc, testCase.profile.ID)
// 			pr := repository.NewProfileRepo(db)
// 			err := pr.UpdateProfile(context.Background(), testCase.profile.ID, testCase.profile)
// 			require.ErrorIs(t, err, testCase.expectedError)
// 		})
// 	}
// }

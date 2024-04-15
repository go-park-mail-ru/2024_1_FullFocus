package usecase_test

import (
	"context"
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	mock_repository "github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/mocks"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

func TestNewProfileUsecase(t *testing.T) {
	t.Run("Check products Usecase creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		pu := usecase.NewProfileUsecase(mock_repository.NewMockProfiles(ctrl))
		require.NotEmpty(t, pu, "product repo not created")
	})
}

func TestGetProfile(t *testing.T) {
	testCases := []struct {
		name           string
		id             uint
		mockBehavior   func(*mock_repository.MockProfiles, uint)
		expectedResult dto.Profile
		expectedErr    error
	}{
		{
			name: "Check single product get",
			id:   1,
			mockBehavior: func(r *mock_repository.MockProfiles, uID uint) {
				r.EXPECT().GetProfile(context.Background(), uID).Return(models.Profile{}, nil)
			},
			expectedResult: dto.Profile{},
			expectedErr:    nil,
		},
		{
			name: "Check no products get",
			id:   1000,
			mockBehavior: func(r *mock_repository.MockProfiles, uID uint) {
				r.EXPECT().GetProfile(context.Background(), uID).Return(models.Profile{}, models.ErrNoProfile)
			},
			expectedResult: dto.Profile{},
			expectedErr:    models.ErrNoProfile,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProfileRepo := mock_repository.NewMockProfiles(ctrl)
			testCase.mockBehavior(mockProfileRepo, testCase.id)
			pu := usecase.NewProfileUsecase(mockProfileRepo)
			prods, err := pu.GetProfile(context.Background(), testCase.id)
			require.Equal(t, testCase.expectedResult, prods)
			require.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestUpdateProfile(t *testing.T) {
	testCases := []struct {
		name         string
		id           uint
		profile      dto.Profile
		mockBehavior func(r *mock_repository.MockProfiles, uID uint, profile dto.Profile)
		expectedErr  error
		valid        bool
	}{
		{
			name: "get profile success",
			id:   1,
			profile: dto.Profile{
				ID:          1,
				FullName:    "testing",
				Email:       "my@mail.com",
				PhoneNumber: "70000000000",
				ImgSrc:      "test",
			},
			mockBehavior: func(r *mock_repository.MockProfiles, uID uint, profile dto.Profile) {
				r.EXPECT().UpdateProfile(context.Background(), uID, dto.ConvertProfileToProfileData(profile)).Return(nil)
			},
			expectedErr: nil,
			valid:       true,
		},
		{
			name: "get profile fail",
			id:   1000,
			profile: dto.Profile{
				ID:          10000,
				FullName:    "testing",
				Email:       "my@mail.com",
				PhoneNumber: "79037783633",
				ImgSrc:      "test",
			},
			mockBehavior: func(r *mock_repository.MockProfiles, uID uint, profile dto.Profile) {
				r.EXPECT().UpdateProfile(context.Background(), uID, dto.ConvertProfileToProfileData(profile)).Return(models.ErrNoProfile)
			},
			expectedErr: models.ErrNoProfile,
			valid:       true,
		},
		{
			name: "get profile fail valid email",
			id:   1,
			profile: dto.Profile{
				ID:          1,
				FullName:    "testingemail",
				Email:       "myemail.com",
				PhoneNumber: "70000000000",
				ImgSrc:      "test",
			},
			mockBehavior: func(r *mock_repository.MockProfiles, uID uint, profile dto.Profile) {
				r.EXPECT().UpdateProfile(context.Background(), uID, dto.ConvertProfileToProfileData(profile)).Return(nil)
			},
			expectedErr: helper.NewValidationError("invalid email input", "Имеил должен содержать @ и ."),
			valid:       false,
		},
		{
			name: "get profile fail valid email",
			id:   1,
			profile: dto.Profile{
				ID:          1,
				FullName:    "te",
				Email:       "my@email.com",
				PhoneNumber: "7000000000",
				ImgSrc:      "test",
			},
			mockBehavior: func(r *mock_repository.MockProfiles, uID uint, profile dto.Profile) {
				r.EXPECT().UpdateProfile(context.Background(), uID, dto.ConvertProfileToProfileData(profile)).Return(nil)
			},
			expectedErr: helper.NewValidationError("invalid fullname input",
				"Имя должно содержать от 4 до 32 букв английского алфавита или цифр"),
			valid: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProfileRepo := mock_repository.NewMockProfiles(ctrl)
			if testCase.valid {
				testCase.mockBehavior(mockProfileRepo, testCase.id, testCase.profile)
			}
			pu := usecase.NewProfileUsecase(mockProfileRepo)
			err := pu.UpdateProfile(context.Background(), testCase.id, testCase.profile)
			require.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestCreateProfile(t *testing.T) {
	testCases := []struct {
		name         string
		id           uint
		profile      dto.Profile
		mockBehavior func(r *mock_repository.MockProfiles, profile dto.Profile)
		expectedErr  error
		valid        bool
	}{
		{
			name: "create profile success",
			id:   1,
			profile: dto.Profile{
				ID:          1,
				FullName:    "testing",
				Email:       "my@mail.com",
				PhoneNumber: "7000000000",
				ImgSrc:      "test",
			},
			mockBehavior: func(r *mock_repository.MockProfiles, profile dto.Profile) {
				r.EXPECT().CreateProfile(context.Background(), dto.ConvertProfileToProfileData(profile)).Return(profile.ID, nil)
			},
			expectedErr: nil,
			valid:       true,
		},
		{
			name: "create profile fail",
			id:   1,
			profile: dto.Profile{
				ID:          1,
				FullName:    "testing",
				Email:       "my@mail.com",
				PhoneNumber: "7000000000",
				ImgSrc:      "test",
			},
			mockBehavior: func(r *mock_repository.MockProfiles, profile dto.Profile) {
				r.EXPECT().CreateProfile(context.Background(), dto.ConvertProfileToProfileData(profile)).Return(uint(0), models.ErrProfileAlreadyExists)
			},
			expectedErr: models.ErrProfileAlreadyExists,
			valid:       true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProfileRepo := mock_repository.NewMockProfiles(ctrl)
			if testCase.valid {
				testCase.mockBehavior(mockProfileRepo, testCase.profile)
			}
			pu := usecase.NewProfileUsecase(mockProfileRepo)
			_, err := pu.CreateProfile(context.Background(), testCase.profile)
			require.Equal(t, testCase.expectedErr, err)
		})
	}
}

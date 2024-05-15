package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/models"
	mockrepo "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/repository/mocks"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/usecase"
)

func TestNewUsecase(t *testing.T) {
	t.Run("Check profile Usecase creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		pu := usecase.NewUsecase(mockrepo.NewMockProfile(ctrl))
		require.NotEmpty(t, pu, "usecase not created")
	})
}

func TestGetProfile(t *testing.T) {
	testCases := []struct {
		name           string
		pID            uint
		mockBehavior   func(*mockrepo.MockProfile, uint)
		expectedResult models.Profile
		expectedErr    error
	}{
		{
			name: "Test profile not found",
			pID:  1,
			mockBehavior: func(r *mockrepo.MockProfile, uID uint) {
				r.EXPECT().GetProfile(context.Background(), uID).Return(models.Profile{}, models.ErrNoProfile)
			},
			expectedResult: models.Profile{},
			expectedErr:    models.ErrNoProfile,
		},
		{
			name: "Test profile found",
			pID:  1000,
			mockBehavior: func(r *mockrepo.MockProfile, uID uint) {
				r.EXPECT().GetProfile(context.Background(), uID).Return(models.Profile{
					ID:         1000,
					FullName:   "tester",
					AvatarName: "avatar",
				}, nil)
			},
			expectedResult: models.Profile{
				ID:         1000,
				FullName:   "tester",
				AvatarName: "avatar",
			},
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProfileRepo := mockrepo.NewMockProfile(ctrl)
			testCase.mockBehavior(mockProfileRepo, testCase.pID)
			pu := usecase.NewUsecase(mockProfileRepo)
			profile, err := pu.GetProfile(context.Background(), testCase.pID)
			require.Equal(t, testCase.expectedResult, profile)
			require.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestUpdateProfile(t *testing.T) {
	testCases := []struct {
		name         string
		id           uint
		profile      models.ProfileUpdateInput
		mockBehavior func(r *mockrepo.MockProfile, uID uint, profile models.ProfileUpdateInput)
		expectedErr  error
		valid        bool
	}{
		{
			name: "Text successful update",
			id:   1,
			profile: models.ProfileUpdateInput{
				FullName: "testing",
			},
			mockBehavior: func(r *mockrepo.MockProfile, uID uint, profile models.ProfileUpdateInput) {
				r.EXPECT().UpdateProfile(context.Background(), uID, profile).Return(nil)
			},
			expectedErr: nil,
			valid:       true,
		},
		{
			name: "Test profile not found",
			id:   1000,
			profile: models.ProfileUpdateInput{
				FullName: "testing",
			},
			mockBehavior: func(r *mockrepo.MockProfile, uID uint, profile models.ProfileUpdateInput) {
				r.EXPECT().UpdateProfile(context.Background(), uID, profile).Return(models.ErrNoProfile)
			},
			expectedErr: models.ErrNoProfile,
			valid:       true,
		},
		{
			name: "Test invalid name",
			id:   1,
			profile: models.ProfileUpdateInput{
				FullName: "",
			},
			mockBehavior: func(r *mockrepo.MockProfile, uID uint, profile models.ProfileUpdateInput) {
				r.EXPECT().UpdateProfile(context.Background(), uID, profile).Return(nil)
			},
			expectedErr: models.ErrInvalidInput,
			valid:       false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProfileRepo := mockrepo.NewMockProfile(ctrl)
			if testCase.valid {
				testCase.mockBehavior(mockProfileRepo, testCase.id, testCase.profile)
			}
			pu := usecase.NewUsecase(mockProfileRepo)
			err := pu.UpdateProfile(context.Background(), testCase.id, testCase.profile)
			require.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestCreateProfile(t *testing.T) {
	testCases := []struct {
		name         string
		id           uint
		profile      models.Profile
		mockBehavior func(r *mockrepo.MockProfile, pID uint)
		expectedErr  error
		valid        bool
	}{
		{
			name: "Test successful create",
			id:   1,
			profile: models.Profile{
				ID:       1,
				FullName: "testing",
			},
			mockBehavior: func(r *mockrepo.MockProfile, pID uint) {
				r.EXPECT().CreateProfile(context.Background(), pID).Return(nil)
			},
			expectedErr: nil,
			valid:       true,
		},
		{
			name: "Test already exists",
			id:   1,
			profile: models.Profile{
				ID:       1,
				FullName: "testing",
			},
			mockBehavior: func(r *mockrepo.MockProfile, pID uint) {
				r.EXPECT().CreateProfile(context.Background(), pID).Return(models.ErrProfileAlreadyExists)
			},
			expectedErr: models.ErrProfileAlreadyExists,
			valid:       true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProfileRepo := mockrepo.NewMockProfile(ctrl)
			if testCase.valid {
				testCase.mockBehavior(mockProfileRepo, testCase.id)
			}
			pu := usecase.NewUsecase(mockProfileRepo)
			err := pu.CreateProfile(context.Background(), testCase.id)
			require.Equal(t, testCase.expectedErr, err)
		})
	}
}

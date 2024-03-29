package usecase

import (
	"io"
	"log"
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	mock_repository "github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewAuthUsecase(t *testing.T) {
	t.Run("Check Auth Usecase creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		au := NewAuthUsecase(mock_repository.NewMockUsers(ctrl), mock_repository.NewMockSessions(ctrl))
		require.NotEmpty(t, au, "auth repo not created")
	})
}

func TestSignUp(t *testing.T) {
	log.SetOutput(io.Discard)

	testCases := []struct {
		name                string
		login               string
		password            string
		userMockBehavior    func(*mock_repository.MockUsers, models.User)
		sessionMockBehavior func(*mock_repository.MockSessions, uint)
		expectedSID         string
		expectedErr         error
		callUserMock        bool
		callSessionMock     bool
	}{
		{
			name:     "Check valid user signup",
			login:    "test123",
			password: "Qa5yAbrLhkwT4Y9u",
			userMockBehavior: func(r *mock_repository.MockUsers, user models.User) {
				r.EXPECT().CreateUser(user).Return(uint(0), nil)
			},
			sessionMockBehavior: func(r *mock_repository.MockSessions, userID uint) {
				r.EXPECT().CreateSession(userID).Return("123")
			},
			expectedSID:     "123",
			expectedErr:     nil,
			callUserMock:    true,
			callSessionMock: true,
		},
		{
			name:     "Check duplicate user signup",
			login:    "test123",
			password: "Qa5yAbrLhkwT4Y9u",
			userMockBehavior: func(r *mock_repository.MockUsers, user models.User) {
				r.EXPECT().CreateUser(user).Return(uint(0), models.ErrUserAlreadyExists)
			},
			expectedSID:     "",
			expectedErr:     models.ErrUserAlreadyExists,
			callUserMock:    true,
			callSessionMock: false,
		},
		{
			name:            "Check short username signup",
			login:           "t",
			password:        "test",
			expectedSID:     "",
			expectedErr:     models.ErrShortUsername,
			callUserMock:    false,
			callSessionMock: false,
		},
		{
			name:            "Check weak password signup",
			login:           "test123",
			password:        "12345",
			expectedSID:     "",
			expectedErr:     models.ErrWeakPassword,
			callUserMock:    false,
			callSessionMock: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepo := mock_repository.NewMockUsers(ctrl)
			mockSessionRepo := mock_repository.NewMockSessions(ctrl)
			testUser := models.User{
				ID:       0,
				Username: testCase.login,
				Password: testCase.password,
			}
			if testCase.callUserMock {
				testCase.userMockBehavior(mockUserRepo, testUser)
				if testCase.callSessionMock {
					testCase.sessionMockBehavior(mockSessionRepo, testUser.ID)
				}
			}
			au := NewAuthUsecase(mockUserRepo, mockSessionRepo)
			sID, _, err := au.Signup(testCase.login, testCase.password)
			require.Equal(t, testCase.expectedErr, err)
			require.Equal(t, testCase.expectedSID, sID)
		})
	}
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		name                string
		login               string
		password            string
		userMockBehavior    func(*mock_repository.MockUsers, string)
		sessionMockBehavior func(*mock_repository.MockSessions, uint)
		expectedSID         string
		expectedErr         error
		callUserMock        bool
		callSessionMock     bool
	}{
		{
			name:     "Check valid user login",
			login:    "test",
			password: "test",
			userMockBehavior: func(r *mock_repository.MockUsers, username string) {
				r.EXPECT().GetUser(username).Return(models.User{ID: 0, Username: "test", Password: "test"}, nil)
			},
			sessionMockBehavior: func(r *mock_repository.MockSessions, userID uint) {
				r.EXPECT().CreateSession(userID).Return("123")
			},
			expectedSID:     "123",
			expectedErr:     nil,
			callUserMock:    true,
			callSessionMock: true,
		},
		{
			name:     "Check invalid user login",
			login:    "test",
			password: "test",
			userMockBehavior: func(r *mock_repository.MockUsers, username string) {
				r.EXPECT().GetUser(username).Return(models.User{}, models.ErrNoUser)
			},
			expectedSID:     "",
			expectedErr:     models.ErrNoUser,
			callUserMock:    true,
			callSessionMock: false,
		},
		{
			name:     "Check wrong password login",
			login:    "test",
			password: "wrongpass",
			userMockBehavior: func(r *mock_repository.MockUsers, username string) {
				r.EXPECT().GetUser(username).Return(models.User{ID: 0, Username: "test", Password: "test"}, nil)
			},
			expectedSID:     "",
			expectedErr:     models.ErrWrongPassword,
			callUserMock:    true,
			callSessionMock: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepo := mock_repository.NewMockUsers(ctrl)
			mockSessionRepo := mock_repository.NewMockSessions(ctrl)
			testUser := models.User{
				ID:       0,
				Username: testCase.login,
				Password: testCase.password,
			}
			if testCase.callUserMock {
				testCase.userMockBehavior(mockUserRepo, testUser.Username)
				if testCase.callSessionMock {
					testCase.sessionMockBehavior(mockSessionRepo, testUser.ID)
				}
			}
			au := NewAuthUsecase(mockUserRepo, mockSessionRepo)
			sID, err := au.Login(testCase.login, testCase.password)
			require.Equal(t, testCase.expectedErr, err)
			require.Equal(t, testCase.expectedSID, sID)
		})
	}
}

func TestIsLogout(t *testing.T) {
	testCases := []struct {
		name                string
		sID                 string
		sessionMockBehavior func(*mock_repository.MockSessions, string)
		expectedErr         error
	}{
		{
			name: "Check existing user logout",
			sID:  "test",
			sessionMockBehavior: func(r *mock_repository.MockSessions, sID string) {
				r.EXPECT().DeleteSession(sID).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Check not existing user logout",
			sID:  "test",
			sessionMockBehavior: func(r *mock_repository.MockSessions, sID string) {
				r.EXPECT().DeleteSession(sID).Return(models.ErrNoSession)
			},
			expectedErr: models.ErrNoSession,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepo := mock_repository.NewMockUsers(ctrl)
			mockSessionRepo := mock_repository.NewMockSessions(ctrl)
			testCase.sessionMockBehavior(mockSessionRepo, testCase.sID)
			au := NewAuthUsecase(mockUserRepo, mockSessionRepo)
			err := au.Logout(testCase.sID)
			require.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestIsLoggedIn(t *testing.T) {
	testCases := []struct {
		name                string
		sID                 string
		sessionMockBehavior func(*mock_repository.MockSessions, string)
		expectedResult      bool
	}{
		{
			name: "Check existing user logged in",
			sID:  "test",
			sessionMockBehavior: func(r *mock_repository.MockSessions, sID string) {
				r.EXPECT().SessionExists(sID).Return(true)
			},
			expectedResult: true,
		},
		{
			name: "Check not existing user logged in",
			sID:  "test",
			sessionMockBehavior: func(r *mock_repository.MockSessions, sID string) {
				r.EXPECT().SessionExists(sID).Return(false)
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepo := mock_repository.NewMockUsers(ctrl)
			mockSessionRepo := mock_repository.NewMockSessions(ctrl)
			testCase.sessionMockBehavior(mockSessionRepo, testCase.sID)
			au := NewAuthUsecase(mockUserRepo, mockSessionRepo)
			ok := au.IsLoggedIn(testCase.sID)
			require.Equal(t, testCase.expectedResult, ok)
		})
	}
}

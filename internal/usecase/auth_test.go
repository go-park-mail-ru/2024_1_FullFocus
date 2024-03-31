package usecase_test

import (
	"context"
	"golang.org/x/crypto/argon2"
	"io"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	mock_repository "github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/mocks"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
)

func PasswordArgon2(plainPassword []byte, salt []byte) []byte {
	return argon2.IDKey(plainPassword, salt, 1, 64*1024, 4, 32)
}

func TestNewAuthUsecase(t *testing.T) {
	t.Run("Check Auth Usecase creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		au := usecase.NewAuthUsecase(mock_repository.NewMockUsers(ctrl), mock_repository.NewMockSessions(ctrl))
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
				r.EXPECT().CreateUser(context.Background(), user).Return(uint(0), nil)
			},
			sessionMockBehavior: func(r *mock_repository.MockSessions, userID uint) {
				r.EXPECT().CreateSession(context.Background(), userID).Return("123")
			},
			expectedSID:     "123",
			expectedErr:     nil,
			callUserMock:    true,
			callSessionMock: true,
		},
		{
			name:     "Check valid user signup",
			login:    "test123",
			password: "testtest1",
			userMockBehavior: func(r *mock_repository.MockUsers, user models.User) {
				r.EXPECT().CreateUser(context.Background(), user).Return(uint(0), nil)
			},
			sessionMockBehavior: func(r *mock_repository.MockSessions, userID uint) {
				r.EXPECT().CreateSession(context.Background(), userID).Return("123")
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
				r.EXPECT().CreateUser(context.Background(), user).Return(uint(0), models.ErrUserAlreadyExists)
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
			expectedErr:     models.NewValidationError("invalid login input", "Логин должен содержать от 4 до 32 букв английского алфавита или цифр"),
			callUserMock:    false,
			callSessionMock: false,
		},
		{
			name:            "Check weak password signup",
			login:           "test123",
			password:        "12345",
			expectedSID:     "",
			expectedErr:     models.NewValidationError("invalid password input", "Пароль должен содержать от 8 до 32 букв английского алфавита или цифр"),
			callUserMock:    false,
			callSessionMock: false,
		},
		{
			name:            "Check invalid login signup",
			login:           "test123!@!@$#@#$%@!#$",
			password:        "12345",
			expectedSID:     "",
			expectedErr:     models.NewValidationError("invalid login input", "Логин должен содержать от 4 до 32 букв английского алфавита или цифр"),
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
			au := usecase.NewAuthUsecase(mockUserRepo, mockSessionRepo)
			sID, _, err := au.Signup(context.Background(), testCase.login, testCase.password)
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
		salt                []byte
		userMockBehavior    func(*mock_repository.MockUsers, string)
		sessionMockBehavior func(*mock_repository.MockSessions, uint)
		expectedSID         string
		expectedErr         error
		callUserMock        bool
		callSessionMock     bool
	}{
		{
			name:     "Check valid user login",
			login:    "test123",
			password: "test12345",
			salt:     []byte{0xd7, 0xc2, 0xf2, 0x51, 0xaa, 0x6a, 0x4e, 0x7b},
			userMockBehavior: func(r *mock_repository.MockUsers, username string) {
				r.EXPECT().GetUser(context.Background(), username).Return(models.User{ID: 0, Username: "test123", Password: string(PasswordArgon2([]byte("test12345"), []byte{0xd7, 0xc2, 0xf2, 0x51, 0xaa, 0x6a, 0x4e, 0x7b})),
					Salt: []byte{0xd7, 0xc2, 0xf2, 0x51, 0xaa, 0x6a, 0x4e, 0x7b}}, nil)
			},
			sessionMockBehavior: func(r *mock_repository.MockSessions, userID uint) {
				r.EXPECT().CreateSession(context.Background(), userID).Return("123")
			},
			expectedSID:     "123",
			expectedErr:     nil,
			callUserMock:    true,
			callSessionMock: true,
		},
		{
			name:            "Check invalid username login",
			login:           "test",
			password:        "test",
			expectedSID:     "",
			expectedErr:     models.NewValidationError("invalid password input", "Пароль должен содержать от 8 до 32 букв английского алфавита или цифр"),
			callUserMock:    false,
			callSessionMock: false,
		},
		{
			name:            "Check invalid password login",
			login:           "test123",
			password:        "test%^",
			expectedSID:     "",
			expectedErr:     models.NewValidationError("invalid password input", "Пароль должен содержать от 8 до 32 букв английского алфавита или цифр"),
			callUserMock:    false,
			callSessionMock: false,
		},
		{
			name:     "Check not existing user login",
			login:    "test123",
			password: "wrongpass",
			userMockBehavior: func(r *mock_repository.MockUsers, username string) {
				r.EXPECT().GetUser(context.Background(), username).Return(models.User{}, models.ErrNoUser)
			},
			expectedSID:     "",
			expectedErr:     models.ErrNoUser,
			callUserMock:    true,
			callSessionMock: false,
		},
		{
			name:     "Check wrong password login",
			login:    "test123",
			password: "wrongpass",
			userMockBehavior: func(r *mock_repository.MockUsers, username string) {
				r.EXPECT().GetUser(context.Background(), username).Return(models.User{ID: 0, Username: "test123", Password: "test"}, nil)
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
			}
			if testCase.callUserMock {
				testCase.userMockBehavior(mockUserRepo, testUser.Username)
				if testCase.callSessionMock {
					testCase.sessionMockBehavior(mockSessionRepo, testUser.ID)
				}
			}
			au := usecase.NewAuthUsecase(mockUserRepo, mockSessionRepo)
			sID, err := au.Login(context.Background(), testCase.login, testCase.password)
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
				r.EXPECT().DeleteSession(context.Background(), sID).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Check not existing user logout",
			sID:  "test",
			sessionMockBehavior: func(r *mock_repository.MockSessions, sID string) {
				r.EXPECT().DeleteSession(context.Background(), sID).Return(models.ErrNoSession)
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
			au := usecase.NewAuthUsecase(mockUserRepo, mockSessionRepo)
			err := au.Logout(context.Background(), testCase.sID)
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
				r.EXPECT().SessionExists(context.Background(), sID).Return(true)
			},
			expectedResult: true,
		},
		{
			name: "Check not existing user logged in",
			sID:  "test",
			sessionMockBehavior: func(r *mock_repository.MockSessions, sID string) {
				r.EXPECT().SessionExists(context.Background(), sID).Return(false)
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
			au := usecase.NewAuthUsecase(mockUserRepo, mockSessionRepo)
			ok := au.IsLoggedIn(context.Background(), testCase.sID)
			require.Equal(t, testCase.expectedResult, ok)
		})
	}
}

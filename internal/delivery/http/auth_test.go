package delivery

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	mock_usecase "github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewAuthHandler(t *testing.T) {
	t.Run("Check new auth handler creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		require.NotEmpty(t, NewAuthHandler(&http.Server{}, mock_usecase.NewMockAuth(ctrl)))
	})
}

func TestSignUp(t *testing.T) {
	testCases := []struct {
		name           string
		mockBehavior   func(*mock_usecase.MockAuth, string, string)
		login          string
		password       string
		expectedStatus int
		expectedErr    string
		expectedCookie string
	}{
		{
			name:     "Successful signup",
			login:    "test",
			password: "test",
			mockBehavior: func(u *mock_usecase.MockAuth, login, password string) {
				u.EXPECT().Signup(login, password).Return("test", "test", nil)
			},
			expectedStatus: 200,
			expectedErr:    "",
			expectedCookie: "test",
		},
		{
			name:     "Empty fileds",
			login:    "",
			password: "",
			mockBehavior: func(u *mock_usecase.MockAuth, login, password string) {
				u.EXPECT().Signup(login, password).Return("", "", models.ErrWrongUsername)
			},
			expectedStatus: 400,
			expectedErr:    "unavailable username",
			expectedCookie: "",
		},
		{
			name:     "Short username",
			login:    "test",
			password: "",
			mockBehavior: func(u *mock_usecase.MockAuth, login, password string) {
				u.EXPECT().Signup(login, password).Return("", "", models.ErrWrongUsername)
			},
			expectedStatus: 400,
			expectedErr:    "unavailable username",
			expectedCookie: "",
		},
		{
			name:     "Invalid password",
			login:    "test",
			password: "12345",
			mockBehavior: func(u *mock_usecase.MockAuth, login, password string) {
				u.EXPECT().Signup(login, password).Return("", "", models.ErrWeakPassword)
			},
			expectedStatus: 400,
			expectedErr:    "unavailable password",
			expectedCookie: "",
		},
		{
			name:     "User exists",
			login:    "test",
			password: "12345",
			mockBehavior: func(u *mock_usecase.MockAuth, login, password string) {
				u.EXPECT().Signup(login, password).Return("", "", models.ErrUserAlreadyExists)
			},
			expectedStatus: 400,
			expectedErr:    "user already exists",
			expectedCookie: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockAuthUsecase := mock_usecase.NewMockAuth(ctrl)
			testCase.mockBehavior(mockAuthUsecase, testCase.login, testCase.password)
			srv := &http.Server{}
			ah := NewAuthHandler(srv, mockAuthUsecase)

			form := url.Values{}
			form.Add("login", testCase.login)
			form.Add("password", testCase.password)
			req := httptest.NewRequest("POST", "/api/auth/signup", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			r := httptest.NewRecorder()
			handler := http.HandlerFunc(ah.Signup)
			handler.ServeHTTP(r, req)

			if testCase.expectedStatus != 200 {
				var errResp models.ErrResponse
				err := json.NewDecoder(r.Body).Decode(&errResp)
				require.Equal(t, nil, err)
				require.Equal(t, testCase.expectedStatus, errResp.Status)
				require.Equal(t, testCase.expectedErr, errResp.Msg)
			} else {
				var successResp models.SuccessResponse
				err := json.NewDecoder(r.Body).Decode(&successResp)
				require.Equal(t, nil, err)
				require.Equal(t, testCase.expectedStatus, successResp.Status)
				cookie := r.Result().Cookies()
				err = cookie[0].Valid()
				require.Equal(t, nil, err)
			}

		})
	}
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		name           string
		mockBehavior   func(*mock_usecase.MockAuth, string, string)
		login          string
		password       string
		expectedStatus int
		expectedErr    string
		expectedCookie string
	}{
		{
			name:     "Successful login",
			login:    "test",
			password: "test",
			mockBehavior: func(u *mock_usecase.MockAuth, login, password string) {
				u.EXPECT().Login(login, password).Return("test", nil)
			},
			expectedStatus: 200,
			expectedErr:    "",
			expectedCookie: "test",
		},
		{
			name:     "Wrong login",
			login:    "test",
			password: "test",
			mockBehavior: func(u *mock_usecase.MockAuth, login, password string) {
				u.EXPECT().Login(login, password).Return("", models.ErrNoUser)
			},
			expectedStatus: 401,
			expectedErr:    "wrong login",
			expectedCookie: "",
		},
		{
			name:     "Wrong password",
			login:    "test",
			password: "test",
			mockBehavior: func(u *mock_usecase.MockAuth, login, password string) {
				u.EXPECT().Login(login, password).Return("", models.ErrWrongPassword)
			},
			expectedStatus: 401,
			expectedErr:    "wrong password",
			expectedCookie: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockAuthUsecase := mock_usecase.NewMockAuth(ctrl)
			testCase.mockBehavior(mockAuthUsecase, testCase.login, testCase.password)
			srv := &http.Server{}
			ah := NewAuthHandler(srv, mockAuthUsecase)

			form := url.Values{}
			form.Add("login", testCase.login)
			form.Add("password", testCase.password)
			req := httptest.NewRequest("POST", "/api/auth/login", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			r := httptest.NewRecorder()
			handler := http.HandlerFunc(ah.Login)
			handler.ServeHTTP(r, req)

			if testCase.expectedStatus != 200 {
				var errResp models.ErrResponse
				err := json.NewDecoder(r.Body).Decode(&errResp)
				require.Equal(t, nil, err)
				require.Equal(t, testCase.expectedStatus, errResp.Status)
				require.Equal(t, testCase.expectedErr, errResp.Msg)
			} else {
				var successResp models.SuccessResponse
				err := json.NewDecoder(r.Body).Decode(&successResp)
				require.Equal(t, nil, err)
				require.Equal(t, testCase.expectedStatus, successResp.Status)
				cookie := r.Result().Cookies()
				err = cookie[0].Valid()
				require.Equal(t, nil, err)
			}

		})
	}
}

func TestLogout(t *testing.T) {
	testCases := []struct {
		name           string
		session        string
		mockBehavior   func(*mock_usecase.MockAuth, string)
		expectedStatus int
		expectedErr    string
		expectedCookie string
		setCookie      bool
	}{
		{
			name:    "Successful logout",
			session: "test",
			mockBehavior: func(u *mock_usecase.MockAuth, sID string) {
				u.EXPECT().Logout(sID).Return(nil)
			},
			expectedStatus: 200,
			expectedErr:    "",
			expectedCookie: "",
			setCookie:      true,
		},
		{
			name:    "No session",
			session: "test",
			mockBehavior: func(u *mock_usecase.MockAuth, sID string) {
				u.EXPECT().Logout(sID).Return(models.ErrNoSession)
			},
			expectedStatus: 401,
			expectedErr:    "no session",
			expectedCookie: "",
			setCookie:      true,
		},
		{
			name:    "No session",
			session: "",
			mockBehavior: func(u *mock_usecase.MockAuth, sID string) {
				u.EXPECT().Logout(sID).Return(nil)
			},
			expectedStatus: 401,
			expectedErr:    "no session",
			expectedCookie: "",
			setCookie:      false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockAuthUsecase := mock_usecase.NewMockAuth(ctrl)
			srv := &http.Server{}
			ah := NewAuthHandler(srv, mockAuthUsecase)
			req := httptest.NewRequest("POST", "/api/auth/logout", nil)
			if testCase.setCookie {
				testCase.mockBehavior(mockAuthUsecase, testCase.session)
				req.AddCookie(&http.Cookie{
					Name:    "session_id",
					Value:   testCase.session,
					Expires: time.Now().AddDate(0, 0, 1),
				})
			}

			r := httptest.NewRecorder()
			handler := http.HandlerFunc(ah.Logout)
			handler.ServeHTTP(r, req)

			if testCase.expectedStatus != 200 {
				var errResp models.ErrResponse
				err := json.NewDecoder(r.Body).Decode(&errResp)
				require.Equal(t, nil, err)
				require.Equal(t, testCase.expectedStatus, errResp.Status)
				require.Equal(t, testCase.expectedErr, errResp.Msg)
			} else {
				var successResp models.SuccessResponse
				err := json.NewDecoder(r.Body).Decode(&successResp)
				require.Equal(t, nil, err)
				require.Equal(t, testCase.expectedStatus, successResp.Status)
				diff := time.Now().AddDate(0, 0, -1).UTC().Sub(r.Result().Cookies()[0].Expires.UTC()).Seconds()
				require.Less(t, diff, float64(1))
			}
		})
	}
}

func TestCheckAuth(t *testing.T) {
	testCases := []struct {
		name           string
		session        string
		mockBehavior   func(*mock_usecase.MockAuth, string)
		expectedStatus int
		expectedErr    string
		setCookie      bool
	}{
		{
			name:    "Check logged user",
			session: "test",
			mockBehavior: func(u *mock_usecase.MockAuth, sID string) {
				u.EXPECT().IsLoggedIn(sID).Return(true)
			},
			expectedStatus: 200,
			expectedErr:    "",
			setCookie:      true,
		},
		{
			name:           "Check no cookie",
			session:        "",
			expectedStatus: 401,
			expectedErr:    "no session",
			setCookie:      false,
		},
		{
			name:    "Check no session",
			session: "test",
			mockBehavior: func(u *mock_usecase.MockAuth, sID string) {
				u.EXPECT().IsLoggedIn(sID).Return(false)
			},
			expectedStatus: 401,
			expectedErr:    "no session",
			setCookie:      true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockAuthUsecase := mock_usecase.NewMockAuth(ctrl)
			srv := &http.Server{}
			ah := NewAuthHandler(srv, mockAuthUsecase)
			req := httptest.NewRequest("POST", "/api/auth/check", nil)
			if testCase.setCookie {
				testCase.mockBehavior(mockAuthUsecase, testCase.session)
				req.AddCookie(&http.Cookie{
					Name:  "session_id",
					Value: testCase.session,
				})
			}

			r := httptest.NewRecorder()
			handler := http.HandlerFunc(ah.CheckAuth)
			handler.ServeHTTP(r, req)

			if testCase.expectedStatus != 200 {
				var errResp models.ErrResponse
				err := json.NewDecoder(r.Body).Decode(&errResp)
				require.Equal(t, nil, err)
				require.Equal(t, testCase.expectedStatus, errResp.Status)
				require.Equal(t, testCase.expectedErr, errResp.Msg)
			} else {
				var successResp models.SuccessResponse
				err := json.NewDecoder(r.Body).Decode(&successResp)
				require.Equal(t, nil, err)
				require.Equal(t, testCase.expectedStatus, successResp.Status)
			}
		})
	}
}

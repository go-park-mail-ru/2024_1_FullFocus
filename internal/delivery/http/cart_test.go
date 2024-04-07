package delivery_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	delivery "github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/http"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	mock_usecase "github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewCartHandler(t *testing.T) {
	t.Run("Check new cart handler creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		cu := mock_usecase.NewMockCarts(ctrl)
		ch := delivery.NewCartHandler(cu)
		require.NotEmpty(t, ch)
	})
}

// TODO нужны моки хелпера
func TestUpdateCartItem(t *testing.T) {
	testCases := []struct {
		name           string
		uID            uint
		prID           uint
		count          uint
		mockBehavior   func(*mock_usecase.MockCarts, uint, uint)
		callMock       bool
		fillCtx        bool
		expectedStatus int
		expectedErr    string
	}{
		{
			name:  "Successful create",
			uID:   1,
			prID:  1,
			count: 0,
			mockBehavior: func(u *mock_usecase.MockCarts, uID, pID uint) {
				u.EXPECT().UpdateCartItem(context.WithValue(context.Background(), helper.UserID{}, uID), uID, pID).Return(uint(1), nil)
			},
			callMock:       true,
			fillCtx:        true,
			expectedStatus: 200,
			expectedErr:    "",
		},
		{
			name:  "Successful update",
			uID:   1,
			prID:  1,
			count: 2,
			mockBehavior: func(u *mock_usecase.MockCarts, uID, pID uint) {
				u.EXPECT().UpdateCartItem(context.WithValue(context.Background(), helper.UserID{}, uID), uID, pID).Return(uint(3), nil)
			},
			callMock:       true,
			fillCtx:        true,
			expectedStatus: 200,
			expectedErr:    "",
		},
		{
			name:  "User not authorized",
			uID:   0,
			prID:  1,
			count: 1,
			mockBehavior: func(u *mock_usecase.MockCarts, uID, pID uint) {
				u.EXPECT().UpdateCartItem(context.WithValue(context.Background(), helper.UserID{}, uID), uID, pID).Return(uint(3), nil)
			},
			callMock:       false,
			fillCtx:        false,
			expectedStatus: 403,
			expectedErr:    models.ErrNoUserID.Error(),
		},
		{
			name:  "Incorrect data input",
			uID:   0,
			prID:  1,
			count: 1,
			mockBehavior: func(u *mock_usecase.MockCarts, uID, pID uint) {
				u.EXPECT().UpdateCartItem(context.WithValue(context.Background(), helper.UserID{}, uID), uID, pID).Return(uint(3), nil)
			},
			callMock:       false,
			fillCtx:        false,
			expectedStatus: 403,
			expectedErr:    models.ErrNoUserID.Error(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCartsUsecase := mock_usecase.NewMockCarts(ctrl)
			if testCase.callMock {
				testCase.mockBehavior(mockCartsUsecase, testCase.uID, testCase.prID)
			}
			ch := delivery.NewCartHandler(mockCartsUsecase)

			data := dto.UpdateCartItemInput{
				ProductId: testCase.prID,
			}
			jsonBody, _ := json.Marshal(data)

			req := httptest.NewRequest("POST", "/api/cart/add", bytes.NewReader(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			if testCase.fillCtx {
				reqCtx := context.WithValue(context.Background(), helper.UserID{}, testCase.uID)
				req = req.WithContext(reqCtx)
			}

			r := httptest.NewRecorder()

			handler := http.HandlerFunc(ch.UpdateCartItem)
			handler.ServeHTTP(r, req)

			if testCase.expectedStatus != 200 {
				var errResp dto.ErrResponse
				err := json.NewDecoder(r.Body).Decode(&errResp)
				require.NoError(t, err)
				require.Equal(t, testCase.expectedStatus, errResp.Status)
				require.Equal(t, testCase.expectedErr, errResp.Msg)
			} else {
				var successResp dto.SuccessResponse
				err := json.NewDecoder(r.Body).Decode(&successResp)
				require.NoError(t, err)
				require.Equal(t, testCase.expectedStatus, successResp.Status)
			}
		})
	}
}

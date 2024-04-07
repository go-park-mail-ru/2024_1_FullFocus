package delivery_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	delivery "github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/http"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	mock_usecase "github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase/mocks"
)

func TestNewProductsHandler(t *testing.T) {
	t.Run("Check new products handler creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		require.NotEmpty(t, delivery.NewProductHandler(mock_usecase.NewMockProducts(ctrl)))
	})
}

func TestGetProducts(t *testing.T) {
	testCases := []struct {
		name           string
		mockBehavior   func(*mock_usecase.MockProducts, int, int)
		lastID         int
		limit          int
		lastIDstr      string
		limitStr       string
		expectedStatus int
	}{
		{
			name:      "Successful request",
			lastID:    3,
			limit:     6,
			lastIDstr: "3",
			limitStr:  "6",
			mockBehavior: func(u *mock_usecase.MockProducts, lastId, limit int) {
				u.EXPECT().GetProducts(context.Background(), lastId, limit).Return([]models.Product{}, nil)
			},
			expectedStatus: 200,
		},
		{
			name:      "Successful request with no params",
			lastID:    1,
			limit:     10,
			lastIDstr: "",
			limitStr:  "",
			mockBehavior: func(u *mock_usecase.MockProducts, lastId, limit int) {
				u.EXPECT().GetProducts(context.Background(), lastId, limit).Return([]models.Product{}, nil)
			},
			expectedStatus: 200,
		},
		{
			name:      "Successful request with wrong params",
			lastID:    1,
			limit:     10,
			lastIDstr: "freferf",
			limitStr:  "3123123dwed",
			mockBehavior: func(u *mock_usecase.MockProducts, lastId, limit int) {
				u.EXPECT().GetProducts(context.Background(), lastId, limit).Return([]models.Product{}, nil)
			},
			expectedStatus: 200,
		},
		{
			name:      "Not found request",
			lastID:    90,
			limit:     10,
			lastIDstr: "90",
			limitStr:  "",
			mockBehavior: func(u *mock_usecase.MockProducts, lastId, limit int) {
				u.EXPECT().GetProducts(context.Background(), lastId, limit).Return(nil, models.ErrNoProduct)
			},
			expectedStatus: 404,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProductsUsecase := mock_usecase.NewMockProducts(ctrl)
			testCase.mockBehavior(mockProductsUsecase, testCase.lastID, testCase.limit)
			ph := delivery.NewProductHandler(mockProductsUsecase)

			req := httptest.NewRequest("GET", "/api/products", nil)
			q := req.URL.Query()
			q.Set("lastid", testCase.lastIDstr)
			q.Set("limit", testCase.limitStr)
			req.URL.RawQuery = q.Encode()

			r := httptest.NewRecorder()
			handler := http.HandlerFunc(ph.GetProducts)
			handler.ServeHTTP(r, req)

			var (
				err         error
				successResp dto.SuccessResponse
				errResp     dto.ErrResponse
			)
			if testCase.expectedStatus != 200 {
				err = json.NewDecoder(r.Body).Decode(&errResp)
				require.Equal(t, testCase.expectedStatus, errResp.Status)
				require.NoError(t, err)
			} else {
				err = json.NewDecoder(r.Body).Decode(&successResp)
				require.NoError(t, err)
				require.Equal(t, testCase.expectedStatus, successResp.Status)
			}
		})
	}
}

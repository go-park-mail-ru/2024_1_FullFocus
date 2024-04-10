package delivery_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	delivery "github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/http"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	mock_usecase "github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewCategoryHandler(t *testing.T) {
	t.Run("Check new category handler creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		cu := mock_usecase.NewMockCategories(ctrl)
		ch := delivery.NewCategoryHandler(cu)
		require.NotEmpty(t, ch, "Categories handler not created")
	})
}

func TestGetAllCategories(t *testing.T) {
	testCases := []struct {
		name           string
		mockBehavior   func(*mock_usecase.MockCategories)
		expectedStatus int
		expectedErr    string
	}{
		{
			name: "Successful get",
			mockBehavior: func(u *mock_usecase.MockCategories) {
				u.EXPECT().GetAllCategories(context.Background()).Return([]models.Category{}, nil)
			},
			expectedStatus: 200,
			expectedErr:    "",
		},
		{
			name: "Get zero categories",
			mockBehavior: func(u *mock_usecase.MockCategories) {
				u.EXPECT().GetAllCategories(context.Background()).Return([]models.Category{}, nil)
			},
			expectedStatus: 200,
			expectedErr:    "",
		},
		{
			name: "Internal error",
			mockBehavior: func(u *mock_usecase.MockCategories) {
				u.EXPECT().GetAllCategories(context.Background()).Return([]models.Category{}, models.ErrInternal)
			},
			expectedStatus: 500,
			expectedErr:    models.ErrInternal.Error(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cu := mock_usecase.NewMockCategories(ctrl)
			testCase.mockBehavior(cu)
			ch := delivery.NewCategoryHandler(cu)

			req := httptest.NewRequest("POST", "/api/category/public/v1", nil)
			r := httptest.NewRecorder()

			handler := http.HandlerFunc(ch.GetAllCategories)
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

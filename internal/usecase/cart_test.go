package usecase_test

import (
	"context"
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	mock_repository "github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/mocks"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewCartUsecase(t *testing.T) {
	t.Run("Test new cart usecase creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		cr := mock_repository.NewMockCarts(ctrl)
		cu := usecase.NewCartUsecase(cr)
		require.NotEmpty(t, cu, "cart usecase not created")
	})
}

func TestGetAllCartItems(t *testing.T) {
	testCases := []struct {
		name          string
		uID           uint
		mockBehavior  func(*mock_repository.MockCarts, uint)
		expectedItems []models.CartItem
		expectedError error
	}{
		{
			name: "Test successfull get",
			uID:  1,
			mockBehavior: func(mc *mock_repository.MockCarts, u uint) {
				mc.EXPECT().GetAllCartItems(context.Background(), u).Return([]models.CartItem{{Count: 1}}, nil)
			},
			expectedItems: []models.CartItem{{Count: 1}},
			expectedError: nil,
		},
		{
			name: "Test empty cart get",
			uID:  1,
			mockBehavior: func(mc *mock_repository.MockCarts, u uint) {
				mc.EXPECT().GetAllCartItems(context.Background(), u).Return(nil, models.ErrEmptyCart)
			},
			expectedItems: nil,
			expectedError: models.ErrEmptyCart,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCartRepo := mock_repository.NewMockCarts(ctrl)
			testCase.mockBehavior(mockCartRepo, testCase.uID)
			cu := usecase.NewCartUsecase(mockCartRepo)

			items, err := cu.GetAllCartItems(context.Background(), testCase.uID)

			require.Equal(t, testCase.expectedItems, items)
			require.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestUpdateCartItem(t *testing.T) {
	testCases := []struct {
		name          string
		uID           uint
		prID          uint
		mockBehavior  func(*mock_repository.MockCarts, uint, uint)
		expectedCount uint
		expectedError error
	}{
		{
			name: "Test successfull count increment",
			uID:  1,
			prID: 1,
			mockBehavior: func(mc *mock_repository.MockCarts, u uint, p uint) {
				mc.EXPECT().UpdateCartItem(context.Background(), u, p).Return(uint(2), nil)
			},
			expectedCount: 2,
			expectedError: nil,
		},
		{
			name: "Test successfull add",
			uID:  1,
			prID: 1,
			mockBehavior: func(mc *mock_repository.MockCarts, u uint, p uint) {
				mc.EXPECT().UpdateCartItem(context.Background(), u, p).Return(uint(1), nil)
			},
			expectedCount: 1,
			expectedError: nil,
		},
		{
			name: "Test update not existing item",
			uID:  1,
			prID: 0,
			mockBehavior: func(mc *mock_repository.MockCarts, u uint, p uint) {
				mc.EXPECT().UpdateCartItem(context.Background(), u, p).Return(uint(0), models.ErrNoProduct)
			},
			expectedCount: 0,
			expectedError: models.ErrNoProduct,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCartRepo := mock_repository.NewMockCarts(ctrl)
			testCase.mockBehavior(mockCartRepo, testCase.uID, testCase.prID)
			cu := usecase.NewCartUsecase(mockCartRepo)

			newCount, err := cu.UpdateCartItem(context.Background(), testCase.uID, testCase.prID)

			require.Equal(t, testCase.expectedCount, newCount)
			require.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestDeleteCartItem(t *testing.T) {
	testCases := []struct {
		name          string
		uID           uint
		prID          uint
		mockBehavior  func(*mock_repository.MockCarts, uint, uint)
		expectedCount uint
		expectedError error
	}{
		{
			name: "Test successfull count decrement",
			uID:  1,
			prID: 1,
			mockBehavior: func(mc *mock_repository.MockCarts, u uint, p uint) {
				mc.EXPECT().DeleteCartItem(context.Background(), u, p).Return(uint(1), nil)
			},
			expectedCount: 1,
			expectedError: nil,
		},
		{
			name: "Test successfull delete",
			uID:  1,
			prID: 1,
			mockBehavior: func(mc *mock_repository.MockCarts, u uint, p uint) {
				mc.EXPECT().DeleteCartItem(context.Background(), u, p).Return(uint(0), nil)
			},
			expectedCount: 0,
			expectedError: nil,
		},
		{
			name: "Test not existing item delete",
			uID:  1,
			prID: 0,
			mockBehavior: func(mc *mock_repository.MockCarts, u uint, p uint) {
				mc.EXPECT().DeleteCartItem(context.Background(), u, p).Return(uint(0), models.ErrNoProduct)
			},
			expectedCount: 0,
			expectedError: models.ErrNoProduct,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCartRepo := mock_repository.NewMockCarts(ctrl)
			testCase.mockBehavior(mockCartRepo, testCase.uID, testCase.prID)
			cu := usecase.NewCartUsecase(mockCartRepo)

			newCount, err := cu.DeleteCartItem(context.Background(), testCase.uID, testCase.prID)

			require.Equal(t, testCase.expectedCount, newCount)
			require.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestDeleteAllCartItems(t *testing.T) {
	testCases := []struct {
		name          string
		uID           uint
		mockBehavior  func(*mock_repository.MockCarts, uint)
		expectedError error
	}{
		{
			name: "Test successfull cart clear",
			uID:  1,
			mockBehavior: func(mc *mock_repository.MockCarts, u uint) {
				mc.EXPECT().DeleteAllCartItems(context.Background(), u).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Test empty cart delete",
			uID:  1,
			mockBehavior: func(mc *mock_repository.MockCarts, u uint) {
				mc.EXPECT().DeleteAllCartItems(context.Background(), u).Return(models.ErrEmptyCart)
			},
			expectedError: models.ErrEmptyCart,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCartRepo := mock_repository.NewMockCarts(ctrl)
			testCase.mockBehavior(mockCartRepo, testCase.uID)
			cu := usecase.NewCartUsecase(mockCartRepo)

			err := cu.DeleteAllCartItems(context.Background(), testCase.uID)

			require.Equal(t, testCase.expectedError, err)
		})
	}
}

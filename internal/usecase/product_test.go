package usecase_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"

// 	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
// 	mock_repository "github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/mocks"
// 	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/usecase"
// )

// func TestNewProductsUsecase(t *testing.T) {
// 	t.Run("Check products Usecase creation", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		defer ctrl.Finish()
// 		pu := usecase.NewProductUsecase(mock_repository.NewMockProducts(ctrl))
// 		require.NotEmpty(t, pu, "product repo not created")
// 	})
// }

// func TestGetProducts(t *testing.T) {
// 	testCases := []struct {
// 		name           string
// 		lastID         int
// 		limit          int
// 		mockBehavior   func(*mock_repository.MockProducts, int, int)
// 		expectedResult []models.Product
// 		expectedErr    error
// 	}{
// 		{
// 			name:   "Check single product get",
// 			lastID: 1,
// 			limit:  1,
// 			mockBehavior: func(r *mock_repository.MockProducts, lastID, limit int) {
// 				r.EXPECT().GetProducts(context.Background(), lastID, limit).Return([]models.Product{{}}, nil)
// 			},
// 			expectedResult: []models.Product{{}},
// 			expectedErr:    nil,
// 		},
// 		{
// 			name:   "Check several products get",
// 			lastID: 1,
// 			limit:  3,
// 			mockBehavior: func(r *mock_repository.MockProducts, lastID, limit int) {
// 				r.EXPECT().GetProducts(context.Background(), lastID, limit).Return([]models.Product{{}, {}, {}}, nil)
// 			},
// 			expectedResult: []models.Product{{}, {}, {}},
// 			expectedErr:    nil,
// 		},
// 		{
// 			name:   "Check no products get",
// 			lastID: 1,
// 			limit:  0,
// 			mockBehavior: func(r *mock_repository.MockProducts, lastID, limit int) {
// 				r.EXPECT().GetProducts(context.Background(), lastID, limit).Return(nil, models.ErrNoProduct)
// 			},
// 			expectedResult: nil,
// 			expectedErr:    models.ErrNoProduct,
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
// 			mockProductRepo := mock_repository.NewMockProducts(ctrl)
// 			testCase.mockBehavior(mockProductRepo, testCase.lastID, testCase.limit)
// 			pu := usecase.NewProductUsecase(mockProductRepo)
// 			prods, err := pu.GetProducts(context.Background(), testCase.lastID, testCase.limit)
// 			require.Equal(t, testCase.expectedResult, prods)
// 			require.Equal(t, testCase.expectedErr, err)
// 		})
// 	}
// }

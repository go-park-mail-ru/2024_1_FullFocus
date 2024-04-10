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

func TestNewCategoryUsecase(t *testing.T) {
	t.Run("Test new category usecase creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		cr := mock_repository.NewMockCategories(ctrl)
		cu := usecase.NewCategoryUsecase(cr)
		require.NotEmpty(t, cu, "category usecase not created")
	})
}

func TestGetAllCategories(t *testing.T) {
	testCases := []struct {
		name               string
		mockBehavior       func(*mock_repository.MockCategories)
		expectedCategories []models.Category
		expectedError      error
	}{
		{
			name: "Test successful get",
			mockBehavior: func(m *mock_repository.MockCategories) {
				m.EXPECT().GetAllCategories(context.Background()).Return([]models.Category{
					{
						ID:   1,
						Name: "Electronics",
					},
					{
						ID:   1,
						Name: "Food",
					},
				}, nil)
			},
			expectedCategories: []models.Category{
				{
					ID:   1,
					Name: "Electronics",
				},
				{
					ID:   1,
					Name: "Food",
				},
			},
			expectedError: nil,
		},
		{
			name: "Test no categories get",
			mockBehavior: func(m *mock_repository.MockCategories) {
				m.EXPECT().GetAllCategories(context.Background()).Return([]models.Category{}, nil)
			},
			expectedCategories: []models.Category{},
			expectedError:      nil,
		},
		{
			name: "Test internal error get",
			mockBehavior: func(m *mock_repository.MockCategories) {
				m.EXPECT().GetAllCategories(context.Background()).Return([]models.Category{}, models.ErrInternal)
			},
			expectedCategories: []models.Category{},
			expectedError:      models.ErrInternal,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cr := mock_repository.NewMockCategories(ctrl)
			testCase.mockBehavior(cr)
			cu := usecase.NewCategoryUsecase(cr)

			res, err := cu.GetAllCategories(context.Background())
			require.Equal(t, testCase.expectedCategories, res)
			require.Equal(t, testCase.expectedError, err)
		})
	}
}

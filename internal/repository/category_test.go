package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	mock_database "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database/mocks"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewCategoryRepo(t *testing.T) {
	t.Run("Check CategoryRepo creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		db := mock_database.NewMockDatabase(ctrl)
		defer ctrl.Finish()
		cr := repository.NewCategoryRepo(db)
		require.NotEmpty(t, cr, "category repo not created")
	})
}

func TestGetAllCategories(t *testing.T) {
	testCases := []struct {
		name         string
		mockBehavior func(*mock_database.MockDatabase, *[]dao.CategoryTable, string)
		// TODO expectedCategories []models.Category
		expectedError error
	}{
		{
			name: "Test successful get",
			mockBehavior: func(d *mock_database.MockDatabase, t *[]dao.CategoryTable, q string) {
				d.EXPECT().Select(context.Background(), t, q).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Test empty cart get",
			mockBehavior: func(d *mock_database.MockDatabase, t *[]dao.CategoryTable, q string) {
				d.EXPECT().Select(context.Background(), t, q).Return(errors.ErrUnsupported)
			},
			expectedError: models.ErrInternal,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mock_database.NewMockDatabase(ctrl)
			defer ctrl.Finish()

			q := `SELECT id, category_name FROM category WHERE parent_id IS NULL;`

			rows := []dao.CategoryTable{}
			testCase.mockBehavior(db, &rows, q)
			cr := repository.NewCategoryRepo(db)

			_, err := cr.GetAllCategories(context.Background())
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	mock_database "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database/mocks"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNewCartRepo(t *testing.T) {
	t.Run("Check CartRepo creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		db := mock_database.NewMockDatabase(ctrl)
		defer ctrl.Finish()
		ur := repository.NewCartRepo(db)
		require.NotEmpty(t, ur, "cart repo not created")
	})
}

func TestGetAllCartProductsID(t *testing.T) {
	testCases := []struct {
		name         string
		uID          uint
		mockBehavior func(*mock_database.MockDatabase, *[]database.CartItemTable, string, uint)
		// TODO expectedItems []models.CartItem
		expectedError error
	}{
		{
			name: "Test successful get",
			uID:  1,
			mockBehavior: func(d *mock_database.MockDatabase, i *[]database.CartItemTable, q string, u uint) {
				d.EXPECT().Select(context.Background(), i, q, u).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Test empty cart get",
			uID:  1,
			mockBehavior: func(d *mock_database.MockDatabase, i *[]database.CartItemTable, q string, u uint) {
				d.EXPECT().Select(context.Background(), i, q, u).Return(sql.ErrNoRows)
			},
			expectedError: models.ErrEmptyCart,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mock_database.NewMockDatabase(ctrl)
			defer ctrl.Finish()

			i := []database.CartItemTable{}
			testCase.mockBehavior(db, &i, "SELECT (product_id, count) FROM cart_item WHERE profile_id = $1;", testCase.uID)
			cr := repository.NewCartRepo(db)

			_, err := cr.GetAllCartProductsId(context.Background(), testCase.uID)
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

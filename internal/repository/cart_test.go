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

func TestGetAllCartItems(t *testing.T) {
	testCases := []struct {
		name         string
		uID          uint
		mockBehavior func(*mock_database.MockDatabase, *[]database.CartProductTable, string, uint)
		// TODO expectedItems []models.CartProduct
		expectedError error
	}{
		{
			name: "Test successful get",
			uID:  1,
			mockBehavior: func(d *mock_database.MockDatabase, t *[]database.CartProductTable, q string, u uint) {
				d.EXPECT().Select(context.Background(), t, q, u).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Test empty cart get",
			uID:  1,
			mockBehavior: func(d *mock_database.MockDatabase, t *[]database.CartProductTable, q string, u uint) {
				d.EXPECT().Select(context.Background(), t, q, u).Return(sql.ErrNoRows)
			},
			expectedError: models.ErrEmptyCart,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mock_database.NewMockDatabase(ctrl)
			defer ctrl.Finish()

			q := `SELECT p.id, p.product_name, p.price, p.imgsrc, c.count
		FROM product AS p JOIN cart_item AS c ON p.id = c.product_id
		WHERE profile_id = $1;`
			rows := []database.CartProductTable{}
			testCase.mockBehavior(db, &rows, q, testCase.uID)
			cr := repository.NewCartRepo(db)

			_, err := cr.GetAllCartItems(context.Background(), testCase.uID)
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

func TestGetAllCartItemsID(t *testing.T) {
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
			mockBehavior: func(d *mock_database.MockDatabase, t *[]database.CartItemTable, q string, u uint) {
				d.EXPECT().Select(context.Background(), t, q, u).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Test empty cart get",
			uID:  1,
			mockBehavior: func(d *mock_database.MockDatabase, t *[]database.CartItemTable, q string, u uint) {
				d.EXPECT().Select(context.Background(), t, q, u).Return(sql.ErrNoRows)
			},
			expectedError: models.ErrEmptyCart,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mock_database.NewMockDatabase(ctrl)
			defer ctrl.Finish()

			q := "SELECT product_id, count FROM cart_item WHERE profile_id = $1;"
			rows := []database.CartItemTable{}
			testCase.mockBehavior(db, &rows, q, testCase.uID)
			cr := repository.NewCartRepo(db)

			_, err := cr.GetAllCartItemsId(context.Background(), testCase.uID)
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

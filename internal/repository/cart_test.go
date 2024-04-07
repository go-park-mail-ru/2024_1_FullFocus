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

func TestUpdateCartItem(t *testing.T) {
	testCases := []struct {
		name          string
		uID           uint
		prID          uint
		mockBehavior  func(*mock_database.MockDatabase, string, uint, uint)
		expectedCount uint
		expectedError error
	}{
		{
			name: "Test successful creation",
			uID:  1,
			prID: 1,
			mockBehavior: func(d *mock_database.MockDatabase, q string, u, p uint) {
				d.EXPECT().Exec(context.Background(), q, u, p).Return(mock_database.MockSqlResult{
					LastInsertedId: 1,
					RowsAffect:     1,
				}, nil)
			},
			expectedCount: 1,
			expectedError: nil,
		},
		{
			name: "Test successfull update",
			uID:  1,
			prID: 1,
			mockBehavior: func(d *mock_database.MockDatabase, q string, u, p uint) {
				d.EXPECT().Exec(context.Background(), q, u, p).Return(mock_database.MockSqlResult{
					LastInsertedId: 2,
					RowsAffect:     1,
				}, nil)
			},
			expectedCount: 2,
			expectedError: nil,
		},
		{
			name: "Test not existing item update",
			uID:  1,
			prID: 1,
			mockBehavior: func(d *mock_database.MockDatabase, q string, u, p uint) {
				d.EXPECT().Exec(context.Background(), q, u, p).Return(mock_database.MockSqlResult{
					LastInsertedId: 0,
					RowsAffect:     0,
				}, sql.ErrNoRows)
			},
			expectedCount: 0,
			expectedError: models.ErrNoProduct,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mock_database.NewMockDatabase(ctrl)
			defer ctrl.Finish()

			q := `INSERT INTO ozon.cart_item(profile_id, product_id) VALUES($1, $2)
	ON CONFLICT (profile_id, product_id)
	DO UPDATE set count = cart_item.count + 1
	returning cart_item.count;`
			testCase.mockBehavior(db, q, testCase.uID, testCase.prID)
			cr := repository.NewCartRepo(db)

			newCount, err := cr.UpdateCartItem(context.Background(), testCase.uID, testCase.prID)
			require.Equal(t, newCount, testCase.expectedCount)
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

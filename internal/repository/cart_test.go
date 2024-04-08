package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	mock_database "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database/mocks"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
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
		mockBehavior func(*mock_database.MockDatabase, *[]dao.CartProductTable, string, uint)
		// TODO expectedItems []models.CartProduct
		expectedError error
	}{
		{
			name: "Test successful get",
			uID:  1,
			mockBehavior: func(d *mock_database.MockDatabase, t *[]dao.CartProductTable, q string, u uint) {
				d.EXPECT().Select(context.Background(), t, q, u).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Test empty cart get",
			uID:  1,
			mockBehavior: func(d *mock_database.MockDatabase, t *[]dao.CartProductTable, q string, u uint) {
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
			rows := []dao.CartProductTable{}
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
		mockBehavior func(*mock_database.MockDatabase, *[]dao.CartItemTable, string, uint)
		// TODO expectedItems []models.CartItem
		expectedError error
	}{
		{
			name: "Test successful get",
			uID:  1,
			mockBehavior: func(d *mock_database.MockDatabase, t *[]dao.CartItemTable, q string, u uint) {
				d.EXPECT().Select(context.Background(), t, q, u).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Test empty cart get",
			uID:  1,
			mockBehavior: func(d *mock_database.MockDatabase, t *[]dao.CartItemTable, q string, u uint) {
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
			rows := []dao.CartItemTable{}
			testCase.mockBehavior(db, &rows, q, testCase.uID)
			cr := repository.NewCartRepo(db)

			_, err := cr.GetAllCartItemsID(context.Background(), testCase.uID)
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

func TestUpdateCartItem(t *testing.T) {
	testCases := []struct {
		name          string
		uID           uint
		prID          uint
		mockBehavior  func(*mock_database.MockDatabase, *dao.CartItemTable, string, uint, uint)
		expectedCount uint
		expectedError error
	}{
		{
			name: "Test successful creation",
			uID:  1,
			prID: 1,
			mockBehavior: func(d *mock_database.MockDatabase, t *dao.CartItemTable, q string, u, p uint) {
				d.EXPECT().Get(context.Background(), t, q, u, p).Return(nil)
			},
			expectedCount: 1,
			expectedError: nil,
		},
		{
			name: "Test successfull update",
			uID:  1,
			prID: 1,
			mockBehavior: func(d *mock_database.MockDatabase, t *dao.CartItemTable, q string, u, p uint) {
				d.EXPECT().Get(context.Background(), t, q, u, p).Return(nil)
			},
			expectedCount: 2,
			expectedError: nil,
		},
		{
			name: "Test not existing item update",
			uID:  1,
			prID: 1,
			mockBehavior: func(d *mock_database.MockDatabase, t *dao.CartItemTable, q string, u, p uint) {
				d.EXPECT().Get(context.Background(), t, q, u, p).Return(sql.ErrNoRows)
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

			q := `INSERT INTO cart_item(profile_id, product_id) VALUES($1, $2)
	ON CONFLICT (profile_id, product_id)
	DO UPDATE set count = cart_item.count + 1
	returning cart_item.count AS count;`
			testItem := dao.CartItemTable{}
			testCase.mockBehavior(db, &testItem, q, testCase.uID, testCase.prID)
			cr := repository.NewCartRepo(db)

			_, err := cr.UpdateCartItem(context.Background(), testCase.uID, testCase.prID)
			// TODO require.Equal(t, testCase.expectedCount, newCount)
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

// TODO fix test
// func TestDeleteCartItem(t *testing.T) {
// 	testCases := []struct {
// 		name               string
// 		uID                uint
// 		prID               uint
// 		mockBehaviorUpdate func(*mock_database.MockDatabase, *dao.CartItemTable, string, uint, uint)
// 		mockBehaviorDelete func(*mock_database.MockDatabase, string, uint, uint)
// 		callMockDelete     bool
// 		expectedCount      uint
// 		expectedError      error
// 	}{
// 		{
// 			name: "Test successful count decrement",
// 			uID:  1,
// 			prID: 1,
// 			mockBehaviorUpdate: func(d *mock_database.MockDatabase, t *dao.CartItemTable, q string, u, p uint) {
// 				d.EXPECT().Get(context.Background(), t, q, u, p).Return(nil)
// 			},
// 			callMockDelete: false,
// 			expectedCount:  1,
// 			expectedError:  nil,
// 		},
// 		{
// 			name: "Test successfull delete",
// 			uID:  1,
// 			prID: 1,
// 			mockBehaviorUpdate: func(d *mock_database.MockDatabase, t *dao.CartItemTable, q string, u, p uint) {
// 				d.EXPECT().Get(context.Background(), t, q, u, p).Return(nil)
// 			},
// 			mockBehaviorDelete: func(d *mock_database.MockDatabase, q string, u, p uint) {
// 				d.EXPECT().Exec(context.Background(), q, u, p).Return(mock_database.MockSQLResult{}, nil)
// 			},
// 			callMockDelete: true,
// 			expectedCount:  0,
// 			expectedError:  nil,
// 		},
// 		{
// 			name: "Test not existing item decrement",
// 			uID:  1,
// 			prID: 1,
// 			mockBehaviorUpdate: func(d *mock_database.MockDatabase, t *dao.CartItemTable, q string, u, p uint) {
// 				d.EXPECT().Get(context.Background(), t, q, u, p).Return(sql.ErrNoRows)
// 			},
// 			callMockDelete: false,
// 			expectedCount:  0,
// 			expectedError:  models.ErrNoProduct,
// 		},
// 		{
// 			name: "Test not existing item delete",
// 			uID:  1,
// 			prID: 1,
// 			mockBehaviorUpdate: func(d *mock_database.MockDatabase, t *dao.CartItemTable, q string, u, p uint) {
// 				d.EXPECT().Get(context.Background(), t, q, u, p).Return(nil)
// 			},
// 			mockBehaviorDelete: func(d *mock_database.MockDatabase, q string, u, p uint) {
// 				d.EXPECT().Exec(context.Background(), q, u, p).Return(mock_database.MockSQLResult{}, sql.ErrNoRows)
// 			},
// 			callMockDelete: true,
// 			expectedCount:  0,
// 			expectedError:  models.ErrNoProduct,
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			db := mock_database.NewMockDatabase(ctrl)
// 			defer ctrl.Finish()

// 			q1 := `UPDATE cart_item SET count = cart_item.count - 1
// 	WHERE profile_id = $1 AND product_id = $2
// 	returning cart_item.count AS count;`

// 			testItem := dao.CartItemTable{}
// 			testCase.mockBehaviorUpdate(db, &testItem, q1, testCase.uID, testCase.prID)
// 			if testCase.callMockDelete {
// 				q2 := `DELETE FROM cart_item WHERE profile_id = $1 AND product_id = $2;`
// 				testCase.mockBehaviorDelete(db, q2, testCase.uID, testCase.prID)
// 			}
// 			cr := repository.NewCartRepo(db)

// 			newCount, err := cr.DeleteCartItem(context.Background(), testCase.uID, testCase.prID)
// 			require.Equal(t, testCase.expectedCount, newCount)
// 			require.ErrorIs(t, err, testCase.expectedError)
// 		})
// 	}
// }

func TestDeleteAllCartItems(t *testing.T) {
	testCases := []struct {
		name          string
		uID           uint
		mockBehavior  func(*mock_database.MockDatabase, string, uint)
		expectedError error
	}{
		{
			name: "Test successful cart clear",
			uID:  1,
			mockBehavior: func(d *mock_database.MockDatabase, q string, u uint) {
				d.EXPECT().Exec(context.Background(), q, u).Return(mock_database.MockSQLResult{
					LastInsertedID: 1,
					RowsAffect:     1,
				}, nil)
			},
			expectedError: nil,
		},
		{
			name: "Test empty cart clear",
			uID:  1,
			mockBehavior: func(d *mock_database.MockDatabase, q string, u uint) {
				d.EXPECT().Exec(context.Background(), q, u).Return(mock_database.MockSQLResult{}, sql.ErrNoRows)
			},
			expectedError: models.ErrEmptyCart,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mock_database.NewMockDatabase(ctrl)
			defer ctrl.Finish()

			q := `DELETE FROM cart_item WHERE profile_id = $1;`
			testCase.mockBehavior(db, q, testCase.uID)
			cr := repository.NewCartRepo(db)

			err := cr.DeleteAllCartItems(context.Background(), testCase.uID)
			require.ErrorIs(t, err, testCase.expectedError)
		})
	}
}

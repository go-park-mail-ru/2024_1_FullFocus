package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
)

type CartRepo struct {
	storage db.Database
}

func NewCartRepo(dbClient db.Database) *CartRepo {
	return &CartRepo{
		storage: dbClient,
	}
}

func (r *CartRepo) GetAllCartItems(ctx context.Context, uID uint) ([]models.CartProduct, error) {
	q := `SELECT p.id, p.product_name, p.price, p.imgsrc, c.count
	FROM product AS p JOIN cart_item AS c ON p.id = c.product_id
	WHERE profile_id = $1;`

	cartProductRows := []dao.CartProductTable{}
	err := r.storage.Select(ctx, &cartProductRows, q, uID)
	if err != nil {
		logger.Error(ctx, "Error: %s", err.Error())
		return nil, models.ErrEmptyCart
	}
	if len(cartProductRows) == 0 {
		logger.Info(ctx, "users cart is empty")
		return nil, models.ErrEmptyCart
	}
	return dao.ConvertTablesToCartProducts(cartProductRows), nil
}

func (r *CartRepo) GetAllCartItemsID(ctx context.Context, uID uint) ([]models.CartItem, error) {
	q := `SELECT product_id, count FROM cart_item WHERE profile_id = $1;`

	cartItemRows := []dao.CartItemTable{}
	err := r.storage.Select(ctx, &cartItemRows, q, uID)
	if err != nil {
		logger.Error(ctx, "Error: %s", err.Error())
		return nil, models.ErrEmptyCart
	}
	if len(cartItemRows) == 0 {
		logger.Info(ctx, "users cart is empty")
		return nil, models.ErrEmptyCart
	}
	return dao.ConvertTablesToCartItems(cartItemRows), nil
}

func (r *CartRepo) UpdateCartItem(ctx context.Context, uID, prID uint) (uint, error) {
	q := `INSERT INTO cart_item(profile_id, product_id) VALUES($1, $2)
	ON CONFLICT (profile_id, product_id)
	DO UPDATE set count = cart_item.count + 1
	returning cart_item.count AS count;`

	resRow := dao.CartItemTable{}
	err := r.storage.Get(ctx, &resRow, q, uID, prID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Info(ctx, "users cart is empty")
			return 0, models.ErrNoProduct
		}
		logger.Error(ctx, "Error: %s", err.Error())
		// TODO ErrNoProduct
		return 0, models.ErrEmptyCart
	}
	return resRow.Count, nil
}

func (r *CartRepo) DeleteCartItem(ctx context.Context, uID, prID uint) (uint, error) {
	q := `UPDATE cart_item SET count = cart_item.count - 1
	WHERE profile_id = $1 AND product_id = $2
	returning cart_item.count AS count;`

	resRow := dao.CartItemTable{}
	err := r.storage.Get(ctx, &resRow, q, uID, prID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Info(ctx, "users cart is empty")
			return 0, models.ErrNoProduct
		}
		logger.Error(ctx, "Error: %s", err.Error())
		// TODO ErrNoProduct
		return 0, models.ErrEmptyCart
	}

	if resRow.Count == 0 {
		q = `DELETE FROM cart_item WHERE profile_id = $1 AND product_id = $2;`

		if _, err = r.storage.Exec(ctx, q, uID, prID); err != nil {
			return 0, models.ErrNoProduct
		}
		return 0, nil
	}
	return resRow.Count, nil
}

func (r *CartRepo) DeleteAllCartItems(ctx context.Context, uID uint) error {
	q := `DELETE FROM cart_item WHERE profile_id = $1;`

	res, err := r.storage.Exec(ctx, q, uID)
	if err != nil {
		logger.Error(ctx, "Error: %s", err.Error())
		return models.ErrEmptyCart
	}
	if count, _ := res.RowsAffected(); count == 0 {
		logger.Error(ctx, "users cart is empty")
		return models.ErrEmptyCart
	}
	return nil
}

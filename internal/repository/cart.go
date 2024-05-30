package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
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
	q := `SELECT p.id, p.product_name, p.price, p.imgsrc, p.on_sale, c.count
	FROM product AS p JOIN cart_item AS c ON p.id = c.product_id
	WHERE profile_id = $1;`

	var cartProductRows []dao.CartProductTable
	err := r.storage.Select(ctx, &cartProductRows, q, uID)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, models.ErrEmptyCart
	}
	if len(cartProductRows) == 0 {
		return nil, models.ErrEmptyCart
	}
	return dao.ConvertTablesToCartProducts(cartProductRows), nil
}

func (r *CartRepo) GetAllCartItemsInfo(ctx context.Context, uID uint) ([]models.CartItem, error) {
	q := `SELECT product_id, count, p.price, p.on_sale
	FROM cart_item
	JOIN product p ON cart_item.product_id = p.id
	WHERE profile_id = $1;`

	var cartItemRows []dao.CartItem
	err := r.storage.Select(ctx, &cartItemRows, q, uID)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, models.ErrEmptyCart
	}
	if len(cartItemRows) == 0 {
		return nil, models.ErrEmptyCart
	}
	return dao.ConvertTablesToCartItems(cartItemRows), nil
}

func (r *CartRepo) GetCartItemsAmount(ctx context.Context, uID uint) (uint, error) {
	q := `SELECT count(*) FROM cart_item ci WHERE ci.profile_id = ?;`

	var amount uint
	if err := r.storage.Get(ctx, &amount, q, uID); err != nil {
		logger.Error(ctx, err.Error())
		return 0, models.ErrNoProfile
	}
	return amount, nil
}

func (r *CartRepo) UpdateCartItem(ctx context.Context, uID, prID uint) (uint, error) {
	q := `INSERT INTO cart_item(profile_id, product_id) VALUES($1, $2)
	ON CONFLICT (profile_id, product_id)
	DO UPDATE set count = cart_item.count + 1
	returning cart_item.count AS count;`

	resRow := dao.CartItem{}
	err := r.storage.Get(ctx, &resRow, q, uID, prID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrNoProduct
		}
		logger.Error(ctx, err.Error())
		// TODO ErrNoProduct
		return 0, models.ErrEmptyCart
	}
	return resRow.Count, nil
}

func (r *CartRepo) DeleteCartItem(ctx context.Context, uID, prID uint) (uint, error) {
	q := `UPDATE cart_item SET count = cart_item.count - 1
	WHERE profile_id = $1 AND product_id = $2
	returning cart_item.count AS count;`

	resRow := dao.CartItem{}
	err := r.storage.Get(ctx, &resRow, q, uID, prID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrNoProduct
		}
		logger.Error(ctx, err.Error())
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
		logger.Error(ctx, err.Error())
		return models.ErrEmptyCart
	}
	if count, _ := res.RowsAffected(); count == 0 {
		return models.ErrEmptyCart
	}
	return nil
}

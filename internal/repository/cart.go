package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

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

	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", uID)))
	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("queried in %s", time.Since(start)))
	}()

	cartProductRows := []dao.CartProductTable{}
	err := r.storage.Select(ctx, &cartProductRows, q, uID)
	if errors.Is(err, sql.ErrNoRows) {
		logger.Error(ctx, "users cart is empty")
		return nil, models.ErrEmptyCart
	}

	return dao.ConvertTablesToCartProducts(cartProductRows), nil
}

func (r *CartRepo) GetAllCartItemsID(ctx context.Context, uID uint) ([]models.CartItem, error) {
	q := `SELECT product_id, count FROM cart_item WHERE profile_id = $1;`

	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", uID)))
	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("queried in %s", time.Since(start)))
	}()

	cartItemRows := []dao.CartItemTable{}
	if err := r.storage.Select(ctx, &cartItemRows, q, uID); errors.Is(err, sql.ErrNoRows) {
		logger.Error(ctx, "users cart is empty")
		return nil, models.ErrEmptyCart
	}

	return dao.ConvertTablesToCartItems(cartItemRows), nil
}

func (r *CartRepo) UpdateCartItem(ctx context.Context, uID, prID uint) (uint, error) {
	q := `INSERT INTO cart_item(profile_id, product_id) VALUES($1, $2)
	ON CONFLICT (profile_id, product_id)
	DO UPDATE set count = cart_item.count + 1
	returning cart_item.count AS count;`

	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d $2 = %d", uID, prID)))
	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("updated in %s", time.Since(start)))
	}()

	resRow := dao.CartItemTable{}
	err := r.storage.Get(ctx, &resRow, q, uID, prID)
	if err != nil {
		return 0, models.ErrNoProduct
	}
	return resRow.Count, nil
}

func (r *CartRepo) DeleteCartItem(ctx context.Context, uID, prID uint) (uint, error) {
	q := `UPDATE cart_item SET count = cart_item.count - 1
	WHERE profile_id = $1 AND product_id = $2
	returning cart_item.count AS count;`

	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d $2 = %d", uID, prID)))
	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("updated in %s", time.Since(start)))
	}()

	resRow := dao.CartItemTable{}
	err := r.storage.Get(ctx, &resRow, q, uID, prID)
	if err != nil {
		logger.Info(ctx, fmt.Sprintf("\n\nLEVEL GET %s\n\n", err.Error()))
		return 0, models.ErrNoProduct
	}

	if resRow.Count == 0 {
		q = `DELETE FROM cart_item WHERE profile_id = $1 AND product_id = $2;`

		logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d $2 = %d", uID, prID)))
		start = time.Now()
		defer func() {
			logger.Info(ctx, fmt.Sprintf("deleted in %s", time.Since(start)))
		}()

		if _, err = r.storage.Exec(ctx, q, uID, prID); err != nil {
			return 0, models.ErrNoProduct
		}
		return 0, nil
	}
	return resRow.Count, nil
}

func (r *CartRepo) DeleteAllCartItems(ctx context.Context, uID uint) error {
	q := `DELETE FROM cart_item WHERE profile_id = $1;`

	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", uID)))
	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("deleted in %s", time.Since(start)))
	}()

	if _, err := r.storage.Exec(ctx, q, uID); err != nil {
		return models.ErrEmptyCart
	}

	return nil
}

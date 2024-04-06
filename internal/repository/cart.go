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
)

type CartRepo struct {
	storage db.Database
}

func NewCartRepo(dbClient db.Database) *CartRepo {
	return &CartRepo{
		storage: dbClient,
	}
}

func (r *CartRepo) GetAllCartItems(ctx context.Context, uID uint) ([]models.CartItem, error) {
	return nil, nil
}

func (r *CartRepo) GetAllCartProductsId(ctx context.Context, uID uint) ([]models.CartItem, error) {
	q := `SELECT (product_id, count) FROM cart_item WHERE profile_id = $1;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", uID)))
	start := time.Now()
	defer func() {
		logger.Info(ctx, fmt.Sprintf("queried in %s", time.Since(start)))
	}()

	cartItemRows := []db.CartItemTable{}
	if err := r.storage.Select(ctx, &cartItemRows, q, uID); errors.Is(err, sql.ErrNoRows) {
		logger.Error(ctx, "users cart is empty")
		return nil, models.ErrEmptyCart
	}

	return db.ConvertTablesToCartItems(cartItemRows), nil
}

func (r *CartRepo) UpdateCartItem(ctx context.Context, uID, prID uint) (uint, error) {
	return 0, nil
}

func (r *CartRepo) DeleteCartItem(ctx context.Context, uID, orID uint) (uint, error) {
	return 0, nil
}

func (r *CartRepo) DeleteAllCartItems(ctx context.Context, uID uint) error {
	return nil
}

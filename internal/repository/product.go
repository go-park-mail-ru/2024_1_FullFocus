package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
)

const (
	selectAllProductsQuery = `SELECT `
)

type ProductRepo struct {
	storage db.Database
}

func NewProductRepo(dbClient db.Database) *ProductRepo {
	return &ProductRepo{
		storage: dbClient,
	}
}

func (r *ProductRepo) GetAllProducts(ctx context.Context, lastID, limit int) ([]models.Product, error) {
	logger.Info(ctx, createOrderingQuery, slog.String("args", fmt.Sprintf("$1 = %d $2 = %s", userID, "created")))
	start := time.Now()
	result, err := r.storage.Exec(ctx, createOrderingQuery, userID, "created")
	if err != nil {
		logger.Error(ctx, "error while creating order: "+err.Error())
		return 0, err
	}
	logger.Info(ctx, fmt.Sprintf("inserted in %s", time.Since(start)))
}

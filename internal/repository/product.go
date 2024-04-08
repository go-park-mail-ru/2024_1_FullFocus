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

type ProductRepo struct {
	storage db.Database
}

func NewProductRepo(dbClient db.Database) *ProductRepo {
	return &ProductRepo{
		storage: dbClient,
	}
}

func (r *ProductRepo) GetAllProductCards(ctx context.Context, pageNum uint, perPage uint) ([]models.ProductCard, error) {
	q := `SELECT p.id, p.product_name, p.product_description, p.price, p.imgsrc, p.seller, p.rating
		  FROM ozon.product p
		  WHERE p.id - (SELECT MIN(id) from ozon.product) < 20
		  ORDER BY p.id DESC LIMIT 10;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d $2 = %d", pageNum, perPage)))
	start := time.Now()
	var products []db.ProductCard

	if err := r.storage.Get(ctx, &products, q, pageNum*perPage, perPage); err != nil {
		logger.Error(ctx, "error while selecting products: "+err.Error())
		return nil, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return db.ConvertProductCardsFromTable(products), nil
}

func (r *ProductRepo) GetProductById(ctx context.Context, productID uint) (models.Product, error) {
	q := `SELECT p.id, p.product_description, p.product_name, p.price, p.imgsrc, p.seller, p.rating
		  FROM ozon.product p
		  WHERE p.id = ?;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", productID)))
	start := time.Now()
	var product db.Product
	if err := r.storage.Get(ctx, &product, q, productID); err != nil {
		logger.Error(ctx, "error while selecting product: "+err.Error())
		return models.Product{}, err
	} else if product.ID == 0 {
		logger.Error(ctx, "no product found")
		return models.Product{}, models.ErrNoProduct
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))

	q = `SELECT c.category_name
		  FROM ozon.product_category pc
    	  	  INNER JOIN ozon.category c ON c.id = pc.category_id
		  WHERE pc.product_id = ?;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", productID)))
	start = time.Now()
	var categories []string
	if err := r.storage.Get(ctx, &categories, q, productID); err != nil {
		logger.Error(ctx, "error while selecting categories: "+err.Error())
		return models.Product{}, err
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return db.ConvertProductFromTable(categories, product), nil
}

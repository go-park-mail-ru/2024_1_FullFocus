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

// OK

func (r *ProductRepo) GetAllProductCards(ctx context.Context, input models.GetAllProductsInput) ([]models.ProductCard, error) {
	q := `WITH products_info AS (
			SELECT p.id, p.product_name, p.price, p.imgsrc, p.seller, p.rating,
				   CASE
					   WHEN cart_query.in_cart IS NULL THEN 0
					   ELSE 1
				   END AS in_cart
			FROM product p
				LEFT JOIN (
					SELECT i.product_id, i.profile_id AS in_cart
					FROM cart_item i
					WHERE i.profile_id = ?
			) cart_query ON p.id = cart_query.product_id
		 )
		 SELECT * FROM products_info pi
		 WHERE pi.id - (SELECT MIN(id) from product) < ?
		 ORDER BY pi.id DESC LIMIT ?;`
	offset := input.PageNum * input.PageSize
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d $2 = %d $3 = %d", input.ProfileID, offset, input.PageSize)))
	start := time.Now()
	var products []db.ProductCard
	if err := r.storage.Select(ctx, &products, q, input.ProfileID, offset, input.PageSize); err != nil {
		logger.Info(ctx, "error while selecting: "+err.Error())
		return nil, models.ErrNoRowsFound
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return db.ConvertProductCardsFromTable(products), nil
}

// OK

func (r *ProductRepo) GetProductById(ctx context.Context, profileID uint, productID uint) (models.Product, error) {
	q := `SELECT id, product_description, product_name, price, imgsrc, seller, rating,
   			CASE
       			WHEN ci.product_id IS NULL THEN 0
       			ELSE 1
    		END AS in_cart
		  FROM (SELECT *
          		FROM product p
      			WHERE p.id = ?
		  ) subquery
    		LEFT JOIN cart_item ci ON ci.product_id = subquery.id AND ci.profile_id = ?;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d $2 = %d", productID, profileID)))
	start := time.Now()
	var product db.Product
	if err := r.storage.Get(ctx, &product, q, productID, profileID); err != nil {
		logger.Error(ctx, "error while selecting product: "+err.Error())
		return models.Product{}, models.ErrNoRowsFound
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))

	q = `SELECT c.category_name
		  FROM product_category pc
    	  	  INNER JOIN category c ON c.id = pc.category_id
		  WHERE pc.product_id = ?;`
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d", productID)))
	start = time.Now()
	var categories []string
	if err := r.storage.Select(ctx, &categories, q, productID); err != nil {
		logger.Info(ctx, "error while selecting categories: "+err.Error())
		return models.Product{}, models.ErrNoRowsFound
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return db.ConvertProductFromTable(categories, product), nil
}

func (r *ProductRepo) GetProductsByCategoryId(ctx context.Context, input models.GetProductsByCategoryIDInput) ([]models.ProductCard, error) {
	q := `WITH products_info AS (
				SELECT p.id, p.product_name, p.price, p.imgsrc, p.seller, p.rating,
					   CASE
						   WHEN cart_query.in_cart IS NULL THEN 0
						   ELSE 1
						   END AS in_cart
				FROM product p
						 INNER JOIN (
					SELECT p.id
					FROM product p
							 INNER JOIN product_category pc ON pc.product_id = p.id
							 INNER JOIN category c ON c.id = pc.category_id
					WHERE pc.category_id = ?
				) subquery ON p.id = subquery.id
						 LEFT JOIN (
					SELECT i.product_id, i.profile_id AS in_cart
					FROM cart_item i
					WHERE i.profile_id = ?
				) cart_query ON p.id = cart_query.product_id
			)
			SELECT * FROM products_info pi
			OFFSET ? LIMIT ?;`
	offset := (input.PageNum - 1) * input.PageSize
	logger.Info(ctx, q, slog.String("args", fmt.Sprintf("$1 = %d $2 = %d $3 = %d $4 = %d", input.CategoryID, input.ProfileID, offset, input.PageSize)))
	start := time.Now()
	var products []db.ProductCard
	if err := r.storage.Select(ctx, &products, q, input.CategoryID, input.ProfileID, offset, input.PageSize); err != nil {
		logger.Info(ctx, "error while selecting: "+err.Error())
		return nil, models.ErrNoRowsFound
	}
	logger.Info(ctx, fmt.Sprintf("selected in %s", time.Since(start)))
	return db.ConvertProductCardsFromTable(products), nil
}

package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type ProductRepo struct {
	storage db.Database
}

func NewProductRepo(dbClient db.Database) *ProductRepo {
	return &ProductRepo{
		storage: dbClient,
	}
}

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
		 ORDER BY pi.id DESC
		 OFFSET ?
		 LIMIT ?;`
	offset := input.PageNum * input.PageSize
	var products []dao.ProductCard
	if err := r.storage.Select(ctx, &products, q, input.ProfileID, offset, input.PageSize); err != nil {
		logger.Info(ctx, "error while selecting: "+err.Error())
		return nil, models.ErrNoRowsFound
	}
	return dao.ConvertProductCardsFromTable(products), nil
}

func (r *ProductRepo) GetProductByID(ctx context.Context, profileID uint, productID uint) (models.Product, error) {
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
	var product dao.Product
	if err := r.storage.Get(ctx, &product, q, productID, profileID); err != nil {
		logger.Error(ctx, "error while selecting product: "+err.Error())
		return models.Product{}, models.ErrNoRowsFound
	}

	q = `SELECT c.category_name
		  FROM product_category pc
    	  	  INNER JOIN category c ON c.id = pc.category_id
		  WHERE pc.product_id = ?;`
	var categories []string
	if err := r.storage.Select(ctx, &categories, q, productID); err != nil {
		logger.Info(ctx, "error while selecting categories: "+err.Error())
		return models.Product{}, models.ErrNoRowsFound
	}
	return dao.ConvertProductFromTable(categories, product), nil
}

func (r *ProductRepo) GetProductsByCategoryID(ctx context.Context, input models.GetProductsByCategoryIDInput) ([]models.ProductCard, error) {
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
			OFFSET ?
			LIMIT ?;`
	offset := (input.PageNum - 1) * input.PageSize
	var products []dao.ProductCard
	if err := r.storage.Select(ctx, &products, q, input.CategoryID, input.ProfileID, offset, input.PageSize); err != nil {
		logger.Info(ctx, "error while selecting: "+err.Error())
		return nil, models.ErrNoRowsFound
	}
	return dao.ConvertProductCardsFromTable(products), nil
}

func (r *ProductRepo) GetProductsByQuery(ctx context.Context, input models.GetProductsByQueryInput) ([]models.ProductCard, error) {
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
		WHERE pi.product_name ILIKE '%%%s%%'
		ORDER BY pi.product_name
		OFFSET ?
		LIMIT ?;`
	offset := (input.PageNum - 1) * input.PageSize
	var products []dao.ProductCard
	if err := r.storage.Select(ctx, &products, fmt.Sprintf(q, input.Query), input.ProfileID, offset, input.PageSize); err != nil {
		logger.Info(ctx, "error while selecting: "+err.Error())
		return nil, models.ErrNoRowsFound
	}
	return dao.ConvertProductCardsFromTable(products), nil
}

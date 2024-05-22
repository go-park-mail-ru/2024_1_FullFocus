package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
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
	q := `SELECT p.id, p.product_name, p.price, p.imgsrc, p.seller, p.rating, COALESCE(ci.count, 0) AS count
			FROM product p
				LEFT JOIN cart_item ci
					ON p.id = ci.product_id
					   AND ci.profile_id = ?
			WHERE on_sale = FALSE
			%s
			OFFSET ?
			LIMIT ?;`
	offset := (input.PageNum - 1) * input.PageSize
	var products []dao.ProductCard
	q = helper.ApplySorting(q, input.Sorting.QueryPart)
	if err := r.storage.Select(ctx, &products, q, input.ProfileID, offset, input.PageSize); err != nil {
		logger.Info(ctx, err.Error())
		return nil, models.ErrNoRowsFound
	}
	return dao.ConvertProductCardsFromTable(products), nil
}

func (r *ProductRepo) GetProductByID(ctx context.Context, profileID uint, productID uint) (models.Product, error) {
	q := `SELECT p.id, p.product_description, p.product_name, p.price, p.imgsrc, p.seller, p.rating, COALESCE(ci.count, 0) AS count
			FROM product p
				 LEFT JOIN cart_item ci ON ci.product_id = p.id AND ci.profile_id = ?
			WHERE p.id = ?;`

	var product dao.Product
	if err := r.storage.Get(ctx, &product, q, profileID, productID); err != nil {
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

func (r *ProductRepo) GetProductCardByID(ctx context.Context, productID uint) (models.ProductCard, error) {
	q := `SELECT id, product_description, product_name, price, imgsrc, seller, rating
	FROM product
	WHERE id = ?;`

	var data dao.ProductCard
	if err := r.storage.Get(ctx, &data, q, productID); err != nil {
		logger.Error(ctx, err.Error())
		return models.ProductCard{}, models.ErrNoProduct
	}
	return dao.ConvertProductCardToModel(data), nil
}

func (r *ProductRepo) GetProductsByCategoryID(ctx context.Context, input models.GetProductsByCategoryIDInput) ([]models.ProductCard, error) {
	q := `SELECT p.id, p.product_name, p.price, p.imgsrc, p.seller, p.rating, COALESCE(ci.count, 0) AS count
			FROM product p
    			INNER JOIN product_category pc
        			ON p.id = pc.product_id
        			AND pc.category_id = ?
				LEFT JOIN cart_item ci
					ON p.id = ci.product_id
				   	AND ci.profile_id = ?
			WHERE on_sale = FALSE
			%s
			OFFSET ?
			LIMIT ?;`
	offset := (input.PageNum - 1) * input.PageSize
	var products []dao.ProductCard
	q = helper.ApplySorting(q, input.Sorting.QueryPart)
	if err := r.storage.Select(ctx, &products, q, input.CategoryID, input.ProfileID, offset, input.PageSize); err != nil {
		logger.Info(ctx, "error while selecting: "+err.Error())
		return nil, models.ErrNoRowsFound
	}
	return dao.ConvertProductCardsFromTable(products), nil
}

func (r *ProductRepo) GetProductsByQuery(ctx context.Context, input models.GetProductsByQueryInput) ([]models.ProductCard, error) {
	q := `SELECT p.id, p.product_name, p.price, p.imgsrc, p.seller, p.rating, COALESCE(ci.count, 0) AS count
			FROM product p
    			LEFT JOIN cart_item ci
        			ON p.id = ci.product_id
        			AND ci.profile_id = ?
		  WHERE p.product_name ILIKE '%%%s%%' AND on_sale = FALSE`
	q1 := `%s
		OFFSET ?
		LIMIT ?;`
	offset := (input.PageNum - 1) * input.PageSize
	var products []dao.ProductCard
	q = fmt.Sprintf(q, input.Query)
	q1 = helper.ApplySorting(q1, input.Sorting.QueryPart)
	q = q + q1
	if err := r.storage.Select(ctx, &products, q, input.ProfileID, offset, input.PageSize); err != nil {
		logger.Info(ctx, "error while selecting: "+err.Error())
		return nil, models.ErrNoRowsFound
	}
	return dao.ConvertProductCardsFromTable(products), nil
}

func (r *ProductRepo) GetProductsByIDs(ctx context.Context, IDs []uint) ([]models.ProductCard, error) {
	q := `SELECT id, product_name, price, imgsrc, seller, rating
	FROM product
	WHERE id = ANY(?)
	ORDER BY array_position(?, id);`

	products := make([]dao.ProductCard, 0, len(IDs))
	if err := r.storage.Select(ctx, &products, q, IDs, IDs); err != nil {
		logger.Error(ctx, err.Error())
		return nil, models.ErrInternal
	}
	if len(products) == 0 {
		return nil, models.ErrNoProduct
	}
	return dao.ConvertProductCardsFromTable(products), nil
}

func (r *ProductRepo) GetProductPriceByID(ctx context.Context, ID uint) (uint, error) {
	q := `SELECT price FROM product WHERE id = ?;`

	var price uint
	if err := r.storage.Get(ctx, &price, q, ID); err != nil {
		logger.Error(ctx, err.Error())
		return 0, models.ErrNoProduct
	}
	return price, nil
}

func (r *ProductRepo) MarkProduct(ctx context.Context, ID uint, promo bool) error {
	q := `UPDATE product SET on_sale = ? WHERE id = ?;`

	if _, err := r.storage.Exec(ctx, q, promo, ID); err != nil {
		logger.Error(ctx, err.Error())
		return models.ErrInternal
	}
	return nil
}

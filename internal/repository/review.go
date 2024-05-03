package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type ReviewRepo struct {
	storage database.Database
}

func NewReviewRepo(st database.Database) *ReviewRepo {
	return &ReviewRepo{
		storage: st,
	}
}

func (r *ReviewRepo) GetProductReviews(ctx context.Context, input models.GetProductReviewsInput) ([]models.ProductReview, error) {
	q := `SELECT p.full_name , p.imgsrc, r.rating, DATE(r.created_at) AS created_at, r.comments, r.advantages, r.disadvantages
	FROM review r
	JOIN user_profile p ON r.profile_id = p.id
	WHERE r.product_id = $1
	ORDER BY r.created_at 
	LIMIT $2 OFFSET $3;`

	reviews := make([]dao.ProductReviewTable, 0)
	if err := r.storage.Select(ctx, &reviews, q, input.ProductID, input.PageSize, input.LastReviewID); err != nil {
		logger.Info(ctx, "Error:"+err.Error())
		return []models.ProductReview{}, models.ErrInternal
	}
	if len(reviews) == 0 {
		return []models.ProductReview{}, models.ErrNoReviews
	}
	return dao.ConvertReviewsToModels(reviews), nil
}

func (r *ReviewRepo) CreateProductReview(ctx context.Context, uID uint, input models.ProductReview) error {
	q := `INSERT INTO review(profile_id, product_id, rating, comments, advantages, disadvantages)
	VALUES($1, $2, $3, $4, $5, $6);`

	if _, err := r.storage.Exec(ctx, q, uID, input.ProductID, input.Rating, input.Comment, input.Advanatages, input.Disadvantages); err != nil {
		logger.Info(ctx, "Error:"+err.Error())
		var sqlErr *pgconn.PgError
		if errors.As(err, &sqlErr) {
			if strings.Contains(sqlErr.Message, "duplicate") {
				return models.ErrReviewAlreadyExists
			}
			if strings.Contains(sqlErr.Message, "foreign") {
				return models.ErrNoProduct
			}
		}
		return models.ErrInternal
	}
	return nil
}

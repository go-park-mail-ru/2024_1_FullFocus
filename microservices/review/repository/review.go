package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"

	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/review/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/review/repository/dao"
	commonError "github.com/go-park-mail-ru/2024_1_FullFocus/pkg/error"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
)

type ReviewRepo struct {
	storage db.Database
}

func NewReviewRepo(st db.Database) *ReviewRepo {
	return &ReviewRepo{
		storage: st,
	}
}

func (r *ReviewRepo) GetProductReviews(ctx context.Context, input models.GetProductReviewsInput) ([]models.ProductReviewData, error) {
	q := `SELECT id, profile_id, rating, DATE(created_at) AS created_at, comments, advantages, disadvantages
	FROM review
	WHERE product_id = $1
	ORDER BY created_at 
	LIMIT $2 OFFSET $3;`

	reviews := make([]dao.ProductReviewTable, 0)
	if err := r.storage.Select(ctx, &reviews, q, input.ProductID, input.Limit, input.LastReviewID); err != nil {
		logger.Info(ctx, "Error:"+err.Error())
		return nil, commonError.ErrInternal
	}
	if len(reviews) == 0 {
		return nil, models.ErrNoReviews
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
		return commonError.ErrInternal
	}
	return nil
}

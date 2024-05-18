package repository

import (
	"context"
	"errors"
	"strings"

	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/repository/dao"
	commonError "github.com/go-park-mail-ru/2024_1_FullFocus/pkg/error"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repo struct {
	storage db.Database
}

func NewRepo(st db.Database) *Repo {
	return &Repo{
		storage: st,
	}
}

func (r *Repo) CreatePromoProductInfo(ctx context.Context, input models.PromoData) error {
	q := `INSERT INTO promo_product(product_id, benefit_type, benefit_value) VALUES(?, ?, ?);`

	if _, err := r.storage.Exec(ctx, q, input.ProductID, input.BenefitType, input.BenefitValue); err != nil {
		logger.Info(ctx, "Error:"+err.Error())
		var sqlErr *pgconn.PgError
		if errors.As(err, &sqlErr) {
			if strings.Contains(sqlErr.Message, "duplicate") {
				return models.ErrPromoProductAlreadyExists
			}
			if strings.Contains(sqlErr.Message, "foreign") {
				return models.ErrProductNotFound
			}
		}
		return commonError.ErrInternal
	}
	return nil
}

func (r *Repo) GetPromoProductsInfo(ctx context.Context, amount uint32) ([]models.PromoData, error) {
	q := `SELECT product_id, benefit_type, benefit_value FROM promo_product LIMIT ?;`

	promoData := make([]dao.PromoProductTable, 0, amount)
	if err := r.storage.Select(ctx, promoData, q, amount); err != nil {
		logger.Info(ctx, "Error:"+err.Error())
		return nil, commonError.ErrInternal
	}
	if len(promoData) == 0 {
		return nil, models.ErrProductNotFound
	}
	return dao.ConvertPromoProductTablesToModels(promoData), nil
}

func (r *Repo) DeletePromoProductInfo(ctx context.Context, pID uint32) error {
	q := `DELETE FROM promo_product WHERE product_id = ?;`

	//TODO not found
	if _, err := r.storage.Exec(ctx, q, pID); err != nil {
		logger.Info(ctx, "Error:"+err.Error())
		return commonError.ErrInternal
	}
	return nil
}

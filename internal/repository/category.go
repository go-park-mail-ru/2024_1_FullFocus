package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	db "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
)

type CategoryRepo struct {
	storage db.Database
}

func NewCategoryRepo(dbClient db.Database) *CategoryRepo {
	return &CategoryRepo{
		storage: dbClient,
	}
}

func (r *CategoryRepo) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	q := `SELECT id, category_name FROM category WHERE parent_id IS NULL;`

	categoryRows := []dao.CategoryTable{}
	err := r.storage.Select(ctx, &categoryRows, q)
	if err != nil {
		logger.Error(ctx, "Error: %s", err.Error())
		return nil, models.ErrInternal
	}
	return dao.ConvertTablesToCategories(categoryRows), nil
}
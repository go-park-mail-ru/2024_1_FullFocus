package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type CategoryUsecase struct {
	categoryRepo repository.Categories
}

func NewCategoryUsecase(cr repository.Categories) *CategoryUsecase {
	return &CategoryUsecase{
		categoryRepo: cr,
	}
}

func (u *CategoryUsecase) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	return u.categoryRepo.GetAllCategories(ctx)
}

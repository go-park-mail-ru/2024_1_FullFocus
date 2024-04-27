package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type SuggestUsecase struct {
	suggestRepo repository.Suggests
}

func NewSuggestUsecase(sr repository.Suggests) *SuggestUsecase {
	return &SuggestUsecase{
		suggestRepo: sr,
	}
}

func (u *SuggestUsecase) GetSuggestions(ctx context.Context, query string) (models.Suggestion, error) {
	categories, err := u.suggestRepo.GetCategorySuggests(ctx, query)
	if err != nil {
		return models.Suggestion{}, err
	}
	products, err := u.suggestRepo.GetProductSuggests(ctx, query)
	if err != nil {
		return models.Suggestion{}, err
	}
	return models.Suggestion{
		Categories: categories,
		Products:   products,
	}, nil
}

package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type GetProductReviewsInput struct {
	ProductID    uint `json:"productID"`
	LastReviewID uint `json:"lastReviewID"`
	PageSize     uint `json:"limit"`
}

func ConvertGetProductReviewInputToModel(d GetProductReviewsInput) models.GetProductReviewsInput {
	return models.GetProductReviewsInput{
		ProductID:    d.ProductID,
		LastReviewID: d.LastReviewID,
		PageSize:     d.PageSize,
	}
}

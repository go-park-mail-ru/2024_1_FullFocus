package dto

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type PromocodeActivationTerms struct {
	ID               uint   `json:"id"`
	MinSumActivation uint   `json:"min_sum_activation"`
	BenefitType      string `json:"benefit_type"`
	Value            uint   `json:"value"`
}

func ConvertTerms(terms models.PromocodeActivationTerms) PromocodeActivationTerms {
	return PromocodeActivationTerms{
		ID:               terms.ID,
		MinSumActivation: terms.MinSumActivation,
		BenefitType:      terms.BenefitType,
		Value:            terms.Value,
	}
}

type PromocodeItem struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Code             string `json:"code"`
	MinSumActivation uint   `json:"minSumActivation"`
	BenefitType      string `json:"benefitType"`
	Value            uint   `json:"value"`
	TimeLeft         string `json:"timeLeft"`
}

func ConvertPromocodes(promos []models.PromocodeItem) []PromocodeItem {
	result := make([]PromocodeItem, 0, len(promos))
	for _, promo := range promos {
		result = append(result, PromocodeItem{
			ID:               promo.ID,
			Name:             promo.Name,
			Description:      promo.Description,
			Code:             promo.Code,
			MinSumActivation: promo.MinSumActivation,
			BenefitType:      promo.BenefitType,
			Value:            promo.Value,
			TimeLeft:         promo.TimeLeft,
		})
	}
	return result
}

type Promocode struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	MinSumGive       uint   `json:"minSumGive"`
	MinSumActivation uint   `json:"minSumActivation"`
	BenefitType      string `json:"benefitType"`
	Value            uint   `json:"value"`
	TTL              int    `json:"ttl"`
}

func ConvertPromocode(promo models.Promocode) Promocode {
	return Promocode{
		ID:               promo.ID,
		Name:             promo.Name,
		Description:      promo.Description,
		MinSumGive:       promo.MinSumGive,
		MinSumActivation: promo.MinSumActivation,
		BenefitType:      promo.BenefitType,
		Value:            promo.Value,
		TTL:              promo.TTL,
	}
}

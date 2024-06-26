package dao

import (
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type Promocode struct {
	ID               uint   `db:"id"`
	Description      string `db:"description"`
	MinSumGive       uint   `db:"min_sum_give"`
	MinSumActivation uint   `db:"min_sum_activation"`
	BenefitType      string `db:"benefit_type"`
	Value            uint   `db:"value"`
	TTL              int    `db:"ttl_hours"`
}

func ConvertPromocode(code Promocode) models.Promocode {
	return models.Promocode{
		ID:               code.ID,
		Description:      code.Description,
		MinSumGive:       code.MinSumGive,
		MinSumActivation: code.MinSumActivation,
		BenefitType:      code.BenefitType,
		Value:            code.Value,
		TTL:              code.TTL,
	}
}

type PromocodeItemInfo struct {
	ID               uint   `db:"id"`
	Description      string `db:"description"`
	TimeLeft         string `db:"time_left"`
	Code             string `db:"code"`
	MinSumActivation uint   `db:"min_sum_activation"`
	BenefitType      string `db:"benefit_type"`
	Value            uint   `db:"value"`
}

func ConvertPromocodeItem(item PromocodeItemInfo) models.PromocodeItem {
	return models.PromocodeItem{
		ID:               item.ID,
		Description:      item.Description,
		TimeLeft:         item.TimeLeft,
		Code:             item.Code,
		MinSumActivation: item.MinSumActivation,
		BenefitType:      item.BenefitType,
		Value:            item.Value,
	}
}

func ConvertPromocodeItems(item []PromocodeItemInfo) []models.PromocodeItem {
	promos := make([]models.PromocodeItem, len(item))
	for i := range item {
		promos[i] = ConvertPromocodeItem(item[i])
	}
	return promos
}

type PromocodeBenefit struct {
	MinSum    uint      `db:"min_sum_activation"`
	Type      string    `db:"benefit_type"`
	Value     uint      `db:"value"`
	ExpiresAt time.Time `db:"expires_at"`
}

type PromocodeActivationTerms struct {
	ID               uint      `db:"id"`
	Description      string    `db:"description"`
	ExpiresAt        time.Time `db:"expires_at"`
	MinSumActivation uint      `db:"min_sum_activation"`
	BenefitType      string    `db:"benefit_type"`
	Value            uint      `db:"value"`
}

func ConvertTerms(terms PromocodeActivationTerms) models.PromocodeActivationTerms {
	return models.PromocodeActivationTerms{
		ID:               terms.ID,
		Description:      terms.Description,
		MinSumActivation: terms.MinSumActivation,
		BenefitType:      terms.BenefitType,
		Value:            terms.Value,
	}
}

package dao

import (
	"time"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type Promocode struct {
	ID               uint      `db:"id"`
	Name             string    `db:"name"`
	Description      string    `db:"description"`
	MinSumGive       uint      `db:"min_sum_give"`
	MinSumActivation uint      `db:"min_sum_activation"`
	BenefitType      string    `db:"benefit_type"`
	Value            uint      `db:"value"`
	TimeLeft         time.Time `db:"ttl_hours"`
}

func ConvertPromocode(code Promocode) models.Promocode {
	return models.Promocode{
		ID:               code.ID,
		Name:             code.Name,
		Description:      code.Description,
		MinSumGive:       code.MinSumGive,
		MinSumActivation: code.MinSumActivation,
		BenefitType:      code.BenefitType,
		Value:            code.Value,
		TimeLeft:         code.TimeLeft,
	}
}

type PromocodeItemInfo struct {
	ID               uint      `db:"id"`
	Name             string    `db:"name"`
	Description      string    `db:"description"`
	TimeLeft         time.Time `db:"time_left"`
	Code             string    `db:"code"`
	MinSumActivation uint      `db:"min_sum_activation"`
	BenefitValue     string    `db:"benefit_value"`
	Value            uint      `db:"value"`
}

func ConvertPromocodeItem(item PromocodeItemInfo) models.PromocodeItem {
	return models.PromocodeItem{
		ID:               item.ID,
		Name:             item.Name,
		Description:      item.Description,
		TimeLeft:         item.TimeLeft,
		Code:             item.Code,
		MinSumActivation: item.MinSumActivation,
		BenefitType:      item.BenefitValue,
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
	Type  string `db:"benefit_type"`
	Value uint   `db:"value"`
}

type PromocodeActivationTerms struct {
	ID               uint      `db:"id"`
	ExpiresAt        time.Time `db:"expires_at"`
	MinSumActivation uint      `db:"min_sum_activation"`
	BenefitType      string    `db:"benefit_type"`
	Value            uint      `db:"value"`
}

func ConvertTerms(terms PromocodeActivationTerms) models.PromocodeActivationTerms {
	return models.PromocodeActivationTerms{
		ID:               terms.ID,
		MinSumActivation: terms.MinSumActivation,
		BenefitType:      terms.BenefitType,
		Value:            terms.Value,
	}
}

package dao

import "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/models"

type PromoProductTable struct {
	ProductID    uint   `db:"product_id"`
	BenefitType  string `db:"benefit_type"`
	BenefitValue uint   `db:"benefit_value"`
}

func ConvertPromoProductTableToModel(t PromoProductTable) models.PromoData {
	return models.PromoData{
		ProductID:    t.ProductID,
		BenefitType:  t.BenefitType,
		BenefitValue: t.BenefitValue,
	}
}

func ConvertPromoProductTablesToModels(tt []PromoProductTable) []models.PromoData {
	data := make([]models.PromoData, 0, len(tt))
	for _, t := range tt {
		data = append(data, ConvertPromoProductTableToModel(t))
	}
	return data
}

package models

import "time"

type PromoProduct struct {
	ProductData  ProductCard
	BenefitType  string
	BenefitValue uint
	Deadline     time.Time
}

type PromoData struct {
	ProductID    uint
	BenefitType  string
	BenefitValue uint
	Deadline     time.Time
}

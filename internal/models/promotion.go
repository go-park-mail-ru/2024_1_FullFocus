package models

import "time"

type PromoProduct struct {
	ProductData  Product
	BenefitType  string
	BenefitValue uint
	Deadline     time.Time
}

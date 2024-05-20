package models

type Promocode struct {
	ID               uint
	Name             string
	Description      string
	MinSumGive       uint
	MinSumActivation uint
	BenefitType      string
	Value            uint
	TTL              int
}

type PromocodeItem struct {
	ID               uint
	Name             string
	Description      string
	Code             string
	MinSumActivation uint
	BenefitType      string
	Value            uint
	TimeLeft         string
}

type CreatePromocodeItemInput struct {
	PromocodeID uint
	ProfileID   uint
	Code        string
}

type PromocodeActivationTerms struct {
	ID               uint
	MinSumActivation uint
	BenefitType      string
	Value            uint
}

type ApplyPromocodeInput struct {
	Sum       uint
	PromoID   uint
	ProfileID uint
}

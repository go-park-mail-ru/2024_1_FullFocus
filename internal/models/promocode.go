package models

import "time"

type Promocode struct {
	ID               uint
	Name             string
	Description      string
	MinSumGive       uint
	MinSumActivation uint
	BenefitType      string
	Value            uint
	TimeLeft         time.Time
}

type PromocodeItem struct {
	ID               uint
	Name             string
	Description      string
	Code             string
	MinSumActivation uint
	BenefitType      string
	Value            uint
	TimeLeft         time.Time
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

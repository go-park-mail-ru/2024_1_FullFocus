package models

type Poll struct {
	ID    uint
	Title string
}

type CreatePollRate struct {
	profileID uint
	pollID    uint
	rate      uint
}

package models

type PollStats struct {
	PollTitle string
	Stats     StatData
}

type StatData struct {
	Amount  uint
	Above70 uint
	Rates   []uint
}

type StatRate struct {
	Rate   uint
	Amount uint
}

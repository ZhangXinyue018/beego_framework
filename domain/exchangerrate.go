package domain

import "time"

type ExchangerRate struct {
	Id                  int       `xorm:"pk" json:"id"`
	Symbol              string    `json:"symbol"`
	Exchanger           string    `json:"exchanger"`
	Price               float64   `json:"price"`
	IsActivated         int       `json:"is_activated"`
	PercentageChange24H float64   `xorm:"'percentage_change_24h'" json:"percentage_change_24h"`
	UpdateTime          time.Time `json:"update_time"`
}

type SimpleExchangerRate struct {
	Symbol              string  `json:"symbol"`
	Exchanger           string  `json:"exchanger"`
	Price               float64 `json:"price"`
	PercentageChange24H float64 `json:"percentage_change_24h"`
}

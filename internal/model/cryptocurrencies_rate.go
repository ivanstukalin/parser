package model

import "time"

type CryptoRate struct {
	ID                       int       `json:"id"`
	Cryptocurrencies_rate_id int       `json:"cryptocurrencies_rate_id"`
	Rate                     string    `json:"rate"`
	CreatedAt                time.Time `json:"created_at"`
}

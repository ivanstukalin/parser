package provider

import (
	"context"
	"parser/internal/client"
	"parser/internal/model"
)

type PgProvider struct {
	client *client.PGClient
}

func NewPgProvider(client *client.PGClient) *PgProvider {
	return &PgProvider{client: client}
}

func (p *PgProvider) GetDomains(ctx context.Context) ([]model.Domain, error) {
	rows, err := p.client.Query("SELECT id, name, url, created_at FROM domains")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Domain
	for rows.Next() {
		var d model.Domain
		if err := rows.Scan(&d.ID, &d.Name, &d.URL, &d.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, d)
	}
	return result, nil
}

func (p *PgProvider) GetCryptos(ctx context.Context) ([]model.Crypto, error) {
	rows, err := p.client.Query("SELECT id, name, code, created_at FROM cryptocurrencies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Crypto
	for rows.Next() {
		var c model.Crypto
		if err := rows.Scan(&c.ID, &c.Name, &c.Code, &c.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}

func (p *PgProvider) GetCryptoRate(ctx context.Context, cryptocurrenciesID int) ([]model.CryptoRate, error) {
	rows, err := p.client.Query(`
        SELECT id, cryptocurrencies_rate_id, rate, created_at
        FROM cryptocurrencies_rate
        WHERE cryptocurrencies_rate_id = $1
        ORDER BY created_at DESC`,
		cryptocurrenciesID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rates []model.CryptoRate
	for rows.Next() {
		var cryptoRate model.CryptoRate
		if err := rows.Scan(&cryptoRate.ID, &cryptoRate.Cryptocurrencies_rate_id, &cryptoRate.Rate, &cryptoRate.CreatedAt); err != nil {
			return nil, err
		}
		rates = append(rates, cryptoRate)
	}

	return rates, nil
}

func (p *PgProvider) InsertCryptoRate(ctx context.Context, cryptocurrenciesID int, rate string) (*model.CryptoRate, error) {
	var newRate model.CryptoRate

	err := p.client.QueryRow(`
        INSERT INTO cryptocurrencies_rate(cryptocurrencies_rate_id, rate)
        VALUES ($1, $2)
        RETURNING id, cryptocurrencies_rate_id, rate, created_at`,
		cryptocurrenciesID,
		rate,
	).Scan(&newRate.ID, &newRate.Cryptocurrencies_rate_id, &newRate.Rate, &newRate.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &newRate, nil
}

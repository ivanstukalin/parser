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

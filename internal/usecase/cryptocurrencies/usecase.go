package cryptocurrencies

import (
	"context"
	"parser/internal/config"
	"parser/internal/model"
	"parser/internal/provider"
)

type CryptoUseCase struct {
	provider      *provider.PgProvider
	binanceAPIURL string
	config        *config.AppConfig
}

func NewCryptoUseCase(provider *provider.PgProvider, cfg *config.AppConfig) *CryptoUseCase {
	return &CryptoUseCase{provider: provider, binanceAPIURL: cfg.BinanceAPIURL, config: cfg}
}

type UseCase interface {
	GetAllCryptos(ctx context.Context) ([]model.Crypto, error)
	GetCryptoRate(ctx context.Context, code string) (*model.CryptoRate, error)
}

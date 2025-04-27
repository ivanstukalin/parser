package cryptocurrencies

import (
	"context"
	"parser/internal/model"
	"parser/internal/provider"
)

type CryptoUseCase struct {
	provider      *provider.PgProvider
	binanceAPIURL string
}

func NewCryptoUseCase(provider *provider.PgProvider, binanceAPIURL string) *CryptoUseCase {
	return &CryptoUseCase{provider: provider, binanceAPIURL: binanceAPIURL}
}

type UseCase interface {
	GetAllCryptos(ctx context.Context) ([]model.Crypto, error)
	GetCryptoRate(ctx context.Context, code string) (*model.CryptoRate, error)
}

package cryptocurrencies

import (
	"context"
	"parser/internal/model"
	"parser/internal/provider"
)

type CryptoUseCase struct {
	provider *provider.PgProvider
}

func NewCryptoUseCase(provider *provider.PgProvider) *CryptoUseCase {
	return &CryptoUseCase{provider: provider}
}

func (uc *CryptoUseCase) GetAllCryptos(ctx context.Context) ([]model.Crypto, error) {
	return uc.provider.GetCryptos(ctx)
}

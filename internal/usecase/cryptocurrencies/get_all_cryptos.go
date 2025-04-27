package cryptocurrencies

import (
	"context"
	"parser/internal/model"
)

func (uc *CryptoUseCase) GetAllCryptos(ctx context.Context) ([]model.Crypto, error) {
	return uc.provider.GetCryptos(ctx)
}

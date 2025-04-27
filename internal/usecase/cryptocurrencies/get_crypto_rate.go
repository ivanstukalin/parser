package cryptocurrencies

import (
	"context"
	"fmt"
	"parser/internal/model"
	"strings"
	"time"
)

const usdtPair = "USDT"

func (uc *CryptoUseCase) GetCryptoRate(ctx context.Context, code string) (*model.CryptoRate, error) {

	cryptoID := uc.provider.GetCryptoID(code)
	if cryptoID == 0 {
		return nil, fmt.Errorf("cryptocurrency with code '%s' not found", code)
	}

	latestRate := uc.provider.GetLatestCryptoRate(cryptoID)
	if latestRate != nil && time.Since(latestRate.CreatedAt) < time.Minute*15 {
		return latestRate, nil
	}

	binanceSymbole := strings.ToUpper(code) + usdtPair

	rate, err := uc.fetchFromBinance(binanceSymbole)
	if err != nil {
		return nil, fmt.Errorf("failed from Binance: %w", err)
	}

	savedRate, err := uc.provider.InsertCryptoRate(ctx, cryptoID, rate.Rate)
	if err != nil {
		return nil, fmt.Errorf("failed to save rate: %w", err)
	}
	return savedRate, nil
}

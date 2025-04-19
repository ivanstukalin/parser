package cryptocurrencies

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"parser/internal/model"
	"parser/internal/provider"
	"strings"
	"time"
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

func (uc *CryptoUseCase) GetCryptoRate(ctx context.Context, code string) (*model.CryptoRate, error) {
	cryptos, err := uc.provider.GetCryptos(ctx)
	if err != nil {
		return nil, err
	}

	var cryptocurrenciesID int
	for _, c := range cryptos {
		if c.Code == code {
			cryptocurrenciesID = c.ID
			break
		}
	}

	if cryptocurrenciesID == 0 {
		return nil, err
	}

	cashRates, err := uc.provider.GetCryptoRate(ctx, cryptocurrenciesID)
	if err != nil {
		return nil, err
	}

	if len(cashRates) > 0 {
		latest := cashRates[0]
		if time.Since(latest.CreatedAt) < time.Minute {
			return &latest, nil
		}
	}

	binanceSymbole := strings.ToUpper(code) + "USDT"

	nr, err := fetchFromBinance(binanceSymbole)
	if err != nil {
		return nil, err
	}

	sr, err := uc.provider.InsertCryptoRate(ctx, cryptocurrenciesID, nr.Rate)
	if err != nil {
		return nil, err
	}
	return sr, nil
}

func fetchFromBinance(symbol string) (*model.CryptoRate, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &model.CryptoRate{
		Rate: result.Price,
	}, nil
}

package cryptocurrencies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"parser/internal/model"
)

type BinanceResult struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (uc *CryptoUseCase) fetchFromBinance(symbol string) (*model.CryptoRate, error) {
	url := fmt.Sprintf(uc.binanceAPIURL, symbol)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result BinanceResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &model.CryptoRate{
		Rate: result.Price,
	}, nil
}

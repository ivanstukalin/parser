package controller

import (
	"encoding/json"
	"net/http"

	"parser/internal/usecase/cryptocurrencies"
)

type CryptoController struct {
	useCase *cryptocurrencies.CryptoUseCase
}

func NewCryptoController(useCase *cryptocurrencies.CryptoUseCase) *CryptoController {
	return &CryptoController{useCase: useCase}
}

func (c *CryptoController) GetCryptos(w http.ResponseWriter, r *http.Request) {
	cryptos, err := c.useCase.GetAllCryptos(r.Context())
	if err != nil {
		http.Error(w, "Failed to get cryptocurrencies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cryptos)
}

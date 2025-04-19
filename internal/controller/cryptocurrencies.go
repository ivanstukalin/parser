package controller

import (
	"encoding/json"
	"log/slog"
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
	ctx := r.Context()
	cryptos, err := c.useCase.GetAllCryptos(r.Context())
	if err != nil {
		slog.ErrorContext(ctx, "Failed cryptocurrencies", "path", r.URL.Path)
		http.Error(w, "Failed to get cryptocurrencies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cryptos)
}

func (c *CryptoController) GetCryptoRate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "missing code", http.StatusBadRequest)
		return
	}

	rate, err := c.useCase.GetCryptoRate(ctx, code)
	if err != nil {
		slog.ErrorContext(ctx, "Failed crypto rate", "error", err)
		http.Error(w, "Failed to get rate", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rate)
}

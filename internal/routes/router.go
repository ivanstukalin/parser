package routes

import (
	"parser/internal/controller"

	"github.com/go-chi/chi/v5"
)

func NewRouter(controller *controller.DomainController, cryptosController *controller.CryptoController) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/domains", controller.GetDomains)

	r.Get("/api/cryptocurrencies", cryptosController.GetCryptos)

	return r
}

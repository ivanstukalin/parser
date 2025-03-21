package routes

import (
	"github.com/go-chi/chi/v5"
	"parser/internal/controller"
)

func NewRouter(controller *controller.DomainController) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/domains", controller.GetDomains)

	return r
}

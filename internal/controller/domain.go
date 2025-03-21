package controller

import (
	"encoding/json"
	"net/http"

	"parser/internal/usecase/domain"
)

type DomainController struct {
	useCase *domain.DomainUseCase
}

func NewDomainController(useCase *domain.DomainUseCase) *DomainController {
	return &DomainController{useCase: useCase}
}

func (c *DomainController) GetDomains(w http.ResponseWriter, r *http.Request) {
	// обращение к юз кейсу
	domains, err := c.useCase.GetAllDomains(r.Context())
	if err != nil {
		// если ошибка, возвращаем 500
		http.Error(w, "failed to fetch domains", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domains)
}

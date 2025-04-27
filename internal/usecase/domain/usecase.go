package domain

import (
	"context"
	"parser/internal/model"
	"parser/internal/provider"
)

type DomainUseCase struct {
	provider *provider.PgProvider
}

func NewDomainUseCase(provider *provider.PgProvider) *DomainUseCase {
	return &DomainUseCase{provider: provider}
}

type UseCase interface {
	GetAllDomains(ctx context.Context) ([]model.Domain, error)
}

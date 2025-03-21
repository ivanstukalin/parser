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

func (u *DomainUseCase) GetAllDomains(ctx context.Context) ([]model.Domain, error) {
	return u.provider.GetDomains(ctx)
}

package domain

import (
	"context"
	"parser/internal/model"
)

func (u *DomainUseCase) GetAllDomains(ctx context.Context) ([]model.Domain, error) {
	return u.provider.GetDomains(ctx)
}

package port

import (
	"context"

	"elsa-coding-challange/internal/core/domain"
)

type UserRepository interface {
	GetByIDs(ctx context.Context, ids []string) ([]domain.User, error)
}

package storage

import (
	"context"

	"elsa-coding-challange/internal/core/domain"
)

type mockUserRepository struct{}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (r *mockUserRepository) GetByIDs(ctx context.Context, ids []string) ([]domain.User, error) {
	users := make([]domain.User, 0)
	for _, id := range ids {
		users = append(users, domain.User{
			ID:       id,
			UserName: "User " + id,
		})
	}
	return users, nil
}

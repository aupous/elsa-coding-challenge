package port

import (
	"context"

	"elsa-coding-challange/internal/core/domain"
)

type SortedSetRepository interface {
	ZAdd(ctx context.Context, setKey string, item domain.SortedSetItem) error
	ZRevRangeWithScores(ctx context.Context, setKey string, start, stop int64) ([]domain.SortedSetItem, error)
	ZRevRankWithScore(ctx context.Context, setKey string, member string) (int64, float64, error)
	ZCard(ctx context.Context, setKey string) (int64, error)
}

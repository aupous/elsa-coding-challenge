package storage

import (
	"context"

	"elsa-coding-challange/internal/core/domain"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *redisCache {
	return &redisCache{
		client: client,
	}
}

func (r *redisCache) ZAdd(ctx context.Context, setKey string, item domain.SortedSetItem) error {
	return r.client.ZAdd(ctx, setKey, redis.Z{Member: item.Member, Score: item.Score}).Err()
}

func (r *redisCache) ZRevRangeWithScores(ctx context.Context, setKey string, start, stop int64) ([]domain.SortedSetItem, error) {
	result, err := r.client.ZRevRangeWithScores(ctx, setKey, start, stop).Result()
	if err != nil {
		return nil, err
	}

	res := make([]domain.SortedSetItem, 0)
	for _, item := range result {
		res = append(res, domain.SortedSetItem{Member: item.Member.(string), Score: item.Score})
	}
	return res, nil
}

func (r *redisCache) ZRevRankWithScore(ctx context.Context, setKey string, member string) (int64, float64, error) {
	rank, err := r.client.ZRevRankWithScore(ctx, setKey, member).Result()
	if err != nil {
		return 0, 0, err
	}

	return rank.Rank, rank.Score, nil
}

func (r *redisCache) ZCard(ctx context.Context, setKey string) (int64, error) {
	return r.client.ZCard(ctx, setKey).Result()
}

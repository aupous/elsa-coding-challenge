package main

import (
	"context"
	"fmt"
	"math/rand"

	"elsa-coding-challange/internal/adapter/storage"
	"elsa-coding-challange/internal/core/domain"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

const (
	QUIZ_KEY = "test_key"
	NO_USERS = 10000
	USER_ID  = "user-100"
)

func main() {
	redisClient := redis.NewClient((&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}))

	sortedSetCache := storage.NewRedis(redisClient)

	ctx := context.Background()

	// Store sample data to redis
	log.Info().Msg("Starting add data to redis sorted set")
	for i := 0; i < NO_USERS; i++ {
		sortedSetCache.ZAdd(ctx, QUIZ_KEY, domain.SortedSetItem{
			Member: fmt.Sprintf("user-%d", i),
			Score:  rand.Float64() * 10,
		})
	}
	log.Info().Msg("Add data to redis sorted set completed!")

	// For testing
	topTen, err := sortedSetCache.ZRevRangeWithScores(ctx, QUIZ_KEY, 0, 10)
	if err != nil {
		panic(err)
	}

	rank, score, err := sortedSetCache.ZRevRankWithScore(ctx, QUIZ_KEY, USER_ID)
	if err != nil {
		panic(err)
	}

	log.Info().Fields(map[string]interface{}{
		"topTen": topTen,
		"user": map[string]interface{}{
			"id":    USER_ID,
			"rank":  rank,
			"score": score,
		},
	}).Msg("Result: ")
}

package main

import (
	"os"

	"elsa-coding-challange/internal/adapter/handler/http"
	"elsa-coding-challange/internal/adapter/storage"
	"elsa-coding-challange/internal/core/service"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func main() {
	redisClient := redis.NewClient((&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}))

	sortedSetCache := storage.NewRedis(redisClient)

	userRepo := storage.NewMockUserRepository()

	leaderboardSvc := service.NewLeaderboardService(sortedSetCache, userRepo)
	leaderboardHandler := http.NewLeaderboardHandler(leaderboardSvc)

	router := http.NewRouter(leaderboardHandler)
	err := router.Serve("localhost:9000")
	if err != nil {
		log.Error().Err(err).Msg("Error starting the HTTP server")
		os.Exit(1)
	}
}

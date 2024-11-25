package port

import (
	"context"

	"elsa-coding-challange/internal/core/domain"
)

type LeaderboardService interface {
	GetLeaderboardByQuiz(ctx context.Context, key string, start, stop int64) (items []domain.LeaderBoardItem, total int64, err error)
	GetUserRankByQuiz(ctx context.Context, id string, userId string) (rank int64, score float64, err error)
}

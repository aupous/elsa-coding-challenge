package service

import (
	"context"

	"elsa-coding-challange/internal/core/domain"
	"elsa-coding-challange/internal/core/port"
)

type leaderboardService struct {
	sortedSetCache port.SortedSetRepository
	userRepo       port.UserRepository
}

func NewLeaderboardService(
	sortedSetCache port.SortedSetRepository,
	userRepo port.UserRepository,
) *leaderboardService {
	return &leaderboardService{
		sortedSetCache: sortedSetCache,
		userRepo:       userRepo,
	}
}

func (s *leaderboardService) GetLeaderboardByQuiz(ctx context.Context, key string, start, stop int64) ([]domain.LeaderBoardItem, int64, error) {
	count, err := s.sortedSetCache.ZCard(ctx, key)
	if err != nil {
		return nil, 0, err
	}

	items, err := s.sortedSetCache.ZRevRangeWithScores(ctx, key, start, stop)
	if err != nil {
		return nil, 0, err
	}

	userIDs := make([]string, 0)
	for _, item := range items {
		userIDs = append(userIDs, item.Member)
	}
	users, err := s.userRepo.GetByIDs(ctx, userIDs)
	if err != nil {
		return nil, 0, err
	}
	userMap := make(map[string]domain.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	results := make([]domain.LeaderBoardItem, 0)
	for index, item := range items {
		results = append(results, domain.LeaderBoardItem{
			Rank:  start + int64(index),
			Score: item.Score,
			User:  userMap[item.Member],
		})
	}

	return results, count, nil
}

func (s *leaderboardService) GetUserRankByQuiz(ctx context.Context, key string, userId string) (rank int64, score float64, err error) {
	return s.sortedSetCache.ZRevRankWithScore(ctx, key, userId)
}

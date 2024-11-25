package http

import (
	"elsa-coding-challange/internal/core/port"

	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_LEADERBOARD_LIMIT = 100
	MAX_LEADERBOARD_LIMIT     = 1000
)

type leaderboardHandler struct {
	svc port.LeaderboardService
}

func NewLeaderboardHandler(svc port.LeaderboardService) *leaderboardHandler {
	return &leaderboardHandler{
		svc: svc,
	}
}

type getLeaderboardRequest struct {
	Limit int64 `form:"limit"`
	Skip  int64 `form:"skip"`
}

func (h *leaderboardHandler) GetLeaderboardByQuiz(ctx *gin.Context) {
	quizKey := ctx.Param("key")
	var req getLeaderboardRequest
	if err := ctx.ShouldBind(&req); err != nil {
		validationError(ctx, err)
		return
	}
	if req.Limit == 0 {
		req.Limit = DEFAULT_LEADERBOARD_LIMIT
	}

	items, total, err := h.svc.GetLeaderboardByQuiz(ctx, quizKey, req.Skip, req.Skip+req.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	meta := newMeta(total, req.Limit, req.Skip)
	rsp := map[string]any{
		"meta": meta,
		"data": items,
	}

	handleSuccess(ctx, rsp)
}

type getUserRankRequest struct {
	UserId string `form:"userId" binding:"required" example:"user1"`
}

func (h *leaderboardHandler) GetUserRankByQuiz(ctx *gin.Context) {
	quizKey := ctx.Param("key")
	var req getUserRankRequest
	if err := ctx.ShouldBind(&req); err != nil {
		validationError(ctx, err)
		return
	}

	rank, score, err := h.svc.GetUserRankByQuiz(ctx, quizKey, req.UserId)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := map[string]any{
		"rank":  rank,
		"score": score,
	}
	handleSuccess(ctx, map[string]any{
		"data": rsp,
	})
}

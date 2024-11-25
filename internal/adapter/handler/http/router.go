package http

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(leaderboardHandler *leaderboardHandler) *Router {
	router := gin.New()
	router.Use(gin.Recovery(), loggerMiddleware())

	router.GET("/v1/leaderboard/:key", leaderboardHandler.GetLeaderboardByQuiz)
	router.GET("/v1/leaderboard/:key/user-rank", leaderboardHandler.GetUserRankByQuiz)

	return &Router{router}
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}

package http

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func loggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		jsonData, _ := io.ReadAll(ctx.Request.Body)

		log.Info().Fields(map[string]interface{}{
			"method":       ctx.Request.Method,
			"url":          ctx.Request.URL.String(),
			"resquestBody": jsonData,
			"errors":       ctx.Errors.Errors(),
		}).Msg("HTTP request info")
	}
}

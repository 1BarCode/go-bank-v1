package api

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) setupPingRoutes(rg *gin.RouterGroup) {
	pingRoutes := rg.Group("/ping")
	{
		pingRoutes.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


func (s *Server) setupUserRoutes(apiRoutes *gin.RouterGroup) {
	userRoutes := apiRoutes.Group("/users")
	{
		userRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "users get enpoint",
			})
		})
		userRoutes.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			ctx.JSON(200, gin.H{
				"message": fmt.Sprintf("user with id %s", id),
			})
		})
		userRoutes.POST("/", nil)
		userRoutes.PUT("/:id", nil)
		userRoutes.DELETE("/:id", nil)
	}
}
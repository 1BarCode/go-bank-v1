package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


func (s *Server) setupUserRoutes(rg *gin.RouterGroup) {
	userRoutes := rg.Group("/users")
	{
		userRoutes.GET("", listUsers)
		userRoutes.GET(":id", getUser)
		userRoutes.POST("", nil)
		userRoutes.PUT(":id", nil)
		userRoutes.DELETE(":id", nil)
	}
}

func listUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "users get enpoint",
	})
}

func getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("user with id %s", id),
	})
}
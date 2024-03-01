package api

import "github.com/gin-gonic/gin"

func (s *Server) setupAccountRoutes(rg *gin.RouterGroup) {
	accountRoutes := rg.Group("/accounts")
	{
		accountRoutes.POST("", s.createAccount)
		accountRoutes.GET(":id", s.getAccount)
		accountRoutes.GET("", s.listAccounts)
		// accountRoutes.PUT(":id", nil)
		accountRoutes.DELETE(":id", s.deleteAccount)
	}
}
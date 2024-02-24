package app

import (
	"net/http"

	"github.com/1BarCode/go-bank-v1/app/services"
	"github.com/gin-gonic/gin"
)

type Server struct {
	services services.Services // interface so no pointer needed
	router *gin.Engine
}

func NewServer(s services.Services) *Server {
	server := &Server{services: s}
	router := gin.Default()
	server.router = router

	
	// user routes
	apiRoutes := router.Group("/api")
	{
		router.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		server.setupUserRoutes(apiRoutes)
	}

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}


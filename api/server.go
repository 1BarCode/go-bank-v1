package api

import (
	"fmt"
	"time"

	"github.com/1BarCode/go-bank-v1/services"
	"github.com/gin-gonic/gin"
)

type Server struct {
	services services.Services // interface so no pointer needed
	router *gin.Engine
}

// NewServer creates a new server with all the necessary routes
func NewServer(s services.Services) *Server {
	server := &Server{services: s}
	router := gin.Default()
	server.router = router

	server.setUpRoutes()
	
	return server
}

// Start runs the server on a specific address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func (s *Server) setUpRoutes() {
	v1Routes := s.router.Group("/v1")
	{
		// ping routes
		s.setupPingRoutes(v1Routes)

		// account routes
		s.setupAccountRoutes(v1Routes)

		// user routes
		s.setupUserRoutes(v1Routes)

		v1Routes.GET("/concurrent", doConcurrentStuff)
	}
}

func doConcurrentStuff(ctx *gin.Context) {
	start := time.Now()
	// ch := make(chan string)

	for i := 1; i < 5; i++ {
		task2(i)
	}

	res := []string{}

	// for i := 1; i < 5; i++ {
	// 	res = append(res, <-ch)
	// }

	elapsed := time.Since(start).Seconds()

	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("concurrent stuff done in %f seconds", elapsed),
		"results": res,
	})
}

// func task1(ch chan string, delay int) {
// 	time.Sleep(time.Duration(delay) * time.Second)
// 	ch <- fmt.Sprintf("task %d done", delay)
// }

func task2(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func intServerErrorResponse() gin.H {
	return gin.H{"error": "internal server error"}
}

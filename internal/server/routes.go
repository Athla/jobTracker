package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	api := r.Group("/api/jobs")
	{
		api.POST("/", s.CreateJobHandler)        // Create a job
		api.GET("/", s.JobHandler)               // Get all, queryied to get a specific one
		api.GET("/:id", s.GetSpecificJobHandler) // Get all, queryied to get a specific one
		api.PUT("/:id", s.EditJobHandler)        // Update, idempotently, a job
		api.DELETE("/:id", s.DeleteJobHandler)   // Delete a job based on it's ide
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetSpecificJobHandler(c *gin.Context) {

}

func (s *Server) JobHandler(c *gin.Context) {

}

func (s *Server) CreateJobHandler(c *gin.Context) {

}

func (s *Server) DeleteJobHandler(c *gin.Context) {

}

func (s *Server) EditJobHandler(c *gin.Context) {

}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(CORSMiddlware())

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	api := r.Group("/api/jobs", AuthMiddleware())
	{
		api.POST("/", s.CreateJobHandler)                // Create a job
		api.GET("/", s.JobHandler)                       // Get all, queryied to get a specific one
		api.GET("/:id", s.GetSpecificJobHandler)         // Get all, queryied to get a specific one
		api.PUT("/:id", s.EditJobHandler)                // Update, idempotently, a job
		api.DELETE("/:id", s.DeleteJobHandler)           // Delete a job based on it's ide
		api.DELETE("/deleteAll", s.DeleteAllJobsHandler) // Delete a job based on it's idel
		api.PATCH("/:id/status", s.UpdateJobStatusHandler)
	}
	users := r.Group("/users/")
	{
		users.POST("/register", s.RegisterUserHandler)
		users.POST("/login", s.LoginUserHandler)
		users.POST("/logout", s.LogoutUserHandler)
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

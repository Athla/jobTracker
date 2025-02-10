package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(CORSMiddlware())
	r.POST("/login", s.LoginHandler)
	r.POST("/logout", s.LogoutHandler, AuthMiddleware())
	r.GET("/health", s.healthHandler)
	r.GET("/", s.HelloWorldHandler)

	api := r.Group("/api", AuthMiddleware())
	{
		jobs := api.Group("/jobs")
		{
			jobs.POST("/", s.CreateJobHandler)
			jobs.GET("/", s.JobHandler)
			jobs.GET("/:id", s.GetSpecificJobHandler)
			jobs.PUT("/:id", s.EditJobHandler)
			jobs.DELETE("/:id", s.DeleteJobHandler)
			jobs.DELETE("/deleteAll", s.DeleteAllJobsHandler)
			jobs.PATCH("/:id/status", s.UpdateJobStatusHandler)
		}
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

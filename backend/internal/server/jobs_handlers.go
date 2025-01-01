package server

import (
	"fmt"
	"jobTracker/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetSpecificJobHandler(c *gin.Context) {

}
func (s *Server) JobHandler(c *gin.Context) {
	// Return a slice of jobs to the front-end
	var jobs []models.Job
	if err := s.GetAllJobs(&jobs); err != nil {
		log.Errorf("Unable to get data due: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Failure": "Unable to get data."})
		return
	}
	c.JSON(http.StatusOK, jobs)

}

func (s *Server) CreateJobHandler(c *gin.Context) {
	// Receive a job struct as a payload
	// Check if it does not exist -> later, for now make it work then add redundancy for that later, maybe a local cache upon init.
	// Add the job after data validation to the database.
	var newJob models.Job
	if err := c.ShouldBindJSON(&newJob); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failure": "Unable to create new job."})
		return
	}

	createdAt, err := time.Parse(time.RFC3339, newJob.CreatedAt)
	if err != nil {
		log.Fatalf("Current err: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Failure": "Invalid date format"})
		return
	}

	newJob.CreatedAt = createdAt.Format(time.RFC3339)
	log.Infof("Job to be created: %v", newJob)
	if err := s.CreateNewJob(newJob); err != nil {
		log.Errorf("Error during creation: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Failure": "Unable to create new job."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Job created successfully!"})
}

func (s *Server) DeleteJobHandler(c *gin.Context) {
	rawId := c.Param("id")
	if rawId != "" {
		id, err := strconv.Atoi(rawId)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Failure": "Unable to delete job."})
			return
		}

		if err := s.DeleteJob(&id); err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Failure": "Unable to delete job."})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Success": fmt.Sprintf("Successfully deleted of id: %v", id)})
		return
	}
	log.Error("Unable to parse id.")
	// This shouldn't happen since ID will be forced to be a number in the creation.
	c.JSON(http.StatusBadRequest, gin.H{"Failure": "No ID found."})
	return
}
func (s *Server) DeleteAllJobsHandler(c *gin.Context) {
	if err := s.DeleteAllJobs(); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Failure": "Unable to delete job."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success": fmt.Sprintf("Successfully deleted all jobs")})
	return
}

func (s *Server) EditJobHandler(c *gin.Context) {

}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.Health())
}

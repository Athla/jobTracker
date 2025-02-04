package server

import (
	"database/sql"
	"fmt"
	"jobTracker/internal/database"
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
	id := c.Param("id")
	version, err := strconv.Atoi(c.Query("version"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Version parameter required."})
		return
	}

	tx, err := s.db.Beginx()
	if err != nil {
		log.Errorf("Failed to start transaction due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	defer tx.Rollback()

	result, err := tx.Exec(database.DeleteIdQuery, id, version)
	if err != nil {
		log.Errorf("Unable to delete job due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job."})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Errorf("Unable get rows affected due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if rowsAffected == 0 {
		log.Error("No rows affected or deleted.")
		c.JSON(http.StatusConflict, gin.H{"error": "Job has been modified or deleted."})
		return
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Unable to commit transaction due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully!"})
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

func (s *Server) UpdateJobStatusHandler(c *gin.Context) {
	id := c.Param("id")
	var statusUpdate struct {
		Status  string `json:"status"`
		Version int    `json:"version"`
	}

	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if statusUpdate.Status != string(models.StatusPending) &&
		statusUpdate.Status != string(models.StatusInProgress) &&
		statusUpdate.Status != string(models.StatusCompleted) &&
		statusUpdate.Status != string(models.StatusRejected) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status passed."})
		return
	}

	tx, err := s.db.Beginx()
	if err != nil {
		log.Errorf("Failed to start transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer tx.Rollback()

	var currentJob models.Job
	err = tx.Get(&currentJob, database.GetJobByIDQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
			return
		}
		log.Errorf("Failed to get job: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if currentJob.Version != statusUpdate.Version {
		c.JSON(http.StatusConflict, gin.H{"error": "Job has been modified by another request"})
		return
	}

	var newVersion int
	err = tx.QueryRow(
		database.UpdateJobStatusQuery,
		statusUpdate.Status,
		id,
		statusUpdate.Version,
	).Scan(&newVersion)

	if err != nil {
		log.Errorf("Failed to update job status: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job status"})
		return
	}

	if err = tx.Commit(); err != nil {
		log.Errorf("Failed to commit transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var updatedJob models.Job
	err = s.db.Get(&updatedJob, database.GetJobByIDQuery, id)
	if err != nil {
		log.Errorf("Failed to get updated job: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, updatedJob)
}

func (s *Server) EditJobHandler(c *gin.Context) {
	id := c.Param("id")
	var update models.JobUpdate

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return
	}

	tx, err := s.db.Beginx()
	defer tx.Rollback()

	var currentJob models.Job
	err = tx.Get(&currentJob, database.GetJobByIDQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job not found."})
			return
		}

		log.Errorf("Unable to get job due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	if currentJob.Version != update.Version {
		c.JSON(http.StatusConflict, gin.H{"error": "Job has been modified by another request."})
		return
	}

	var newVersion int

	err = tx.QueryRow(
		database.UpdateJobQuery,
		update.Name,
		update.Source,
		update.Description,
		update.Status,
		id,
		update.Version,
	).Scan(&newVersion)

	if err != nil {
		log.Errorf("Failed to update job due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	if err = tx.Commit(); err != nil {
		log.Errorf("Unable to commit transaction due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server error.",
		})
		return
	}

	var updatedJob models.Job
	err = s.db.Get(&updatedJob, database.GetJobByIDQuery, id)
	if err != nil {
		log.Errorf("Unable to get job due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	c.JSON(http.StatusOK, updatedJob)

}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.Health())
}

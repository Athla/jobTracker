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
	id := c.Param("id")

	var job models.Job

	err := s.db.Get(&job, database.GetJobByIDQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job not found."})
			return
		}

		log.Errorf("Unable to get job of id '%s' due: %s", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	c.JSON(http.StatusOK, job)
}
func (s *Server) JobHandler(c *gin.Context) {
	status := c.Query("status")

	var jobs []models.Job
	var err error

	if status != "" {
		err = s.db.Select(&jobs, database.GetJobByStatusQuery, status)
	} else {
		err = s.db.Select(&jobs, database.GetAllJobs)
	}

	if err != nil {
		log.Errorf("Unable to get jobs due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve jobs."})
		return
	}

	if jobs == nil {
		jobs = []models.Job{}
	}

	c.JSON(http.StatusOK, jobs)
}

func (s *Server) CreateJobHandler(c *gin.Context) {
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		log.Errorf("Unable to bind JSON due: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	job.Status = models.Wishlist
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
	job.Version = 1

	tx, err := s.db.Beginx()
	if err != nil {
		log.Errorf("Unable to start transaction due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer tx.Rollback()

	result, err := tx.NamedExec(database.CreateJobQuery, job)
	if err != nil {
		log.Errorf("Unable to create job due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job"})
		return
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Unable to commit transaction due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	id, _ := result.LastInsertId()
	job.ID = strconv.Itoa(int(id))

	c.JSON(http.StatusCreated, job)
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

	result, err := tx.Exec(database.DeleteJobQuery, id, version)
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
	var update struct {
		Status  models.JobStatus `json:"status"`
		Version int              `json:"version"`
	}

	if err := c.ShouldBindJSON(&update); err != nil {
		log.Errorf("Unable to bind JSON due: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	tx, err := s.db.Beginx()
	if err != nil {
		log.Errorf("Failed to start transaction due: %s", err)
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
		log.Errorf("Failed to get job due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if err := currentJob.ValidateStatus(update.Status); err != nil {
		log.Errorf("Invalid status transition due: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newVersion int
	err = tx.QueryRow(
		database.UpdateJobStatusQuery,
		update.Status,
		id,
		update.Version,
	).Scan(&newVersion)

	if newVersion == 0 {
		newVersion = update.Version + 1
		err = nil
	}
	if err != nil {
		log.Errorf("Failed to update job status due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job status"})
		return
	}

	if err = tx.Commit(); err != nil {
		log.Errorf("Failed to commit transaction due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Get updated job
	var updatedJob models.Job
	err = s.db.Get(&updatedJob, database.GetJobByIDQuery, id)
	if err != nil {
		log.Errorf("Failed to get updated job due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, updatedJob)
}

func (s *Server) EditJobHandler(c *gin.Context) {
	id := c.Param("id")
	var update models.Job

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	tx, err := s.db.Beginx()
	if err != nil {
		log.Errorf("Failed to start transaction due: %s", err)
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
		log.Errorf("Failed to get job due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var newVersion int
	err = tx.QueryRow(
		database.UpdateJobQuery,
		update.Name,
		update.Company,
		update.Source,
		update.Description,
		update.JobType,
		update.Status,
		id,
		update.Version,
	).Scan(&newVersion)

	if err != nil {
		log.Errorf("Failed to update job due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job"})
		return
	}

	if err = tx.Commit(); err != nil {
		log.Errorf("Failed to commit transaction due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var updatedJob models.Job
	err = s.db.Get(&updatedJob, database.GetJobByIDQuery, id)
	if err != nil {
		log.Errorf("Failed to get updated job due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, updatedJob)
}

func (s *Server) GetBoardHandler(c *gin.Context) {
	tx, err := s.db.Beginx()
	if err != nil {
		log.Errorf("Unable to begin transaction due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer tx.Rollback()

	appliedJobs, err := tx.Queryx(database.GetJobsByBoardColumnQuery, "APPLIED")
	inProgressJobs, err := tx.Queryx(database.GetJobsByBoardColumnQuery, "IN_PROGRESS")
	finishedJobs, err := tx.Queryx(database.GetJobsByBoardColumnQuery, "FINISHED")

	if err != nil {
		log.Errorf("Unable to fetch board data due: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch board data"})
		return
	}

	board := gin.H{
		"applied":    appliedJobs,
		"inProgress": inProgressJobs,
		"finished":   finishedJobs,
	}
	c.JSON(http.StatusOK, board)
}
func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.Health())
}

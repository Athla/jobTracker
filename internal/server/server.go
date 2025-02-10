package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"

	"jobTracker/internal/database"
	"jobTracker/internal/models"
)

type Server struct {
	port int
	db   *sqlx.DB
}

func New(db *sqlx.DB, port int) *http.Server {
	NewServer := &Server{
		port: port,
		db:   db,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Infof("Server running on address: %s", server.Addr)

	return server
}

func (s *Server) DeleteAllJobs() error {
	if s.db == nil {
		log.Fatal("Null database, unable to interact and create data.")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(database.DeleteAllJobsQuery); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (s *Server) DeleteJob(id *int) error {
	if s.db == nil {
		log.Fatal("Null database, unable to interact and create data.")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(database.DeleteJobQuery, id); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Server) CreateNewJob(job models.Job) error {
	if s.db == nil {
		log.Fatal("Null database, unable to interact and create data.")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	if _, err := tx.NamedExec(database.CreateJobQuery, &job); err != nil {
		log.Errorf("Unable to create entry due: %v", err)
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Errorf("Unable to commit the transaction due: %v", err)
		return err
	}

	log.Info("Commited the job %v to the database", job)

	return nil
}

func (s *Server) GetAllJobs(jobs *[]models.Job) error {
	if s.db == nil {
		log.Fatal("Null database, unable to interact and create data.")
	}

	if err := s.db.Select(jobs, database.GetAllJobs); err != nil {
		return err
	}

	return nil
}

func (s *Server) updateJobStatus(id string, status string) error {
	if s.db == nil {
		return errors.New("Database conn not initialized.")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		log.Errorf("Unable to create transaction due: %s", err)
		return err
	}

	if _, err := tx.Exec(database.UpdateJobStatusQuery, status, id); err != nil {
		log.Errorf("Unable to execute transaction due: %s", err)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *Server) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

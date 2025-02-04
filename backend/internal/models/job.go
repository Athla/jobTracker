package models

import (
	"errors"

	"github.com/charmbracelet/log"
)

type Job struct {
	ID          string `json:"id" db:"id" form:"id"`
	Name        string `json:"name" db:"name" form:"name"`
	Source      string `json:"source" db:"source" form:"source"`
	Description string `json:"description" db:"description" form:"description"`
	Status      string `json:"status" db:"status" form:"status"`
	Version     int    `json:"version" db:"version"`
	CreatedAt   string `json:"created_at" db:"created_at" form:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}

// JobUpdate represents the fields that can be updated
type JobUpdate struct {
	Name        *string `json:"name,omitempty"`
	Source      *string `json:"source,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
	Version     int     `json:"version"` // Required for optimistic locking
}

type JobStatus string

const (
	StatusPending    JobStatus = "pending"
	StatusInProgress JobStatus = "in_progress"
	StatusCompleted  JobStatus = "completed"
	StatusRejected   JobStatus = "rejected"
)

func NewJob(name, source, description string, date string) (*Job, error) {
	// Parse date into RFC.3339Z
	return &Job{
		Name:        name,
		Source:      source,
		Description: description,
		Status:      string(StatusPending),
		CreatedAt:   date,
	}, nil
}

func (j *Job) ValidateStatus(status JobStatus) error {
	if status != StatusPending &&
		status != StatusInProgress &&
		status != StatusCompleted &&
		status != StatusRejected {
		return errors.New("Invalid status.")
	}
	return nil
}

func (j *Job) UpdateStatus(status JobStatus) error {
	if err := j.ValidateStatus(status); err != nil {
		log.Errorf("Unable to validate status due: %s", err)
		return err
	}

	j.Status = string(status)
	return nil
}

func (j *Job) Edit() error {

	return nil
}

func (j *Job) Delete() error {

	return nil
}

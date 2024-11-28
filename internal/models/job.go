package models

import (
	"jobTracker/internal/utils"
	"time"
)

type Job struct {
	ID          string    `json:"id" db:"id" form:"id"`
	Name        string    `json:"name" db:"name" form:"name"`
	Source      string    `json:"source" db:"source" form:"source"`
	Description string    `json:"description" db:"description" form:"description"`
	CreatedAt   time.Time `json:"createdat" db:"created_at" form:"created_at"`
}

func NewJob(name, source, description string, date string) (*Job, error) {
	// Parse date into RFC.3339Z
	cDate, err := utils.ConvertDate(date)
	if err != nil {
		return nil, err
	}
	return &Job{
		Name:        name,
		Source:      source,
		Description: description,
		CreatedAt:   cDate,
	}, nil
}

func (j *Job) Edit() error {

	return nil
}

func (j *Job) Delete() error {

	return nil
}

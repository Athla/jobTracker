package models

import (
	"jobTracker/internal/utils"
	"time"
)

type Job struct {
	ID          string
	Name        string
	Source      string
	Description string
	CreateAt    time.Time
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
		CreateAt:    cDate,
	}, nil
}

func (j *Job) Edit() error {

	return nil
}

func (j *Job) Delete() error {

	return nil
}

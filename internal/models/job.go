package models

import (
	"time"

	"github.com/charmbracelet/log"
)

type JobType string
type JobStatus string

type Job struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Company     string    `json:"company" db:"company"`
	Source      string    `json:"source" db:"source"`
	Description string    `json:"description,omitempty" db:"description"`
	JobType     JobType   `json:"job_type" db:"job_type"`
	Status      JobStatus `json:"status" db:"status"`
	Version     int       `json:"version" db:"version"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// JobUpdate represents the fields that can be updated
type JobUpdate struct {
	Name                *string    `json:"name,omitempty"`
	Source              *string    `json:"source,omitempty"`
	Description         *string    `json:"description,omitempty"`
	Company             *string    `json:"company,omitempty"`
	Location            *string    `json:"location,omitempty"`
	SalaryRange         *string    `json:"salary_range,omitempty"`
	JobType             *JobType   `json:"job_type,omitempty"`
	Status              *JobStatus `json:"status,omitempty"`
	ApplicationLink     *string    `json:"application_link,omitempty"`
	RejectionReason     *string    `json:"rejection_reason,omitempty"`
	Notes               *string    `json:"notes,omitempty"`
	InterviewNotes      *string    `json:"interview_notes,omitempty"`
	NextInterviewDate   *time.Time `json:"next_interview_date,omitempty"`
	LastInteractionDate *time.Time `json:"last_interaction_date,omitempty"`
	Version             int        `json:"version"` // Required for optimistic locking
}

const (
	// Job Types
	FullTime   JobType = "FULL_TIME"
	PartTime   JobType = "PART_TIME"
	Contract   JobType = "CONTRACT"
	Internship JobType = "INTERNSHIP"
	Remote     JobType = "REMOTE"

	// Job Status
	Wishlist           JobStatus = "WISHLIST"
	Applied            JobStatus = "APPLIED"
	PhoneScreen        JobStatus = "PHONE_SCREEN"
	TechnicalInterview JobStatus = "TECHNICAL_INTERVIEW"
	Onsite             JobStatus = "ONSITE"
	Offer              JobStatus = "OFFER"
	Rejected           JobStatus = "REJECTED"
	Accepted           JobStatus = "ACCEPTED"
	Withdrawn          JobStatus = "WITHDRAWN"
)

func NewJob(name, company, source, description string, date string) (*Job, error) {
	if name == "" || company == "" {
		return nil, ErrNoNameOrCompany
	}
	return &Job{
		Name:      name,
		Company:   company,
		Source:    source,
		Status:    Wishlist,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}, nil
}

func (j *Job) ValidateStatus(status JobStatus) error {
	validStatuses := map[JobStatus]bool{
		Wishlist:           true,
		Applied:            true,
		PhoneScreen:        true,
		TechnicalInterview: true,
		Onsite:             true,
		Offer:              true,
		Rejected:           true,
		Accepted:           true,
		Withdrawn:          true,
	}
	if !validStatuses[status] {
		return ErrInvalidStatus
	}

	return nil
}

func (j *Job) UpdateStatus(status JobStatus) error {
	if err := j.ValidateStatus(status); err != nil {
		log.Errorf("Unable to validate status due: %s", err)
		return err
	}

	j.Status = status

	return nil
}

func (j *Job) ValidateStatusTranstion(newStatus JobStatus) error {
	validTransitions := map[JobStatus][]JobStatus{
		Wishlist:           {Applied},
		Applied:            {PhoneScreen, Rejected, Withdrawn},
		PhoneScreen:        {TechnicalInterview, Rejected, Withdrawn},
		TechnicalInterview: {Onsite, Rejected, Withdrawn},
		Onsite:             {Offer, Rejected, Withdrawn},
		Offer:              {Accepted, Rejected, Withdrawn},
		// Terminal states
		Accepted:  {},
		Rejected:  {},
		Withdrawn: {},
	}

	allowedTransitions, exists := validTransitions[j.Status]
	if !exists {
		return ErrInvalidCurrentStatus
	}

	for _, allowed := range allowedTransitions {
		if newStatus == allowed {
			return nil
		}
	}

	return ErrInvalidStatusTransition
}

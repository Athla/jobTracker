package models

type Job struct {
	ID          string `json:"id" db:"id" form:"id"`
	Name        string `json:"name" db:"name" form:"name"`
	Source      string `json:"source" db:"source" form:"source"`
	Description string `json:"description" db:"description" form:"description"`
	CreatedAt   string `json:"created_at" db:"created_at" form:"created_at"` // Match JSON with DB field
}

func NewJob(name, source, description string, date string) (*Job, error) {
	// Parse date into RFC.3339Z
	return &Job{
		Name:        name,
		Source:      source,
		Description: description,
		CreatedAt:   date,
	}, nil
}

func (j *Job) Edit() error {

	return nil
}

func (j *Job) Delete() error {

	return nil
}

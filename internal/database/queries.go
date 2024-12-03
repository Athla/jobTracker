package database

const (
	CreateJobQuery = `INSERT INTO Jobs (name, source, description, created_at) VALUES (:name, :source, :description, :created_at)`
	GetAllJobs     = `SELECT id, name, source, description, created_at FROM Jobs`
	DeleteIdQuery  = `DELETE FROM Jobs WHERE ID = $1`
	DeleteAllQuery = `DELETE FROM Jobs`
)

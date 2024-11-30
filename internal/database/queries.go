package database

const (
	CreateJobQuery = `INSERT INTO Jobs (id, name, source, description, created_at) VALUES (:id, :name, :source, :description, :created_at)`
	GetAllJobs     = `SELECT id, name, source, description, created_at FROM Jobs`
	DeleteIdQuery  = `DELETE FROM Jobs WHERE ID = $1`
	DeleteAllQuery = `DELETE FROM Jobs`
)

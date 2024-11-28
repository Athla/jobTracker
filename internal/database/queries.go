package database

var (
	CreateJobQuery = `INSERT INTO Jobs (id, name, source, description, created_at) VALUES (:id, :name, :source, :description, :created_at)`
)

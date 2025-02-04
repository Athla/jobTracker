package database

const (
	CreateJobQuery = `
        INSERT INTO Jobs (name, source, description, status, created_at)
        VALUES (:name, :source, :description, :status, :created_at)
        RETURNING id, version`

	GetAllJobs = `
        SELECT id, name, source, description, status, version, created_at, updated_at
        FROM Jobs`

	GetJobByIDQuery = `
        SELECT id, name, source, description, status, version, created_at, updated_at
        FROM Jobs
        WHERE id = $1`

	UpdateJobQuery = `
        UPDATE Jobs
        SET name = COALESCE($1, name),
            source = COALESCE($2, source),
            description = COALESCE($3, description),
            status = COALESCE($4, status)
        WHERE id = $5 AND version = $6
        RETURNING version`

	DeleteIdQuery = `
        DELETE FROM Jobs
        WHERE ID = $1 AND version = $2`

	DeleteAllQuery = `DELETE FROM Jobs`

	UpdateJobStatusQuery = `
		UPDATE Jobs
			SET status = COALESCE($1, status)
			WHERE id = $2 and VERSION = $3
			RETURNING version`
)

package database

const (
	CreateJobQuery = `
        INSERT INTO Jobs (
            name, company, source, description,
            job_type, status, created_at
        ) VALUES (
            :name, :company, :source, :description,
            :job_type, :status, :created_at
        ) RETURNING id, version`

	GetAllJobs = `
        SELECT * FROM Jobs ORDER BY created_at DESC`

	GetJobByIDQuery = `
        SELECT * FROM Jobs WHERE id = $1`

	GetJobByStatusQuery = `
        SELECT * FROM Jobs
        WHERE status = $1
        ORDER BY created_at DESC`

	UpdateJobQuery = `
        UPDATE Jobs
        SET name = COALESCE($1, name),
            company = COALESCE($2, company),
            source = COALESCE($3, source),
            description = COALESCE($4, description),
            job_type = COALESCE($5, job_type),
            status = COALESCE($6, status),
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $7 AND version = $8
        RETURNING version`

	DeleteJobQuery = `
        DELETE FROM Jobs
        WHERE id = $1 AND version = $2`

	DeleteAllJobsQuery = `
        DELETE FROM Jobs`

	UpdateJobStatusQuery = `
        UPDATE Jobs
        SET status = COALESCE($1, status),
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $2 AND version = $3
        RETURNING version`

	GetJobsByBoardColumnQuery = `
        SELECT * FROM Jobs
        WHERE CASE
            WHEN $1 = 'applied' THEN status IN ('WISHLIST', 'APPLIED')
            WHEN $1 = 'in-progress' THEN status IN ('PHONE_SCREEN', 'TECHNICAL_INTERVIEW', 'ONSITE')
            WHEN $1 = 'finished' THEN status IN ('OFFER', 'ACCEPTED', 'REJECTED', 'WITHDRAWN')
        END
        ORDER BY created_at DESC`

	UpdateJobBoardPositionQuery = `
                UPDATE Jobs
                SET board_column = $1,
                    board_position = $2,
                    updated_at = CURRENT_TIMESTAMP,
                    version = version + 1
                WHERE id = $3 AND version = $4
                RETURNING version`
)

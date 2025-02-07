package database

const (
	CreateJobQuery = `
        INSERT INTO Jobs (
            name, company, source, description, location, salary_range,
            job_type, status, application_link, notes, created_at
        ) VALUES (
            :name, :company, :source, :description, :location, :salary_range,
            :job_type, :status, :application_link, :notes, :created_at
        ) RETURNING id, version`

	GetAllJobs = `
        SELECT * FROM Jobs ORDER BY created_at DESC`

	GetJobByIDQuery = `
        SELECT * FROM Jobs WHERE id = $1`

	GetJobByStatusQuery = `
            SELECT * FROM Jobs WHERE status = $1 created_at DESC`

	UpdateJobQuery = `
        UPDATE Jobs
        SET name = COALESCE($1, name),
            company = COALESCE($2, company),
            source = COALESCE($3, source),
            description = COALESCE($4, description),
            location = COALESCE($5, location),
            salary_range = COALESCE($6, salary_range),
            job_type = COALESCE($7, job_type),
            status = COALESCE($8, status),
            application_link = COALESCE($9, application_link),
            rejection_reason = COALESCE($10, rejection_reason),
            notes = COALESCE($11, notes),
            interview_notes = COALESCE($12, interview_notes),
            next_interview_date = COALESCE($13, next_interview_date),
            last_interaction_date = COALESCE($14, last_interaction_date)
        WHERE id = $15 AND version = $16
        RETURNING version`

	DeleteIdQuery = `
        DELETE FROM Jobs
        WHERE id = $1 AND version = $2`

	DeleteAllQuery = `
        DELETE FROM Jobs`

	UpdateJobStatusQuery = `
        UPDATE Jobs
        SET status = COALESCE($1, status),
            last_interaction_date = CURRENT_TIMESTAMP
        WHERE id = $2 AND version = $3
        RETURNING version`

	GetJobsByBoardColumnQuery = `
                SELECT * FROM Jobs
                WHERE board_column = $1
                ORDER BY board_position ASC`

	UpdateJobBoardPositionQuery = `
                UPDATE Jobs
                SET board_column = $1,
                    board_position = $2,
                    updated_at = CURRENT_TIMESTAMP,
                    version = version + 1
                WHERE id = $3 AND version = $4
                RETURNING version`
)

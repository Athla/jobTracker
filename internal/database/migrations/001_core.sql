-- Up
CREATE TABLE IF NOT EXISTS jobs (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    company TEXT NOT NULL,
    source TEXT NOT NULL,
    description TEXT DEFAULT NULL,
    job_type TEXT CHECK (
        job_type IN (
            'FULL_TIME',
            'PART_TIME',
            'CONTRACT',
            'INTERNSHIP',
            'REMOTE'
        )
    ),
    status TEXT CHECK (
        status IN (
            'WISHLIST',
            'APPLIED',
            'PHONE_SCREEN',
            'TECHNICAL_INTERVIEW',
            'ONSITE',
            'OFFER',
            'REJECTED',
            'ACCEPTED',
            'WITHDRAWN'
        )
    ) DEFAULT 'WISHLIST',
    version INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add trigger for updated_at
CREATE TRIGGER IF NOT EXISTS update_jobs_timestamp AFTER
UPDATE ON jobs BEGIN
UPDATE jobs
SET
    updated_at = CURRENT_TIMESTAMP,
    version = version + 1
WHERE
    id = NEW.id;

END;

-- Down
DROP TABLE IF EXISTS jobs;

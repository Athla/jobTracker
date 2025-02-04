-- Up Migration
CREATE TABLE IF NOT EXISTS jobs (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    source TEXT NOT NULL,
    description TEXT,
    status TEXT DEFAULT 'pending' CHECK (
        status IN ('pending', 'in_progress', 'completed', 'rejected')
    ),
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

-- Down Migration
DROP TABLE IF EXISTS jobs;

DROP TABLE IF EXISTS users;

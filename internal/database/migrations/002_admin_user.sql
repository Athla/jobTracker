-- Up
CREATE TABLE IF NOT EXISTS admin_user (
    id INTEGER PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP
);

-- Create the initial admin user with environment variables
INSERT INTO
    admin_user (username, password_hash)
SELECT
    COALESCE(?, 'admin'), -- Will be replaced with ADMIN_USERNAME from env
    ? -- Will be replaced with hashed ADMIN_PASSWORD from env
WHERE
    NOT EXISTS (
        SELECT
            1
        FROM
            admin_user
        LIMIT
            1
    );

-- Down
DROP TABLE IF EXISTS admin_user;

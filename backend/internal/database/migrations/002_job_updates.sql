ALTER TABLE jobs
ADD COLUMN board_position INTEGER DEFAULT 0,
ADD COLUMN board_column TEXT CHECK (
    board_column IN ('APPLIED', 'IN_PROGRESS', 'FINISHED')
) DEFAULT 'APPLIED';

CREATE INDEX idx_jobs_board_column ON jobs (board_column);

CREATE INDEX idx_jobs_board_position ON jobs (board_position);

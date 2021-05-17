BEGIN;

ALTER TABLE analysis ALTER COLUMN created_at TYPE timestamp USING created_at::timestamp;

ALTER TABLE analysis ALTER COLUMN finished_at TYPE timestamp USING finished_at::timestamp;

COMMIT;

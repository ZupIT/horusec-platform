BEGIN;

ALTER TABLE analysis ALTER COLUMN created_at TYPE date USING created_at::date;

ALTER TABLE analysis ALTER COLUMN finished_at TYPE date USING finished_at::date;

COMMIT;

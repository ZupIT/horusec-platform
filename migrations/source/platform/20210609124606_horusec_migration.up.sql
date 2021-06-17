BEGIN;

CREATE TABLE IF NOT EXISTS "horusec_migrations"
(
    "name" TEXT NOT NULL,
    PRIMARY KEY (name)
);

COMMIT;

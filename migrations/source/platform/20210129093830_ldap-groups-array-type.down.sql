BEGIN;

ALTER TABLE IF EXISTS workspaces
    DROP COLUMN IF EXISTS authz_member,
    DROP COLUMN IF EXISTS authz_admin;

ALTER TABLE IF EXISTS repositories
    DROP COLUMN IF EXISTS authz_member,
    DROP COLUMN IF EXISTS authz_admin,
    DROP COLUMN IF EXISTS authz_supervisor;

COMMIT;
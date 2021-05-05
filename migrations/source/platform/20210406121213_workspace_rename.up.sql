BEGIN;

ALTER TABLE companies
    RENAME COLUMN company_id TO workspace_id;

ALTER TABLE companies
    RENAME TO workspaces;

ALTER TABLE account_company
    RENAME COLUMN company_id TO workspace_id;

ALTER TABLE account_company
    RENAME TO account_workspace;

ALTER TABLE analysis
    RENAME COLUMN company_id TO workspace_id;

ALTER TABLE analysis
    RENAME COLUMN company_name TO workspace_name;

ALTER TABLE webhooks
    RENAME COLUMN company_id TO workspace_id;

ALTER TABLE tokens
    RENAME COLUMN company_id TO workspace_id;

ALTER TABLE repositories
    RENAME COLUMN company_id TO workspace_id;

ALTER TABLE account_repository
    RENAME COLUMN company_id TO workspace_id;

COMMIT;

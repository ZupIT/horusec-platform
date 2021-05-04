BEGIN;

ALTER TABLE workspaces
    RENAME COLUMN workspace_id TO company_id;

ALTER TABLE workspaces
    RENAME TO companies;

ALTER TABLE account_workspace
    RENAME COLUMN workspace_id TO company_id;

ALTER TABLE account_workspace
    RENAME TO account_company;

ALTER TABLE analysis
    RENAME COLUMN workspace_id TO company_id;

ALTER TABLE analysis
    RENAME COLUMN workspace_name TO company_name;

ALTER TABLE webhooks
    RENAME COLUMN workspace_id TO company_id;

ALTER TABLE tokens
    RENAME COLUMN workspace_id TO company_id;

ALTER TABLE repositories
    RENAME COLUMN workspace_id TO company_id;

ALTER TABLE account_repository
    RENAME COLUMN workspace_id TO company_id;

COMMIT;

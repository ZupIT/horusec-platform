BEGIN;

DROP TABLE IF EXISTS vulnerabilities_by_author CASCADE;
DROP TABLE IF EXISTS vulnerabilities_by_language CASCADE;
DROP TABLE IF EXISTS vulnerabilities_by_repository CASCADE;
DROP TABLE IF EXISTS vulnerabilities_by_time CASCADE;

COMMIT;
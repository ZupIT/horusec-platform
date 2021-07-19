@echo off
setlocal enableextensions 
set POSTGRES_USER=root
set POSTGRES_PASSWORD=root
set POSTGRES_HOST=localhost
set POSTGRES_PORT=5432
set POSTGRES_SSL_MODE=disable
set HORUSEC_PLATFORM_DB_NAME=horusec_db
set HORUSEC_ANALYTIC_DB_NAME=horusec_analytic_db
set HORUSEC_DEFAULT_DATABASE_SQL_URI="postgres://%POSTGRES_USER%:%POSTGRES_PASSWORD%@%POSTGRES_HOST%:%POSTGRES_PORT%/%HORUSEC_PLATFORM_DB_NAME%?sslmode=%POSTGRES_SSL_MODE%"
set HORUSEC_ANALYTIC_DATABASE_SQL_URI="postgres://%POSTGRES_USER%:%POSTGRES_PASSWORD%@%POSTGRES_HOST%:%POSTGRES_PORT%/%HORUSEC_ANALYTIC_DB_NAME%?sslmode=%POSTGRES_SSL_MODE%"
SET CURRENTDIR="%cd%"

cd migrations

make build || goto :eof

cd ..


for /F "tokens=* USEBACKQ" %%g in (
    `docker exec -it horusec_postgresql psql -U %POSTGRES_USER% -W %POSTGRES_PASSWORD% --no-password -d postgres -c "select datname from pg_database WHERE datname = '%HORUSEC_PLATFORM_DB_NAME%'"`
) do (
    set "exists_horusec_db=%%g"
)

for /F "tokens=* USEBACKQ" %%g in (
    `docker exec -it horusec_postgresql psql -U %POSTGRES_USER% -W %POSTGRES_PASSWORD% --no-password -d postgres -c "select datname from pg_database WHERE datname = '%HORUSEC_ANALYTIC_DB_NAME%'"`
) do (
    set "exists_horusec_analytic_db=%%g"
)

if "()" == "%exists_horusec_db:0 rows=%" (
    echo "Creating database %HORUSEC_PLATFORM_DB_NAME%..."
    docker exec -it horusec_postgresql createdb %HORUSEC_PLATFORM_DB_NAME% -U %POSTGRES_USER% -W %POSTGRES_PASSWORD% --no-password
) else (
    echo "%HORUSEC_PLATFORM_DB_NAME% was found"
)

if "()" == "%exists_horusec_analytic_db:0 rows=%" (
    echo "Creating database %HORUSEC_ANALYTIC_DB_NAME%..."
    docker exec -it horusec_postgresql createdb %HORUSEC_ANALYTIC_DB_NAME% -U %POSTGRES_USER% -W %POSTGRES_PASSWORD% --no-password
) else (
    echo "%HORUSEC_ANALYTIC_DB_NAME% was found"
)

echo "Aplicando migracoes para o horusec platform..."
docker run --name migrate --rm -v "%CURRENTDIR%\migrations\source":"/horusec-migrations" -e HORUSEC_DATABASE_SQL_URI=%HORUSEC_DEFAULT_DATABASE_SQL_URI% -e MIGRATION_NAME=platform --network=container:horusec_postgresql horuszup/horusec-migrations:local "%1"

echo "Aplicando migracoes para o horusec analytic..."
docker run --name migrate --rm -v "%CURRENTDIR%\migrations\source":"/horusec-migrations" -e HORUSEC_DATABASE_SQL_URI=%HORUSEC_ANALYTIC_DATABASE_SQL_URI% -e MIGRATION_NAME=analytic --network=container:horusec_postgresql horuszup/horusec-migrations:local "%1"

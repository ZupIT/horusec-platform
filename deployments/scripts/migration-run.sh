#!/bin/bash
# Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

COMMAND=$@
POSTGRES_USER="root"
POSTGRES_PASSWORD="root"
POSTGRES_HOST="localhost"
POSTGRES_PORT="5432"
POSTGRES_SSL_MODE="disable"
HORUSEC_PLATFORM_DB_NAME="horusec_db"
HORUSEC_ANALYTIC_DB_NAME="horusec_analytic_db"
HORUSEC_DEFAULT_DATABASE_SQL_URI="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$HORUSEC_PLATFORM_DB_NAME?sslmode=$POSTGRES_SSL_MODE"
HORUSEC_ANALYTIC_DATABASE_SQL_URI="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$HORUSEC_ANALYTIC_DB_NAME?sslmode=$POSTGRES_SSL_MODE"

cd ./migrations

if ! make build; then
    exit 1
fi

cd ..

exists_horusec_db=$(docker exec -i horusec_postgresql psql -U $POSTGRES_USER -W $POSTGRES_PASSWORD --no-password -d postgres -c "select datname from pg_database WHERE datname = '$HORUSEC_PLATFORM_DB_NAME'")
if [[ $exists_horusec_db == *"0 rows"* ]]; then
    echo "Creating database $HORUSEC_PLATFORM_DB_NAME..."
    docker exec -i horusec_postgresql createdb $HORUSEC_PLATFORM_DB_NAME -U $POSTGRES_USER -W $POSTGRES_PASSWORD --no-password
fi

exists_horusec_analytic_db=$(docker exec -i horusec_postgresql psql -U $POSTGRES_USER -W $POSTGRES_PASSWORD --no-password -d postgres -c "select datname from pg_database WHERE datname = '$HORUSEC_ANALYTIC_DB_NAME'")
if [[ $exists_horusec_analytic_db == *"0 rows"* ]]; then
    echo "Creating database $HORUSEC_ANALYTIC_DB_NAME..."
    docker exec -i horusec_postgresql createdb $HORUSEC_ANALYTIC_DB_NAME -U $POSTGRES_USER -W $POSTGRES_PASSWORD --no-password
fi

echo ""

echo "Aplicando migrações para o horusec platform..."
docker run --name migrate --rm \
    -v "$(pwd)/migrations/source:/horusec-migrations" \
    -e HORUSEC_DATABASE_SQL_URI=$HORUSEC_DEFAULT_DATABASE_SQL_URI \
    -e MIGRATION_NAME=platform \
    --network=container:horusec_postgresql \
    horuszup/horusec-migrations:local \
    "$COMMAND"

echo ""

echo "Aplicando migrações para o horusec analytic..."
docker run --name migrate --rm \
    -v "$(pwd)/migrations/source:/horusec-migrations" \
    -e HORUSEC_DATABASE_SQL_URI=$HORUSEC_ANALYTIC_DATABASE_SQL_URI \
    -e MIGRATION_NAME=analytic \
    --network=container:horusec_postgresql \
    horuszup/horusec-migrations:local \
    "$COMMAND"

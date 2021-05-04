#!/bin/sh
# Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

POSTGRES_USER="root"
POSTGRES_PASSWORD="root"
POSTGRES_HOST="localhost"
POSTGRES_PORT="5432"
POSTGRES_SSL_MODE="disable"
HORUSEC_DEFAULT_DATABASE_SQL_URI="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/horusec_db?sslmode=$POSTGRES_SSL_MODE"
HORUSEC_ANALYTIC_DATABASE_SQL_URI="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/horusec_analytic_db?sslmode=$POSTGRES_SSL_MODE"

cd ./migrations

if ! make build; then
    exit 1
fi

cd ..

docker run --name migrate --rm \
    -v "$(pwd)/migrations/source:/horusec-migrations" \
    -e HORUSEC_DATABASE_SQL_URI=$HORUSEC_DEFAULT_DATABASE_SQL_URI \
    -e MIGRATION_NAME=platform \
    --network=container:postgresql \
    horuszup/horusec-migrations:local \
    "$@"

docker run --name migrate --rm \
    -v "$(pwd)/migrations/source:/horusec-migrations" \
    -e HORUSEC_DATABASE_SQL_URI=$HORUSEC_ANALYTIC_DATABASE_SQL_URI \
    -e MIGRATION_NAME=analytic \
    --network=container:postgresql \
    horuszup/horusec-migrations:local \
    "$@"
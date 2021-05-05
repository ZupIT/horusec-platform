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

/usr/local/bin/dockerd-entrypoint.sh &
sleep 5
docker run -d --hostname horusec-rabbit --name horusec-rabbit -e RABBITMQ_DEFAULT_USER=guest -e RABBITMQ_DEFAULT_PASS=guest -p 5672:5672 rabbitmq:3
docker run --name horusec-postgres -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root -e POSTGRES_DB=horusec_db -p 5432:5432 -d postgres
sleep 5
docker run -v "$(pwd)/migrations/source/platform:/migrations" --network host migrate/migrate -path=/migrations/ -database postgresql://root:root@127.0.0.1:5432/horusec_db?sslmode=disable up
/bin/horusec-auth-main

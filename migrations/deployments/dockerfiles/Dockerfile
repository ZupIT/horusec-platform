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

FROM golang:1.17-alpine AS builder

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:3.15.0

COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

COPY ./source /horusec-migrations

COPY ./scripts/migrate.sh /usr/local/bin

RUN chmod +x /usr/local/bin/migrate.sh
RUN chmod +x /usr/local/bin/migrate

ENTRYPOINT [ "migrate.sh" ]

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

RUN apk add build-base

ADD . .

RUN cd auth && go get -t -v -d ./...
RUN cd auth && OOS=linux go build -a -o /bin/horusec-auth-main ./cmd/app/main.go

RUN cd core && go get -t -v -d ./...
RUN cd core && OOS=linux go build -a -o /bin/horusec-core-main ./cmd/app/main.go

RUN cd api && go get -t -v -d ./...
RUN cd api && OOS=linux go build -a -o /bin/horusec-api-main ./cmd/app/main.go

RUN cd analytic && go get -t -v -d ./...
RUN cd analytic && OOS=linux go build -a -o /bin/horusec-analytic-main ./cmd/app/main.go

RUN cd vulnerability && go get -t -v -d ./...
RUN cd vulnerability && OOS=linux go build -a -o /bin/horusec-vulnerability-main ./cmd/app/main.go

RUN cd webhook && go get -t -v -d ./...
RUN cd webhook && OOS=linux go build -a -o /bin/horusec-webhook-main ./cmd/app/main.go

FROM node:16.13.1-alpine3.14 AS builder-manager

WORKDIR /usr/src/app

ADD . .

RUN cd manager && yarn
RUN cd manager && yarn build

FROM docker:dind

ENV HORUSEC_DISABLE_EMAILS=true
ENV HORUSEC_ENABLE_DEFAULT_USER=true

RUN apk add nginx
RUN mkdir -p /run/nginx

COPY --from=builder-backend /bin/horusec-auth-main /bin/horusec-auth-main
COPY --from=builder-backend /bin/horusec-core-main /bin/horusec-core-main
COPY --from=builder-backend /bin/horusec-api-main /bin/horusec-api-main
COPY --from=builder-backend /bin/horusec-analytic-main /bin/horusec-analytic-main
COPY --from=builder-backend /bin/horusec-vulnerability-main /bin/horusec-vulnerability-main
COPY --from=builder-backend /bin/horusec-webhook-main /bin/horusec-webhook-main

COPY deployments/dockerfiles/all-in-one/entrypoint.sh /usr/local/bin
RUN chmod +x /usr/local/bin/entrypoint.sh

COPY deployments/scripts/migration-run.sh /usr/local/bin
RUN chmod +x /usr/local/bin/migration-run.sh

COPY migrations/source /migrations/source

COPY --from=builder-manager /usr/src/app/manager/build /var/www
COPY ./manager/deployments/nginx.conf /etc/nginx/conf.d/default.conf

ENTRYPOINT [ "entrypoint.sh" ]

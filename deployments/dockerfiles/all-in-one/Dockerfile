FROM golang:alpine AS builder-backend

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

FROM node:alpine AS builder-manager

WORKDIR /usr/src/app

ADD . .

RUN cd manager && npm install
RUN cd manager && npm run build

FROM docker:dind

ENV HORUSEC_DISABLE_BROKER=true

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

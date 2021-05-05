FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git build-base

ADD . .

RUN cd auth && go get -t -v -d ./...

RUN cd auth && OOS=linux go build -a -o /bin/horusec-auth-main ./cmd/app/main.go

FROM docker:dind

COPY --from=builder /bin/horusec-auth-main /bin/horusec-auth-main

COPY deployments/dockerfiles/all-in-one/entrypoint.sh /usr/local/bin
RUN chmod +x /usr/local/bin/entrypoint.sh

COPY deployments/scripts/migration-run.sh /usr/local/bin
RUN chmod +x /usr/local/bin/migration-run.sh

COPY migrations/source /migrations/source

ENTRYPOINT [ "entrypoint.sh" ]

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

RUN apk update && apk add --no-cache git build-base

ADD . /analytic

WORKDIR /analytic

RUN go get -t -v -d ./...

RUN GOOS=linux go build -a -o horusec-analytic-main ./cmd/app/main.go
RUN GOOS=linux go build -a -o horusec-analytic-v1-to-v2-migrate ./cmd/migration/v2/main.go

FROM alpine:3.14.2

COPY --from=builder /analytic/horusec-analytic-main .
COPY --from=builder /analytic/horusec-analytic-v1-to-v2-migrate .

ENTRYPOINT ["./horusec-analytic-main"]

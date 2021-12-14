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

ADD . /vulnerability

WORKDIR /vulnerability

RUN go get -t -v -d ./...

RUN GOOS=linux go build -a -o horusec-vulnerability-main ./cmd/app/main.go

FROM alpine:3.15.0

COPY --from=builder /vulnerability/horusec-vulnerability-main .

ENTRYPOINT ["./horusec-vulnerability-main"]

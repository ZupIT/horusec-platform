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

FROM node:16.8.0-alpine3.14 AS builder

COPY . /manager

WORKDIR /manager

RUN npm install
RUN npm run build

FROM nginx:1.21.1-alpine

COPY --from=builder /manager/build /var/www
COPY ./deployments/nginx.conf /etc/nginx/conf.d/default.conf
COPY ./deployments/scripts/run-nginx.sh /var/www/run-nginx.sh

CMD ["sh", "/var/www/run-nginx.sh"]

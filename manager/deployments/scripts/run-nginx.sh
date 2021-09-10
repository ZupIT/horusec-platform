#! /bin/bash
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

# Horusec v2
sed -i -e "s/window.HORUSEC_ENDPOINT_API=\"\"/window.HORUSEC_ENDPOINT_API=\"$HORUSEC_ENDPOINT_API\"/g" "/var/www/index.html"
sed -i -e "s/window.HORUSEC_ENDPOINT_ANALYTIC=\"\"/window.HORUSEC_ENDPOINT_ANALYTIC=\"$HORUSEC_ENDPOINT_ANALYTIC\"/g" "/var/www/index.html"
sed -i -e "s/window.HORUSEC_ENDPOINT_CORE=\"\"/window.HORUSEC_ENDPOINT_CORE=\"$HORUSEC_ENDPOINT_CORE\"/g" "/var/www/index.html"
sed -i -e "s/window.HORUSEC_ENDPOINT_AUTH=\"\"/window.HORUSEC_ENDPOINT_AUTH=\"$HORUSEC_ENDPOINT_AUTH\"/g" "/var/www/index.html"
sed -i -e "s/window.HORUSEC_ENDPOINT_VULNERABILITY=\"\"/window.HORUSEC_ENDPOINT_VULNERABILITY=\"$HORUSEC_ENDPOINT_VULNERABILITY\"/g" "/var/www/index.html"
sed -i -e "s/window.HORUSEC_ENDPOINT_WEBHOOK=\"\"/window.HORUSEC_ENDPOINT_WEBHOOK=\"$HORUSEC_ENDPOINT_WEBHOOK\"/g" "/var/www/index.html"

# Horusec V1
sed -i -e "s/window.REACT_APP_HORUSEC_ENDPOINT_API=\"\"/window.REACT_APP_HORUSEC_ENDPOINT_API=\"$REACT_APP_HORUSEC_ENDPOINT_API\"/g" "/var/www/index.html"
sed -i -e "s/window.REACT_APP_HORUSEC_ENDPOINT_ANALYTIC=\"\"/window.REACT_APP_HORUSEC_ENDPOINT_ANALYTIC=\"$REACT_APP_HORUSEC_ENDPOINT_ANALYTIC\"/g" "/var/www/index.html"
sed -i -e "s/window.REACT_APP_HORUSEC_ENDPOINT_CORE=\"\"/window.REACT_APP_HORUSEC_ENDPOINT_CORE=\"$REACT_APP_HORUSEC_ENDPOINT_CORE\"/g" "/var/www/index.html"
sed -i -e "s/window.REACT_APP_HORUSEC_ENDPOINT_AUTH=\"\"/window.REACT_APP_HORUSEC_ENDPOINT_AUTH=\"$REACT_APP_HORUSEC_ENDPOINT_AUTH\"/g" "/var/www/index.html"
sed -i -e "s/window.REACT_APP_HORUSEC_ENDPOINT_VULNERABILITY=\"\"/window.REACT_APP_HORUSEC_ENDPOINT_VULNERABILITY=\"$REACT_APP_HORUSEC_ENDPOINT_VULNERABILITY\"/g" "/var/www/index.html"
sed -i -e "s/window.REACT_APP_HORUSEC_ENDPOINT_WEBHOOK=\"\"/window.REACT_APP_HORUSEC_ENDPOINT_WEBHOOK=\"$REACT_APP_HORUSEC_ENDPOINT_WEBHOOK\"/g" "/var/www/index.html"


sed -i -e "s/window.REACT_APP_HORUSEC_MANAGER_PATH=\"\"/window.REACT_APP_HORUSEC_MANAGER_PATH=\"$REACT_APP_HORUSEC_MANAGER_PATH\"/g" "/var/www/index.html"

sed -i -e "s/window.REACT_APP_KEYCLOAK_CLIENT_ID=\"\"/window.REACT_APP_KEYCLOAK_CLIENT_ID=\"$REACT_APP_KEYCLOAK_CLIENT_ID\"/g" "/var/www/index.html"
sed -i -e "s/window.REACT_APP_KEYCLOAK_REALM=\"\"/window.REACT_APP_KEYCLOAK_REALM=\"$REACT_APP_KEYCLOAK_REALM\"/g" "/var/www/index.html"
sed -i -e "s/window.REACT_APP_KEYCLOAK_BASE_PATH=\"\"/window.REACT_APP_KEYCLOAK_BASE_PATH=\"$REACT_APP_KEYCLOAK_BASE_PATH\"/g" "/var/www/index.html"

nginx -g "daemon off;"

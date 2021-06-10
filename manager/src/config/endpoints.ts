/**
 * Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
declare global {
  interface Window {
    HORUSEC_ENDPOINT_CORE: string;
    HORUSEC_ENDPOINT_ANALYTIC: string;
    HORUSEC_ENDPOINT_AUTH: string;
    HORUSEC_ENDPOINT_VULNERABILITY: string;
    HORUSEC_ENDPOINT_WEBHOOK: string;
  }
}

const SERVICE_VULNERABILITY =
  window.HORUSEC_ENDPOINT_VULNERABILITY || 'http://127.0.0.1:8001';

const SERVICE_CORE = window.HORUSEC_ENDPOINT_CORE || 'http://127.0.0.1:8003';

const SERVICE_WEBHOOK =
  window.HORUSEC_ENDPOINT_WEBHOOK || 'http://127.0.0.1:8004';

const SERVICE_ANALYTIC =
  window.HORUSEC_ENDPOINT_ANALYTIC || 'http://127.0.0.1:8005';

const SERVICE_AUTH = window.HORUSEC_ENDPOINT_AUTH || 'http://127.0.0.1:8006';

const isLocalHost = (endpoint: string) =>
  endpoint.includes('localhost') || endpoint.includes('127.0.0.1');

if (
  isLocalHost(SERVICE_AUTH) ||
  isLocalHost(SERVICE_CORE) ||
  isLocalHost(SERVICE_VULNERABILITY) ||
  isLocalHost(SERVICE_ANALYTIC) ||
  isLocalHost(SERVICE_WEBHOOK)
) {
  console.warn(`📡 One or more addresses of Horusec services have not been defined
or have been defined as localhost.
If this is not the scenario for your application, visit the guide:
How to run the web application in other host?
https://horusec.io/docs/tutorials/how-to-run-the-web-application-on-other-host`);
}

export {
  SERVICE_CORE,
  SERVICE_ANALYTIC,
  SERVICE_AUTH,
  SERVICE_VULNERABILITY,
  SERVICE_WEBHOOK,
};

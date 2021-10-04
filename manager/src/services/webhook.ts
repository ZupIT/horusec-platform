/**
 * Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

import { AxiosResponse } from 'axios';
import http from 'config/axios';
import { WebhookHeader } from 'helpers/interfaces/Webhook';
import { SERVICE_WEBHOOK } from '../config/endpoints';

const getAll = (workspaceID: string): Promise<AxiosResponse<any>> => {
  return http.get(`${SERVICE_WEBHOOK}/webhook/webhook/${workspaceID}`);
};

const create = (
  workspaceID: string,
  repositoryID: string,
  url: string,
  method: string,
  headers: WebhookHeader[],
  description: string
): Promise<AxiosResponse<any>> => {
  return http.post(`${SERVICE_WEBHOOK}/webhook/webhook/${workspaceID}`, {
    url,
    method,
    headers,
    description,
    repositoryID,
    workspaceID,
  });
};

const update = (
  workspaceID: string,
  repositoryID: string,
  webhookID: string,
  url: string,
  method: string,
  headers: WebhookHeader[],
  description: string
): Promise<AxiosResponse<any>> => {
  return http.put(
    `${SERVICE_WEBHOOK}/webhook/webhook/${workspaceID}/${webhookID}`,
    {
      url,
      method,
      headers,
      description,
      workspaceID,
      repositoryID,
    }
  );
};

const remove = (
  webhookID: string,
  workspaceID: string
): Promise<AxiosResponse<any>> => {
  return http.delete(
    `${SERVICE_WEBHOOK}/webhook/webhook/${workspaceID}/${webhookID}`
  );
};

export default {
  getAll,
  remove,
  create,
  update,
};

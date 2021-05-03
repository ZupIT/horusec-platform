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

import http from 'config/axios';
import { LDAPGroups } from 'helpers/interfaces/LDAPGroups';
import { SERVICE_CORE } from '../config/endpoints';

const getAll = () => {
  return http.get(`${SERVICE_CORE}/core/workspaces`);
};

const getOne = (workspaceID: string) => {
  return http.get(`${SERVICE_CORE}/core/workspaces/${workspaceID}`);
};

const create = (
  name: string,
  description?: string,
  ldapGroups?: LDAPGroups
) => {
  return http.post(`${SERVICE_CORE}/core/workspaces`, {
    name,
    description,
    ...ldapGroups,
  });
};

const update = (
  workspaceID: string,
  name: string,
  description?: string,
  ldapGroups?: LDAPGroups
) => {
  return http.patch(`${SERVICE_CORE}/core/workspaces/${workspaceID}`, {
    name,
    description,
    ...ldapGroups,
  });
};

const remove = (workspaceID: string) => {
  return http.delete(`${SERVICE_CORE}/core/workspaces/${workspaceID}`);
};

const getUsersInWorkspace = (workspaceID: string) => {
  return http.get(`${SERVICE_CORE}/core/workspaces/${workspaceID}/roles`);
};

const createUserInWorkspace = (
  workspaceID: string,
  email: string,
  role: string
) => {
  return http.post(`${SERVICE_CORE}/core/workspaces/${workspaceID}/roles`, {
    email,
    role,
  });
};

const editUserInWorkspace = (
  workspaceID: string,
  accountId: string,
  role: string
) => {
  return http.patch(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/roles/${accountId}`,
    {
      role,
    }
  );
};

const removeUserInWorkspace = (workspaceID: string, accountId: string) => {
  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/roles/${accountId}`
  );
};

const getAllTokens = (workspaceID: string) => {
  return http.get(`${SERVICE_CORE}/core/workspaces/${workspaceID}/tokens`);
};

const createToken = (
  workspaceID: string,
  data: {
    description: string;
    isExpirable?: boolean;
    expiredAt?: string;
  }
) => {
  return http.post(`${SERVICE_CORE}/core/workspaces/${workspaceID}/tokens`, {
    ...data,
  });
};

const removeToken = (workspaceID: string, tokenId: string) => {
  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/tokens/${tokenId}`
  );
};

export default {
  getAll,
  create,
  update,
  remove,
  getOne,
  getUsersInWorkspace,
  createUserInWorkspace,
  editUserInWorkspace,
  removeUserInWorkspace,
  createToken,
  removeToken,
  getAllTokens,
};

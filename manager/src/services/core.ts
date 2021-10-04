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
import { APIResponse } from 'helpers/interfaces/APIResponse';
import { LDAPGroups } from 'helpers/interfaces/LDAPGroups';
import { Repository } from 'helpers/interfaces/Repository';
import { SERVICE_CORE } from '../config/endpoints';

const getAllWorkspaces = (): Promise<AxiosResponse<any>> => {
  return http.get(`${SERVICE_CORE}/core/workspaces`);
};

const getOneWorkspace = (workspaceID: string): Promise<AxiosResponse<any>> => {
  return http.get(`${SERVICE_CORE}/core/workspaces/${workspaceID}`);
};

const createWorkspace = (
  name: string,
  description?: string,
  ldapGroups?: LDAPGroups
): Promise<AxiosResponse<any>> => {
  return http.post(`${SERVICE_CORE}/core/workspaces`, {
    name,
    description,
    ...ldapGroups,
  });
};

const updateWorkspace = (
  workspaceID: string,
  name: string,
  description?: string,
  ldapGroups?: LDAPGroups
): Promise<AxiosResponse<any>> => {
  return http.patch(`${SERVICE_CORE}/core/workspaces/${workspaceID}`, {
    name,
    description,
    ...ldapGroups,
  });
};

const deleteWorkspace = (workspaceID: string): Promise<AxiosResponse<any>> => {
  return http.delete(`${SERVICE_CORE}/core/workspaces/${workspaceID}`);
};

const getUsers = (
  workspaceID: string,
  repositoryID: string,
  notBelong?: string
): Promise<AxiosResponse<any>> => {
  const path = repositoryID ? `/repositories/${repositoryID}/roles` : '/roles';
  return http.get(`${SERVICE_CORE}/core/workspaces/${workspaceID}${path}`, {
    params: {
      notBelong,
    },
  });
};

const inviteUser = (
  workspaceID: string,
  email: string,
  role: string,
  repositoryId?: string,
  accountID?: string,
  username?: string
): Promise<AxiosResponse<any>> => {
  const path = repositoryId ? `/repositories/${repositoryId}/` : '/';

  return http.post(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}${path}roles`,
    {
      email,
      role,
      accountID,
      username,
    }
  );
};

const updateUserRole = (
  workspaceID: string,
  repositoryId: string,
  accountId: string,
  role: string
): Promise<AxiosResponse<any>> => {
  const path = repositoryId ? `/repositories/${repositoryId}/` : '/';

  return http.patch(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}${path}roles/${accountId}`,
    {
      role,
    }
  );
};

const removeUser = (
  workspaceID: string,
  repositoryId: string,
  accountId: string
): Promise<AxiosResponse<any>> => {
  const path = repositoryId ? `/repositories/${repositoryId}/` : '/';

  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}${path}roles/${accountId}`
  );
};

const getAllTokens = (
  workspaceID: string,
  repositoryId?: string
): Promise<AxiosResponse<any>> => {
  const path = repositoryId ? `/repositories/${repositoryId}` : '';
  return http.get(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}${path}/tokens`
  );
};

const removeToken = (
  data: { workspaceID: string; repositoryID?: string },
  tokenId: string
): Promise<AxiosResponse<any>> => {
  const path = data.repositoryID ? `/repositories/${data.repositoryID}` : '';
  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${data.workspaceID}${path}/tokens/${tokenId}`
  );
};

const getAllRepositories = (
  workspaceID: string,
  page?: number,
  search?: string
): Promise<AxiosResponse<any>> => {
  return http.get(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories`,
    {
      params: { page, size: 15, search },
    }
  );
};

const getOneRepository = (
  workspaceID: string,
  repositoryID: string
): Promise<AxiosResponse<any>> => {
  return http.get(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryID}`
  );
};

const createRepository = (
  workspaceID: string,
  name: string,
  description: string,
  ldapGroups?: LDAPGroups
): Promise<AxiosResponse<any>> => {
  return http.post(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories`,
    {
      name,
      description,
      ...ldapGroups,
    }
  );
};

const updateRepository = (
  workspaceID: string,
  repositoryId: string,
  name: string,
  description: string,
  ldapGroups?: LDAPGroups
): Promise<AxiosResponse<any>> => {
  return http.patch(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}`,
    { name, description, ...ldapGroups }
  );
};

const deleteRepository = (
  workspaceID: string,
  repositoryId: string
): Promise<AxiosResponse<any>> => {
  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}`
  );
};

const createToken = (
  params: { workspaceID: string; repositoryID?: string },
  data: {
    description: string;
    isExpirable?: boolean;
    expiredAt?: string;
  }
): Promise<AxiosResponse<any>> => {
  const path = params.repositoryID
    ? `/repositories/${params.repositoryID}`
    : '';
  return http.post(
    `${SERVICE_CORE}/core/workspaces/${params.workspaceID}${path}/tokens`,
    {
      ...data,
    }
  );
};

export default {
  getAllWorkspaces,
  createWorkspace,
  updateWorkspace,
  deleteWorkspace,
  getOneWorkspace,
  getAllRepositories,
  getOneRepository,
  createRepository,
  updateRepository,
  deleteRepository,
  getUsers,
  removeUser,
  inviteUser,
  updateUserRole,
  getAllTokens,
  createToken,
  removeToken,
};

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
import { APIResponse } from 'helpers/interfaces/APIResponse';
import { LDAPGroups } from 'helpers/interfaces/LDAPGroups';
import { Repository } from 'helpers/interfaces/Repository';
import { SERVICE_CORE } from '../config/endpoints';

const getAllWorkspaces = () => {
  return http.get(`${SERVICE_CORE}/core/workspaces`);
};

const getOneWorkspace = (workspaceID: string) => {
  return http.get(`${SERVICE_CORE}/core/workspaces/${workspaceID}`);
};

const createWorkspace = (
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

const updateWorkspace = (
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

const deleteWorkspace = (workspaceID: string) => {
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

const getAllTokens = (workspaceID: string, repositoryId?: string) => {
  const path = repositoryId ? `/repositories/${repositoryId}` : '';
  return http.get(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}${path}/tokens`
  );
};

const getAllTokensOfWorkspace = (workspaceID: string) => {
  return http.get(`${SERVICE_CORE}/core/workspaces/${workspaceID}/tokens`);
};

const createTokenInWorkspace = (
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

const removeTokenOfWorkspace = (workspaceID: string, tokenId: string) => {
  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/tokens/${tokenId}`
  );
};

const removeToken = (
  data: { workspaceID: string; repositoryID?: string },
  tokenId: string
) => {
  const path = data.repositoryID ? `/repositories/${data.repositoryID}` : '';
  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${data.workspaceID}${path}/tokens/${tokenId}`
  );
};

const getAllRepositories = (workspaceID: string) => {
  return http.get(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories`
  );
};

const getOneRepository = (workspaceID: string, repositoryID: string) => {
  return http.get(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryID}`
  );
};

const createRepository = (
  workspaceID: string,
  name: string,
  description: string,
  ldapGroups?: LDAPGroups
) => {
  return http.post<APIResponse<Repository>>(
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
) => {
  return http.patch<APIResponse<Repository>>(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}`,
    { name, description, ...ldapGroups }
  );
};

const deleteRepository = (workspaceID: string, repositoryId: string) => {
  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}`
  );
};

const getAllTokensOfRepository = (
  workspaceID: string,
  repositoryId: string
) => {
  return http.get(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}/tokens`
  );
};

const createTokenInRepository = (
  workspaceID: string,
  repositoryId: string,
  data: {
    description: string;
    isExpirable?: boolean;
    expiredAt?: string;
  }
) => {
  return http.post(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}/tokens`,
    {
      ...data,
    }
  );
};

const removeTokenOfRepository = (
  workspaceID: string,
  repositoryId: string,
  tokenId: string
) => {
  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}/tokens/${tokenId}`
  );
};

const getUsersInRepository = (workspaceID: string, repositoryId: string) => {
  return http.get(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}/roles`
  );
};

const includeUserInRepository = (
  workspaceID: string,
  repositoryId: string,
  email: string,
  role: string,
  accountID: string,
  username: string
) => {
  return http.post(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}/roles`,
    {
      email,
      role,
      accountID,
      username,
    }
  );
};
const removeUserOfRepository = (
  workspaceID: string,
  repositoryId: string,
  accountId: string
) => {
  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}/roles/${accountId}`
  );
};

const updateUserRoleInRepository = (
  workspaceID: string,
  repositoryId: string,
  accountId: string,
  role: string
) => {
  return http.patch(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}/roles/${accountId}`,
    {
      role,
    }
  );
};

export default {
  getAllWorkspaces,
  createWorkspace,
  updateWorkspace,
  deleteWorkspace,
  getOneWorkspace,
  getUsersInWorkspace,
  createUserInWorkspace,
  editUserInWorkspace,
  removeUserInWorkspace,
  createTokenInWorkspace,
  removeTokenOfWorkspace,
  getAllTokensOfWorkspace,
  getAllTokens,
  getAllRepositories,
  getOneRepository,
  createRepository,
  updateRepository,
  deleteRepository,
  getAllTokensOfRepository,
  createTokenInRepository,
  removeTokenOfRepository,
  removeToken,
  getUsersInRepository,
  includeUserInRepository,
  removeUserOfRepository,
  updateUserRoleInRepository,
};

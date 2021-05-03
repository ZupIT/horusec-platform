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

import { SERVICE_CORE } from '../config/endpoints';
import { FilterVuln } from 'helpers/interfaces/FIlterVuln';
import { PaginationInfo } from 'helpers/interfaces/Pagination';
import { LDAPGroups } from 'helpers/interfaces/LDAPGroups';

const getAll = (workspaceID: string) => {
  return http.get(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories`
  );
};

const create = (
  workspaceID: string,
  name: string,
  description: string,
  ldapGroups?: LDAPGroups
) => {
  return http.post(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories`,
    {
      name,
      description,
      ...ldapGroups,
    }
  );
};

const update = (
  workspaceID: string,
  repositoryId: string,
  name: string,
  description: string,
  ldapGroups?: LDAPGroups
) => {
  return http.patch(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}`,
    { name, description, ...ldapGroups }
  );
};

const remove = (workspaceID: string, repositoryId: string) => {
  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}`
  );
};

const getAllTokens = (workspaceID: string, repositoryId: string) => {
  return http.get(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}/tokens`
  );
};

const createToken = (
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

const removeToken = (
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

const includeUser = (
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
const removeUser = (
  workspaceID: string,
  repositoryId: string,
  accountId: string
) => {
  return http.delete(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}/roles/${accountId}`
  );
};

const updateUserRole = (
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

const getAllVulnerabilities = (
  filters: FilterVuln,
  pagination: PaginationInfo
) => {
  return http.get(
    `${SERVICE_CORE}/core/workspaces/${filters.workspaceID}/repositories/${filters.repositoryID}/management`,
    {
      params: {
        page: pagination.currentPage,
        size: pagination.pageSize,
        vulnSeverity: filters.vulnSeverity,
        vulnHash: filters.vulnHash,
        vulnType: filters.vulnType,
      },
    }
  );
};

const updateVulnerabilityType = (
  workspaceID: string,
  repositoryId: string,
  vulnerabilityId: string,
  type: string
) => {
  return http.put(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}/management/${vulnerabilityId}/type`,
    {
      type,
    }
  );
};

const updateVulnerabilitySeverity = (
  workspaceID: string,
  repositoryId: string,
  vulnerabilityId: string,
  severity: string
) => {
  return http.put(
    `${SERVICE_CORE}/core/workspaces/${workspaceID}/repositories/${repositoryId}/management/${vulnerabilityId}/severity`,
    {
      severity,
    }
  );
};

export default {
  getAll,
  create,
  update,
  remove,
  getAllTokens,
  createToken,
  removeToken,
  getUsersInRepository,
  includeUser,
  removeUser,
  updateUserRole,
  getAllVulnerabilities,
  updateVulnerabilityType,
  updateVulnerabilitySeverity,
};

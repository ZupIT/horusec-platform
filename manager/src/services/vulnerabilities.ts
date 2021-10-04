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
import { FilterVuln } from 'helpers/interfaces/FIlterVuln';
import { PaginationInfo } from 'helpers/interfaces/Pagination';
import { SERVICE_VULNERABILITY } from '../config/endpoints';

const getAllVulnerabilities = (
  filters: FilterVuln,
  type: 'workspace' | 'repository',
  pagination: PaginationInfo
): Promise<AxiosResponse<any>> => {
  const path =
    type === 'repository'
      ? `workspace/${filters.workspaceID}/repository/${filters.repositoryID}`
      : `workspace/${filters.workspaceID}`;

  return http.get(`${SERVICE_VULNERABILITY}/vulnerability/management/${path}`, {
    params: {
      page: pagination.currentPage,
      size: pagination.pageSize,
      vulnSeverity: filters.vulnSeverity,
      vulnHash: filters.vulnHash,
      vulnType: filters.vulnType,
    },
  });
};

const updateVulnerability = (
  workspaceID: string,
  repositoryID: string,
  analysisID: string,
  vulnerabilities: {
    severity: string;
    type: string;
    vulnerabilityID: string;
  }[],
  overviewType: 'repository' | 'workspace'
): Promise<AxiosResponse<any>> => {
  const route = `${SERVICE_VULNERABILITY}/vulnerability/management/workspace/${workspaceID}${
    overviewType === 'repository'
      ? `/repository/${repositoryID}/vulnerabilities`
      : '/vulnerabilities'
  }`;

  return http.patch(route, {
    analysisID,
    vulnerabilities,
  });
};

export default {
  getAllVulnerabilities,
  updateVulnerability,
};

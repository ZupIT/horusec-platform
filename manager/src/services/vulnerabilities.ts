import http from 'config/axios';
import { FilterVuln } from 'helpers/interfaces/FIlterVuln';
import { PaginationInfo } from 'helpers/interfaces/Pagination';
import { SERVICE_VULNERABILITY } from '../config/endpoints';

const getAllVulnerabilities = (
  filters: FilterVuln,
  type: 'workspace' | 'repository',
  pagination: PaginationInfo
) => {
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
  }[]
) => {
  return http.patch(
    `${SERVICE_VULNERABILITY}/vulnerability/management/workspace/${workspaceID}/repository/${repositoryID}/vulnerabilities`,
    {
      analysisID,
      vulnerabilities,
    }
  );
};

export default {
  getAllVulnerabilities,
  updateVulnerability,
};

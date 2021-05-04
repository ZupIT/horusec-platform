import http from 'config/axios';
import { FilterVuln } from 'helpers/interfaces/FIlterVuln';
import { PaginationInfo } from 'helpers/interfaces/Pagination';
import { SERVICE_VULNERABILITY } from '../config/endpoints';

const getAllVulnerabilities = (
  filters: FilterVuln,
  pagination: PaginationInfo
) => {
  const path = filters.repositoryID
    ? `repository/${filters.workspaceID}/${filters.repositoryID}`
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
  vulnerabilityID: string,
  type: string,
  severity: string
) => {
  return http.patch(
    `${SERVICE_VULNERABILITY}/vulnerability/management/${workspaceID}/${vulnerabilityID}`,
    {
      type,
      severity,
    }
  );
};

export default {
  getAllVulnerabilities,
  updateVulnerability,
};

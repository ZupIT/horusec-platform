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

import React, { useEffect, useState } from 'react';
import Styled from './styled';
import { SearchBar, Select, Icon, Datatable, DataSource } from 'components';
import { useTranslation } from 'react-i18next';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import vulnerabilitiesService from 'services/vulnerabilities';
import { PaginationInfo } from 'helpers/interfaces/Pagination';
import { Vulnerability } from 'helpers/interfaces/Vulnerability';
import { cloneDeep, debounce, get } from 'lodash';
import Details from './Details';
import { FilterVuln } from 'helpers/interfaces/FIlterVuln';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import { useTheme } from 'styled-components';
import { AxiosError, AxiosResponse } from 'axios';
import useParamsRoute from 'helpers/hooks/useParamsRoute';
import usePermissions from 'helpers/hooks/usePermissions';

const INITIAL_PAGE = 1;
interface RefreshInterface {
  filter: FilterVuln;
  page: PaginationInfo;
}

interface KeyValueVuln {
  vulnerabilityID: string;
  type: string;
  severity: string;
}

const Vulnerabilities: React.FC = () => {
  const { workspaceId, repositoryId } = useParamsRoute();

  const overviewType = repositoryId ? 'repository' : 'workspace';

  const { t } = useTranslation();
  const { colors } = useTheme();
  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();
  const { ACTIONS, isAuthorizedAction } = usePermissions();

  const [isLoading, setLoading] = useState(false);
  const [vulnerabilities, setVulnerabilities] = useState<Vulnerability[]>([]);
  const [selectedVuln, setSelectedVuln] = useState<Vulnerability>(null);
  const [updateVulnIds, setUpdateVulnIds] = useState<KeyValueVuln[]>([]);

  const [filters, setFilters] = useState<FilterVuln>({
    workspaceID: workspaceId,
    repositoryID: repositoryId,
    vulnHash: '',
    vulnSeverity: 'ALL',
    vulnType: 'ALL',
  });

  const [pagination, setPagination] = useState<PaginationInfo>({
    currentPage: INITIAL_PAGE,
    totalItems: 0,
    pageSize: 10,
    totalPages: 10,
  });

  const [refresh, setRefresh] = useState<RefreshInterface>({
    filter: filters,
    page: pagination,
  });

  const vulnTypes = [
    {
      label: t('VULNERABILITIES_SCREEN.ALL_STATUS'),
      value: 'ALL',
    },
    {
      label: t('VULNERABILITIES_SCREEN.STATUS.VULNERABILITY'),
      value: 'Vulnerability',
    },
    {
      label: t('VULNERABILITIES_SCREEN.STATUS.RISKACCEPTED'),
      value: 'Risk Accepted',
    },
    {
      label: t('VULNERABILITIES_SCREEN.STATUS.FALSEPOSITIVE'),
      value: 'False Positive',
    },
    {
      label: t('VULNERABILITIES_SCREEN.STATUS.CORRECTED'),
      value: 'Corrected',
    },
  ];

  const severities = [
    {
      label: t('VULNERABILITIES_SCREEN.ALL_SEVERITIES'),
      value: 'ALL',
    },
    {
      label: 'CRITICAL',
      value: 'CRITICAL',
    },
    {
      label: 'HIGH',
      value: 'HIGH',
    },
    {
      label: 'MEDIUM',
      value: 'MEDIUM',
    },
    {
      label: 'LOW',
      value: 'LOW',
    },
    {
      label: 'INFO',
      value: 'INFO',
    },
    {
      label: 'UNKNOWN',
      value: 'UNKNOWN',
    },
  ];

  const handleSearch = debounce((searchString: string) => {
    setRefresh((state) => ({
      ...state,
      filter: { ...state.filter, vulnHash: searchString },
    }));
  }, 800);

  const handleUpdateVulnerability = () => {
    vulnerabilitiesService
      .updateVulnerability(
        filters.workspaceID,
        filters.repositoryID,
        vulnerabilities[0].analysisID,
        updateVulnIds,
        overviewType
      )
      .then(() => {
        resetUpdateVuln();
        showSuccessFlash(t('VULNERABILITIES_SCREEN.SUCCESS_UPDATE'));
      })
      .catch((err: AxiosError) => {
        setRefresh((state) => state);
        dispatchMessage(err?.response?.data);
      });
  };

  useEffect(() => {
    let isCancelled = false;

    const fetchVulnerabilities = () => {
      setLoading(true);

      const page = refresh.page;
      const filter = refresh.filter;

      if (page.pageSize !== pagination.pageSize) {
        page.currentPage = INITIAL_PAGE;
      }

      const filterAux = {
        ...filter,
        vulnSeverity: filter.vulnHash ? null : filter.vulnSeverity,
        vulnType: filter.vulnHash ? null : filter.vulnType,
      };

      setFilters(filter);

      vulnerabilitiesService
        .getAllVulnerabilities(filterAux, overviewType, page)
        .then((result: AxiosResponse) => {
          if (!isCancelled) {
            const response = result.data?.content;

            const data: Vulnerability[] = response?.data;

            for (const row of data) {
              const { type = row.type, severity = row.severity } =
                updateVulnIds.find(
                  (x) => x.vulnerabilityID === row.vulnerabilityID
                ) || {};
              row.type = type;
              row.severity = severity;
            }

            setVulnerabilities(data);
            const totalItems = response?.totalItems;

            let totalPages = totalItems
              ? Math.ceil(totalItems / page.pageSize)
              : 1;

            if (totalPages <= 0) {
              totalPages = 1;
            }

            setPagination({ ...page, totalPages, totalItems });
          }
        })
        .catch((err: AxiosError) => {
          if (!isCancelled) {
            dispatchMessage(err?.response?.data);
            setVulnerabilities([]);
          }
        })
        .finally(() => {
          if (!isCancelled) {
            setLoading(false);
          }
        });

      if (!isCancelled) {
        setLoading(false);
      }
    };

    fetchVulnerabilities();
    return () => {
      isCancelled = true;
    };
    // eslint-disable-next-line
  }, [refresh, pagination.pageSize]);

  function resetUpdateVuln() {
    setUpdateVulnIds([]);
    setLoading(true);
    vulnerabilitiesService
      .getAllVulnerabilities(refresh.filter, overviewType, refresh.page)
      .then((result: AxiosResponse) => {
        const response = result.data?.content;
        const data: Vulnerability[] = response?.data;
        setVulnerabilities(data);
      })
      .catch((err: AxiosError) => {
        dispatchMessage(err?.response?.data);
        setVulnerabilities([]);
      })
      .finally(() => {
        setLoading(false);
      });
  }

  function updateVulnerability(
    row: Vulnerability,
    severity: string,
    type: string
  ) {
    setUpdateVulnIds((state) => {
      const data = cloneDeep(state);
      const index = data.findIndex(
        (x) => x.vulnerabilityID === row.vulnerabilityID
      );
      const record: KeyValueVuln = {
        vulnerabilityID: row.vulnerabilityID,
        severity: severity,
        type: type,
      };

      if (index >= 0) {
        data[index] = record;
      } else {
        data.push(record);
      }

      setVulnerabilities((state) =>
        state.map((row) => {
          const { severity = row.severity, type = row.type } =
            data.find((x) => x.vulnerabilityID === row.vulnerabilityID) || {};
          row.severity = severity;
          row.type = type;
          return row;
        })
      );

      return data;
    });
  }

  function isTouched(row: Vulnerability): boolean {
    const index = updateVulnIds.findIndex(
      (x) => x.vulnerabilityID === row.vulnerabilityID
    );

    return index > -1;
  }

  const isDisabled = !isAuthorizedAction(
    repositoryId
      ? ACTIONS.HANDLE_VULNERABILITIES_REPOSITORY
      : ACTIONS.HANDLE_VULNERABILITIES_WORKSPACE
  );

  return (
    <Styled.Wrapper>
      <Styled.Options>
        <SearchBar
          placeholder={t('VULNERABILITIES_SCREEN.SEARCH')}
          onSearch={(value) => handleSearch(value)}
        />

        <Select
          width="250px"
          placeholder={t('VULNERABILITIES_SCREEN.ALL_SEVERITIES')}
          disabled={!!filters.vulnHash}
          options={severities}
          value={filters.vulnSeverity}
          label={t('VULNERABILITIES_SCREEN.SEVERITY')}
          onChangeValue={(item) => {
            setFilters((state) => ({ ...state, vulnSeverity: item }));
            setRefresh({
              filter: { ...filters, vulnSeverity: item },
              page: { ...pagination, currentPage: INITIAL_PAGE },
            });
          }}
        />

        <Styled.Select
          width="250px"
          placeholder={t('VULNERABILITIES_SCREEN.ALL_STATUS')}
          disabled={!!filters.vulnHash}
          options={vulnTypes}
          label={t('VULNERABILITIES_SCREEN.STATUS_TITLE')}
          value={filters.vulnType}
          onChangeValue={(item) => {
            setFilters((state) => ({ ...state, vulnType: item }));
            setRefresh({
              filter: { ...filters, vulnType: item },
              page: { ...pagination, currentPage: INITIAL_PAGE },
            });
          }}
        />
      </Styled.Options>

      <Styled.Content>
        <Datatable
          buttons={[
            {
              title: `${t('VULNERABILITIES_SCREEN.UPDATE_VULNERABILITY')} (${
                updateVulnIds.length
              })`,
              function: handleUpdateVulnerability,
              icon: 'success',
              disabled: !!updateVulnIds.length,
              show: updateVulnIds.length > 0,
            },
            {
              title: t('GENERAL.CANCEL'),
              function: resetUpdateVuln,
              icon: 'error',
              disabled: !!updateVulnIds.length,
              show: updateVulnIds.length > 0,
            },
          ]}
          columns={[
            {
              label: t('VULNERABILITIES_SCREEN.TABLE.SEVERITY'),
              property: 'severity',
              type: 'custom',
              cssClass: ['center'],
            },
            {
              label: t('VULNERABILITIES_SCREEN.TABLE.STATUS'),
              property: 'status',
              type: 'custom',
            },
            {
              label: t('VULNERABILITIES_SCREEN.TABLE.HASH'),
              property: 'hash',
              type: 'text',
            },
            {
              label: t('VULNERABILITIES_SCREEN.TABLE.DETAILS'),
              property: 'details',
              type: 'custom',
            },
          ]}
          dataSource={vulnerabilities.map((row) => {
            const repo: DataSource = {
              ...row,
              id: row.vulnerabilityID,
              hash: row.vulnHash,
              description: `${row.file} ( ${row.line || ' - '} : ${
                row.column || ' - '
              })`,
              highlight: isTouched(row),
              severity: (
                <Select
                  style={{
                    backgroundColor: get(
                      colors.vulnerabilities,
                      row.severity,
                      colors.vulnerabilities.DEFAULT
                    ),
                    color: isDisabled ? '#F4F4F4' : '',
                  }}
                  variant="filled"
                  width="150px"
                  value={row.severity}
                  options={severities.slice(1)}
                  disabled={isDisabled}
                  onChangeValue={(value) => {
                    updateVulnerability(row, value, row.type);
                  }}
                />
              ),
              status: (
                <Select
                  value={row.type}
                  options={vulnTypes.slice(1)}
                  width="200px"
                  variant="filled"
                  style={{ color: isDisabled ? '#F4F4F4' : '' }}
                  disabled={isDisabled}
                  onChangeValue={(value) => {
                    updateVulnerability(row, row.severity, value);
                  }}
                />
              ),
              details: (
                <Icon
                  name="info"
                  size="20px"
                  onClick={() => setSelectedVuln(row)}
                />
              ),
            };
            return repo;
          })}
          isLoading={isLoading}
          emptyListText={t('VULNERABILITIES_SCREEN.TABLE.EMPTY')}
          fixed={false}
          paginate={{
            pagination,
            onChange: (page) => setRefresh({ filter: filters, page }),
          }}
        />
      </Styled.Content>

      <Details
        isOpen={!!selectedVuln}
        onClose={() => setSelectedVuln(null)}
        vulnerability={selectedVuln}
      />
    </Styled.Wrapper>
  );
};

export default Vulnerabilities;

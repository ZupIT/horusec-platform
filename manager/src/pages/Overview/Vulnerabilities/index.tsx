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

import React, { useState, useEffect } from 'react';
import Styled from './styled';
import { Button, Icon, Pagination, Select } from 'components';
import { FilterVuln } from 'helpers/interfaces/FIlterVuln';
import useParamsRoute from 'helpers/hooks/useParamsRoute';
import { useTranslation } from 'react-i18next';
import { findIndex, get, cloneDeep } from 'lodash';
import vulnerabilitiesService from 'services/vulnerabilities';
import { PaginationInfo } from 'helpers/interfaces/Pagination';
import { VulnerableFile } from 'helpers/interfaces/VulnerableFile';
import { formatToDistanceDate } from 'helpers/formatters/date';
import { Vulnerability } from 'helpers/interfaces/Vulnerability';
import Linkify from 'react-linkify';
import usePermissions from 'helpers/hooks/usePermissions';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import { AxiosError } from 'axios';
import useLanguage from 'helpers/hooks/useLanguage';

const INITIAL_PAGE = 1;

const Vulnerabilities: React.FC = () => {
  const { workspaceId, repositoryId } = useParamsRoute();
  const { ACTIONS, isAuthorizedAction } = usePermissions();
  const { dispatchMessage } = useResponseMessage();
  const { showSuccessFlash } = useFlashMessage();
  const { t } = useTranslation();
  const { currentLocale } = useLanguage();

  const [isLoading, setLoading] = useState(false);
  const [isLoadingUpdate, setLoadingUpdated] = useState(false);
  const [isSearching, setSearching] = useState(false);
  const [vulOpened, setVulOpened] = useState<string>();

  const [pagination, setPagination] = useState<PaginationInfo>({
    currentPage: INITIAL_PAGE,
    totalItems: 0,
    pageSize: 10,
    totalPages: 10,
  });

  const [vulnerableFiles, setVulnerableFiles] = useState<VulnerableFile[]>([]);
  const [selectedFile, setSelectedFile] = useState<VulnerableFile>(null);

  const [vulnerabilitiesOfFile, setVulnerabilitiesOfFile] = useState<
    Vulnerability[]
  >([]);

  const [vulnerabilitiesToUpdate, setVulnerabilitiesToUpdate] = useState<
    Vulnerability[]
  >([]);

  const overviewType = repositoryId ? 'repository' : 'workspace';

  const [filters, setFilters] = useState<FilterVuln>({
    workspaceID: workspaceId,
    repositoryID: repositoryId,
    vulnHash: '',
    vulnSeverity: 'ALL',
    vulnType: 'Vulnerability',
  });

  const vulTypes = [
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

  const extensionColors = {
    GO: '#66d1dd',
    SUM: '#66d1dd',
    MOD: '#66d1dd',
    CS: '#8823ec',
    C: '#28348e',
    RB: '#970f03',
    PY: '#366b97',
    JAVA: '#f0931e',
    KT: '#746dda',
    KTS: '#746dda',
    KTM: '#746dda',
    JS: '#f0d81e',
    TS: '#2f74c0',
    CPP: '#005697',
    PHP: '#7377ad',
    HTML: '#e96229',
    DART: '#ccd7dd',
    SH: '#7f8c8d',
    JSON: '#cc9d1d',
    XML: '#e98629',
    YAML: '#72a250',
    YML: '#72a250',
    TF: '#7f8c8d',
    SWIFT: '#f57836',
    NGINX: '#008e36',
    TXT: '#CDCDCD',
    PERF: '#CDCDCD',
    EXS: '#351350',
    EX: '#351350',
    ELF: '#7f8c8d',
    LOCK: '#86523c',
  };

  const isDisabled = !isAuthorizedAction(
    repositoryId
      ? ACTIONS.HANDLE_VULNERABILITIES_REPOSITORY
      : ACTIONS.HANDLE_VULNERABILITIES_WORKSPACE
  );

  const getFileExtension = (path: string) => {
    if (path) {
      if (path?.split('.').length > 1) {
        const extension = path?.split('.')?.slice(-1)?.pop()?.toUpperCase();
        const color = get(extensionColors, extension, '#FFF');
        return { extension, color };
      }
    }

    return { extension: '?', color: '#FFF' };
  };

  const handleSearchValue = (value: string) => {
    setSearching(!!value);
    setFilters((state) => ({ ...state, vulnHash: value }));
    getFilesWithVul();
  };

  const handlePagination = (totalItems: number) => {
    let totalPages = totalItems
      ? Math.ceil(totalItems / pagination.pageSize)
      : 1;

    if (totalPages <= 0) {
      totalPages = 1;
    }

    const paginationToSet = {
      ...pagination,
      currentPage: INITIAL_PAGE,
      totalPages,
      totalItems,
    };

    if (paginationToSet.totalItems !== pagination.totalItems) {
      setPagination(paginationToSet);
    }
  };

  const setVulnerabilityToUpdate = (vul: Vulnerability) => {
    setVulnerabilitiesOfFile((state) =>
      state.map((el) => (el.vulnerabilityID === vul.vulnerabilityID ? vul : el))
    );

    const index = findIndex(vulnerabilitiesToUpdate, {
      vulnerabilityID: vul.vulnerabilityID,
    });

    let vulsToUpdateClone = cloneDeep(vulnerabilitiesToUpdate);

    if (index > -1) {
      vulsToUpdateClone[index] = vul;
    } else {
      vulsToUpdateClone = vulsToUpdateClone.concat(vul);
    }

    setVulnerabilitiesToUpdate(vulsToUpdateClone);
  };

  const applyCurrentUpdatesVulnerabilities = (
    vulnerabilities: Vulnerability[]
  ) => {
    vulnerabilitiesToUpdate.map((item) => {
      const index = findIndex(vulnerabilities, {
        vulnerabilityID: item.vulnerabilityID,
      });

      if (index > -1) vulnerabilities[index] = item;
    });

    return vulnerabilities;
  };

  const getFilesWithVul = () => {
    setLoading(true);
    setSelectedFile(null);

    vulnerabilitiesService
      .getFilesWithVulnerabilities(filters, overviewType, pagination)
      .then((result) => {
        const files: VulnerableFile[] = result?.data?.content?.data;
        const totalItems: number = result?.data?.content?.totalItems;

        setVulnerableFiles(files);
        handlePagination(totalItems);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
        setVulnerableFiles([]);
        handlePagination(0);
      })
      .finally(() => setLoading(false));
  };

  const getVulOfFile = (vulFile: VulnerableFile) => {
    setLoading(true);
    setSelectedFile(vulFile);

    vulnerabilitiesService
      .getVulnerabilitiesOfFile(filters, pagination, vulFile.file)
      .then((result) => {
        let vulnerabilities: Vulnerability[] = result?.data?.content?.data;
        vulnerabilities = applyCurrentUpdatesVulnerabilities(vulnerabilities);

        const totalItems: number = result?.data?.content?.totalItems;

        setVulnerabilitiesOfFile(vulnerabilities);
        handlePagination(totalItems);
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
        setVulnerabilitiesOfFile([]);
        handlePagination(0);
      })
      .finally(() => setLoading(false));
  };

  const getVulCodeDetail = (vuln: Vulnerability) => {
    let message = `Line: ${vuln.line || '-'}, Column: ${vuln.column || '-'}`;

    if (vuln?.commitAuthor && vuln?.commitAuthor !== '-') {
      message = `${message} - Committed by: ${vuln.commitAuthor} <${vuln.commitEmail}>`;
    }

    return message;
  };

  const handleUpdateVulnerability = () => {
    setLoadingUpdated(true);

    vulnerabilitiesService
      .updateVulnerability(
        filters.workspaceID,
        filters.repositoryID,
        vulnerableFiles[0].analysisID,
        vulnerabilitiesToUpdate,
        overviewType
      )
      .then(() => {
        showSuccessFlash(t('VULNERABILITIES_SCREEN.SUCCESS_UPDATE'));
        setVulnerabilitiesToUpdate([]);
      })
      .catch((err: AxiosError) => {
        dispatchMessage(err?.response?.data);
      })
      .finally(() => setLoadingUpdated(false));
  };

  const renderFilter = () => (
    <Styled.Options>
      <Styled.SelectsWrapper isSearching={isSearching}>
        <Select
          width="300px"
          className="filter"
          wrapperStyle={{ marginRight: '50px', marginBottom: '10px' }}
          placeholder={t('VULNERABILITIES_SCREEN.ALL_SEVERITIES')}
          disabled={!!filters.vulnHash}
          options={severities}
          value={filters.vulnSeverity}
          label={t('VULNERABILITIES_SCREEN.SEVERITY')}
          onChangeValue={(item) =>
            setFilters((state) => ({ ...state, vulnSeverity: item }))
          }
        />

        <Select
          width="300px"
          className="filter"
          wrapperStyle={{ marginBottom: '10px' }}
          placeholder={t('VULNERABILITIES_SCREEN.ALL_STATUS')}
          disabled={!!filters.vulnHash}
          options={vulTypes}
          label={t('VULNERABILITIES_SCREEN.STATUS_TITLE')}
          value={filters.vulnType}
          onChangeValue={(item) =>
            setFilters((state) => ({ ...state, vulnType: item }))
          }
        />
      </Styled.SelectsWrapper>

      <Styled.SearchWrapper isSearching={isSearching}>
        <Styled.SearchBar
          placeholder={t('VULNERABILITIES_SCREEN.SEARCH')}
          onSearch={handleSearchValue}
        />
      </Styled.SearchWrapper>
    </Styled.Options>
  );

  const renderMenuUpdating = () => {
    if (vulnerabilitiesToUpdate.length > 0) {
      return (
        <Styled.UpdateContent>
          <Styled.UpdateCount>
            <b>{vulnerabilitiesToUpdate.length}</b>
            {t('VULNERABILITIES_SCREEN.UPDATE_VULNERABILITY')}
          </Styled.UpdateCount>

          <Styled.UpdateBtns>
            <Button
              text={t('GENERAL.SAVE')}
              isLoading={isLoadingUpdate}
              width={100}
              outlinePrimary
              onClick={handleUpdateVulnerability}
              className="save-vulnerabilities"
            />

            <Button
              text={t('GENERAL.CANCEL')}
              width={100}
              isDisabled={isLoadingUpdate}
              outline
              style={{ marginLeft: '15px' }}
              onClick={() => {
                setVulnerabilitiesToUpdate([]);
                getFilesWithVul();
              }}
            />
          </Styled.UpdateBtns>
        </Styled.UpdateContent>
      );
    }
  };

  const renderFileItem = (vulnerableFile: VulnerableFile, index: number) => (
    <Styled.File key={index} onClick={() => getVulOfFile(vulnerableFile)}>
      <Styled.FileRow>
        <Styled.FileColumn>
          <Styled.FileRow>
            <Styled.FileLanguage
              color={get(
                extensionColors,
                getFileExtension(vulnerableFile.file).extension,
                '#FFF'
              )}
            >
              {vulnerableFile?.languages?.join('/')}
            </Styled.FileLanguage>
            <Styled.FileName>{vulnerableFile.file}</Styled.FileName>
          </Styled.FileRow>
        </Styled.FileColumn>

        <Styled.FileRow>
          <Styled.Date>
            {formatToDistanceDate(vulnerableFile.createdAt, currentLocale)}
          </Styled.Date>
          <Styled.View className="view">⮕</Styled.View>
        </Styled.FileRow>
      </Styled.FileRow>

      <Styled.FileInfo>
        <Styled.FileRow>
          <Styled.FileInfoText>
            <Styled.FileVulCount>
              {vulnerableFile.totalVulnerabilities}
            </Styled.FileVulCount>
            {t('VULNERABILITIES_SCREEN.TITLE').toLowerCase()}.
          </Styled.FileInfoText>
        </Styled.FileRow>

        {overviewType === 'workspace' && (
          <Styled.FileRow>
            <Styled.FileInfoText>
              {vulnerableFile.repositoryName}
            </Styled.FileInfoText>
          </Styled.FileRow>
        )}
      </Styled.FileInfo>
    </Styled.File>
  );

  const renderTableMessage = (message: string, icon?: string) => {
    return (
      <Styled.LoadingWrapper>
        {icon && <Icon name={icon} size="50px" />}
        <Styled.LoadingText>{message}</Styled.LoadingText>
      </Styled.LoadingWrapper>
    );
  };

  const renderFilesList = () => (
    <Styled.Content className="file-list">
      <Styled.ScrollList>
        {vulnerableFiles.length > 0
          ? vulnerableFiles.map((file, index) => renderFileItem(file, index))
          : renderTableMessage(t('VULNERABILITIES_SCREEN.TABLE.EMPTY'))}
      </Styled.ScrollList>

      <Pagination
        pagination={pagination}
        onChange={(value) => setPagination(value)}
      />
    </Styled.Content>
  );

  const renderVulnerability = (vulnerability: Vulnerability) => (
    <Styled.Vulnerability key={vulnerability.vulnerabilityID}>
      <Styled.VulDetailWrapper>
        <Styled.VulDetail isOpen={vulOpened === vulnerability.vulnerabilityID}>
          <Linkify>{vulnerability.details}</Linkify>

          <Styled.Info>
            {t('VULNERABILITIES_SCREEN.DETAILS.SECURITY_TOOL')}:{' '}
            {vulnerability.securityTool} -{' '}
            {t('VULNERABILITIES_SCREEN.DETAILS.CONFIDENCE')}:{' '}
            {vulnerability.confidence}
          </Styled.Info>

          <Styled.Code>{vulnerability.code}</Styled.Code>
          <Styled.CodeInfoWrapper>
            <Styled.CodeInfo>{getVulCodeDetail(vulnerability)}</Styled.CodeInfo>
            <Styled.CodeInfo>Hash: {vulnerability.vulnHash}</Styled.CodeInfo>
          </Styled.CodeInfoWrapper>
        </Styled.VulDetail>

        <Styled.Ellipsis
          onClick={() =>
            vulOpened === vulnerability.vulnerabilityID
              ? setVulOpened(null)
              : setVulOpened(vulnerability.vulnerabilityID)
          }
        >
          ...
        </Styled.Ellipsis>
      </Styled.VulDetailWrapper>

      <Styled.SelectOptionsWrapper>
        <Icon name={vulnerability.severity} size="18px" />

        <Select
          variant="filled"
          width="max-content"
          className="severity-dropdown"
          value={vulnerability.severity}
          options={severities.slice(1)}
          disabled={isDisabled}
          style={{
            color: isDisabled ? '#F4F4F4' : '',
            marginRight: '20px',
            fontSize: '12px',
          }}
          onChangeValue={(value) => {
            setVulnerabilityToUpdate({
              ...vulnerability,
              severity: value,
            });
          }}
        />

        <Select
          value={vulnerability.type}
          options={vulTypes}
          width="max-content"
          className="status-dropdown"
          variant="filled"
          style={{ color: isDisabled ? '#F4F4F4' : '', fontSize: '12px' }}
          disabled={isDisabled}
          onChangeValue={(value) => {
            setVulnerabilityToUpdate({
              ...vulnerability,
              type: value,
            });
          }}
        />
      </Styled.SelectOptionsWrapper>
    </Styled.Vulnerability>
  );

  const renderVulnerabilitiesOfFile = () => (
    <Styled.Content className="vulnerabilities-box">
      <Styled.HeaderVulList>
        <Styled.FileTitle>
          {selectedFile.file}
          {overviewType === 'workspace' && (
            <Styled.Info>{selectedFile.repositoryName}</Styled.Info>
          )}
        </Styled.FileTitle>
        <Styled.Back onClick={() => getFilesWithVul()}>
          <Icon name="back" size="20px" />
          <Styled.BackText>{t('GENERAL.BACK')}</Styled.BackText>
        </Styled.Back>
      </Styled.HeaderVulList>

      <Styled.ScrollList>
        {vulnerabilitiesOfFile.length > 0
          ? vulnerabilitiesOfFile.map((vuln) => renderVulnerability(vuln))
          : renderTableMessage(t('VULNERABILITIES_SCREEN.TABLE.EMPTY'))}
      </Styled.ScrollList>

      <Pagination
        pagination={pagination}
        onChange={(value) => setPagination(value)}
      />
    </Styled.Content>
  );

  const renderContent = () => {
    if (isLoading) {
      return (
        <Styled.Content>
          {renderTableMessage(t('VULNERABILITIES_SCREEN.LOADING'), 'loading')}
        </Styled.Content>
      );
    } else if (selectedFile) {
      return renderVulnerabilitiesOfFile();
    } else {
      return renderFilesList();
    }
  };

  useEffect(() => {
    if (selectedFile) getVulOfFile(selectedFile);
    else getFilesWithVul();
  }, [filters, pagination]);

  return (
    <Styled.Wrapper>
      {renderFilter()}

      {renderMenuUpdating()}

      {renderContent()}
    </Styled.Wrapper>
  );
};

export default Vulnerabilities;

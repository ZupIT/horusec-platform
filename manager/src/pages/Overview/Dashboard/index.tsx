/* eslint-disable no-sparse-arrays */
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
import Filters from './Filters';
import { FilterValues } from 'helpers/interfaces/FilterValues';
import { DashboardData } from 'helpers/interfaces/DashboardData';
import analyticService from 'services/analytic';
import { AxiosResponse } from 'axios';

import TotalDevelopers from './TotalDevelopers';
import TotalRepositories from './TotalRepositories';
import AllVulnerabilities from './AllVulnerabilities';
import VulnerabilitiesByDeveloper from './VulnerabilitiesByDeveloper';
import VulnerabilitiesByLanguage from './VulnerabilitiesByLanguage';
import VulnerabilitiesByRepository from './VulnerabilitiesByRepository';
import VulnerabilitiesTimeLine from './VulnerabilitiesTimeLine';
import VulnerabilitiesDetails from './VulnerabilitiesDetails';
import useFlashMessage from 'helpers/hooks/useFlashMessage';
import { createReportWorkSheet } from 'helpers/formatters/xlsx';

import { Button } from 'components';
import { Menu, MenuItem } from '@material-ui/core';
import exportFromJSON, { ExportType } from 'export-from-json';
import { jsPDF } from 'jspdf';
import * as htmlToImage from 'html-to-image';
import { useTranslation } from 'react-i18next';
import XLSX from 'xlsx';
import download from 'downloadjs';
import useParamsRoute from 'helpers/hooks/useParamsRoute';
import { useParams } from 'react-router-dom';

interface Props {
  type: 'workspace' | 'repository';
}

const Dashboard: React.FC<Props> = ({ type }) => {
  const { workspaceId, repositoryId } =
    useParams<{ workspaceId: string; repositoryId: string }>();

  const [filters, setFilters] = useState<FilterValues>(null);

  const [dashboardData, setDashboardData] = useState<DashboardData>();
  const [isLoading, setLoading] = useState(false);
  const { showSuccessFlash } = useFlashMessage();
  const { t } = useTranslation();
  const [anchorElExport, setAnchorElExport] =
    React.useState<null | HTMLElement>(null);

  const handleExportOpen = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorElExport(event.currentTarget);
  };

  const handleExportClose = () => {
    setAnchorElExport(null);
  };

  useEffect(() => {
    let isCancelled = false;
    if (filters) {
      setLoading(true);
      setDashboardData(null);

      analyticService
        .getDashboardData(filters)
        .then((result: AxiosResponse) => {
          if (!isCancelled) {
            const data = result?.data?.content as DashboardData;
            setDashboardData(data);
          }
        })
        .finally(() => {
          if (!isCancelled) {
            setLoading(false);
          }
        });
    }

    return () => {
      isCancelled = true;
    };
  }, [filters]);

  function downloadExport(exportType: ExportType) {
    showSuccessFlash(t('GENERAL.LOADING'), 1000);
    const fileName = 'horusec_dashboard_' + new Date().toLocaleString();

    if (exportType === 'xls' || exportType === 'csv') {
      const workbook = XLSX.utils.book_new();
      const workSheetData = createReportWorkSheet(dashboardData);
      const workSheet = XLSX.utils.aoa_to_sheet(workSheetData);
      workbook.SheetNames.push('Report');
      workbook.Sheets['Report'] = workSheet;
      XLSX.writeFile(workbook, fileName + '.' + exportType);
    } else {
      exportFromJSON({ data: dashboardData, fileName, exportType });
    }
  }

  function downloadExportPdf(exportType: 'pdf' | 'image') {
    showSuccessFlash(t('GENERAL.LOADING'), 1000);
    const printHtml = window.document.getElementById('wrapper-graphic');
    const fileName = 'horusec_dashboard_' + new Date().toLocaleString();
    const imgHeight = 1505;
    const imgWidth = 1755;
    htmlToImage
      .toJpeg(printHtml, {
        height: imgHeight,
        width: imgWidth,
      })
      .then(function (dataUrl) {
        if (exportType === 'image') {
          download(dataUrl, fileName);
        }

        if (exportType === 'pdf') {
          const doc = new jsPDF({
            orientation: 'landscape',
            format: [imgWidth, imgHeight],
            unit: 'px',
          });

          doc.addImage(dataUrl, 'JPEG', 25, 25, imgWidth - 50, imgHeight - 50);
          doc.save(fileName, { returnPromise: true });
        }
      });
  }

  return (
    <Styled.Wrapper id="wrapper-graphic">
      <Styled.FilterWrapper>
        <Filters type={type} onApply={(values) => setFilters(values)} />
        <Button
          text={t('GENERAL.EXPORT')}
          style={{ margin: 20 }}
          onClick={handleExportOpen}
        />
        <Menu
          id="export-menu"
          anchorEl={anchorElExport}
          open={Boolean(anchorElExport)}
          onClose={handleExportClose}
        >
          <MenuItem
            onClick={() => {
              downloadExportPdf('image');
              handleExportClose();
            }}
          >
            {t('GENERAL.DOWNLOAD_AS') + ' JPG'}
          </MenuItem>
          <MenuItem
            onClick={() => {
              downloadExportPdf('pdf');
              handleExportClose();
            }}
          >
            {t('GENERAL.DOWNLOAD_AS') + ' PDF'}
          </MenuItem>
          <MenuItem
            onClick={() => {
              downloadExport('json');
              handleExportClose();
            }}
          >
            {t('GENERAL.DOWNLOAD_AS') + ' JSON'}
          </MenuItem>
          <MenuItem
            onClick={() => {
              downloadExport('xls');
              handleExportClose();
            }}
          >
            {t('GENERAL.DOWNLOAD_AS') + ' XLS'}
          </MenuItem>
          <MenuItem
            onClick={() => {
              downloadExport('csv');
              handleExportClose();
            }}
          >
            {t('GENERAL.DOWNLOAD_AS') + ' CSV'}
          </MenuItem>
          <MenuItem
            onClick={() => {
              downloadExport('xml');
              handleExportClose();
            }}
          >
            {t('GENERAL.DOWNLOAD_AS') + ' XML'}
          </MenuItem>
        </Menu>
      </Styled.FilterWrapper>

      <AllVulnerabilities
        data={dashboardData?.vulnerabilityBySeverity}
        isLoading={isLoading}
      />

      <TotalDevelopers
        isLoading={isLoading}
        data={dashboardData?.totalAuthors}
      />

      <VulnerabilitiesByDeveloper
        isLoading={isLoading}
        data={dashboardData?.vulnerabilitiesByAuthor}
      />

      {type === 'workspace' ? (
        <>
          <TotalRepositories
            data={dashboardData?.totalRepositories}
            isLoading={isLoading}
          />

          <VulnerabilitiesByRepository
            isLoading={isLoading}
            data={dashboardData?.vulnerabilitiesByRepository}
          />
        </>
      ) : null}

      <VulnerabilitiesByLanguage
        isLoading={isLoading}
        data={dashboardData?.vulnerabilitiesByLanguage}
      />

      <VulnerabilitiesTimeLine
        isLoading={isLoading}
        data={dashboardData?.vulnerabilityByTime}
      />

      <VulnerabilitiesDetails filters={filters} />
    </Styled.Wrapper>
  );
};

export default Dashboard;

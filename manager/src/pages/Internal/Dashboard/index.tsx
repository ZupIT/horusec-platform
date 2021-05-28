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
import { Button } from 'components';
import { Menu, MenuItem } from '@material-ui/core';
import exportFromJSON, { ExportType } from 'export-from-json';
import { jsPDF } from 'jspdf';
import * as htmlToImage from 'html-to-image';
const download = require('downloadjs');
interface Props {
  type: 'workspace' | 'repository';
}

const Dashboard: React.FC<Props> = ({ type }) => {
  const [filters, setFilters] = useState<FilterValues>(null);
  const [dashboardData, setDashboardData] = useState<DashboardData>();
  const [isLoading, setLoading] = useState(false);

  const [
    anchorElExport,
    setAnchorElExport,
  ] = React.useState<null | HTMLElement>(null);

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
    let data;
    data = dashboardData;
    if (exportType === 'csv' || exportType === 'xls') {
      for (const x of data.vulnerabilitiesByAuthor) {
        data.vulnerabilitiesByAuthor = [x];
      }
      data = [dashboardData];
    }
    console.log([dashboardData]);
    const fileName = 'horusec_dashboard_' + new Date().toLocaleString();
    exportFromJSON({ data, fileName, exportType });
  }

  function downloadExportPdf(exportType: 'pdf' | 'image') {
    const printHtml = window.document.getElementById('print');
    const fileName = 'horusec_dashboard_' + new Date().toLocaleString();
    htmlToImage.toJpeg(printHtml).then(function (dataUrl) {
      if (exportType === 'image') {
        download(dataUrl, fileName);
      }

      if (exportType === 'pdf') {
        const doc = new jsPDF({
          orientation: 'landscape',
          format: [1950, 1200],
          unit: 'px',
        });

        doc.addImage(dataUrl, 'JPEG', 25, 25, 1900, 1150);
        doc.save(fileName);
      }
    });
  }

  return (
    <Styled.Wrapper>
      <Styled.FilterWrapper>
        <Filters type={type} onApply={(values) => setFilters(values)} />
        <Button
          text="Export"
          style={{ margin: 20 }}
          onClick={handleExportOpen}
        ></Button>
        <Menu
          id="export-menu"
          anchorEl={anchorElExport}
          keepMounted
          open={Boolean(anchorElExport)}
          onClose={handleExportClose}
        >
          <MenuItem onClick={() => downloadExportPdf('image')}>
            Download JPEG
          </MenuItem>
          <MenuItem onClick={() => downloadExportPdf('pdf')}>
            Download PDF
          </MenuItem>
          <MenuItem onClick={() => downloadExport('json')}>
            Download JSON
          </MenuItem>
          <MenuItem onClick={() => downloadExport('csv')}>
            Download CSV
          </MenuItem>
          <MenuItem onClick={() => downloadExport('xls')}>
            Download XLS
          </MenuItem>
          <MenuItem onClick={() => downloadExport('xml')}>
            Download XML
          </MenuItem>
        </Menu>
      </Styled.FilterWrapper>
      <div id="print">
        <Styled.Row>
          <TotalDevelopers
            isLoading={isLoading}
            data={dashboardData?.totalAuthors}
          />

          {type === 'workspace' ? (
            <TotalRepositories
              data={dashboardData?.totalRepositories}
              isLoading={isLoading}
            />
          ) : null}

          <AllVulnerabilities
            data={dashboardData?.vulnerabilityBySeverity}
            isLoading={isLoading}
          />
        </Styled.Row>

        <Styled.Row>
          <VulnerabilitiesByDeveloper
            isLoading={isLoading}
            data={dashboardData?.vulnerabilitiesByAuthor}
          />

          <VulnerabilitiesByLanguage
            isLoading={isLoading}
            data={dashboardData?.vulnerabilitiesByLanguage}
          />

          {type === 'workspace' ? (
            <VulnerabilitiesByRepository
              isLoading={isLoading}
              data={dashboardData?.vulnerabilitiesByRepository}
            />
          ) : null}
        </Styled.Row>

        <Styled.Row>
          <VulnerabilitiesTimeLine
            isLoading={isLoading}
            data={dashboardData?.vulnerabilityByTime}
          />
        </Styled.Row>
      </div>
      <Styled.Row>
        <VulnerabilitiesDetails filters={filters} />
      </Styled.Row>
    </Styled.Wrapper>
  );
};

export default Dashboard;

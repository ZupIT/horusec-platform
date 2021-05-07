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
import { useTranslation } from 'react-i18next';
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

import NewVulnerabilitiesByDeveloper from './NewVulnerabilitiesByDeveloper';

interface Props {
  type: 'workspace' | 'repository';
}

const Dashboard: React.FC<Props> = ({ type }) => {
  const [filters, setFilters] = useState<FilterValues>(null);
  const [dashboardData, setDashboardData] = useState<DashboardData>();
  const [isLoading, setLoading] = useState(false);

  const { t } = useTranslation();

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

  return (
    <Styled.Wrapper>
      <Styled.AriaTitle>
        {type === 'workspace'
          ? t('DASHBOARD_SCREEN.ARIA_TITLE_WORKSPACE')
          : t('DASHBOARD_SCREEN.ARIA_TITLE_REPOSITORY')}
      </Styled.AriaTitle>

      <Filters type={type} onApply={(values) => setFilters(values)} />

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

        {type === 'workspace' ? (
          <VulnerabilitiesByRepository
            isLoading={isLoading}
            data={dashboardData?.vulnerabilitiesByRepository}
          />
        ) : null}
      </Styled.Row>

      <Styled.Row>
        <NewVulnerabilitiesByDeveloper
          isLoading={isLoading}
          data={dashboardData?.vulnerabilitiesByAuthor}
        />

        {type === 'workspace' ? (
          <VulnerabilitiesByRepository
            isLoading={isLoading}
            data={dashboardData?.vulnerabilitiesByRepository}
          />
        ) : null}
      </Styled.Row>

      <Styled.Row>
        <VulnerabilitiesByLanguage
          isLoading={isLoading}
          data={dashboardData?.vulnerabilitiesByLanguage}
        />

        <VulnerabilitiesTimeLine
          isLoading={isLoading}
          data={dashboardData?.vulnerabilityByTime}
        />
      </Styled.Row>

      <Styled.Row>
        <VulnerabilitiesDetails filters={filters} />
      </Styled.Row>
    </Styled.Wrapper>
  );
};

export default Dashboard;

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
import { SERVICE_ANALYTIC } from '../config/endpoints';
import { FilterValues } from 'helpers/interfaces/FilterValues';
import { formatInitialAndFinalDate } from 'helpers/formatters/date';

import analyticMock from './analytic-mock.json';

const getDashboardData = (filters: FilterValues) => {
  const path = filters.repositoryID
    ? `/${filters.workspaceID}/${filters.repositoryID}`
    : `/${filters.workspaceID}`;

  return http.get(`${SERVICE_ANALYTIC}/analytic/dashboard/${path}`, {
    params: formatInitialAndFinalDate(filters.initialDate, filters.finalDate),
  });
};

const getDashboardDataMock = (filters: FilterValues) => {
  return analyticMock;
};

export default {
  getDashboardData,
};

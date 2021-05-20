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

import { ChartBarStacked } from 'helpers/interfaces/ChartData';
import { orderBy } from 'lodash';

const formatChartStacked = (
  listOfData: any[],
  labelKey: string,
  labeIsDate?: boolean
) => {
  const formattedData: ChartBarStacked = {
    series: [],
    categories: [],
  };
  const critical: number[] = [];
  const high: number[] = [];
  const medium: number[] = [];
  const low: number[] = [];
  const info: number[] = [];
  const unknown: number[] = [];

  if (!listOfData || typeof listOfData === 'undefined') {
    return formattedData;
  }

  if (labeIsDate) {
    listOfData = orderBy(listOfData, (item) => item[labelKey]);
  }

  listOfData.forEach((item) => {
    formattedData.categories.push(item[labelKey]);
    critical.push(item?.critical?.count);
    high.push(item?.high?.count);
    medium.push(item?.medium?.count);
    low.push(item?.low?.count);
    info.push(item?.info?.count);
    unknown.push(item?.unknown?.count);
  });

  formattedData.series = [
    { name: 'CRITICAL', data: critical },
    { name: 'HIGH', data: high },
    { name: 'MEDIUM', data: medium },
    { name: 'LOW', data: low },
    { name: 'INFO', data: info },
    { name: 'UNKNOWN', data: unknown },
  ];

  return formattedData;
};

export { formatChartStacked };

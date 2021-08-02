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

import {
  DashboardCriticality,
  DashboardData,
} from 'helpers/interfaces/DashboardData';
import lodash from 'lodash';

const setSeverityRow = (
  data: DashboardCriticality,
  instanceWorkSheet: (string | number)[][]
) => {
  instanceWorkSheet.push(['Severity', 'Type', 'Total']);
  Object.entries(data).forEach(([key, value]) =>
    Object.entries(value.types).forEach(([keyType, valueType]) => {
      instanceWorkSheet.push([
        lodash.upperCase(key),
        lodash.upperFirst(keyType),
        valueType as string,
      ]);
    })
  );
  instanceWorkSheet.push(['']);
  instanceWorkSheet.push([
    'Total',
    '',
    lodash.sumBy(Object.values(data), (vuln) => vuln.count),
  ]);
  instanceWorkSheet.push(['']);
  instanceWorkSheet.push(['']);
};

export function createReportWorkSheet(data: DashboardData) {
  const workSheetData = [
    ['Total Authors'],
    ['Total', data.totalAuthors],
    [''],
    [''],
    ['Total Repositories'],
    ['Total', data.totalRepositories],
    [''],
    [''],
  ];

  if (data.vulnerabilityBySeverity) {
    workSheetData.push(['Vulnerabilities by Severity']);
    setSeverityRow(data.vulnerabilityBySeverity, workSheetData);
  }

  if (data.vulnerabilitiesByRepository) {
    workSheetData.push(['Vulnerabilities by Repository']);
    data.vulnerabilitiesByRepository.forEach((vuln) => {
      workSheetData.push(['Repository', vuln.repositoryName]);
      delete vuln.repositoryName;
      setSeverityRow(vuln, workSheetData);
    });
  }

  if (data.vulnerabilitiesByAuthor) {
    workSheetData.push(['Vulnerabilities by Author']);
    data.vulnerabilitiesByAuthor.forEach((vuln) => {
      workSheetData.push(['Author', vuln.author]);
      delete vuln.author;
      setSeverityRow(vuln, workSheetData);
    });
  }

  if (data.vulnerabilitiesByLanguage) {
    workSheetData.push(['Vulnerabilities by Language']);
    data.vulnerabilitiesByLanguage.forEach((vuln) => {
      workSheetData.push(['Language', vuln.language]);
      delete vuln.language;
      setSeverityRow(vuln, workSheetData);
    });
  }

  if (data.vulnerabilityByTime) {
    workSheetData.push(['Vulnerabilities by Time']);
    data.vulnerabilityByTime.forEach((vuln) => {
      workSheetData.push(['Time', vuln.time]);
      delete vuln.time;
      setSeverityRow(vuln, workSheetData);
    });
  }

  return workSheetData;
}

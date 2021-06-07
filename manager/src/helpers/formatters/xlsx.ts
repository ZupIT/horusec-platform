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

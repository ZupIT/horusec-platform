import {
  DashboardCriticality,
  DashboardData,
} from 'helpers/interfaces/DashboardData';
import lodash from 'lodash';

const setSeverityRow = (
  data: DashboardCriticality,
  instance: (string | number)[][]
) => {
  instance.push(['Severity', 'Type', 'Total']);
  Object.entries(data).forEach(([key, value]) =>
    Object.entries(value.types).forEach(([keyType, valueType]) => {
      instance.push([
        lodash.upperCase(key),
        lodash.upperFirst(keyType),
        valueType as string,
      ]);
    })
  );
  instance.push(['']);
  instance.push([
    'Total',
    '',
    lodash.sumBy(Object.values(data), (el) => el.count),
  ]);
  instance.push(['']);
  instance.push(['']);
};

export function createReportWS(data: DashboardData) {
  const ws_data = [
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
    ws_data.push(['Vulnerabilities by Severity']);
    setSeverityRow(data.vulnerabilityBySeverity, ws_data);
  }

  if (data.vulnerabilitiesByRepository) {
    ws_data.push(['Vulnerabilities by Repository']);
    data.vulnerabilitiesByRepository.forEach((el) => {
      ws_data.push(['Repository', el.repositoryName]);
      delete el.repositoryName;
      setSeverityRow(el, ws_data);
    });
  }

  if (data.vulnerabilitiesByAuthor) {
    ws_data.push(['Vulnerabilities by Author']);
    data.vulnerabilitiesByAuthor.forEach((el) => {
      ws_data.push(['Author', el.author]);
      delete el.author;
      setSeverityRow(el, ws_data);
    });
  }

  if (data.vulnerabilitiesByLanguage) {
    ws_data.push(['Vulnerabilities by Language']);
    data.vulnerabilitiesByLanguage.forEach((el) => {
      ws_data.push(['Language', el.language]);
      delete el.language;
      setSeverityRow(el, ws_data);
    });
  }

  if (data.vulnerabilityByTime) {
    ws_data.push(['Vulnerabilities by Time']);
    data.vulnerabilityByTime.forEach((el) => {
      ws_data.push(['Time', el.time]);
      delete el.time;
      setSeverityRow(el, ws_data);
    });
  }

  return ws_data;
}

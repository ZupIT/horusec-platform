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
import { useTranslation } from 'react-i18next';
import { BarCharRow } from 'helpers/interfaces/BarChartRow';
import { BarChart } from 'components';
import { get } from 'lodash';
import { useTheme } from 'styled-components';
import { VulnerabilitiesByLanguageData } from 'helpers/interfaces/DashboardData';

interface Props {
  data: VulnerabilitiesByLanguageData[];
  isLoading: boolean;
}

const VulnerabilitiesByDeveloper: React.FC<Props> = ({ isLoading, data }) => {
  const { t } = useTranslation();
  const { colors } = useTheme();

  const [layeredLanguage, setLayeredLanguage] = useState<string>('');
  const [isLastLayer, setLastLayer] = useState(false);

  const [allData, setAllData] = useState<VulnerabilitiesByLanguageData[]>([]);
  const [chatData, setChartData] = useState<BarCharRow[]>([]);

  const formatFirstLayer = (data: VulnerabilitiesByLanguageData[]) => {
    const formatted = (data || []).map((item) => {
      let value = 0;

      Object.values(item).forEach((i) => {
        if (i?.count) {
          value = value + i.count;
        }
      });

      const color = get(
        colors.languages,
        item.language.toUpperCase(),
        colors.languages.UNKNOWN
      );

      return {
        value,
        color,
        legend: item.language,
      };
    });

    setLayeredLanguage(null);
    setLastLayer(false);
    setChartData(formatted);
  };

  const formatSecondLayer = (rowKey: string) => {
    setLayeredLanguage(rowKey);

    const data = allData.find((item) => item.language === rowKey);

    const formatted = Object.entries(data).map((item) => {
      const value = item[1]?.count;
      const legend = item[0].toUpperCase();
      const color = get(
        colors.vulnerabilities,
        legend,
        colors.vulnerabilities.DEFAULT
      );

      return {
        value,
        legend,
        color,
      };
    });

    delete formatted[0];
    setChartData(formatted);
  };

  const formatThirdLayer = (rowKey: string) => {
    if (!isLastLayer) {
      setLastLayer(true);

      const authorData = allData.find(
        (item) => item.language === layeredLanguage
      );

      const data = get(authorData, rowKey.toLocaleLowerCase(), { types: [] })
        ?.types;

      const formatted = Object.entries(data).map((item) => {
        const legend = item[0].toUpperCase();
        const value = (item[1] as number) || 0;
        const color = get(
          colors.vulnerabilitiesStatus,
          legend,
          colors.vulnerabilitiesStatus.DEFAULT
        );

        return {
          value,
          legend,
          color,
        };
      });

      setChartData(formatted);
    }
  };

  useEffect(() => {
    formatFirstLayer(data);
    setAllData(data);
    //eslint-disable-next-line
  }, [data]);

  return (
    <div className="block max-space">
      <BarChart
        hasSmallLegend
        isLoading={isLoading}
        data={chatData}
        title={
          layeredLanguage
            ? `${t(
                'DASHBOARD_SCREEN.VULNERABILITIES_BY_LANG'
              )}: ${layeredLanguage}`
            : t('DASHBOARD_SCREEN.VULNERABILITIES_BY_LANG')
        }
        onClickRow={(row) =>
          !layeredLanguage
            ? formatSecondLayer(row.legend)
            : formatThirdLayer(row.legend)
        }
        onClickBack={() => formatFirstLayer(allData)}
        showBackOption={!!layeredLanguage}
      />
    </div>
  );
};

export default VulnerabilitiesByDeveloper;

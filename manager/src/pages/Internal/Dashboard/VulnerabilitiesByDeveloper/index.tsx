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
import { VulnerabilitiesByAuthor } from 'helpers/interfaces/DashboardData';
import { orderBy } from 'lodash';
import Styled from './styled';

interface Props {
  data: VulnerabilitiesByAuthor[];
  isLoading: boolean;
}

const VulnerabilitiesByDeveloper: React.FC<Props> = ({ isLoading, data }) => {
  const { t } = useTranslation();
  const { colors } = useTheme();

  const [layeredDeveloper, setLayeredDeveloper] = useState<string>('');
  const [isLastLayer, setLastLayer] = useState(false);

  const [allData, setAllData] = useState<VulnerabilitiesByAuthor[]>([]);
  const [chatData, setChartData] = useState<BarCharRow[]>([]);

  const formatFirstLayer = (data: VulnerabilitiesByAuthor[]) => {
    const formatted = (data || []).map((item) => {
      let value = 0;

      Object.values(item).forEach((i) => {
        if (i?.count) {
          value = value + i.count;
        }
      });

      return {
        value,
        legend: item.author,
      };
    });

    setLayeredDeveloper(null);
    setLastLayer(false);
    setChartData(orderBy(formatted, (item) => item.value, 'desc'));
  };

  const formatSecondLayer = (rowKey: string) => {
    setLayeredDeveloper(rowKey);

    const data = allData.find((item) => item.author === rowKey);

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
        (item) => item.author === layeredDeveloper
      );

      const data = get(authorData, rowKey.toLocaleLowerCase(), {
        types: [],
      })?.types;

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
  }, [data]);

  return (
    <Styled.Wrapper>
      <BarChart
        isLoading={isLoading}
        data={chatData}
        title={
          layeredDeveloper
            ? `${t(
                'DASHBOARD_SCREEN.VULNERABILITIES_BY_DEV'
              )}: ${layeredDeveloper}`
            : t('DASHBOARD_SCREEN.VULNERABILITIES_BY_DEV')
        }
        onClickRow={(row) =>
          !layeredDeveloper
            ? formatSecondLayer(row.legend)
            : formatThirdLayer(row.legend)
        }
        onClickBack={() => formatFirstLayer(allData)}
        showBackOption={!!layeredDeveloper}
      />
    </Styled.Wrapper>
  );
};

export default VulnerabilitiesByDeveloper;

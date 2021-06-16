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
import { VulnerabilityBySeverity } from 'helpers/interfaces/DashboardData';
import Styled from './styled';

interface Props {
  data: VulnerabilityBySeverity;
  isLoading: boolean;
}

const VulnerabilitiesByRepository: React.FC<Props> = ({ isLoading, data }) => {
  const { t } = useTranslation();
  const { colors } = useTheme();

  const [layeredVuln, setLayeredVuln] = useState<string>('');
  const [isLastLayer, setLastLayer] = useState(false);

  const [allData, setAllData] = useState<VulnerabilityBySeverity>();
  const [chatData, setChartData] = useState<BarCharRow[]>([]);

  const formatFirstLayer = (data: VulnerabilityBySeverity) => {
    let isEmpty = true;

    const formatted = Object.entries(data).map((item) => {
      const value = item[1].count;
      const legend = item[0].toUpperCase();
      const color = get(
        colors.vulnerabilities,
        legend,
        colors.vulnerabilities.DEFAULT
      );

      if (value > 0) isEmpty = false;

      return {
        value,
        legend,
        color,
      };
    });

    setChartData(isEmpty ? [] : formatted);
    setLastLayer(false);
    setLayeredVuln(null);
  };

  const formatSecondLayer = (rowKey: string) => {
    if (!isLastLayer) {
      setLayeredVuln(rowKey);

      setLastLayer(true);

      const data = get(allData, rowKey.toLocaleLowerCase(), {
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
    setAllData(data);
    if (data) formatFirstLayer(data);
    //eslint-disable-next-line
  }, [data]);

  return (
    <Styled.Wrapper>
      <BarChart
        isVertical
        isLoading={isLoading}
        data={chatData}
        title={
          layeredVuln
            ? `${t('DASHBOARD_SCREEN.ALL_VULNERABILITIES')}: ${layeredVuln}`
            : t('DASHBOARD_SCREEN.ALL_VULNERABILITIES')
        }
        onClickRow={(row) => formatSecondLayer(row.legend)}
        onClickBack={() => formatFirstLayer(allData)}
        showBackOption={!!layeredVuln}
      />
    </Styled.Wrapper>
  );
};

export default React.memo(VulnerabilitiesByRepository);

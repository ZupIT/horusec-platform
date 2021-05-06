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

import React, { useState, useEffect, useCallback } from 'react';
import ReactApexChart from 'react-apexcharts';
import { ApexOptions } from 'apexcharts';
import Styled from './styled';
import { useTranslation } from 'react-i18next';
import { useTheme } from 'styled-components';
import { Icon } from 'components';
import { get } from 'lodash';
import { VulnerabilitiesByLanguageData } from 'helpers/interfaces/DashboardData';

interface Props {
  data: VulnerabilitiesByLanguageData[];
  isLoading: boolean;
}

const VulnerabilitiesByLanguage: React.FC<Props> = ({ data, isLoading }) => {
  const { t } = useTranslation();
  const { colors, metrics } = useTheme();

  const [chartValues, setChartValues] = useState<number[]>([]);
  const [chartLabels, setChartLabels] = useState<string[]>([]);
  const [chartColors, setChartColors] = useState<string[]>([]);

  const options: ApexOptions = {
    markers: {
      size: 0,
    },
    noData: {
      text: t('DASHBOARD_SCREEN.CHART_NO_DATA'),
      style: {
        color: colors.chart.legend,
        fontSize: metrics.fontSize.large,
      },
    },
    chart: {
      type: 'donut',
      fontFamily: 'SFRegular',
      animations: {
        enabled: true,
      },
      toolbar: {
        show: false,
      },
    },
    legend: {
      position: 'top',
      horizontalAlign: 'left',
      formatter: (name, opts) =>
        `${name}: ${opts?.w?.config?.series[opts?.seriesIndex]}`,
      labels: {
        colors: colors.chart.legend,
      },
    },
    dataLabels: {
      enabled: true,
      style: {
        fontSize: metrics.fontSize.small,
      },
    },
    stroke: {
      show: false,
    },
    plotOptions: {
      pie: {
        donut: {
          size: '25px',
        },
      },
    },
  };

  useEffect(() => {
    const formatData = (data: VulnerabilitiesByLanguageData[]) => {
      const itemColors: string[] = [];
      const labels: string[] = [];
      const values: number[] = [];

      data.forEach((item) => {
        const total =
          item.critical.count +
          item.high.count +
          item.info.count +
          item.medium.count +
          item.low.count +
          item.unknown.count;

        labels.push(item.language);
        values.push(total);
        itemColors.push(
          get(
            colors.languages,
            item.language.toUpperCase(),
            colors.languages.UNKNOWN
          )
        );
      });

      setChartColors(itemColors);
      setChartLabels(labels);
      setChartValues(values);
    };

    if (data) formatData(data);
  }, [colors.languages, data]);

  return (
    <div className="block max-space">
      <Styled.Wrapper tabIndex={0}>
        <Styled.Title>
          {t('DASHBOARD_SCREEN.VULNERABILITIES_BY_LANG')}
        </Styled.Title>

        {isLoading ? (
          <Styled.LoadingWrapper>
            <Icon name="loading" size="200px" className="loading" />
          </Styled.LoadingWrapper>
        ) : (
          <ReactApexChart
            height={250}
            width="100%"
            series={chartValues}
            options={{ ...options, colors: chartColors, labels: chartLabels }}
            type="donut"
          />
        )}
      </Styled.Wrapper>
    </div>
  );
};

export default VulnerabilitiesByLanguage;

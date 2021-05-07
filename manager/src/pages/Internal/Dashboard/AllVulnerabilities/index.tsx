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
import ReactApexChart from 'react-apexcharts';
import { ApexOptions } from 'apexcharts';
import Styled from './styled';
import { useTranslation } from 'react-i18next';
import { useTheme } from 'styled-components';
import { Icon } from 'components';
import { get } from 'lodash';
import { VulnerabilityBySeverity } from 'helpers/interfaces/DashboardData';

interface Props {
  isLoading: boolean;
  data: VulnerabilityBySeverity[];
}

const AllVulnerabilities: React.FC<Props> = ({ isLoading, data }) => {
  const { t } = useTranslation();
  const { colors, metrics } = useTheme();

  const [chartValues, setChartValues] = useState<number[]>([]);
  const [chartLabels, setChartLabels] = useState<string[]>([]);
  const [chartColors, setChartColors] = useState<string[]>([]);

  const options: ApexOptions = {
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
        fontSize: metrics.fontSize.xsmall,
        fontWeight: 100,
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
    const formatData = (data: VulnerabilityBySeverity[]) => {
      const itemColors: string[] = [];
      const labels: string[] = [];
      const values: number[] = [];

      Object.keys(data).forEach((key) => {
        const label = key.toUpperCase();
        const value = get(data, key, { count: 0 }).count;

        labels.push(key.toUpperCase());
        values.push(value);
        itemColors.push(
          get(colors.vulnerabilities, label, colors.vulnerabilities.DEFAULT)
        );
      });

      const total = values.reduce((item, current) => item + current);

      if (total > 0) {
        setChartColors(itemColors);
        setChartLabels(labels);
        setChartValues(values);
      } else {
        setChartColors([]);
        setChartLabels([]);
        setChartValues([]);
      }
    };

    if (data) formatData(data);
  }, [colors.vulnerabilities, data]);

  return (
    <div className="block max-space">
      <Styled.Wrapper tabIndex={0}>
        <Styled.Title>{t('DASHBOARD_SCREEN.ALL_VULNERABILITIES')}</Styled.Title>

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

export default AllVulnerabilities;

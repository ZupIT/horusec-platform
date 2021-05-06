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
import { ChartBarStacked } from 'helpers/interfaces/ChartData';
import { VulnerabilityByTime } from 'helpers/interfaces/DashboardData';
import { formatChartStacked } from 'helpers/formatters/chartData';

interface Props {
  data: VulnerabilityByTime[];
  isLoading: boolean;
}

const VulnerabilitiesTimeLine: React.FC<Props> = ({ data, isLoading }) => {
  const { t } = useTranslation();
  const { colors, metrics } = useTheme();

  const [chartData, setChartData] = useState<ChartBarStacked>(
    formatChartStacked(data, 'time', true)
  );

  const options: ApexOptions = {
    markers: {
      size: 0,
    },
    colors: Object.values(colors.vulnerabilities),
    noData: {
      text: t('DASHBOARD_SCREEN.CHART_NO_DATA'),
      style: {
        color: colors.chart.legend,
        fontSize: metrics.fontSize.large,
      },
    },
    legend: {
      position: 'top',
      horizontalAlign: 'left',
      offsetX: 40,
      labels: {
        colors: colors.chart.legend,
      },
    },
    chart: {
      fontFamily: 'SFRegular',
      stacked: true,
      animations: {
        enabled: false,
      },
      toolbar: {
        show: false,
      },
    },
    plotOptions: {
      bar: {
        horizontal: true,
      },
    },
    xaxis: {
      labels: {
        style: {
          colors: colors.chart.legend,
          fontSize: metrics.fontSize.small,
        },
      },
      categories: [],
    },
    yaxis: {
      title: {
        text: undefined,
      },
      labels: {
        style: {
          colors: colors.chart.legend,
          fontSize: metrics.fontSize.small,
        },
      },
    },
  };

  useEffect(() => {
    setChartData(formatChartStacked(data, 'time', true));
  }, [data]);

  return (
    <div className="block max-space">
      <Styled.Wrapper tabIndex={0}>
        <Styled.Title>
          {t('DASHBOARD_SCREEN.VULNERABILITY_TIMELINE')}
        </Styled.Title>

        <Styled.LoadingWrapper isLoading={isLoading}>
          <Icon name="loading" size="200px" className="loading" />
        </Styled.LoadingWrapper>

        <ReactApexChart
          height={250}
          width="100%"
          options={{
            ...options,
            xaxis: { ...options.xaxis, categories: chartData.categories },
          }}
          series={chartData.series}
          type="line"
        />
      </Styled.Wrapper>
    </div>
  );
};

export default VulnerabilitiesTimeLine;

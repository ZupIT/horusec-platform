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
import { format } from 'date-fns';
import { get, orderBy } from 'lodash';
import { ptBR, enUS, es } from 'date-fns/locale';
import useLanguage from 'helpers/hooks/useLanguage';

interface Props {
  data: VulnerabilityByTime[];
  isLoading: boolean;
}

const VulnerabilitiesTimeLine: React.FC<Props> = ({ data, isLoading }) => {
  const { t } = useTranslation();
  const { colors, metrics } = useTheme();
  const { currentLanguage } = useLanguage();

  const [ariaLabel, setAriaLabel] = useState<string>('');
  const [chartData, setChartData] = useState<ChartBarStacked>(
    formatChartStacked(data, 'time', true)
  );

  const options: ApexOptions = {
    markers: {
      size: 4,
    },
    stroke: {
      curve: 'smooth',
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
      offsetX: -20,
      offsetY: -5,
      labels: {
        colors: colors.chart.legend,
      },
    },
    chart: {
      fontFamily: 'SFRegular',
      stacked: false,
      animations: {
        enabled: true,
      },
      toolbar: {
        show: true,
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
          fontSize: metrics.fontSize.xsmall,
          fontWeight: 200,
        },
      },
      categories: [],
      axisBorder: {
        show: false,
      },
      axisTicks: {
        show: true,
      },
      type: 'datetime',
    },
    yaxis: {
      title: {
        text: undefined,
      },
      labels: {
        style: {
          colors: colors.chart.legend,
          fontSize: metrics.fontSize.xsmall,
          fontWeight: 200,
        },
      },
      axisBorder: {
        show: false,
      },
      axisTicks: {
        show: false,
      },
    },
  };

  const makeAriaLabel = () => {
    if (data) {
      let ariaText = t('DASHBOARD_SCREEN.ARIA_CHART_TIMELINE');
      const listOfData = orderBy(data, (item) => item.time, 'desc');

      const locale = get(
        {
          es,
          ptBR,
          enUS,
        },
        currentLanguage.i18nValue,
        enUS
      );

      listOfData.forEach((item, index) => {
        if (index <= 2) {
          const legend = format(new Date(item.time), 'PPP', { locale });
          const value =
            item.critical.count +
            item.high.count +
            item.medium.count +
            item.low.count +
            item.info.count +
            item.unknown.count;

          ariaText =
            ariaText +
            `, ${t('DASHBOARD_SCREEN.ARIA_CHART_TIMELINE_ITEM', {
              legend,
              value,
            })}`;
        }
      });

      setAriaLabel(`${ariaText}.`);
    }
  };

  useEffect(() => {
    setChartData(formatChartStacked(data, 'time', true));
    makeAriaLabel();
    // eslint-disable-next-line
  }, [data]);

  return (
    <Styled.Wrapper tabIndex={0} aria-label={ariaLabel}>
      <Styled.Title>
        {t('DASHBOARD_SCREEN.VULNERABILITY_TIMELINE')}
      </Styled.Title>

      <Styled.LoadingWrapper isLoading={isLoading}>
        <Icon name="loading" size="200px" className="loading" />
      </Styled.LoadingWrapper>

      <ReactApexChart
        height={250}
        width="95%"
        options={{
          ...options,
          xaxis: { ...options.xaxis, categories: chartData.categories },
        }}
        series={chartData.series}
        type="line"
      />
    </Styled.Wrapper>
  );
};

export default VulnerabilitiesTimeLine;

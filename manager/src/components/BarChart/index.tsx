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
import Styled from './styled';
import { generateRandomColor } from 'helpers/colors';
import { BarCharRow } from 'helpers/interfaces/BarChartRow';
import { useTranslation } from 'react-i18next';
import { Icon } from 'components';

interface BarChartProps {
  data: BarCharRow[];
  title: string;
  isLoading: boolean;
  showBackOption?: boolean;
  isVertical?: boolean;
  hasSmallLegend?: boolean;
  onClickRow?: (row: BarCharRow) => any;
  onClickBack?: () => any;
}

const BarChart: React.FC<BarChartProps> = ({
  data,
  isLoading,
  title,
  onClickRow,
  onClickBack,
  showBackOption,
  isVertical,
  hasSmallLegend,
}) => {
  const { t } = useTranslation();
  const [ariaLabel, setAriaLabel] = useState<string>('');

  const calculatePercentageOfBar = (value: number) => {
    const total = data.reduce((a, b) => {
      return { legend: null, value: a.value + b.value };
    });
    return `${(value * 100) / total.value}%`;
  };

  const renderRow = ({ value, legend, color }: BarCharRow) => (
    <Styled.Row
      isVertical={isVertical}
      key={legend}
      onClick={() => onClickRow({ value, legend })}
      id={`${title}_${legend}`.replaceAll(' ', '_')}
    >
      <Styled.Value isVertical={isVertical}>{value}</Styled.Value>

      <Styled.Bar
        color={color || generateRandomColor()}
        size={calculatePercentageOfBar(value)}
        isVertical={isVertical}
      />
      <Styled.Legend hasSmallLegend={hasSmallLegend} isVertical={isVertical}>
        {legend}
      </Styled.Legend>
    </Styled.Row>
  );

  const renderLoading = () => {
    return (
      <Styled.LoadingWrapper>
        <Icon name="loading" size="130px" />;
      </Styled.LoadingWrapper>
    );
  };

  const mountAriaLabel = () => {
    const direction = isVertical ? 'vertical' : 'horizontal';
    let ariaText = t('DASHBOARD_SCREEN.ARIA_CHART_BAR', { direction, title });

    data.forEach((item) => {
      const { legend, value } = item;
      ariaText =
        ariaText +
        `, ${t('DASHBOARD_SCREEN.ARIA_CHART_BAR_ITEM', { legend, value })}`;
    });

    setAriaLabel(`${ariaText}.`);
  };

  useEffect(() => {
    if (data) {
      mountAriaLabel();
    }
    // eslint-disable-next-line
  }, [data]);

  return (
    <Styled.Wrapper tabIndex={0} aria-label={ariaLabel} role="figure">
      <Styled.Header>
        <Styled.Title>{title}</Styled.Title>

        {showBackOption ? (
          <Styled.BackWrapper onClick={onClickBack}>
            <Icon name="left-arrow" size="20px" />

            <Styled.Back>{t('GENERAL.BACK')}</Styled.Back>
          </Styled.BackWrapper>
        ) : null}
      </Styled.Header>

      <Styled.WrapperChart isVertical={isVertical}>
        {data.length <= 0 && !isLoading ? (
          <Styled.Empty>{t('DASHBOARD_SCREEN.CHART_NO_DATA')}</Styled.Empty>
        ) : null}

        {!isLoading ? data.map((item) => renderRow(item)) : renderLoading()}
      </Styled.WrapperChart>
    </Styled.Wrapper>
  );
};

export default BarChart;

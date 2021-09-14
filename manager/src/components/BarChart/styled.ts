/**
 * Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

import styled, { keyframes, css } from 'styled-components';

interface VerticalProps {
  isVertical?: boolean;
}

interface BarProps extends VerticalProps {
  color?: string;
  size: string;
}
interface LegendProps extends VerticalProps {
  hasSmallLegend?: boolean;
}

const Wrapper = styled.div`
  background-color: ${({ theme }) => theme.colors.background.secundary};
  border-radius: 4px;
  width: 100%;
  min-width: 170px;
  height: auto;
  min-height: 330px;
  display: flex;
  flex-direction: column;
  padding: 25px;
`;

const Header = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 30px;
`;

const Title = styled.h2`
  color: ${({ theme }) => theme.colors.chart.title};
  font-size: ${({ theme }) => theme.metrics.fontSize.xlarge};
  border-radius: 4px;
  font-weight: normal;
  display: block;
`;

const BackWrapper = styled.span`
  display: flex;
  align-items: center;
  cursor: pointer;

  :hover {
    transform: scale(1.2);
  }
`;

const Back = styled.span`
  display: block;
  color: ${({ theme }) => theme.colors.chart.title};
  margin-left: 5px;
`;

const Empty = styled.h2`
  color: ${({ theme }) => theme.colors.dataTable.column.text};
  font-size: ${({ theme }) => theme.metrics.fontSize.large};
  font-weight: normal;
  text-align: center;
  display: block;
  line-height: 170px;
`;

const WrapperChart = styled.ul<VerticalProps>`
  list-style: none;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  height: 200px;

  ${({ isVertical, theme }) =>
    isVertical &&
    css`
      display: flex;
      flex-direction: row;
    `}
`;

const Row = styled.li<VerticalProps>`
  width: 100%;
  display: flex;
  align-items: center;
  position: relative;

  ${({ isVertical, theme }) =>
    isVertical &&
    css`
      height: 100%;
      display: flex;
      flex-direction: column;
      align-itens: center;
    `}
`;

const Value = styled.span<VerticalProps>`
  color: ${({ theme }) => theme.colors.chart.legend};
  margin-right: 15px;
  display: block;
  min-width: 20px;
  text-align: start;

  ${({ isVertical, theme }) =>
    isVertical &&
    css`
      text-align: center;
      margin-right: 0px;
      margin-bottom: 0px;
    `}
`;

const resizeBar = (perc: string) => keyframes`
  from {
    width: 0px;
  }
  to {
    width: ${perc};
  }
`;

const resizeBarVertical = (perc: string) => keyframes`
  from {
    height: 0px;
  }
  to {
    height: ${perc};
  }
`;

const Bar = styled.div<BarProps>`
  display: block;
  width: 100%;
  height: 25px;
  background-color: ${({ theme }) => theme.colors.chart.background};
  position: relative;
  cursor: pointer;

  :hover {
    box-shadow: 0 0 6px rgba(33, 33, 33, 0.8);
  }

  ::before {
    content: '';
    display: block;
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    background-color: ${({ color }) => color};
    width: ${({ size }) => size};
    animation: ${({ size }) => css`
      ${resizeBar(size)} 1s ease-in-out
    `};
  }

  ${({ isVertical, color, size }) =>
    isVertical &&
    css`
      width: 100%;
      max-width: 60%;
      height: 100%;

      ::before {
        content: '';
        display: block;
        position: absolute;
        top: 0;
        left: 0;
        height: ${size};
        background-color: ${color};
        width: 100%;
        animation: ${resizeBarVertical(size)} 1s ease-in-out};
        transform: rotate(180deg);
      }
    `}
`;

const Legend = styled.span<LegendProps>`
  color: ${({ theme }) => theme.colors.chart.legend};
  margin-left: 20px;
  min-width: 150px;
  text-align: start;
  white-space: nowrap;
  overflow: auto;
  text-overflow: ellipsis;

  ${({ hasSmallLegend }) =>
    hasSmallLegend &&
    css`
      min-width: 90px;
    `}

  ${({ isVertical, theme }) =>
    isVertical &&
    css`
      min-width: 80px;
      text-align: center;
      margin-left: 0px;
      margin-top: 8px;
      font-size: ${theme.metrics.fontSize.xsmall};
    `}
`;

const LoadingWrapper = styled.div`
  width: 100%;
  height: 200px;
  display: flex;
  align-items: center;
  justify-content: center;
`;

export default {
  Wrapper,
  Title,
  WrapperChart,
  Bar,
  Row,
  Legend,
  Value,
  Empty,
  Header,
  Back,
  BackWrapper,
  LoadingWrapper,
};

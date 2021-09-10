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

import styled from 'styled-components';
import { Icon as IconComponent } from 'components';

const Card = styled.li`
  position: relative;
  background-color: ${({ theme }) => theme.colors.background.highlight};
  margin-right: 60px;
  margin-bottom: 30px;
  padding: 20px;
  border-radius: 5px;
  width: 290px;
  min-height: 140px;

  :hover {
    footer {
      height: 40px;

      * {
        visibility: visible;
        opacity: 1;
      }
    }
  }
`;

const Title = styled.span`
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.xxlarge};
  display: block;
  margin-bottom: 8px;
  max-width: 250px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
`;

const Icon = styled(IconComponent)`
  margin-right: 10px;
`;

const Description = styled.span`
  color: ${({ theme }) => theme.colors.text.secundary};
  font-size: ${({ theme }) => theme.metrics.fontSize.medium};
  height: 36px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
`;

const Info = styled.div`
  margin-top: 15px;
  display: flex;
  align-items: center;
  justify-content: space-between;
`;

const InfoItem = styled.span`
  color: ${({ theme }) => theme.colors.text.secundary};
  font-size: ${({ theme }) => theme.metrics.fontSize.xsmall};
  display: flex;
  align-items: center;
`;

const InfoIcon = styled(IconComponent)`
  margin-right: 8px;
`;

const OptionsBar = styled.footer`
  position: absolute;
  bottom: 0;
  left: 0;
  background-color: ${({ theme }) => theme.colors.primary};
  width: 100%;
  height: 10px;
  border-bottom-left-radius: 5px;
  border-bottom-right-radius: 5px;
  transition: 0.7s ease;
  display: flex;
  justify-content: space-around;
  align-items: center;

  * {
    transition: 0.5s ease;
    visibility: hidden;
    opacity: 0;
  }
`;

const Option = styled.button`
  border: none;
  background: none;
  outline: none;
  cursor: pointer;
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.medium};
  height: 100%;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;

  :hover {
    font-weight: bold;
    transform: scale(1.1);
  }
`;

export default {
  Card,
  Title,
  Icon,
  OptionsBar,
  Description,
  Info,
  InfoItem,
  InfoIcon,
  Option,
};

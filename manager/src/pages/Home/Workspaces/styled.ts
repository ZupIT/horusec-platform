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

import styled from 'styled-components';
import HomeStyles from '../styled';

const { Title, SearchWrapper, ListWrapper, Phrase, Message, MessageText } =
  HomeStyles;

const Subtitle = styled.h2`
  color: ${({ theme }) => theme.colors.chart.title};
  font-size: ${({ theme }) => theme.metrics.fontSize.xxlarge};
  margin-top: 40px;
`;

const List = styled.ul`
  display: flex;
  margin-top: 40px;
  list-style: none;
`;

export default {
  Title,
  SearchWrapper,
  ListWrapper,
  Subtitle,
  Phrase,
  List,
  Message,
  MessageText,
};

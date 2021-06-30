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
import { Icon as IconComponent } from 'components';

const {
  Title,
  SearchWrapper,
  ListWrapper,
  Phrase,
  Message,
  MessageText,
  List,
} = HomeStyles;

const Head = styled.div`
  display: flex;
  justify-content: space-between;
`;

const Icon = styled(IconComponent)`
  margin-right: 10px;
`;

const TitleWrapper = styled.div`
  display: flex;
  align-items: center;
`;

export default {
  Title,
  TitleWrapper,
  SearchWrapper,
  ListWrapper,
  Head,
  Phrase,
  List,
  Message,
  MessageText,
  Icon,
};

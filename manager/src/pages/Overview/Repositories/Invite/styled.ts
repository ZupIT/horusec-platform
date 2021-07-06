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

import styled, { css } from 'styled-components';
import { Icon } from 'components';

const Header = styled.div`
  background-color: ${({ theme }) => theme.colors.background.secundary};
  border-radius: 4px;
  padding: 22px;
  display: flex;
  align-items: center;
  justify-content: space-between;
`;

const Content = styled.div`
  padding: 25px 15px;
  background-color: ${({ theme }) => theme.colors.background.secundary};
  border-radius: 4px;
  overflow: hidden;
  max-width: 95vw;
  height: 100%;
`;

const Wrapper = styled.section`
  padding: 35px 15px;
  width: 100%;
  height: 95%;
  display: flex;
  flex-direction: column;
`;

const TitleContent = styled.div`
  display: flex;
  gap: 15px;
  align-items: center;
`;

const Title = styled.div`
  color: ${({ theme }) => theme.colors.dialog.text};
  font-size: ${({ theme }) => theme.metrics.fontSize.xxlarge};
  line-height: 22px;
`;

const Close = styled(Icon)`
  transition-duration: 0.5s;
  transition-property: transform;

  :hover {
    transform: rotate(90deg);
    -webkit-transform: rotate(90deg);
    cursor: pointer;
  }
`;

const SubTitle = styled.div`
  color: ${({ theme }) => theme.colors.text.secundary};
  font-size: ${({ theme }) => theme.metrics.fontSize.small};
  margin-bottom: 20px;
`;

const HelpIcon = styled(Icon)`
  cursor: pointer;

  :hover {
    transform: scale(1.2);
  }
`;

export default {
  Header,
  Wrapper,
  Title,
  TitleContent,
  Close,
  Content,
  SubTitle,
  HelpIcon,
};

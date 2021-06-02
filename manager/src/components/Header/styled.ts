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

import { Icon } from 'components';
import styled, { css } from 'styled-components';

interface ItemProps {
  active?: boolean;
}

const Wrapper = styled.div`
  padding-left: 15px;
  position: sticky;
`;

const Header = styled.header`
  background-color: ${({ theme }) => theme.colors.background.secundary};
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  margin-bottom: 5px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 45px 10px 15px;
`;
const Title = styled.div`
  display: flex;
  align-items: center;
`;

const Text = styled.h1`
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.xlarge};
  padding: 10px 0;
  margin-left: 10px;
  font-weight: normal;
`;

const List = styled.ul`
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
`;

const Item = styled.li<ItemProps>`
  margin-left: 30px;
  display: flex;
  align-items: center;
  cursor: pointer;

  i {
    margin: 0;
  }

  ${({ active }) =>
    active &&
    css`
      span,
      svg,
      * {
        color: ${({ theme }) => theme.colors.active};
      }
    `};

  :hover {
    span,
    svg,
    * {
      color: ${({ theme }) => theme.colors.active};
    }
  }
`;

const ConfigIcon = styled(Icon)`
  cursor: pointer;
  color: ${({ theme }) => theme.colors.text.opaque};
`;

const ConfigText = styled.span`
  color: ${({ theme }) => theme.colors.text.opaque};
  transition: all ease 0.2s;
  font-size: ${({ theme }) => theme.metrics.fontSize.small};
  margin-left: 5px;
`;

export default {
  Header,
  Title,
  Text,
  Wrapper,
  List,
  Item,
  ConfigIcon,
  ConfigText,
};

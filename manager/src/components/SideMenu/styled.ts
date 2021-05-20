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

interface RouterItemProps {
  isActive: boolean;
}

const SideMenu = styled.div`
  background-color: ${({ theme }) => theme.colors.background.primary};
  min-width: 165px;
  max-width: 165px;
  display: flex;
  flex-direction: column;
  z-index: 2;
`;

const WrapperLogoRoutes = styled.div`
  flex: 1;
`;

const OptionsList = styled.ul`
  display: flex;
  flex-direction: column;
  padding: 20px 0px 20px 7.5px;
  list-style: none;
`;

const Logo = styled.img`
  display: block;
  margin: 24px 17px;
  width: 100px;
  height: 22px;
`;

const RoutesList = styled.ul`
  margin-top: 20px;
`;

const Nav = styled.nav``;

const RouteItem = styled.li<RouterItemProps>`
  cursor: pointer;
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.small};
  padding: 17px;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  position: relative;
  transition: background-color 0.6s;

  :hover {
    background-color: ${({ theme }) => theme.colors.background.secundary};
    border-left: 3px solid ${({ theme }) => theme.colors.active};
  }

  ${({ isActive }) =>
    isActive &&
    css`
      color: ${({ theme }) => theme.colors.active};
      background-color: ${({ theme }) => theme.colors.background.secundary};
      border-left: 3px solid ${({ theme }) => theme.colors.active};

      svg {
        color: ${({ theme }) => theme.colors.active};
      }
    `};
`;

const RouteName = styled.span`
  display: block;
  margin-left: 13px;
`;

const SelectWrapper = styled.div`
  margin-left: 17px;
`;

const SubRoutes = styled.ul`
  margin-bottom: 20px;
`;

const SubRouteItem = styled.li<RouterItemProps>`
  cursor: pointer;
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.small};
  padding: 10px;
  padding-left: 35px;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  position: relative;
  transition: background-color 0.6s;

  :hover {
    color: ${({ theme }) => theme.colors.active};
  }

  ${({ isActive }) =>
    isActive &&
    css`
      color: ${({ theme }) => theme.colors.active};

      &::before {
        content: '↠';
        margin-right: 5px;
      }
    `};
`;

export default {
  Nav,
  SideMenu,
  Logo,
  RoutesList,
  RouteItem,
  RouteName,
  WrapperLogoRoutes,
  OptionsList,
  SelectWrapper,
  SubRoutes,
  SubRouteItem,
};

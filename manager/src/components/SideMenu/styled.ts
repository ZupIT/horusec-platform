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

import styled, { css } from 'styled-components';

interface MinimizedProps {
  isMinimized?: boolean;
}
interface RouterItemProps extends MinimizedProps {
  isActive: boolean;
}

const SideMenu = styled.div<MinimizedProps>`
  will-change: transform, box-shadow, z-index;
  background-color: ${({ theme }) => theme.colors.background.primary};
  min-width: 165px;
  max-width: 165px;
  display: flex;
  flex-direction: column;
  z-index: 2;
  position: relative;
  transition: all ease 0.6s;
  transition: transform 300ms linear;

  ${({ isMinimized }) =>
    isMinimized &&
    css`
      min-width: 50px;
      max-width: 50px;
    `};
`;

const SizeHandler = styled.button`
  z-index: 2;
  position: absolute;
  background: none;
  cursor: pointer;
  background-color: ${({ theme }) => theme.colors.background.secundary};
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 5px;
  border: none;
  right: -10px;
  bottom: 20px;
  border-radius: 50%;
  transition: ease all 0.3s;

  &:hover {
    transform: scale(1.15);
  }
`;

const WrapperLogoRoutes = styled.div`
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const Logo = styled.img`
  display: block;
  margin: 24px 0px;
  width: 100px;
  height: 22px;
`;

const NameWrapper = styled.div`
  display: block;
  width: 100%;
  padding: 0 5px;
`;

const NameTitle = styled.span`
  display: block;
  color: ${({ theme }) => theme.colors.text.secundary};
  font-size: ${({ theme }) => theme.metrics.fontSize.small};
  text-align: center;
  margin-bottom: 5px;
`;

const NameText = styled.span`
  display: block;
  color: ${({ theme }) => theme.colors.text.primary};
  font-size: ${({ theme }) => theme.metrics.fontSize.large};
  text-align: center;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
`;

const RoutesList = styled.ul`
  margin-top: 20px;
`;

const Nav = styled.nav`
  width: 100%;
`;

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

const RouteName = styled.span<MinimizedProps>`
  display: block;
  margin-left: 13px;

  ${({ isMinimized }) =>
    isMinimized &&
    css`
      display: none;
    `};
`;

const SelectWrapper = styled.div``;

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

  ${({ isMinimized }) =>
    isMinimized &&
    css`
      display: none;
    `};

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
  SelectWrapper,
  SubRoutes,
  SubRouteItem,
  SizeHandler,
  NameWrapper,
  NameText,
  NameTitle,
};

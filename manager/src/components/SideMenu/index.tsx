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

import React, { useState } from 'react';
import Styled from './styled';
import HorusecLogo from 'assets/logos/horusec.svg';
import HorusecLogoMin from 'assets/logos/horusec_minimized.svg';
import { useTranslation } from 'react-i18next';
import { Icon } from 'components';
import { useHistory, useRouteMatch, Link } from 'react-router-dom';
import { InternalRoute } from 'helpers/interfaces/InternalRoute';
import ReactTooltip from 'react-tooltip';
import useParamsRoute from 'helpers/hooks/useParamsRoute';

const SideMenu: React.FC = () => {
  const history = useHistory();

  const { t } = useTranslation();
  const [isMinimized, setIsMinimized] = useState<boolean>(false);

  const { workspace, workspaceId, repositoryId, repository } = useParamsRoute();

  const routes: InternalRoute[] = [
    {
      name: t('SIDE_MENU.DASHBOARD'),
      icon: 'pie',
      path: '/dashboard',
      roles: ['admin', 'member'],
    },
    {
      name: t('SIDE_MENU.VULNERABILITIES'),
      icon: 'shield',
      path: '/vulnerabilities',
      roles: ['admin', 'member'],
    },
    {
      name: t('SIDE_MENU.TOKENS'),
      icon: 'lock',
      path: '/tokens',
      roles: ['admin', 'member'],
    },
    {
      name: t('SIDE_MENU.USERS'),
      icon: 'users',
      path: '/users',
      roles: ['admin', 'member'],
    },
    {
      name: t('SIDE_MENU.WEBHOOK'),
      icon: 'webhook',
      path: '/webhooks',
      roles: ['admin', 'member'],
    },
  ];

  const handleSelectedRoute = (route: InternalRoute) => {
    const fullPath = `/overview/workspace/${workspaceId}${
      repositoryId ? `/repository/${repositoryId}${route.path}` : route.path
    }`;

    history.push(fullPath);
  };

  const renderRoute = (route: InternalRoute, index: number) => {
    if (route.roles.includes(workspace?.role)) {
      if (!route?.rule || (route?.rule && route?.rule())) {
        return (
          <div key={index}>
            <Styled.RouteItem
              tabIndex={0}
              isActive={window.location.pathname.includes(route?.path)}
              onClick={() => handleSelectedRoute(route)}
              onKeyPress={() => handleSelectedRoute(route)}
            >
              <Icon name={route.icon} size="15" />

              <Styled.RouteName isMinimized={isMinimized}>
                {route.name}
              </Styled.RouteName>
            </Styled.RouteItem>

            <Styled.SubRoutes>
              {route?.subRoutes?.map((subRoute, index) => {
                if (subRoute?.roles?.includes(workspace?.role)) {
                  return (
                    <Styled.SubRouteItem
                      isActive={window.location.pathname.includes(
                        subRoute.path
                      )}
                      key={index}
                      onClick={() => handleSelectedRoute(subRoute)}
                      isMinimized={isMinimized}
                    >
                      {subRoute?.name}
                    </Styled.SubRouteItem>
                  );
                }
                return null;
              })}
            </Styled.SubRoutes>
          </div>
        );
      }
    }
  };

  return (
    <>
      <Styled.SideMenu isMinimized={isMinimized}>
        <Styled.SizeHandler
          aria-label={
            isMinimized ? t('SIDE_MENU.MAX_MENU') : t('SIDE_MENU.MIN_MENU')
          }
          data-tip={
            isMinimized ? t('SIDE_MENU.MAX_MENU') : t('SIDE_MENU.MIN_MENU')
          }
          onClick={() => setIsMinimized(!isMinimized)}
        >
          <Icon
            name={isMinimized ? 'page-next' : 'page-previous'}
            size="18px"
          />
        </Styled.SizeHandler>

        <Styled.WrapperLogoRoutes>
          <Link to="/home" about="Horusec Logo">
            <Styled.Logo
              src={isMinimized ? HorusecLogoMin : HorusecLogo}
              alt="Horusec Logo"
            />
          </Link>

          {isMinimized ? null : (
            <Styled.NameWrapper>
              <Styled.NameTitle>Overview: </Styled.NameTitle>
              <Styled.NameText>
                {repository?.name || workspace?.name}
              </Styled.NameText>
            </Styled.NameWrapper>
          )}

          <Styled.Nav aria-label={t('SIDE_MENU.ARIA_TITLE')}>
            <Styled.RoutesList>
              {routes.map((route, index) => renderRoute(route, index))}
            </Styled.RoutesList>
          </Styled.Nav>
        </Styled.WrapperLogoRoutes>
      </Styled.SideMenu>

      <ReactTooltip />
    </>
  );
};

export default SideMenu;

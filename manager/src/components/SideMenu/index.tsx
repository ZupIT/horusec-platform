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
import { useHistory, Link } from 'react-router-dom';
import { InternalRoute } from 'helpers/interfaces/InternalRoute';
import ReactTooltip from 'react-tooltip';
import useParamsRoute from 'helpers/hooks/useParamsRoute';
import usePermissions from 'helpers/hooks/usePermissions';

const SideMenu: React.FC = () => {
  const history = useHistory();

  const { t } = useTranslation();
  const [isMinimized, setIsMinimized] = useState<boolean>(false);

  const { workspace, workspaceId, repositoryId, repository } = useParamsRoute();
  const { isAuthorizedRoute } = usePermissions();

  const isRepositoryOverview = !!repositoryId;

  const routes: InternalRoute[] = [
    {
      name: t('SIDE_MENU.DASHBOARD'),
      icon: 'pie',
      path: 'dashboard',
    },
    {
      name: t('SIDE_MENU.VULNERABILITIES'),
      icon: 'shield',
      path: 'vulnerabilities',
    },
    {
      name: t('SIDE_MENU.TOKENS'),
      icon: 'lock',
      path: 'tokens',
    },
    {
      name: t('SIDE_MENU.USERS'),
      icon: 'users',
      path: 'users',
    },
    {
      name: t('SIDE_MENU.WEBHOOK'),
      icon: 'webhook',
      path: 'webhooks',
    },
  ];

  const handleSelectedRoute = (route: InternalRoute) => {
    const fullPath = `/overview/workspace/${workspaceId}/${
      isRepositoryOverview
        ? `repository/${repositoryId}/${route.path}`
        : route.path
    }`;

    history.push(fullPath);
  };

  const renderRoute = (route: InternalRoute, index: number) => {
    if (isAuthorizedRoute(route.path)) {
      return (
        <Styled.RouteItem
          key={index}
          tabIndex={0}
          isActive={window.location.pathname.includes(route?.path)}
          onClick={() => handleSelectedRoute(route)}
          onKeyPress={() => handleSelectedRoute(route)}
        >
          <Icon name={route.icon} size="15px" />

          <Styled.RouteName isMinimized={isMinimized}>
            {route.name}
          </Styled.RouteName>
        </Styled.RouteItem>
      );
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
              <Styled.NameTitle>
                {isRepositoryOverview ? 'Repository:' : 'Workspace:'}
              </Styled.NameTitle>
              <Styled.NameText>
                {isRepositoryOverview ? repository?.name : workspace?.name}
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

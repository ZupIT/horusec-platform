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
import { useTranslation } from 'react-i18next';
import { Icon } from 'components';
import { useHistory } from 'react-router-dom';
import { InternalRoute } from 'helpers/interfaces/InternalRoute';
import { find } from 'lodash';
import useWorkspace from 'helpers/hooks/useWorkspace';
import SelectMenu from 'components/SelectMenu';
import { Workspace } from 'helpers/interfaces/Workspace';

const SideMenu: React.FC = () => {
  const history = useHistory();
  const {
    currentWorkspace,
    allWorkspaces,
    handleSetCurrentWorkspace,
  } = useWorkspace();
  const { t } = useTranslation();
  const [selectedRoute, setSelectedRoute] = useState<InternalRoute>();
  const [selectedSubRoute, setSelectedSubRoute] = useState<InternalRoute>();

  const routes: InternalRoute[] = [
    {
      name: t('SIDE_MENU.DASHBOARD'),
      icon: 'pie',
      type: 'route',
      path: '/home/dashboard',
      roles: ['admin', 'member'],
      subRoutes: [
        {
          name: t('SIDE_MENU.WORKSPACE'),
          icon: 'grid',
          path: '/home/dashboard/workspace',
          type: 'subRoute',
          roles: ['admin'],
        },
        {
          name: t('SIDE_MENU.REPOSITORIES'),
          icon: 'columns',
          path: '/home/dashboard/repositories',
          type: 'subRoute',
          roles: ['admin', 'member'],
        },
      ],
    },
    {
      name: t('SIDE_MENU.VULNERABILITIES'),
      icon: 'shield',
      path: '/home/vulnerabilities',
      type: 'route',
      roles: ['admin', 'member'],
    },
    {
      name: t('SIDE_MENU.REPOSITORIES'),
      icon: 'columns',
      path: '/home/repositories',
      type: 'route',
      roles: ['admin', 'member'],
    },
    {
      name: t('SIDE_MENU.WEBHOOK'),
      icon: 'webhook',
      path: '/home/webhooks',
      type: 'route',
      roles: ['admin'],
    },
  ];

  const handleSelectedRoute = (route: InternalRoute) => {
    if (route.type === 'route') {
      setSelectedRoute((state) => {
        if (
          state &&
          state?.subRoutes &&
          route?.subRoutes &&
          !selectedSubRoute
        ) {
          return null;
        }
        return route;
      });
      setSelectedSubRoute(null);

      if (!route?.subRoutes) {
        history.push(route.path);
      } else {
        setTimeout(() => {
          const firstItemSubRoute = document.getElementById('sub-route-0');
          firstItemSubRoute?.focus();
        }, 1000);
      }
    } else {
      setSelectedSubRoute(route);
      history.push(route.path);
    }
  };

  const renderRoute = (route: InternalRoute, index: number) => {
    if (route.roles.includes(currentWorkspace?.role)) {
      if (!route?.rule || (route?.rule && route?.rule())) {
        return (
          <Styled.RouteItem
            key={index}
            tabIndex={0}
            isActive={route.path === selectedRoute?.path}
            onClick={() => handleSelectedRoute(route)}
            onKeyPress={() => handleSelectedRoute(route)}
          >
            <Icon name={route.icon} size="15" />

            <Styled.RouteName>{route.name}</Styled.RouteName>
          </Styled.RouteItem>
        );
      }
    }
  };

  const fetchSubRoutes = () =>
    find(routes, { path: selectedRoute?.path })?.subRoutes || [];

  const renderSubRoute = (subRoute: InternalRoute, index: number) => {
    if (subRoute.roles.includes(currentWorkspace?.role)) {
      return (
        <Styled.SubRouteItem
          id={`sub-route-${index}`}
          key={index}
          tabIndex={0}
          isActive={subRoute.path === selectedSubRoute?.path}
          onClick={() => handleSelectedRoute(subRoute)}
        >
          <Icon name={subRoute.icon} size="15" />

          <Styled.RouteName>{subRoute.name}</Styled.RouteName>
        </Styled.SubRouteItem>
      );
    }
  };

  const handleSelectedWorkspace = (workspace: Workspace) => {
    handleSetCurrentWorkspace(workspace);
    history.replace('/home/dashboard');
    setSelectedRoute(null);
    setSelectedSubRoute(null);
  };

  return (
    <>
      <Styled.SideMenu>
        <Styled.WrapperLogoRoutes>
          <Styled.Logo src={HorusecLogo} alt="Horusec Logo" />

          {allWorkspaces && allWorkspaces.length > 0 ? (
            <Styled.SelectWrapper>
              <SelectMenu
                title={'WORKSPACE'}
                value={currentWorkspace?.name}
                options={allWorkspaces.map((el) => ({
                  title: el.name,
                  action: () => handleSelectedWorkspace(el),
                }))}
                fixItem={{
                  title: t('SIDE_MENU.MANAGE_WORKSPACES'),
                  action: () => history.push('/home/workspaces'),
                }}
              />
            </Styled.SelectWrapper>
          ) : null}

          <nav aria-label={t('SIDE_MENU.ARIA_TITLE')}>
            <Styled.RoutesList>
              {routes.map((route, index) => renderRoute(route, index))}
            </Styled.RoutesList>
          </nav>
        </Styled.WrapperLogoRoutes>
      </Styled.SideMenu>

      <Styled.SubMenu isActive={!!selectedRoute?.subRoutes}>
        <Styled.SubRoutesList>
          {fetchSubRoutes().map((subRoute, index) =>
            renderSubRoute(subRoute, index)
          )}
        </Styled.SubRoutesList>
      </Styled.SubMenu>
    </>
  );
};

export default SideMenu;

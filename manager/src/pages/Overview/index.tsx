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

import React, { useEffect } from 'react';
import { Redirect, Switch, useRouteMatch, useLocation } from 'react-router-dom';
import { PrivateRoute } from 'components';
import InternalLayout from 'layouts/Internal';
import { setCurrentPath } from 'helpers/localStorage/currentPage';

import Dashboard from 'pages/Overview/Dashboard';
import Webhooks from 'pages/Overview/Webhooks';
import Tokens from './Tokens';
import Vulnerabilities from './Vulnerabilities';
import Users from './Users';

function InternalRoutes() {
  const { path } = useRouteMatch();
  const location = useLocation();

  useEffect(() => {
    setCurrentPath(location?.pathname);
  }, [location?.pathname]);

  return (
    <InternalLayout>
      <Switch>
        <PrivateRoute
          exact
          path={`${path}/workspace/:workspaceId/users`}
          component={() => <Users />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspace/:workspaceId/repository/:repositoryId/users`}
          component={() => <Users />}
        />

        <PrivateRoute
          path={`${path}/workspace/:workspaceId/dashboard`}
          exact
          component={() => <Dashboard type="workspace" />}
        />

        <PrivateRoute
          path={`${path}/workspace/:workspaceId/repository/:repositoryId/dashboard`}
          exact
          component={() => <Dashboard type="repository" />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspace/:workspaceId/vulnerabilities`}
          component={() => <Vulnerabilities />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspace/:workspaceId/repository/:repositoryId/vulnerabilities`}
          component={() => <Vulnerabilities />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspace/:workspaceId/tokens`}
          component={() => <Tokens />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspace/:workspaceId/repository/:repositoryId/tokens`}
          component={() => <Tokens />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspace/:workspaceId/webhooks`}
          component={() => <Webhooks type="workspace" />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspace/:workspaceId/repository/:repositoryId/webhooks`}
          component={() => <Webhooks type="repository" />}
        />

        <Redirect from="*" to="/home" />
      </Switch>
    </InternalLayout>
  );
}

export default InternalRoutes;

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

import React from 'react';
import { Redirect, Switch, useRouteMatch } from 'react-router-dom';
import { PrivateRoute } from 'components';
import InternalLayout from 'layouts/Internal';

import Dashboard from 'pages/Overview/Dashboard';
import Webhooks from 'pages/Overview/Webhooks';
import RepositoryTokens from './Repositories/Tokens';
import RepositoryInvite from './Repositories/Invite';
import WorkspaceTokens from './Workspaces/Tokens';
import WorkspaceUsers from './Workspaces/Users';
import Vulnerabilities from './Vulnerabilities';
import useParamsRoute from 'helpers/hooks/useParamsRoute';

function InternalRoutes() {
  const { path } = useRouteMatch();
  const {
    workspace: isAdminOfWorkspace,
    workspaceId,
    repositoryId,
  } = useParamsRoute();

  return (
    <InternalLayout>
      <Switch>
        <PrivateRoute
          exact
          path={`${path}/workspaces/:workspaceId/tokens`}
          component={() => <WorkspaceTokens />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspaces/:workspaceId/users`}
          component={() => <WorkspaceUsers />}
        />

        <Redirect
          exact
          from={`${path}/workspaces/:workspaceId`}
          to={`${path}/workspaces/${workspaceId}/dashboard`}
        />

        <PrivateRoute
          path={`${path}/workspaces/:workspaceId/dashboard`}
          exact
          component={() => <Dashboard type="workspace" />}
        />

        <PrivateRoute
          path={`${path}/workspaces/:workspaceId/repository/:repositoryId/dashboard`}
          exact
          component={() => <Dashboard type="repository" />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspaces/:workspaceId/repository/:repositoryId/invite`}
          component={() => <RepositoryInvite />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspaces/:workspaceId/repository/:repositoryId/tokens`}
          component={() => <RepositoryTokens />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspaces/:workspaceId/repository/:repositoryId/vulnerabilities`}
          component={() => <Vulnerabilities />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspaces/:workspaceId/webhooks`}
          component={() => <Webhooks />}
        />
        <Redirect from="*" to="/home" />
      </Switch>
    </InternalLayout>
  );
}

export default InternalRoutes;

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
import useWorkspace from 'helpers/hooks/useWorkspace';

import Dashboard from 'pages/Overview/Dashboard';
import Repositories from 'pages/Overview/Repositories';
import Vulnerabilities from 'pages/Overview/Vulnerabilities';
import Webhooks from 'pages/Overview/Webhooks';
import AddWorkspace from 'pages/Overview/AddWorkspace';
import Workspaces from 'pages/Overview/Workspaces';
import RepositoryTokens from './Repositories/Tokens';
import RepositoryInvite from './Repositories/Invite';
import WorkspaceTokens from './Workspaces/Tokens';
import WorkspaceUsers from './Workspaces/Users';

function InternalRoutes() {
  const { path } = useRouteMatch();
  const { isAdminOfWorkspace } = useWorkspace();

  return (
    <InternalLayout>
      <Switch>
        <PrivateRoute
          exact
          path={`${path}/add-workspace`}
          component={() => <AddWorkspace />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspaces`}
          component={() => <Workspaces />}
        />

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
          from={`${path}/dashboard`}
          to={
            isAdminOfWorkspace
              ? `${path}/dashboard/workspace`
              : `${path}/dashboard/repositories`
          }
        />

        <PrivateRoute
          path={`${path}/dashboard/workspace`}
          exact
          component={() => <Dashboard type="workspace" />}
        />

        <PrivateRoute
          path={`${path}/dashboard/repositories`}
          exact
          component={() => <Dashboard type="repository" />}
        />

        <PrivateRoute
          exact
          path={`${path}/vulnerabilities`}
          component={() => <Vulnerabilities />}
        />

        <PrivateRoute
          exact
          path={`${path}/repositories`}
          component={() => <Repositories />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspaces/:workspaceId/repositories/:repositoryId/invite`}
          component={() => <RepositoryInvite />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspaces/:workspaceId/repositories/:repositoryId/tokens`}
          component={() => <RepositoryTokens />}
        />

        <PrivateRoute
          exact
          path={`${path}/workspaces/:workspaceId/webhooks`}
          component={() => <Webhooks />}
        />
      </Switch>
    </InternalLayout>
  );
}

export default InternalRoutes;

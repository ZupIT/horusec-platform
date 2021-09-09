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

import React, { useState, useEffect } from 'react';
import { Workspace } from 'helpers/interfaces/Workspace';
import { RouteParams } from 'helpers/interfaces/RouteParams';
import {
  matchPath,
  useHistory,
  useParams,
  useRouteMatch,
} from 'react-router-dom';
import coreService from 'services/core';

interface WorkspaceCtx {
  workspaceId: string;
  workspace: Workspace;
  getWorkspace: () => Promise<Workspace>;
}

const WorkspaceContext = React.createContext<WorkspaceCtx>({
  workspaceId: null,
  workspace: null,
  getWorkspace: () => null,
});

const WorkspaceProvider = ({ children }: { children: JSX.Element }) => {
  const history = useHistory();
  const { path } = useRouteMatch();
  const params = useParams<RouteParams>();

  const { workspaceId = '' } = matchPath<{ workspaceId: string }>(
    history.location.pathname,
    {
      path: `${path}/workspace/:workspaceId`,
    }
  )?.params || { workspaceId: params.workspaceId } || { workspaceId: '' };

  const [workspace, setWorkspace] = useState<Workspace>();

  async function getWorkspace(id = workspaceId) {
    try {
      const { data } = await coreService.getOneWorkspace(id);
      setWorkspace(data.content);
      return data.content as Workspace;
    } catch (error) {
      history.push('/home');
    }
  }

  useEffect(() => {
    if (workspaceId) getWorkspace();
    // eslint-disable-next-line
  }, [workspaceId]);

  return (
    <WorkspaceContext.Provider
      value={{
        workspaceId,
        workspace,
        getWorkspace,
      }}
    >
      {children}
    </WorkspaceContext.Provider>
  );
};

export { WorkspaceContext, WorkspaceProvider };

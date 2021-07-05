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

import { Repository } from 'helpers/interfaces/Repository';
import { Workspace } from 'helpers/interfaces/Workspace';
import { useEffect, useState } from 'react';
import { matchPath, useHistory, useRouteMatch } from 'react-router-dom';
import core from 'services/core';

const useParamsRoute = () => {
  const history = useHistory();
  const { path } = useRouteMatch();

  const [workspace, setWorkspace] = useState<Workspace>({} as Workspace);
  const [repository, setRepository] = useState<Repository>({} as Repository);

  const { workspaceId = '' } = matchPath<{ workspaceId: string }>(
    history.location.pathname,
    {
      path: `${path}/workspace/:workspaceId`,
    }
  )?.params || { workspaceId: '' };

  const { repositoryId = '' } = matchPath<{ repositoryId: string }>(
    history.location.pathname,
    {
      path: `${path}/workspace/:workspaceId/repository/:repositoryId`,
    }
  )?.params || { repositoryId: '' };

  async function getWorkspace(workspace = workspaceId) {
    try {
      const { data } = await core.getOneWorkspace(workspace);
      setWorkspace(data.content);
      return data.content;
    } catch (error) {
      history.push('/home');
    }
  }

  async function getRepository(
    workspace = workspaceId,
    repository = repositoryId
  ) {
    try {
      const { data } = await core.getOneRepository(workspace, repository);
      setRepository(data.content);
      return data.content;
    } catch (error) {
      history.push('/home');
    }
  }

  useEffect(() => {
    let isCancelled = false;
    if (!isCancelled) {
      if (workspaceId) getWorkspace();
      if (workspaceId && repositoryId) getRepository();
    }
    return () => {
      isCancelled = true;
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [workspaceId, repositoryId]);

  return {
    workspaceId,
    repositoryId,
    workspace,
    repository,
    getWorkspace,
    getRepository,
  };
};

export default useParamsRoute;

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
      path: `${path}/workspaces/:workspaceId`,
    }
  )?.params || { workspaceId: '' };

  const { repositoryId = '' } = matchPath<{ repositoryId: string }>(
    history.location.pathname,
    {
      path: `${path}/workspaces/:workspaceId/repositories/:repositoryId`,
    }
  )?.params || { repositoryId: '' };

  useEffect(() => {
    function getWorkspace() {
      core.getOneWorkspace(workspaceId).then((result) => {
        setWorkspace(result.data.content);
      });
    }

    function getRepository() {
      core.getOneRepository(workspaceId, repositoryId).then((result) => {
        setRepository(result.data.content);
      });
    }
    if (workspaceId) getWorkspace();
    if (workspaceId && repositoryId) getRepository();
  }, [workspaceId, repositoryId]);

  return { workspaceId, repositoryId, workspace, repository };
};

export default useParamsRoute;

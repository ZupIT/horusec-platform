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

import React, { useState, useEffect } from 'react';
import { Repository } from 'helpers/interfaces/Repository';
import { RouteParams } from 'helpers/interfaces/RouteParams';
import {
  matchPath,
  useHistory,
  useParams,
  useRouteMatch,
} from 'react-router-dom';
import coreService from 'services/core';

interface RepositoryCtx {
  repositoryId: string;
  repository: Repository;
  getRepository: (
    workspaceId?: string,
    repositoryId?: string
  ) => Promise<Repository>;
}

const RepositoryContext = React.createContext<RepositoryCtx>({
  repositoryId: null,
  repository: null,
  getRepository: () => null,
});

const RepositoryProvider = ({ children }: { children: JSX.Element }) => {
  const history = useHistory();
  const { path } = useRouteMatch();
  const params = useParams<RouteParams>();

  const { repositoryId = '' } = matchPath<{ repositoryId: string }>(
    history.location.pathname,
    {
      path: `${path}/workspace/:workspaceId/repository/:repositoryId`,
    }
  )?.params || { repositoryId: params.repositoryId } || { repositoryId: '' };

  const [repository, setRepository] = useState();

  async function getRepository(
    workspace = params.workspaceId,
    repository = repositoryId
  ) {
    try {
      const { data } = await coreService.getOneRepository(
        workspace,
        repository
      );
      setRepository(data.content);
      return data.content;
    } catch (error) {
      history.push('/home');
    }
  }

  useEffect(() => {
    if (repositoryId) getRepository();
    // eslint-disable-next-line
  }, [repositoryId]);

  return (
    <RepositoryContext.Provider
      value={{
        repositoryId,
        repository,
        getRepository,
      }}
    >
      {children}
    </RepositoryContext.Provider>
  );
};

export { RepositoryContext, RepositoryProvider };

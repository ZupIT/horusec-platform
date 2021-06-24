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
import coreService from 'services/core';
import { Repository } from 'helpers/interfaces/Repository';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import useWorkspace from 'helpers/hooks/useWorkspace';
import { isString } from 'lodash';

interface RepositoryCtxInterface {
  currentRepository: Repository;
  allRepositories: Repository[];
  isMemberOfRepository: boolean;
  setAllRepositories: (repositories: Repository[]) => void;
  setCurrentRepository: (repository: string | Repository) => void;
  getOneRepository: (repositoryID: string) => void;
  fetchAllRepositories: () => void;
}

const RepositoryContext = React.createContext<RepositoryCtxInterface>({
  currentRepository: null,
  allRepositories: [],
  setAllRepositories: null,
  isMemberOfRepository: false,
  setCurrentRepository: null,
  fetchAllRepositories: null,
  getOneRepository: null,
});

const RepositoryProvider = ({ children }: { children: JSX.Element }) => {
  const [currentRepository, setRepository] = useState<Repository>(null);
  const [allRepositories, setAllRepositories] = useState<Repository[]>([]);
  const [isMemberOfRepository, setIsMemberOfRepository] =
    useState<boolean>(false);

  const { dispatchMessage } = useResponseMessage();
  const { currentWorkspace } = useWorkspace();

  const setCurrentRepository = (repository: string | Repository) => {
    if (isString(repository)) {
      const repo = allRepositories.find(
        (item) => item.repositoryID === repository
      );

      setRepository(repo);
      setIsMemberOfRepository(repo.role === 'member');
    } else {
      setRepository(repository);
      setIsMemberOfRepository(repository?.role === 'member' || false);
    }
  };

  const fetchAllRepositories = () => {
    if (currentWorkspace?.workspaceID) {
      coreService
        .getAllRepositories(currentWorkspace?.workspaceID)
        .then((result) => {
          const repositories = (result?.data?.content as Repository[]) || [];
          setAllRepositories(repositories);

          const hasCurrent = repositories.some(
            (repo) => repo.repositoryID === currentRepository?.repositoryID
          );

          if (hasCurrent) {
            setCurrentRepository(currentRepository);
          } else {
            repositories.length > 0
              ? setCurrentRepository(repositories[0])
              : setCurrentRepository(null);
          }
        })
        .catch((err) => {
          dispatchMessage(err?.response?.data);
          setCurrentRepository(null);
        });
    }
  };

  const getOneRepository = (repositoryID: string) => {
    if (currentWorkspace?.workspaceID) {
      coreService
        .getOneRepository(currentWorkspace?.workspaceID, repositoryID)
        .then((result) => {
          const repository = result?.data?.content as Repository;
          setCurrentRepository(repository);
        })
        .catch((err) => {
          dispatchMessage(err?.response?.data);
          setCurrentRepository(null);
        });
    }
  };

  useEffect(() => {
    fetchAllRepositories();
    // eslint-disable-next-line
  }, [currentWorkspace]);

  return (
    <RepositoryContext.Provider
      value={{
        currentRepository,
        allRepositories,
        getOneRepository,
        setCurrentRepository,
        fetchAllRepositories,
        isMemberOfRepository,
        setAllRepositories,
      }}
    >
      {children}
    </RepositoryContext.Provider>
  );
};

export { RepositoryProvider, RepositoryContext };

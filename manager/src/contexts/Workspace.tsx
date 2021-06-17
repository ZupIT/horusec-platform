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
import { Workspace } from 'helpers/interfaces/Workspace';
import useResponseMessage from 'helpers/hooks/useResponseMessage';
import { useHistory } from 'react-router-dom';
import { roles } from 'helpers/enums/roles';
import { isLogged } from 'helpers/localStorage/tokens';
import { getCurrentUser } from 'helpers/localStorage/currentUser';
import {
  getFavoriteWorkspace,
  setFavoriteWorkspace,
} from 'helpers/localStorage/favorite';

interface WorkspaceCtx {
  currentWorkspace: Workspace;
  favoriteWorkspace: Workspace;
  allWorkspaces: Workspace[];
  isAdminOfWorkspace: boolean;
  handleSetCurrentWorkspace: (workspace: Workspace) => void;
  fetchAllWorkspaces: () => void;
  setAsFavoriteWorkspace: (workspace: Workspace) => void;
}

const WorkspaceContext = React.createContext<WorkspaceCtx>({
  currentWorkspace: null,
  isAdminOfWorkspace: false,
  allWorkspaces: [],
  handleSetCurrentWorkspace: null,
  fetchAllWorkspaces: null,
  setAsFavoriteWorkspace: null,
  favoriteWorkspace: null,
});

const WorkspaceProvider = ({ children }: { children: JSX.Element }) => {
  const [currentWorkspace, setCurrentWorkspace] = useState<Workspace>(null);
  const [favoriteWorkspace, setFavorite] = useState<Workspace>(null);
  const [allWorkspaces, setAllWorkspaces] = useState<Workspace[]>([]);
  const [isAdminOfWorkspace, setIsAdminOfWorkspace] = useState<boolean>(false);

  const { dispatchMessage } = useResponseMessage();
  const history = useHistory();

  const setAsFavoriteWorkspace = (workspace: Workspace) => {
    setFavoriteWorkspace(workspace.workspaceID);
    setFavorite(workspace);
  };

  const handleSetCurrentWorkspace = (workspace: Workspace) => {
    if (workspace) {
      const currentUser = getCurrentUser();
      workspace = {
        ...workspace,
        role: currentUser.isApplicationAdmin ? roles.ADMIN : workspace.role,
      };
      setCurrentWorkspace(workspace);

      const isAdmin = workspace?.role === roles.ADMIN;
      setIsAdminOfWorkspace(isAdmin);
    } else {
      setCurrentWorkspace(null);
      setIsAdminOfWorkspace(false);
    }
  };

  const fetchAll = () => {
    coreService
      .getAllWorkspaces()
      .then((result) => {
        const workspaces = (result?.data?.content as Workspace[]) || [];

        setAllWorkspaces(workspaces);

        if (workspaces && workspaces.length > 0) {
          let favorite = getFavoriteWorkspace();
          favorite = workspaces.find((item) => item.workspaceID === favorite);

          if (favorite) {
            handleSetCurrentWorkspace(favorite);
            setFavorite(favorite);
          } else {
            handleSetCurrentWorkspace(workspaces[0]);
          }
        } else {
          handleSetCurrentWorkspace(null);
        }
      })
      .catch((err) => {
        dispatchMessage(err?.response?.data);
      });
  };

  useEffect(() => {
    if (isLogged()) fetchAll();
    // eslint-disable-next-line
  }, []);

  return (
    <WorkspaceContext.Provider
      value={{
        currentWorkspace,
        allWorkspaces,
        isAdminOfWorkspace,
        handleSetCurrentWorkspace,
        fetchAllWorkspaces: fetchAll,
        favoriteWorkspace,
        setAsFavoriteWorkspace,
      }}
    >
      {children}
    </WorkspaceContext.Provider>
  );
};

export { WorkspaceProvider, WorkspaceContext };

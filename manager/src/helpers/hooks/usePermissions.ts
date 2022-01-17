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
import { Repository } from 'helpers/interfaces/Repository';
import { Workspace } from 'helpers/interfaces/Workspace';
import { isLogged } from 'helpers/localStorage/tokens';
import { get } from 'lodash';
import { useHistory } from 'react-router-dom';
import useParamsRoute from './useParamsRoute';

const usePermissions = () => {
  const history = useHistory();
  const { repository, workspace, repositoryId, workspaceId } = useParamsRoute();

  const contexts = {
    REPOSITORY: repository,
    WORKSPACE: workspace,
  };

  enum ROLES {
    APP_ADMIN = 'applicationAdmin',
    ADMIN = 'admin',
    MEMBER = 'member',
    SUPERVISOR = 'supervisor',
  }

  enum ACTIONS {
    CREATE_WORKSPACE = 'CREATE_WORKSPACE',
    VIEW_WORKSPACE = 'VIEW_WORKSPACE',
    HANDLE_WORKSPACE = 'HANDLE_WORKSPACE',
    HANDLE_VULNERABILITIES_WORKSPACE = 'HANDLE.VULNERABILITIES_WORKSPACE',
    CREATE_REPOSITORY = 'CREATE.REPOSITORY_WORKSPACE',
    VIEW_REPOSITORY = 'VIEW_REPOSITORY',
    HANDLE_REPOSITORY = 'HANDLE_REPOSITORY',
    HANDLE_VULNERABILITIES_REPOSITORY = 'HANDLE.VULNERABILITIES_REPOSITORY',
  }

  const actionsPermissions = {
    [ACTIONS.CREATE_WORKSPACE]: Object.values(ROLES),
    [ACTIONS.VIEW_WORKSPACE]: [ROLES.ADMIN, ROLES.APP_ADMIN],
    [ACTIONS.HANDLE_WORKSPACE]: [ROLES.ADMIN, ROLES.APP_ADMIN],
    [ACTIONS.HANDLE_VULNERABILITIES_WORKSPACE]: [ROLES.ADMIN, ROLES.APP_ADMIN],
    [ACTIONS.CREATE_REPOSITORY]: [ROLES.ADMIN, ROLES.APP_ADMIN],
    [ACTIONS.VIEW_REPOSITORY]: Object.values(ROLES),
    [ACTIONS.HANDLE_REPOSITORY]: [ROLES.ADMIN, ROLES.APP_ADMIN],
    [ACTIONS.HANDLE_VULNERABILITIES_REPOSITORY]: [
      ROLES.ADMIN,
      ROLES.APP_ADMIN,
      ROLES.SUPERVISOR,
    ],
  };

  const getPermissionsByRoute = (pathname?: string) => {
    const route = pathname || history.location.pathname;

    const routesPermissions = {
      dashboard: Object.values(ROLES),
      vulnerabilities: Object.values(ROLES),
      tokens: [ROLES.ADMIN, ROLES.APP_ADMIN],
      users: [ROLES.ADMIN, ROLES.APP_ADMIN],
      webhooks: [ROLES.ADMIN, ROLES.APP_ADMIN],
      new: Object.values(ROLES),
    };

    const routeMatched = Object.entries(routesPermissions).find((item) =>
      route.endsWith(item[0])
    );

    if (routeMatched) return routeMatched[1];
    else return null;
  };

  const isOverviewScreen = () => {
    return history.location.pathname.includes('overview');
  };

  const isAuthorizedAction = (
    action: ACTIONS,
    contextItem?: Repository | Workspace
  ): boolean => {
    const currentContext =
      contextItem || get(contexts, action.split('_')[1], null);

    if (!currentContext?.role) return true;
    else return actionsPermissions[action].includes(currentContext.role);
  };

  const isAuthorizedRoute = (pathname?: string) => {
    const allowedPermissions = getPermissionsByRoute(pathname);
    const isOverviewOfRepository = !!repositoryId;
    const isOverviewOfWorkspace =
      isOverviewScreen() && workspaceId && !repositoryId;

    if (!isLogged()) return false;

    if (isOverviewOfWorkspace && workspace?.role === ROLES.MEMBER) {
      return false;
    }

    if (isOverviewOfRepository && repository?.role) {
      return allowedPermissions.includes(repository?.role as ROLES);
    }

    return true;
  };

  return {
    ROLES,
    ACTIONS,
    isAuthorizedAction,
    isAuthorizedRoute,
  };
};

export default usePermissions;

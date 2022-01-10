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

import { keycloakInstance } from 'config/keycloak';
import { clearCurrentUser } from 'helpers/localStorage/currentUser';
import { clearTokens } from 'helpers/localStorage/tokens';
import { MANAGER_BASE_PATH } from 'config/basePath';
import { clearCurrentPath } from 'helpers/localStorage/currentPage';

const redirectUri = `${window.location.origin}${MANAGER_BASE_PATH}auth`;

const login = () => keycloakInstance.login({ redirectUri });

const logout = () => {
  clearCurrentUser();
  clearTokens();
  clearCurrentPath();

  return keycloakInstance.logout({ redirectUri });
};

export default {
  login,
  logout,
};

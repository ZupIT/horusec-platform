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

import { localStorageKeys } from 'helpers/enums/localStorageKeys';
import { content } from '../fixtures/login/horusec/success.json';

const LOCAL_STORAGE_MEMORY = {};

Cypress.Commands.add('saveLocalStorage', () => {
  Object.keys(localStorage).forEach((key) => {
    LOCAL_STORAGE_MEMORY[key] = localStorage[key];
  });
});

Cypress.Commands.add('restoreLocalStorage', () => {
  Object.keys(LOCAL_STORAGE_MEMORY).forEach((key) => {
    localStorage.setItem(key, LOCAL_STORAGE_MEMORY[key]);
  });
});

Cypress.Commands.add('authenticated', () => {
  localStorage.setItem(localStorageKeys.ACCESS_TOKEN, content.accessToken);
  localStorage.setItem(localStorageKeys.REFRESH_TOKEN, content.refreshToken);
  localStorage.setItem(localStorageKeys.TOKEN_EXPIRES, content.expiresAt);
  localStorage.setItem(
    localStorageKeys.USER,
    JSON.stringify({
      username: content.username,
      email: content.email,
      isApplicationAdmin: content.isApplicationAdmin,
    })
  );
});

Cypress.Commands.add('setHorusecAuthConfig', () => {
  cy.intercept(
    {
      method: 'GET',
      url: 'auth/authenticate/config',
    },
    { fixture: 'login/horusec/auth-config', statusCode: 200 }
  );
});

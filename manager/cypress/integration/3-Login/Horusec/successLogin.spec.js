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
import { NEW_WORKSPACE_SCREEN } from '../../../../src/config/i18n/enUS.json';

/* eslint-disable cypress/no-unnecessary-waiting */
describe('Login in the application when a correct username and password.', () => {
  beforeEach(() => {
    cy.restoreLocalStorage();

    cy.setHorusecAuthConfig();

    cy.intercept(
      {
        method: 'POST',
        url: 'auth/authenticate/login',
      },
      { fixture: 'login/horusec/success', statusCode: 200 }
    ).as('authenticate');

    cy.intercept(
      {
        method: 'GET',
        url: 'core/workspaces',
      },
      { fixture: 'workspaces/empty', statusCode: 200 }
    ).as('getWorkspaces');
  });

  afterEach(() => {
    cy.saveLocalStorage();
  });

  it('Fill login form and login.', () => {
    cy.visit('/');
    cy.wait(4200);

    cy.get('#email').click().type('admin@horusec.com');
    cy.get('#password').click().type('secret');
    cy.get('#submit-login').click();

    cy.wait('@authenticate');

    cy.wait('@getWorkspaces');
  });
});

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

import { API_ERRORS } from '../../../src/config/i18n/enUS.json';

/* eslint-disable cypress/no-unnecessary-waiting */
describe('Validation the field of login create account form.', () => {
  beforeEach(() => {
    cy.setHorusecAuthConfig();

    cy.intercept(
      {
        method: 'POST',
        url: 'auth/account/verify-already-used',
      },
      { fixture: 'createAccount/verifyAlreadyUsed/error', statusCode: 400 }
    ).as('verifyAlreadyUsed');
  });

  it('Go to register screen', () => {
    cy.visit('/');
    cy.wait(4200);

    cy.get('#create-account').should('be.visible');
    cy.get('#create-account').click();
  });

  it('Fill the fields and submit', () => {
    cy.get('#username').click().type('testing');
    cy.get('#email').click().type('test@test.com');

    cy.get('#next-step').click();

    cy.wait('@verifyAlreadyUsed');
  });

  it('Check the flash message error', () => {
    cy.get('#flash-message-error').should('be.visible');
    cy.get('#flash-message-error').should(
      'contain.text',
      API_ERRORS.EMAIL_IN_USE
    );
  });
});

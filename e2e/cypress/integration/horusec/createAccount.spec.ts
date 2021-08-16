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
import * as user from '../../fixtures/user.json';

describe('Create account flow', () => {
  before(() => {
    cy.visitHorusecManager();
  });

  it('Render Horusec manager login page', () => {
    cy.get('#create-account').should('be.visible');
  });

  it('Go to create account screen', () => {
    cy.get('#create-account').click();
    cy.get('#email').should('be.visible');
  });

  it('Fill username and email', () => {
    cy.intercept({
      method: 'POST',
      url: 'auth/account/verify-already-used',
    }).as('verifyAlreadyUsed');

    cy.get('#username').click().type(user.username);
    cy.get('#email').click().type(user.email);

    cy.get('#next-step').click();

    cy.wait('@verifyAlreadyUsed').should((xhr) => {
      expect(xhr.response.statusCode).to.equal(204);
    });

    cy.location().should((location) => {
      expect(location.pathname).to.eq('/auth/create-account');
    });
  });

  it('Go to create account screen', () => {
    cy.intercept({
      method: 'POST',
      url: 'auth/account/verify-already-used',
    }).as('c');

    cy.get('#password').click().type(user.password);
    cy.get('#confirm-pass').click().type(user.password);
    cy.get('#register').click();
  });
});

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

import ptBR from '../../../src/config/i18n/ptBR.json';
import enUS from '../../../src/config/i18n/enUS.json';

/* eslint-disable cypress/no-unnecessary-waiting */
describe('In Login screen, alter the language to PTBR and back to ENUS', () => {
  beforeEach(() => {
    cy.restoreLocalStorage();
    cy.setHorusecAuthConfig();
  });

  afterEach(() => {
    cy.saveLocalStorage();
  });

  it('Change language to ptBR', () => {
    cy.visit('/');

    cy.wait(4200);

    cy.get('#submit-login').should('contain.text', enUS.LOGIN_SCREEN.SUBMIT);

    cy.get('#language').click();
    cy.get('#ptBR').click();

    cy.get('#submit-login').should('contain.text', ptBR.LOGIN_SCREEN.SUBMIT);

    cy.get('#language').click();
    cy.get('#enUS').click();

    cy.get('#submit-login').should('contain.text', enUS.LOGIN_SCREEN.SUBMIT);
  });
});

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

/* eslint-disable no-unused-vars */
/* eslint-disable no-use-before-define */
/* eslint-disable cypress/no-unnecessary-waiting */

declare namespace Cypress {
  interface Chainable<Subject> {
    dockerComposeUpHorusecDefaultAuth: typeof dockerComposeUpHorusecDefaultAuth;
    visitHorusecManager: typeof visitHorusecManager;
  }
}

function dockerComposeUpHorusecDefaultAuth(): Cypress.Chainable<Cypress.Exec> {
  return cy.exec('docker-compose -f ./deployments/docker-compose.horusec-default.yaml up -d --build --force-recreate', {
    timeout: 1800000,
    log: true,
    failOnNonZeroExit: false,
  });
}

function visitHorusecManager(): void {
  cy.visit('http://127.0.0.1:8043');
  cy.wait(5000);
}

Cypress.Commands.add('dockerComposeUpHorusecDefaultAuth', dockerComposeUpHorusecDefaultAuth);
Cypress.Commands.add('visitHorusecManager', visitHorusecManager);

import enUS from '../../../src/config/i18n/enUS.json';

/* eslint-disable cypress/no-unnecessary-waiting */
describe('On the home screen, create a new repository', () => {
  beforeEach(() => {
    cy.setHorusecAuthConfig();
    cy.restoreLocalStorage();

    cy.intercept(
      {
        method: 'POST',
        url: 'core/workspaces',
      },
      { fixture: 'workspaces/admin', statusCode: 200 }
    ).as('createWorkspace');
  });

  it('Go to home screen and check if add workspace', () => {
    cy.visit('/home').wait(3000);
    cy.get('#addWorkspace').click();
  });

  it('Check rendering fields.', () => {
    cy.get('#name').should('be.visible');
    cy.get('#description').should('be.visible');
  });

  it('Check if show error message of empty name', () => {
    cy.get('#name').click().type('test').clear().blur();
    cy.get('#name-error').should(
      'contain.text',
      enUS.WORKSPACES_SCREEN.INVALID_WORKSPACE_NAME
    );
  });

  it('Fill in valid data and send form', () => {
    cy.get('#name').clear().click().type('Zup IT');
    cy.get('#description').clear().click().type('Zup It Organization');
    cy.get('#submit-workspace').submit();

    cy.wait('@createWorkspace');
  });
});

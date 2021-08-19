/* eslint-disable cypress/no-unnecessary-waiting */
describe('On the home screen, select a workspace', () => {
  beforeEach(() => {
    cy.setHorusecAuthConfig();
    cy.restoreLocalStorage();
    cy.authenticated();
    cy.intercept(
      {
        method: 'GET',
        url: 'core/workspaces',
      },
      { fixture: 'workspaces/admin', statusCode: 200 }
    ).as('getWorkspaces');

    cy.intercept(
      {
        method: 'GET',
        url: 'core/workspaces/0f4de453-305f-4efb-8d87-879103738efa',
      },
      { fixture: 'workspaces/admin', statusCode: 200 }
    ).as('getWorkspace');

    cy.intercept(
      {
        method: 'PUT',
        url: 'core/workspaces/0f4de453-305f-4efb-8d87-879103738efa',
      },
      { fixture: 'workspaces/admin', statusCode: 200 }
    );

    cy.intercept(
      {
        method: 'GET',
        url: 'analytic/dashboard/0f4de453-305f-4efb-8d87-879103738efa',
      },
      { fixture: 'workspaces/dashboard', statusCode: 200 }
    ).as('getDashboardData');
  });

  it('Go to home screen and select a workspace', () => {
    cy.visit('/home');
    cy.get('ul > li > footer > button').click({ force: true });
  });

  it('Fill in valid data and edit data', () => {
    cy.visit('home/workspace/0f4de453-305f-4efb-8d87-879103738efa').wait(1000);
    cy.get('#handler').click();
    cy.get('#name').clear().click().type('Zup IT Edited');
    cy.get('#description').clear().click().type('Zup It Organization Edited');
    cy.get('#submit-workspace').submit();
  });
});

# Horusec Platform e2e tests

This module will bring together all the **Horusec Platform** e2e tests using the [**cypress**](https://docs.cypress.io) tool.

## Structure

- **Fixture folder** - This folder stores all static data that will be used as a *mock*
- **Integration folder** - Here all test cases will be gathered
- **Plugins folder** - The Plugins API allows you to hook into and extend Cypress behavior. [read more about plugins](https://docs.cypress.io/api/plugins/writing-a-plugin)
- **Support folder** - This files runs before every single spec file. We do this purely as a convenience mechanism so you don't have to import this file in every single one of your spec files. [read more about support files](https://docs.cypress.io/guides/core-concepts/writing-and-organizing-tests#Support-file)

## Guide Lines

This project has guaranteed code standardization through [typescript](https://www.typescriptlang.org/), [eslint](https://eslint.org) and the official [cypress eslint plugin](https://github.com/cypress-io/eslint-plugin-cypress).

---
This project exists thanks to all the contributors. You rock! ❤️🚀

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

import { Requests } from "../../utils/request";
import AnalysisMock from "../../mocks/analysis.json";
import { v4 as uuidv4 } from "uuid";

describe("Horusec tests", () => {
    before(() => {
        cy.exec("make migrate-horusec-postgresql", {log: true}).its("code").should("eq", 0);
    });
    it("Should test all operations horusec", () => {
        // cy.visit("http://localhost:8043")
        // cy.wait(4000);
        // cy.get("#email").type("dev@example.com");
        // cy.get("#password").type("Devpass0*");
        // cy.get("button").contains("Sign in").click();
        // cy.wait(2000);
        // cy.get("li").contains("Company e2e").parent().get("button").contains("Select").click({ force: true });
        // cy.wait(2000);
        // cy.get("li").contains("Core-API").parent().get("button").contains("Overview").click({ force: true });
        // cy.wait(2000);

        LoginWithDefaultAccountAndCheckIfNotExistWorkspace();
        CreateEditDeleteAnWorkspace();
        CreateWorkspaceCheckIfDashboardIsEmpty();
        CreateDeleteWorkspaceTokenAndSendFirstAnalysisMock();
        CheckIfDashboardNotIsEmpty();
        CreateEditDeleteAnRepository();
        CreateRepository();
        CreateDeleteRepositoryTokenAndSendFirstAnalysisMock();
        CheckIfDashboardNotIsEmptyWithTwoRepositories();
        CheckIfExistsVulnerabilitiesAndCanUpdateSeverityAndStatus();
        CreateUserAndInviteToExistingWorkspace();
        CheckIfPermissionsIsEnableToWorkspaceMember();
        InviteUserToRepositoryAndCheckPermissions();
        LoginAndUpdateDeleteAccount();
    });
});

function LoginWithDefaultAccountAndCheckIfNotExistWorkspace(): void {
    cy.visit("http://localhost:8043/auth");
    cy.wait(4000);

    // Login with default account
    cy.get("#email").type("dev@example.com");
    cy.get("#password").type("Devpass0*");
    cy.get("button").contains("Sign in").click();
    cy.wait(1000);

    // Check if not exists workspace
    cy.contains("Add your first workspace to start ...").should("exist");
}

function CreateEditDeleteAnWorkspace(): void {
    cy.wait(1500);

    // Create an workspace
    cy.get("button").contains("Add Workspace").click();
    cy.wait(500);
    cy.get("#name").type("first_workspace");
    cy.get("button").contains("Save").click();

    // Check if exist on list
    cy.contains("first_workspace").should("exist");
    cy.wait(500);

    // Edit existing workspace
    cy.get("li").contains("first_workspace").parent().get("button").contains("Select").click({ force: true });
    cy.wait(1000);
    cy.get("button").contains("Handler").click();
    cy.wait(500);
    cy.get("#name").type("_edited");
    cy.get("button").contains("Save").click();
    cy.contains("first_workspace_edited").should("exist");
    cy.get("li").contains("Home").click();

    // Check if was updated on list
    cy.contains("first_workspace_edited").should("exist");
    cy.wait(500);

    // Delete existing workspace
    cy.get("li").contains("first_workspace_edited").parent().get("button").contains("Select").click({ force: true });
    cy.get("button").contains("Handler").click();
    cy.wait(500);
    cy.get("button").contains("Delete").click();
    cy.wait(500);
    cy.get("button").contains("Yes").click();
    cy.wait(1000);

    // Check if was removed on list
    cy.contains("first_workspace_edited").should("not.exist");
}

function CreateWorkspaceCheckIfDashboardIsEmpty(): void {
    cy.get("button").contains("Add Workspace").click();
    cy.wait(500);
    cy.get("#name").type("Company e2e");
    cy.get("button").contains("Save").click();
    cy.wait(1500);

    // Check if exist new workspace on list and select workspace
    cy.get("li").contains("Company e2e").parent().get("button").contains("Select").click({ force: true });
    cy.wait(1000);

    cy.get("button").contains("Overview").click();
    cy.wait(1000);
    cy.get("button").contains("Apply").click();
    cy.wait(2000);
    cy.get("#total-developers").contains("No results were found").should("exist");
}

function CreateDeleteWorkspaceTokenAndSendFirstAnalysisMock(): void {
    // Go to manage tokens of workspace page
    cy.get("li").contains("Tokens").click();
    cy.wait(500);

    // Disable alert when copy data to clipboard
    cy.window().then(win => {
        cy.stub(win, "prompt").returns("DISABLED WINDOW PROMPT");
    });

    // Create access token
    cy.get("button").contains("Add token").click();
    cy.wait(500);
    cy.get("#description").type("Access Token");
    cy.get("button").contains("To save").click();

    // Copy acceess token to clipboard and create first analysis with this token
    cy.get("[data-testid=\"icon-copy\"").click();
    cy.get("h3").first().then((content) => {
        const _requests: Requests = new Requests();
        const body: any = AnalysisMock;
        const url: any = `${_requests.baseURL}${_requests.services.Api}/api/analysis`;
        _requests
            .setHeadersAllRequests({"X-Horusec-Authorization": content[0].innerText})
            .post(url, body)
            .then((response) => {
                expect(response.status).eq(201, "First Analysis of workspace created with sucess");
            })
            .catch((err) => {
                cy.log("Error on send analysis in token of workspace: ", err).end();
            });
    });
    cy.wait(3000);
    cy.get("button").contains("Ok, I got it.").click();

    // Check if exists access token on list of token
    cy.contains("Access Token").should("exist");
    cy.wait(1000);

    // Remove access token created
    cy.get("[class=\"MuiButtonBase-root MuiIconButton-root\"]").click();
    cy.get("button").contains("Delete").click();
    cy.wait(500);
    cy.get("button").contains("Yes").click();
    cy.wait(1000);

    // Check if not exists access token on list of token
    cy.contains("Access Token").should("not.exist");
}

function CheckIfDashboardNotIsEmpty(): void {
    // Go to dashboard page
    cy.get("li").contains("Dashboard").click();
    cy.wait(2000);

    // Search from begging data
    cy.get("button").contains("Apply").click();
    cy.wait(2000);

    checkDashboardInitialContent(true)

    // Go to Repositories page
    cy.get("li").contains("Repositories").click();
    cy.wait(2000);
    cy.get("li").contains("Register-API").parent().get("button").contains("Overview").click({ force: true });
    cy.wait(2000);

    checkDashboardInitialContent(false)
}

function checkDashboardInitialContent(isWorkspace: boolean) {
    // Check if chart of total developers exist 1 user in selected repository
    cy.get("#total-developers").contains("No results were found").should("not.exist");

    // Check if chart with all vulnerabilities of exists all vulnerabilities
    cy.get("#All_vulnerabilities_CRITICAL").contains("20").should("exist");
    cy.get("#All_vulnerabilities_HIGH").contains("39").should("exist");
    cy.get("#All_vulnerabilities_MEDIUM").contains("48").should("exist");
    cy.get("#All_vulnerabilities_LOW").contains("28").should("exist");
    cy.get("#All_vulnerabilities_INFO").contains("0").should("exist");
    cy.get("#All_vulnerabilities_UNKNOWN").contains("5").should("exist");

    // Check if chart exists total vulnerabilities in chart of vulnerabilities by developer
    cy.get("#Vulnerabilities_by_developer_-").contains("140").should("exist");

    // Check if chart of languages contains vulnerabilities
    cy.get("[id=\"Language_vulnerabilities_C#\"]").contains("13").should("exist");
    cy.get("#Language_vulnerabilities_C").contains("7").should("exist");
    cy.get("#Language_vulnerabilities_Go").contains("5").should("exist");
    cy.get("#Language_vulnerabilities_Dart").contains("3").should("exist");
    cy.get("#Language_vulnerabilities_Elixir").contains("3").should("exist");

    if (isWorkspace) {
        // Check if chart exists total vulnerabilities in chart of vulnerabilities by repository
        cy.get("#Vulnerabilities_by_repository_Register-API").contains("140").should("exist");

        // Check if chart of total developers exist 1 user
        cy.get("#total-repositories").contains("No results were found").should("not.exist");
    }
}

function CreateEditDeleteAnRepository(): void {
    // Go to repositories page
    cy.get("li").contains("Repositories").click();
    cy.wait(1000);

    // Create an repository
    cy.get("button").contains("Add Repository").click();
    cy.get("#name").type("first_repository");
    cy.get("button").contains("Save").click();

    // Check if this repository exists on list
    cy.contains("first_repository").should("exist");
    cy.wait(1500);

    // Edit the new repository
    cy.contains("first_repository").parent().contains("Handler").click({ force: true });
    cy.wait(500);
    cy.get("#name").type("_edited");
    cy.get("button").contains("Save").click();

    // Check if repository was edited with success
    cy.contains("first_repository_edited").should("exist");
    cy.wait(1500);

    // Delete the repository
    cy.contains("first_repository_edited").parent().contains("Handler").click({ force: true });
    cy.wait(500);
    cy.get("button").contains("Delete").click();
    cy.wait(500);
    cy.get("button").contains("Yes").click();
    cy.wait(1000);

    // Check if repository was deleted with success
    cy.contains("first_repository_edited").should("not.exist");
}

function CreateRepository(): void {
    cy.get("button").contains("Add Repository").click();
    cy.get("#name").type("Core-API");
    cy.get("button").contains("Save").click();
    cy.contains("Core-API").should("exist");

    cy.contains("Core-API").parent().contains("Overview").click({ force: true });
    cy.wait(3000);
}

function CreateDeleteRepositoryTokenAndSendFirstAnalysisMock(): void {
    // Get repository created and create new access token
    cy.get("li").contains("Tokens").click();
    cy.wait(500);

    cy.contains("Add token").should("exist");
    cy.get("button").contains("Add token").click();
    cy.wait(500);
    cy.get("#description").type("Access Token");
    cy.get("button").contains("To save").click();

    // Copy acceess token to clipboard and create first analysis with this token into repository
    cy.get("[data-testid=\"icon-copy\"").click();
    cy.get("h3").first().then((content) => {
        const _requests: Requests = new Requests();
        const body: any = AnalysisMock;
        body.repositoryName = "Core-API";
        body.analysis.id = uuidv4();
        body.analysis.repositoryID = "00000000-0000-0000-0000-000000000000"
        body.analysis.repositoryName = ""
        body.analysis.workspaceID = "00000000-0000-0000-0000-000000000000"
        body.analysis.workspaceName = ""
        body.analysis.analysisVulnerabilities = body.analysis.analysisVulnerabilities.map((i) => {
            i.analysisID = body.analysis.id
            i.vulnerabilityID = uuidv4();
            i.vulnerabilities.vulnerabilityID = i.vulnerabilityID;
            return i
        })
        const url: any = `${_requests.baseURL}${_requests.services.Api}/api/analysis`;
        _requests
            .setHeadersAllRequests({"X-Horusec-Authorization": content[0].innerText})
            .post(url, body)
            .then((response) => {
                expect(response.status).eq(201, "First Analysis of repository created with sucess");
            })
            .catch((err) => {
                cy.log("Error on send analysis in token of repository: ", err).end();
            });
    });
    cy.wait(3000);
    cy.get("button").contains("Ok, I got it.").click();

    // Check if access token exist on list of tokens
    cy.contains("Access Token").should("exist");
    cy.wait(1000);

    // Delete access token
    cy.get("[class=\"MuiButtonBase-root MuiIconButton-root\"]").click();
    cy.get("button").contains("Delete").click();
    cy.wait(500);
    cy.get("button").contains("Yes").click();

    // Check if access token was deleted
    cy.contains("Access Token").should("not.exist");
}

function CheckIfDashboardNotIsEmptyWithTwoRepositories(): void {
    // Go to dasboard page
    cy.get("li").contains("Dashboard").click();
    cy.wait(1000);

    cy.get("button").contains("Apply").click();
    cy.wait(1500);

    checkDashboardInitialContent(false)
}

function CheckIfExistsVulnerabilitiesAndCanUpdateSeverityAndStatus(): void {
    // Go to vulnerabilities page
    cy.get("li").contains("Vulnerabilities").click();
    cy.wait(1500);

    cy.get('.MuiTableBody-root > :nth-child(1) > :nth-child(3)').invoke('text')
    .then((firstHash)=>{ 
        console.log(firstHash)
        expect(firstHash).not.to.equal("");
        expect(firstHash).not.to.equal(undefined);
        expect(firstHash).not.to.equal(null);

        // Select first vulnerability and open Severity dropdown
        cy.get(":nth-child(1) > .center > .MuiFormControl-root > .MuiInputBase-root > #select").click();
        cy.wait(500);

        // Change severity to HIGH
        cy.get('[data-value="HIGH"]').click();

        // Select first vulnerability and open status dropdown
        cy.get(":nth-child(1) > :nth-child(2) > .MuiFormControl-root > .MuiInputBase-root > #select").click();

        // Change status to Risk Accepted
        cy.get('[data-value="Risk Accepted"]').click();
        cy.get("button").contains("Update Vulnerabilities").click();
        cy.wait(1500);

        // Open modal of vulnerability and check if details exists
        cy.get("[data-testid=\"icon-info\"]").first().click();
        cy.contains("Vulnerability Details").should("exist")
        cy.wait(1500);

        cy.get("[data-testid=\"icon-close\"]").first().click();

        cy.contains(firstHash).should("not.exist");
    })
}

function CreateUserAndInviteToExistingWorkspace(): void {
    // Go to home page
    // Logout user
    cy.get("li").contains("Logout").click();
    cy.wait(3000);

    // Create new account
    cy.get("button").contains("Don't have an account? Sign up").click();
    cy.get("#username").clear().type("e2e_user");
    cy.get("#email").clear().type("e2e_user@example.com");
    cy.get("button").contains("Next").click();
    cy.get("#password").clear().type("Ch@ng3m3");
    cy.get("#confirm-pass").clear().type("Ch@ng3m3");
    cy.get("button").contains("Register").click();

    // Check if account was created
    cy.contains("Your Horusec account has been created successfully!");
    cy.get("button").contains("Ok, I got it.").click();

    // Login with new account and check if not exists company and logout user
    cy.get("#email").type("e2e_user@example.com");
    cy.get("#password").type("Ch@ng3m3");
    cy.get("button").contains("Sign in").click();
    cy.wait(1500);

    // Check if not exists company to this account and logout user
    cy.contains("Add your first workspace to start ...").should("exist");
    cy.get("li").contains("Logout").click();
    cy.wait(3000);

    // Login with default account
    cy.get("#email").type("dev@example.com");
    cy.get("#password").type("Devpass0*");
    cy.get("button").contains("Sign in").click();
    cy.wait(1500);

    // Go to manage users workspace page
    cy.get("li").contains("Company e2e").parent().get("button").contains("Select").click({ force: true });
    cy.wait(1500);
    cy.get("button").contains("Overview").first().click();

    // Invite user
    cy.get("li").contains("Users").click();
    cy.get("button").contains("Invite").click();
    cy.get("#email").clear().type("e2e_user@example.com");
    cy.get("#select-role").click();
    cy.get("#select-role-option-1").click();
    cy.get("button").contains("Save").click();
}

function CheckIfPermissionsIsEnableToWorkspaceMember(): void {
    // Go to home page
    cy.get("li").contains("Logout").click();
    cy.wait(4000);

    // Login with new account
    cy.get("#email").type("e2e_user@example.com");
    cy.get("#password").type("Ch@ng3m3");
    cy.get("button").contains("Sign in").click();
    cy.wait(1500);

    // Check if exists company and user can create workspace
    cy.get("button").contains("Add Workspace").should("exist");
    cy.get("li").contains("Company e2e").parent().get("button").contains("Select").click({ force: true });
    cy.wait(1000);

    // Check if not exists button for edit workspace and view dashboard
    cy.contains("Add a new repository to proceed ...").should("exist");
    cy.get("#title-workspace-wrapper").contains("Overview").should("not.exist");
    cy.get("#title-workspace-wrapper").contains("Handler").should("not.exist");
    cy.get("button").contains("Add Repository").should("not.exist");
}

function InviteUserToRepositoryAndCheckPermissions(): void {
    // Go to home page
    cy.get("li").contains("Logout").click();
    cy.wait(4000);

    // Login with default user
    cy.get("#email").type("dev@example.com");
    cy.get("#password").type("Devpass0*");
    cy.get("button").contains("Sign in").click();
    cy.wait(1500);

    // Go to repositories page
    cy.get("li").contains("Company e2e").parent().get("button").contains("Select").click({ force: true });
    cy.wait(1000);
    cy.contains("Core-API").parent().contains("Overview").click({ force: true });
    cy.wait(2000);

    // Invite user to repository
    cy.get("li").contains("Users").click();
    cy.get("button").contains("Invite").click();
    cy.wait(500);
    cy.get("#select-user").click();
    cy.get("#select-user-option-0").click();
    cy.get("#select-role").click();
    cy.get("#select-role-option-1").click();
    cy.get("button").contains("Save").click();
    cy.wait(500);
    cy.contains("e2e_user@example.com").should("exist");

    // Logout user
    cy.get("li").contains("Logout").click();
    cy.wait(4000);

    // Login with new user
    cy.get("#email").type("e2e_user@example.com");
    cy.get("#password").type("Ch@ng3m3");
    cy.get("button").contains("Sign in").click();
    cy.wait(1500);

    // Check if dashboard show data to repository
    cy.get("li").contains("Company e2e").parent().get("button").contains("Select").click({ force: true });
    cy.wait(1000);
    cy.contains("Core-API").parent().contains("Overview").click({ force: true });
    cy.wait(2000);
    cy.get("#total-developers").contains("No results were found").should("not.exist");
    cy.get("li").contains("Repositories").click();

    // Check if user not contains permissions
    cy.contains("Core-API").parent().contains("Handler").should("not.exist");
}

function LoginAndUpdateDeleteAccount(): void {
    // Logout user
    cy.get("li").contains("Logout").click();
    cy.wait(4000);

    // Login with new account and go to settings page
    cy.get("#email").type("e2e_user@example.com");
    cy.get("#password").type("Ch@ng3m3");
    cy.get("button").contains("Sign in").click();
    cy.wait(1500);
    cy.get("li").contains("Settings").click();
    cy.wait(1000);

    // Open modal and edit user
    cy.get("button").contains("Edit information").click();
    cy.get("#username").clear().type("user_updated");
    cy.get("#email").clear().type("user_updated@example.com");
    cy.get("button").contains("Save").click();

    // Logout user
    cy.get("li").contains("Logout").click();
    cy.wait(4000);

    // Check if is enable login with new email
    cy.get("#email").type("user_updated@example.com");
    cy.get("#password").type("Ch@ng3m3");
    cy.get("button").contains("Sign in").click();
    cy.wait(1500);

    // Go to config page
    cy.get("li").contains("Settings").click();
    cy.wait(1000);

    // Change password of user
    cy.get("button").contains("Change Password").click();
    cy.get("#password").clear().type("Ch@ng3m3N0w");
    cy.get("#confirmPass").clear().type("Ch@ng3m3N0w");
    cy.get("button").contains("Save").click();

    // Logout user
    cy.get("li").contains("Logout").click();
    cy.wait(4000);

    // Check if is enable login with new password
    cy.get("#email").type("user_updated@example.com");
    cy.get("#password").type("Ch@ng3m3N0w");
    cy.get("button").contains("Sign in").click();
    cy.wait(1500);

    // When login in page check if exist "Version" o system
    cy.get("li").contains("Company e2e").parent().get("button").contains("Select").click({ force: true });
    cy.wait(1000);
    cy.contains("Core-API").parent().contains("Overview").click({ force: true });
    cy.wait(2000);
    cy.contains("Version v").should("exist");

    // Go to config page
    cy.get("li").contains("Settings").click();
    cy.wait(1000);

    // Delete account
    cy.get("button").contains("Delete").click();
    cy.get("button").contains("Yes").click();
    cy.wait(5000);

    // Check if account not exists
    cy.get("#email").type("user_updated@example.com");
    cy.get("#password").type("Ch@ng3m3N0w");
    cy.get("button").contains("Sign in").click();
    cy.wait(1500);

    // Check if login is not authorized
    cy.url().should("eq", "http://localhost:8043/auth");
}

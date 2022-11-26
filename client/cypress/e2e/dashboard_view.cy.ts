Cypress.env("MODE");

describe("Dashboard View Test Case", () => {
  it("visits the app root url", () => {
    cy.visit("/");
    cy.contains("h3", "International Space Station dashboard");
  });
});

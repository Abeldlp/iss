Cypress.env("MODE");

describe("Dashboard View Test Case", () => {
  it("can load dashboard page", () => {
    cy.visit("/");
    cy.contains("h3", "International Space Station dashboard");
  });

  it("can change auto-fetching mode", () => {
    cy.visit("/");
    cy.contains(".q-toggle__label", "Auto-Fetch");
    cy.get(".q-toggle__inner").click();
    cy.contains(".q-toggle__label", "Auto-Fetching");
  });

  it("can display Aggregated datatable", () => {
    cy.visit("/");
    cy.contains('[data-cy="datatable-label"]', "Aggregated data");
  });

  it("can build filter selectors", () => {
    cy.visit("/");
    cy.get('[data-cy="filter-selector"]').first().click();
    cy.get('[data-cy="filter-selector-option"]').first().contains("Etc/GMT-10");
  });
});

Cypress.env("MODE");

describe("Information section Test Case", () => {
  it("can load dashboard page", () => {
    cy.visit("/");
    cy.contains("h3", "International Space Station dashboard");
  });

  it("can open information tree", () => {
    cy.visit("/");
    cy.get(".q-tree > :nth-child(1) > :nth-child(1)").click();
    cy.get(
      '[style=""] > :nth-child(1) > :nth-child(1) > .q-tree__node--link'
    ).contains("Auto-fetching feature");
  });
});

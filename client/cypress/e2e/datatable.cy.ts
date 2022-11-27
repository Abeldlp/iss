Cypress.env("MODE");

describe("Datatable Test Case", () => {
  it("can display Aggregated datatable", () => {
    cy.visit("/");
    cy.contains('[data-cy="datatable-label"]', "Aggregated data");
  });

  it("can build filter selectors", () => {
    cy.visit("/");
    cy.get('[data-cy="filter-selector"]').first().click();
    cy.get('[data-cy="filter-selector-option"]').first().contains("Etc/GMT-10");
  });

  it("can sort order of table", () => {
    cy.visit("/");
    cy.get(".q-table > tbody > :nth-child(1) > :nth-child(1)").contains(
      "Etc/GMT-10"
    );
    cy.get(".q-table > thead > tr > :nth-child(1)").first().click();
    cy.get(".q-table > tbody > :nth-child(1) > :nth-child(1)").contains(
      "Asia/Oral"
    );
  });

  it("can paginate datatable", () => {
    cy.visit("/");
    cy.get(".q-table > tbody > :nth-child(1) > :nth-child(1)").contains(
      "Etc/GMT-10"
    );
    cy.get(
      ".q-table__bottom > :nth-child(3) > .q-btn--actionable > .q-btn__content"
    )
      .first()
      .click();
    cy.get(".q-table > tbody > :nth-child(1) > :nth-child(1)").contains(
      "Asia/Ulaanbaatar"
    );
  });
});

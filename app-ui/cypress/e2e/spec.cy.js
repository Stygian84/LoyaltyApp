describe("displaying page", () => {
  it("should contain the right number of elements ", () => {
    cy.visit("http://localhost:3000");
    cy.get(".card").each(($el, index, $list) => {
      cy.wrap($el).children().children().should("have.length", 5);
    });
  });
  it("should contain a transfer credit button", () => {
    cy.visit("http://localhost:3000");
    cy.get(".card").each(($el, index, $list) => {
      cy.wrap($el).get("div> div>button").contains("Transfer credits");
    });
  });
  it("should display the right number of cards", () => {
    cy.visit("http://localhost:3000");
    cy.request("http://localhost:8080/loyalty").then((response) => {
      cy.get(".card").should("have.length", response.body.length);
    });
  });
});

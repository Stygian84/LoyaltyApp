const userName = "gracecole";
describe("displaying page", () => {
  it("should contain the right number of elements ", () => {
    cy.visit("http://localhost:3000");
    cy.get("#username").type(userName);
    cy.get("#password").type("password");
    cy.get("#loginButton").click();

    cy.get(".offerCard").each(($el,index,$list)=>{
      cy.wrap($el).children().children().should("have.length",6)
    })

    cy.get(".normalCard").each(($el, index, $list) => {
      cy.wrap($el).children().children().should("have.length", 5);
    });
  });
  it("should contain a transfer credit button", () => {
    // cy.visit("http://localhost:3000");
    cy.get(".normalCard").each(($el, index, $list) => {
      cy.wrap($el).get("div> div>button").contains("Transfer credits");
    });
  });
  it("should display the right number of cards", () => {
    // cy.visit("http://localhost:3000");
    cy.request("http://localhost:8080/loyalty").then((response) => {
      cy.get(".card").should("have.length", response.body.length);
    });
  });
});
describe("transfer credit", () => {
  const creditValue = 20;
  it("should navigate to transaction page and validate an invalid membership id", () => {
    cy.get(".normalCard:first>div>button").wait(8000).click()
    
    
   
  });
  it("should validate an valid membership id", () => {
    cy.get("#membershipID").clear().type("1005610");
    cy.get("#membershipIDButton").click();
    cy.get("#membershipValidity").contains("Membership ID is valid");
  });
  it("Should check reward", () => {
    cy.get("#creditInput").type(creditValue);
    cy.get("#creditInputButton").click();
    cy.get("#expectedReward").wait(8000).contains("Your reward expected is: ");
  });
  it("should submit credit request", () => {
    cy.get("#submitCreditRequest").click();
    cy.get("#referenceDoc")
      .children()
      .should(($el) => {
        expect($el).to.have.length(5);
        expect($el.first().text()).to.match(/^Reference Number: \d{2,3}$/);
      });
  });
  it("should update the balance", () => {
    cy.request(`http://localhost:8080/getUserbyUsername/${userName}`).then(
      (response) => {
        cy.get("#balance").should(
          "have.text",
          response.body.credit_balance + creditValue
        );
      }
    );
  });
});

describe("transaction status page", () => {
  it("go to trasaction status page ", () => {
    cy.go('back').wait(8000);
    cy.get("#transactionPageButton").click();

  });

  it("display correct number of status", () => {
    cy.request(`http://localhost:8080/getUserbyUsername/${userName}`).then((response) => {
      cy.request(`http://localhost:8080/transaction_status/${response.body.id}`).then((response) => {
        cy.get(".card").should("have.length", response.body.length);
      }
      
  )});

  })

  it("correct number of elements in status", ()=>{
    cy.get(".card").each(($el, index, $list) => {
      cy.wrap($el).children().children().should("have.length", 8);
    });
  })
})

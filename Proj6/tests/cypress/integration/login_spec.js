import { constants } from '../constants';
const url = constants.URL

describe('Login tests', function() {
    it('Empty data test', function(){
        cy.visit(url)
        cy.contains("Zaloguj").should('exist')
        cy.contains("Zaloguj").click()
        cy.url().should('include', '/login')
        cy.contains("Login").should('exist')
        cy.contains("Login").click()
        assert(cy.contains("Nick jest wymaganym polem"))
        assert(cy.contains("Hasło jest wymaganym polem"))
        cy.contains("Zaloguj").should('exist')
    })
    
    it('Invalid data test', function(){
        cy.visit(url)
        cy.contains("Zaloguj").should('exist')
        cy.contains("Zaloguj").click()
        cy.url().should('include', '/login')
        cy.get('input[name=username]').type('test')
        cy.get('input[name=password]').type('test')
        cy.contains("Login").should('exist')
        cy.contains("Login").click()
        assert(cy.contains("Niepoprawny login lub hasło!"))
        cy.contains("Wyloguj").should('not.exist')
        cy.contains("Zaloguj").should('exist')
    })
    
    it('Valid data test', function(){
        cy.visit(url)
        cy.contains("Zaloguj").should('exist')
        cy.contains("Zaloguj").click()
        cy.url().should('include', '/login')
        cy.get('input[name=username]').type('janusz_nosacz')
        cy.get('input[name=password]').type('janusz123')
        cy.contains("Login").should('exist')
        cy.contains("Login").click()
        cy.contains("Wyloguj")
        cy.contains("Wyloguj").click()
    })
    
    it('Register link test', function(){
        cy.visit(url)
        cy.contains("Zaloguj").should('exist')
        cy.contains("Zaloguj").click()
        cy.url().should('include', '/login')
        cy.contains("Nie masz konta? Zarejestruj się").click()
        cy.url().should('include', '/register')
    })

    it('Logged in exam show', function(){
        cy.visit(url)
        cy.contains("Zaloguj").should('exist')
        cy.contains("Zaloguj").click()
        cy.url().should('include', '/login')
        cy.get('input[name=username]').type('janusz_nosacz')
        cy.get('input[name=password]').type('janusz123')
        cy.contains("Login").should('exist')
        cy.contains("Login").click()
        cy.get(".MuiToolbar-root > :nth-child(2) > .active").click()
        cy.get('[href="/tests"]').click()
        cy.url().should('include', '/tests')
    })
})
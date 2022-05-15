import { constants } from '../constants';
const url = constants.URL

describe('Simple pages', function() {
    it('Main page', function(){
        cy.visit(url)
        cy.contains('Home Screen')
        cy.contains('WPZ-Egzaminy')
        cy.contains('Strona Główna')
        cy.contains('Egzaminy')
        cy.contains('FAQ')
        cy.contains('Zaloguj')
        cy.contains('Zarejestruj')
    })
    
    it('Faq page main', function(){
        cy.visit(url)
        cy.contains('FAQ').click()
        cy.url().should('include', '/about')
        cy.contains('Wkraczamy w nową erę')
    })
    
    it('Faq page click plus', function(){
        cy.visit(url)
        cy.contains('FAQ').click()
        cy.contains('Do założenia konta wystarczy')
        cy.contains('Czy mogę podejść').click()
        cy.contains('Do założenia konta wystarczy').
        should('to.be.hidden')
    })
    
    it('Egzaminy page', function(){
        cy.visit(url)
        cy.get('[href="/tests"]').click()
        cy.url().should('include', '/tests')
    })

})
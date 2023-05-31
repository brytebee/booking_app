package main

import (
	"fmt"
	"strings"
)

const eventTickets = 50
var eventName = "Go Summit"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookings, remainingTickets = bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()

			fmt.Printf("The first names of the bookings: %v.\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("The tickets are sold off. Book earlier next year!")
				break
			}
		} else {
			printValidations(isValidName, isValidEmail, isValidTicketNumber)
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v application.\n", eventName)
	fmt.Printf("There are %v tickets and %v tickets are still available.\n", eventTickets, remainingTickets)
	fmt.Println("Make your bookings below.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var fullName = strings.Fields(booking)
		firstNames = append(firstNames, fullName[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) ([]string, uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email would be sent to %v.\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets are left for %v.\n", remainingTickets, eventName)

	return bookings, remainingTickets
}

func printValidations(isValidName bool, isValidEmail bool, isValidTicketNumber bool) {
	if !isValidName {
		fmt.Println("First or last name is invalid, try again")
	}
	if !isValidEmail {
		fmt.Println("Email is invalid, try again")
	}
	if !isValidTicketNumber {
		fmt.Println("Invalid number of tickets, try again")
	}
}

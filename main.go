package main

import (
	"fmt"
	"time"
)

const eventTickets = 50
var eventName = "Go Summit"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// bookings, remainingTickets = bookTicket(userTickets, firstName, lastName, email)
			bookTicket(userTickets, firstName, lastName, email)
			go sendMails(userTickets, firstName, lastName, email)

			fmt.Printf("List of bookings: %v\n", bookings)

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
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) ([]UserData, uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	
	bookings = append(bookings, userData)

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

func sendMails(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("\n%v for %v %v\nTo %v\n", userTickets, firstName, lastName, email)
	fmt.Println("#############")
	fmt.Printf("Sending tickets: %v\n", ticket)
	fmt.Println("#############")
}

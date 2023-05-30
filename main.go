package main

import (
	"fmt"
	"strings"
)

func main() {
	var eventName = "Go Summit"
	const eventTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v application.\n", eventName)
	fmt.Printf("There are %v tickets and %v tickets are still available.\n", eventTickets, remainingTickets)

	fmt.Println("Make your bookings below.")

	for {
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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email would be sent to %v.\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets are left for %v.\n", remainingTickets, eventName)

			firstNames := []string{}
			for _, booking := range bookings {
				var fullName = strings.Fields(booking)
				firstNames = append(firstNames, fullName[0])
			}

			fmt.Printf("The first names of the bookings: %v.\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("The tickets are sold off. Book earlier next year!")
				break
			}
		} else {
			fmt.Printf("There's only %v remaiming, you can't book %v tickecks\n", remainingTickets, userTickets)
		}
	}
}

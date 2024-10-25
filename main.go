package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conferfence"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcom to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")

	bookings := []string{}

	for {
		// user input
		var firstName string
		var lastName string
		var email string
		var userTickets int

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >=2 && len(lastName) >=2
		isVaidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)



		// check if userTickets is greater than remainingTickets
		if isVaidEmail && isValidName && isValidTicketNumber {
			//booking tickets

			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v who booked %v tickets. You will receive your tickets per email: %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("We have %v tickets left\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstName := names[0]
				firstNames = append(firstNames, firstName)
			}
			fmt.Println("First names of all bookings: ", firstNames)
			fmt.Println("Current bookings: ", bookings)
			noTickets := remainingTickets == 0
			if noTickets {
				fmt.Println("All tickets are sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Invalid first name or last name. Try again\n")
			}
			if !isVaidEmail {
				fmt.Printf("Invalid email. Try again\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Invalid number of tickets. Try again\n")
			}
		}
	}
}

package main

import (
	"fmt"
	"time"
	"sync"
)

const conferenceTickets int = 50

var conferenceName = "Go Conferfence"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		// user input
		firstName, lastName, email, userTickets := getUserInput()

		// check if userTickets is greater than remainingTickets
		isValidName, isValidEmail, isValidTicketNumber := ValidUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//booking tickets
			bookingTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstName()
			fmt.Println("First names of all bookings:", firstNames)
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
			if !isValidEmail {
				fmt.Printf("Invalid email. Try again\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Invalid number of tickets. Try again\n")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcom to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// user input
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookingTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)

	//create map for users
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: uint(userTickets),
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of all bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v who booked %v tickets. You will receive your tickets per email: %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets left\n", remainingTickets)
	wg.Wait()
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// send ticket to user
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", remainingTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket\n %v \nto email address %v.\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}

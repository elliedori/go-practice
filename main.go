package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Freedom Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	var bookings []string

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames(bookings)
			fmt.Printf("All booking first names: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out! Come back next year :)")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name and last name need to be at least 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Emails must contain an '@' sign")
			}
			if !isValidTicketNumber {
				fmt.Printf("The number of tickets you entered is invalid, we have %v tickets available\n", remainingTickets)
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, with %v still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}
func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName, lastName, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets
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

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you'd like")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName, lastName, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
}

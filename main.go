package main

import "fmt"
import "strings"

func main() {
	conferenceName := "Freedom Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remaining is %T\n", conferenceName, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, with %v still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

	var bookings []string

	for {

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

		if userTickets < remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets remaining\n", remainingTickets)

			firstNames := []string{}

			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("All booking first names: %v\n", firstNames)
			if remainingTickets == 0 {
				// end the program
				fmt.Println("All tickets are sold out! Come back next year :)")
				break
			}
		} else {
			fmt.Printf("Sorry, we only have %v tickets available\n", remainingTickets)
		}
	}
}

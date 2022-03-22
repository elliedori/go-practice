package main

import "fmt"

func main() {
	conferenceName := "Freedom Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remaining is %T\n", conferenceName, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, with %v still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

	var userName string
	var userTickets int
	// this is a comment

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}

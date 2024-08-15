package main

import "fmt"

func main() {
	var confrenceName = "GO CONFRENCE"
	const totalTickets = 50
	var remainingTickets = totalTickets
	var bookings []string
	fmt.Printf("\n")
	fmt.Printf("********************\n")
	fmt.Printf("********************\n")
	fmt.Printf(" %v BOOKING\n", confrenceName)
	fmt.Printf("     %v left     \n", remainingTickets)
	fmt.Printf("********************\n")
	fmt.Printf("********************\n\n")

	for {

		var userName string = ""
		var userEmail string = ""
		var userTickets int = 0

		fmt.Printf("Enter your name\n")
		fmt.Scan(&userName)

		fmt.Printf("Enter your email\n")
		fmt.Scan(&userEmail)

		fmt.Printf("Enter the number of tickets to be booked\n")
		fmt.Scan(&userTickets)

		fmt.Printf("Thanks for booking %v tickets %v\n\n", userTickets, userName)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, userName)

		fmt.Printf("****************************************************\n")
		fmt.Printf("check your email (%v) for confirmation\n", userEmail)
		fmt.Printf("****************************************************\n\n")

		fmt.Printf("Bookings :%v\n", bookings)

		fmt.Printf("***** %v tickets left ***** \n", remainingTickets)
		fmt.Printf("______________________________________________________\n\n")
	}

}

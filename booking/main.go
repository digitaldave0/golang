package main

import (
	"booking/helper"
	"fmt"
)


const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

//using structs

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				//end
				fmt.Println("Sorry all our Tickets are all booked out")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address does not include @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}

			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			fmt.Println("Sorry your data was wrong, try again")
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	//blank identifer _
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

	// ask user name

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your surname name: ")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email adresss: ")
	fmt.Scan(&email)
	fmt.Println("Please enter number of tickets ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//slice

	//map
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	//fmt.Printf("The whole Slice: %v\n", bookings)
	//fmt.Printf("The first Slice: %v\n", bookings[0])
	//fmt.Printf("Slice type: %T\n", bookings)
	//fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)

}
